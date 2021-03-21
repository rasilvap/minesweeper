package engine

import (
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

type service struct{}

var (
	gameRepository datasource.GameRepository
	minesWeeper    MinesWeeperService
)

func NewGameService(gameRepositoryImp datasource.GameRepository, minesWeeperService MinesWeeperService) GameService {
	gameRepository = gameRepositoryImp
	minesWeeper = minesWeeperService
	return &service{}
}

func (*service) GetOneGame(id int) (*model.GameResponse, error) {
	game := gameRepository.Get(id)
	gameResponse := model.GameResponse{Rows: game.Rows, Columns: game.Columns, MineAmount: game.MineAmount}
	return &gameResponse, nil
}

func (*service) CreateGame(rows, columns, mineAmount int) (int, error) {
	game := minesWeeper.BuildGame(rows, columns, mineAmount)
	id := gameRepository.Save(game)
	return id, nil
}

func (*service) PlayMove(id int, playRequest model.PlayRequest) (*model.PlayResponse, error) {
	game := gameRepository.Get(id)
	log.Println("PlayRequest", playRequest)

	visibleGame := minesWeeper.Play(playRequest, game)
	gameRepository.Save(game)

	log.Println("show: ", visibleGame)
	playResponse := buildPlayResponse(visibleGame)

	return &playResponse, nil
}

func buildPlayResponse(game minesweeper.Game) model.PlayResponse {
	gameStateDTO := mapStateGame(game.State)
	row := len(game.Board)
	if row == 0 {
		return model.PlayResponse{gameStateDTO,
			model.GameDTO{[][]model.TileDTO{}, game.Rows, game.Columns, game.FlagAmount}}
	}

	boardDTO := make([][]model.TileDTO, row)
	for i := 0; i < row; i++ {
		column := len(game.Board[i])
		boardDTO[i] = make([]model.TileDTO, column)

		for j := 0; j < column; j++ {
			board := game.Board[i][j]
			tileStateDTO := mapTileState(board.State)
			boardDTO[i][j] = model.TileDTO{tileStateDTO, board.Row, board.Column, board.SurroundingMineCount, board.IsMine, -1}
		}
	}

	return model.PlayResponse{gameStateDTO,
		model.GameDTO{boardDTO, game.Rows, game.Columns, game.FlagAmount}}
}
