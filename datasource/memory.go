package datasource

import (
	"minesweeper-API/minesweeper-service/model"
	"sync"
)

type memoryMap struct {
	sync.RWMutex
	m map[int]*model.Game
}

type memory struct {
	memoryMap memoryMap
}

func NewMemory() Spec {
	return &memory{
		memoryMap: memoryMap{
			m: make(map[int]*model.Game),
		},
	}
}

//TODO deberia retornar una copia? si retorna un puntero puede dar problemas de concurrencia?
func (m *memory) FindGame(id int) (*model.Game, error) {
	m.memoryMap.RLock()
	defer m.memoryMap.RUnlock()
	if game, ok := m.memoryMap.m[id]; ok {
		return game, nil
	}
	return nil, nil
}

func (m *memory) InsertGame(game *model.Game) (int, error) {
	m.memoryMap.Lock()
	m.memoryMap.m[len(m.memoryMap.m)] = game
	m.memoryMap.Unlock()
	return len(m.memoryMap.m) - 1, nil
}

func (m *memory) UpdateGame(g *model.Game) error {
	m.InsertGame(g)
	return nil
}
