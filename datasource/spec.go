package datasource

import (
	"minesweeper-API/models"
)

type Spec interface {
	FindGame(id int) (*models.Game, error)
	InsertGame(g *models.Game) (int, error)
	UpdateGame(g *models.Game) error
}
