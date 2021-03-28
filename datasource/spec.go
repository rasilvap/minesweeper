package datasource

import (
	"minesweeper-API/minesweeper-service/model"
)

type Spec interface {
	FindGame(id int) (*model.Game, error)
	SaveGame(g *model.Game) (int, error)
	UpdateGame(g *model.Game) error
}
