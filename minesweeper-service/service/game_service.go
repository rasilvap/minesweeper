package service

import (
	"log"
	"sync"

	"minesweeper-API/minesweeper-service/engine"
	"minesweeper-API/minesweeper-service/model"
)

// used to hold our product list in memory
var gameStorageMap = struct {
	sync.RWMutex
	m map[int]engine.Game
}{m: make(map[int]engine.Game)}

func CreateGame(rows, colums, mineAmount int) (int, error) {
	game := engine.BuildNewGame(rows, colums)
	minedPointTile := [][2]int{{0, 1}, {1, 1}, {1, 0}}
	game.SetUpMines(mineAmount, minedPointTile)
	log.Println(game)

	gameStorageMap.Lock()
	gameStorageMap.m[2] = game
	gameStorageMap.Unlock()
	return 2, nil
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
		log.Println("row", row, "col", column)
		game.PlayMovement(row, column)

		return &model.PlayResponse{game}, nil
		log.Println(game.GetStates())
	}
	return nil, nil
}

func GetOneGame(id int) (*model.GameResponse, error) {
	gameStorageMap.RLock()
	defer gameStorageMap.RUnlock()
	if game, ok := gameStorageMap.m[id]; ok {
		gameResponse := model.GameResponse{game.Rows, game.Columns, 44}
		return &gameResponse, nil
	}
	return nil, nil
}

func DeleteGame(id int) error {
	return nil
}
