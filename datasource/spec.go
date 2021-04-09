package datasource

import (
	"minesweeper-API/models"
)

type Game interface {
	Find(id int) (*models.Game, error)
	Insert(g *models.Game) (int, error)
	Update(g *models.Game) error
}
