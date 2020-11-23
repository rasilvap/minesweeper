package service

import (
	"log"
	"sync"

	"minesweeper-API/minesweeper-service/engine"
	"minesweeper-API/minesweeper-service/model"
)

var gameStorageMap = struct {
	sync.RWMutex
	m map[int]engine.Game
}{m: make(map[int]engine.Game)}

func GetOneGame(id int) (*model.GameResponse, error) {
	gameStorageMap.RLock()
	defer gameStorageMap.RUnlock()
	if game, ok := gameStorageMap.m[id]; ok {
		gameResponse := model.GameResponse{game.Rows, game.Columns, 44}
		return &gameResponse, nil
	}
	return nil, nil
}

func CreateGame(rows, colums, mineAmount int) (int, error) {
	game := engine.BuildNewGame(rows, colums, mineAmount)
	minedPointTile := [][2]int{{1, 1}}
	game.SetUpMines(minedPointTile)
	log.Println(game)

	gameStorageMap.Lock()
	gameStorageMap.m[len(gameStorageMap.m)] = game
	gameStorageMap.Unlock()
	return len(gameStorageMap.m) - 1, nil
}

//TODO use mark for flag ant ???
func MarkTile(id int, row int, column int, mark string) error {
	gameStorageMap.RLock()
	defer gameStorageMap.RUnlock()
	if game, ok := gameStorageMap.m[id]; ok {
		log.Println(game.GetStates())
		log.Println("row", row, "col", column)
		game.MarkFlag(row, column)
		log.Println(game.GetStates())
	}
	return nil
}

func PlayMovement(id int, row int, column int) (*model.PlayResponse, error) {
	gameStorageMap.RLock()
	defer gameStorageMap.RUnlock()
	if game, ok := gameStorageMap.m[id]; ok {
		log.Println(game.GetStates())
		log.Println("Req row", row, "col", column)
		gameState, showableGame := game.PlayMovement(row, column)
		log.Println("show: ", gameState, showableGame)
		playResponse := buildPlayResponse(gameState, showableGame)

		return &playResponse, nil
		log.Println(game.GetStates())
	}
	return nil, nil
}

func buildPlayResponse(gameState engine.StateGame, showableGame engine.Game) model.PlayResponse {
	var gameStateDTO string
	switch gameState {
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

func GetCompleteGame(id int) *model.GameCompleteResponse {
	gameStorageMap.RLock()
	defer gameStorageMap.RUnlock()
	if game, ok := gameStorageMap.m[id]; ok {
		return &model.GameCompleteResponse{game}
	}
	return nil
}

func DeleteGame(id int) error {
	return nil
}
