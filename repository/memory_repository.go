package repository

import (
	"sync"

	"github.com/obarra-dev/minesweeper"
)

type GameRepository interface {
	Save(game *minesweeper.Game) int
	Get(id int) *minesweeper.Game
}

type gameRepositoryMemory struct{}

type gameMemoryMap struct {
	sync.RWMutex
	m map[int]*minesweeper.Game
}

var (
	gameStorageMap gameMemoryMap
)

func NewMemoryRepository() GameRepository {
	gameStorageMap = gameMemoryMap{m: make(map[int]*minesweeper.Game)}
	return &gameRepositoryMemory{}
}

func (*gameRepositoryMemory) Save(game *minesweeper.Game) int {
	gameStorageMap.Lock()
	gameStorageMap.m[len(gameStorageMap.m)] = game
	gameStorageMap.Unlock()
	return len(gameStorageMap.m) - 1
}

//TODO deberia retornar una copia? si retorna un puntero puede dar problemas de concurrencia?
func (*gameRepositoryMemory) Get(id int) *minesweeper.Game {
	gameStorageMap.RLock()
	defer gameStorageMap.RUnlock()
	if game, ok := gameStorageMap.m[id]; ok {
		return game
	}
	return nil
}
