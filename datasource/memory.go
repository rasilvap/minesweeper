package datasource

import (
	"minesweeper-API/models"
)

type gameMemory struct {
	*datasourceMemory
}

func NewGameMemory(ds *datasourceMemory) Game {
	return &gameMemory{ds}
}

//TODO deberia retornar una copia? si retorna un puntero puede dar problemas de concurrencia?
func (gm *gameMemory) Find(id int) (*models.Game, error) {
	gm.RLock()
	defer gm.RUnlock()
	if game, ok := gm.cache[id]; ok {
		return game.(*models.Game), nil
	}
	return &models.Game{}, nil
}

func (gm *gameMemory) Insert(game *models.Game) (int, error) {
	gm.Lock()
	game.GameID = len(gm.cache) + 1
	gm.cache[game.GameID] = game
	gm.Unlock()
	return game.GameID, nil
}

func (gm *gameMemory) Update(g *models.Game) error {
	_, _ = gm.Insert(g)
	return nil
}
