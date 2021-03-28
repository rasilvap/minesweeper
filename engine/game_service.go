package engine

import (
	"encoding/json"
	"log"

	"minesweeper-API/minesweeper-service/datasource"

	"minesweeper-API/minesweeper-service/model"

	"github.com/obarra-dev/minesweeper"
)

type GameService interface {
	GetOneGame(id int) (*model.GameResponse, error)
	CreateGame(rows, columns, mineAmount int) (int, error)
	PlayMove(id int, playRequest model.PlayRequest) (*model.PlayResponse, error)
}

type service struct {
	gameDS      datasource.Spec
	minesWeeper MinesWeeperService
}

func NewGameService(gameDSImp datasource.Spec, minesWeeperService MinesWeeperService) GameService {
	return &service{
		gameDS:      gameDSImp,
		minesWeeper: minesWeeperService,
	}
}

func (s service) CreateGame(rows, columns, mineAmount int) (int, error) {
	g := s.minesWeeper.BuildGame(rows, columns, mineAmount)
	gameDS, err := buildGameDS(g)
	if err != nil {
		log.Printf("Error marshing board, err: %v", err)
		return 0, err
	}

	id, err := s.gameDS.SaveGame(gameDS)
	if err != nil {
		log.Printf("Error creating game, err: %v", err)
		return 0, err
	}

	return id, nil
}

func (s service) GetOneGame(id int) (*model.GameResponse, error) {
	g, err := s.gameDS.FindGame(id)
	if err != nil {
		log.Printf("Error finding game: %d, err: %v", id, err)
		return nil, err
	}

	return &model.GameResponse{
			Rows:       g.Rows,
			Columns:    g.Columns,
			MineAmount: g.MineAmount,
		},
		nil
}

func (s *service) PlayMove(id int, playRequest model.PlayRequest) (*model.PlayResponse, error) {
	log.Println("Playing game", playRequest)
	gameDS, err := s.gameDS.FindGame(id)
	if err != nil {
		log.Printf("Error finding g: %d, err: %v", id, err)
		return nil, err
	}

	var board [][]minesweeper.Tile
	err = json.Unmarshal([]byte(gameDS.Board), &board)
	if err != nil {
		log.Printf("Error unmarshaling board, err: %v", err)
		return nil, err
	}

	g := minesweeper.Game{
		State:      1,
		Rows:       gameDS.Rows,
		Columns:    gameDS.Columns,
		MineAmount: gameDS.MineAmount,
		FlagAmount: gameDS.FlagAmount,
		Board:      board,
	}

	visibleGame := s.minesWeeper.Play(playRequest, &g)

	gameDS, err = buildGameDS(&visibleGame)
	if err != nil {
		log.Printf("Error marshing board, err: %v", err)
		return nil, err
	}

	s.gameDS.UpdateGame(gameDS)

	log.Println("show: ", visibleGame)
	playResponse := buildPlayResponse(visibleGame)
	return &playResponse, nil
}

func buildGameDS(g *minesweeper.Game) (*model.Game, error) {
	j, err := json.Marshal(g.Board)
	if err != nil {
		return nil, err
	}

	return &model.Game{
			State:      string(g.State),
			Columns:    g.Columns,
			Rows:       g.Rows,
			MineAmount: g.MineAmount,
			FlagAmount: g.FlagAmount,
			Board:      string(j),
		},
		nil
}

func buildPlayResponse(game minesweeper.Game) model.PlayResponse {
	gameStateDTO := mapStateGame(game.State)
	row := len(game.Board)
	if row == 0 {
		return model.PlayResponse{
			StateGame: gameStateDTO,
			Game: model.GameDTO{Board: [][]model.TileDTO{},
				Rows:       game.Rows,
				Columns:    game.Columns,
				FlagAmount: game.FlagAmount,
			},
		}
	}

	boardDTO := make([][]model.TileDTO, row)
	for i := 0; i < row; i++ {
		column := len(game.Board[i])
		boardDTO[i] = make([]model.TileDTO, column)

		for j := 0; j < column; j++ {
			board := game.Board[i][j]
			tileStateDTO := mapTileState(board.State)
			boardDTO[i][j] = model.TileDTO{
				State:                tileStateDTO,
				Row:                  board.Row,
				Column:               board.Column,
				SurroundingMineCount: board.SurroundingMineCount,
				Mine:                 board.IsMine,

				ValueTest: -1}
		}
	}

	return model.PlayResponse{
		StateGame: gameStateDTO,
		Game: model.GameDTO{
			Board:      boardDTO,
			Rows:       game.Rows,
			Columns:    game.Columns,
			FlagAmount: game.FlagAmount,
		},
	}
}
