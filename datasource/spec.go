package datasource

import (
	"github.com/obarra-dev/minesweeper"
)

type Spec interface {
	GetGame(id int) (*minesweeper.Game, error)
	SaveGame(g *minesweeper.Game) (int, error)
}
