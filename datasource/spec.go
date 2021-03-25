package datasource

import (
	"minesweeper-API/minesweeper-service/model"
)

type Spec interface {
	GetGame(id int) (*model.Game, error)
	SaveGame(g *model.Game) (int, error)
}
