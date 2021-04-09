package datasource

import (
	"minesweeper-API/models"
)


//TODO deberia retornar una copia? si retorna un puntero puede dar problemas de concurrencia?
func (ds datasourceMemory) Find(id int) (*models.Game, error) {
	ds.RLock()
	defer ds.RUnlock()
	if game, ok := ds.cache[id]; ok {
		return game.(*models.Game), nil
	}
	return &models.Game{}, nil
}

func (ds datasourceMemory) Insert(game *models.Game) (int, error) {
	ds.Lock()
	game.GameId = len(ds.cache) + 1
	ds.cache[game.GameId] = game
	ds.Unlock()
	return game.GameId, nil
}

func (ds datasourceMemory) Update(g *models.Game) error {
	_, _ = ds.Insert(g)
	return nil
}
