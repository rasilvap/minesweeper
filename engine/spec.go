package engine

import "minesweeper-API/models"

type Game interface {
	Create(rows, columns, mineAmount int) (int, error)
	Play(id int, playRequest models.PlayRequest) (*models.PlayResponse, error)
	Get(id int) (*models.GameResponse, error)
}

type MinesWeeper interface {
	BuildGame(rows, columns, mineAmount int) (*models.Game, error)
	Play(playRequest models.PlayRequest, game *models.Game) (*models.Game, *models.PlayResponse, error)
}
