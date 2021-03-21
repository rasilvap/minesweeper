package engine

import (
	"log"

	"minesweeper-API/minesweeper-service/datasource"

	"minesweeper-API/minesweeper-service/model"

	"github.com/obarra-dev/minesweeper"
)

type GameService interface {
	GetOneGame(id int) (*model.GameResponse, error)
	CreateGame(rows, colums, mineAmount int) (int, error)
	PlayMove(id int, playRequest model.PlayRequest) (*model.PlayResponse, error)
}

type service struct{}

var (
	gameRepository datasource.GameRepository
	minesWeeper    MinesWeeperService
)

func NewGameService(gameMomoryRepository datasource.GameRepository, minesWeeperService MinesWeeperService) GameService {
	gameRepository = gameMomoryRepository
	minesWeeper = minesWeeperService
	return &service{}
}

func (*service) GetOneGame(id int) (*model.GameResponse, error) {
	game := gameRepository.Get(id)
	gameResponse := model.GameResponse{game.Rows, game.Columns, game.MineAmount}
	return &gameResponse, nil
}

func (*service) CreateGame(rows, colums, mineAmount int) (int, error) {
	game := minesWeeper.BuildGame(rows, colums, mineAmount)
	id := gameRepository.Save(game)
	return id, nil
}

func (*service) PlayMove(id int, playRequest model.PlayRequest) (*model.PlayResponse, error) {
	game := gameRepository.Get(id)
	log.Println("PlayRequest", playRequest)

	showableGame := minesWeeper.Play(playRequest, game)
	gameRepository.Save(game)

	log.Println("show: ", showableGame)
	playResponse := buildPlayResponse(showableGame)

	return &playResponse, nil
}

func buildPlayResponse(showableGame minesweeper.Game) model.PlayResponse {
	gameStateDTO := mapStateGame(showableGame.State)
	row := len(showableGame.Board)
	if row == 0 {
		return model.PlayResponse{gameStateDTO,
			model.GameDTO{[][]model.TileDTO{}, showableGame.Rows, showableGame.Columns, showableGame.FlagAmount}}
	}

	boardDTO := make([][]model.TileDTO, row)
	for i := 0; i < row; i++ {
		column := len(showableGame.Board[i])
		boardDTO[i] = make([]model.TileDTO, column)

		for j := 0; j < column; j++ {
			board := showableGame.Board[i][j]
			tileStateDTO := mapTileState(board.State)
			boardDTO[i][j] = model.TileDTO{tileStateDTO, board.Row, board.Column, board.SurroundingMineCount, board.IsMine, -1}
		}
	}

	return model.PlayResponse{gameStateDTO,
		model.GameDTO{boardDTO, showableGame.Rows, showableGame.Columns, showableGame.FlagAmount}}
}
