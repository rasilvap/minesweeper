package datasource

import (
	"minesweeper-API/model"
	"sync"
)

type datasourceMemory struct {
	sync.RWMutex
	cache map[int]interface{}
}

func NewDatasourceMemory() Spec {
	return &datasourceMemory{
		cache: make(map[int]interface{}),
	}
}

//TODO deberia retornar una copia? si retorna un puntero puede dar problemas de concurrencia?
func (ds *datasourceMemory) FindGame(id int) (*model.Game, error) {
	ds.RLock()
	defer ds.RUnlock()
	if game, ok := ds.cache[id]; ok {
		return game.(*model.Game), nil
	}
	return &model.Game{}, nil
}

func (ds *datasourceMemory) InsertGame(game *model.Game) (int, error) {
	ds.Lock()
	game.GameId = len(ds.cache) + 1
	ds.cache[game.GameId] = game
	ds.Unlock()
	return game.GameId, nil
}

func (ds *datasourceMemory) UpdateGame(g *model.Game) error {
	_, _ = ds.InsertGame(g)
	return nil
}
