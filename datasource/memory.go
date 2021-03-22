package datasource

import (
	"sync"

	"github.com/obarra-dev/minesweeper"
)

type gameRepositoryMemory struct{}

type gameMemoryMap struct {
	sync.RWMutex
	m map[int]*minesweeper.Game
}

var (
	gameStorageMap gameMemoryMap
)

func NewMemoryRepository() Spec {
	gameStorageMap = gameMemoryMap{m: make(map[int]*minesweeper.Game)}
	return &gameRepositoryMemory{}
}

func (*gameRepositoryMemory) SaveGame(game *minesweeper.Game) (int, error) {
	gameStorageMap.Lock()
	gameStorageMap.m[len(gameStorageMap.m)] = game
	gameStorageMap.Unlock()
	return len(gameStorageMap.m) - 1, nil
}

//TODO deberia retornar una copia? si retorna un puntero puede dar problemas de concurrencia?
func (*gameRepositoryMemory) GetGame(id int) (*minesweeper.Game, error) {
	gameStorageMap.RLock()
	defer gameStorageMap.RUnlock()
	if game, ok := gameStorageMap.m[id]; ok {
		return game, nil
	}
	return nil, nil
}
