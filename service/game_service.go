package service

import (
	"log"
	"minesweeper-API/minesweeper-service/repository"

	"minesweeper-API/minesweeper-service/engine"
	"minesweeper-API/minesweeper-service/model"
)

type GameService interface {
	GetOneGame(id int) (*model.GameResponse, error)
	CreateGame(rows, colums, mineAmount int) (int, error)
	PlayMove(id int, playRequest model.PlayRequest) (*model.PlayResponse, error)
}

type service struct{}

var (
	gameRepository repository.GameRepository
)

func NewGameService(gameMomoryRepository repository.GameRepository) GameService {
	gameRepository = gameMomoryRepository
	return &service{}
}

func (*service) GetOneGame(id int) (*model.GameResponse, error) {
	game := gameRepository.Get(id)
	gameResponse := model.GameResponse{game.Rows, game.Columns, game.MineAmount}
	return &gameResponse, nil
}

func (*service) CreateGame(rows, colums, mineAmount int) (int, error) {
	minedPointTile := engine.GenerateMinedPoints(mineAmount, rows, colums)
	game := engine.BuildNewGame(rows, colums, minedPointTile)
	log.Println(game)

	id := gameRepository.Save(game)
	return id, nil
}

func (*service) PlayMove(id int, playRequest model.PlayRequest) (*model.PlayResponse, error) {
	game := gameRepository.Get(id)
	log.Println("PlayRequest", playRequest)

	showableGame := game.Play(playRequest.Row, playRequest.Column, mapTypeMove(playRequest.Move))
	gameRepository.Save(game)

	log.Println("show: ", showableGame)
	playResponse := buildPlayResponse(showableGame)

	return &playResponse, nil
}

func mapTypeMove(typeMove model.TypeMove) engine.TypeMove {
	var move engine.TypeMove
	switch typeMove {
	case model.TypeMoveFlag:
		move = engine.TypeMoveFlag
	case model.TypeMoveQuestion:
		move = engine.TypeMoveQuestion
	case model.TypeMoveClean:
		move = engine.TypeMoveClean
	}

	return move
}

func buildPlayResponse(showableGame engine.Game) model.PlayResponse {
	var gameStateDTO string
	switch showableGame.State {
	case engine.StateGameRunning:
		gameStateDTO = "RUNNING"
	case engine.StateGameLost:
		gameStateDTO = "LOST"
	case engine.StateGameNew:
		gameStateDTO = "NEW"
	case engine.StateGameWon:
		gameStateDTO = "WON"
	default:
		gameStateDTO = ""
	}
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

			var tileStateDTO string
			switch board.State {
			case engine.StateTileCovered:
				tileStateDTO = "COVERED"
			case engine.StateTileClear:
				tileStateDTO = "CLEAR"
			case engine.StateTileFlagged:
				tileStateDTO = "FLAGGED"
			case engine.StateTileNumberd:
				tileStateDTO = "NUMBERED"
			case engine.StateTileExploted:
				tileStateDTO = "EXPLOTED"
			default:
				tileStateDTO = ""
			}

			boardDTO[i][j] = model.TileDTO{tileStateDTO, board.Row, board.Column, board.SurroundingMineCount, board.IsMine, -1}
		}
	}

	return model.PlayResponse{gameStateDTO,
		model.GameDTO{boardDTO, showableGame.Rows, showableGame.Columns, showableGame.FlagAmount}}
}
