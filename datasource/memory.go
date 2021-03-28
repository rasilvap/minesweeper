package datasource

import (
	"minesweeper-API/minesweeper-service/model"
	"sync"
)

type gameRepositoryMemory struct{}

func (m *gameRepositoryMemory) UpdateGame(g *model.Game) error {
	panic("implement me")
}

type gameMemoryMap struct {
	sync.RWMutex
	m map[int]*model.Game
}

var (
	gameStorageMap gameMemoryMap
)

func NewMemoryRepository() Spec {
	gameStorageMap = gameMemoryMap{m: make(map[int]*model.Game)}
	return &gameRepositoryMemory{}
}

func (*gameRepositoryMemory) SaveGame(game *model.Game) (int, error) {
	gameStorageMap.Lock()
	gameStorageMap.m[len(gameStorageMap.m)] = game
	gameStorageMap.Unlock()
	return len(gameStorageMap.m) - 1, nil
}

//TODO deberia retornar una copia? si retorna un puntero puede dar problemas de concurrencia?
func (*gameRepositoryMemory) FindGame(id int) (*model.Game, error) {
	gameStorageMap.RLock()
	defer gameStorageMap.RUnlock()
	if game, ok := gameStorageMap.m[id]; ok {
		return game, nil
	}
	return nil, nil
}
