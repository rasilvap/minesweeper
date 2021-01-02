package repository

import (
	"minesweeper-API/minesweeper-service/engine"
	"sync"
)

type GameRepository interface {
	Save(game *engine.Game) int
	Get(id int) *engine.Game
}

type gameRepositoryMemory struct{}

type gameMemoryMap struct {
	sync.RWMutex
	m map[int]*engine.Game
}

var (
	gameStorageMap gameMemoryMap
)

func NewMemoryRepository() GameRepository {
	gameStorageMap = gameMemoryMap{m: make(map[int]*engine.Game)}
	return &gameRepositoryMemory{}
}

func (*gameRepositoryMemory) Save(game *engine.Game) int {
	gameStorageMap.Lock()
	gameStorageMap.m[len(gameStorageMap.m)] = game
	gameStorageMap.Unlock()
	return len(gameStorageMap.m) - 1
}

//TODO deberia retornar una copia? si retorna un puntero puede dar problemas de concurrencia?
func (*gameRepositoryMemory) Get(id int) *engine.Game {
	gameStorageMap.RLock()
	defer gameStorageMap.RUnlock()
	if game, ok := gameStorageMap.m[id]; ok {
		return game
	}
	return nil
}
