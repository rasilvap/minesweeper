package datasource

import (
	"minesweeper-API/model"
)

type Spec interface {
	FindGame(id int) (*model.Game, error)
	InsertGame(g *model.Game) (int, error)
	UpdateGame(g *model.Game) error
}
