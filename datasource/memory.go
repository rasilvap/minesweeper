package datasource

import (
	"minesweeper-API/minesweeper-service/model"
	"sync"
)

type memory struct {
	sync.RWMutex
	cache map[int]interface{}
}

func NewMemory() Spec {
	return &memory{
		cache: make(map[int]interface{}),
	}
}

//TODO deberia retornar una copia? si retorna un puntero puede dar problemas de concurrencia?
func (m *memory) FindGame(id int) (*model.Game, error) {
	m.RLock()
	defer m.RUnlock()
	if game, ok := m.cache[id]; ok {
		return game.(*model.Game), nil
	}
	return &model.Game{}, nil
}

func (m *memory) InsertGame(game *model.Game) (int, error) {
	m.Lock()
	game.GameId = len(m.cache) + 1
	m.cache[game.GameId] = game
	m.Unlock()
	return game.GameId, nil
}

func (m *memory) UpdateGame(g *model.Game) error {
	_, _ = m.InsertGame(g)
	return nil
}
