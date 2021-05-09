package engine

import (
	"minesweeper-API/models"
	"minesweeper-API/models/dto"
)

type Game interface {
	Create(rows, columns, mineAmount int) (int, error)
	Play(id int, playRequest dto.PlayRequest) (*dto.PlayResponse, error)
	Get(id int) (*dto.GameResponse, error)
}

type MinesWeeper interface {
	BuildGame(rows, columns, mineAmount int) (*models.Game, error)
	Play(playRequest dto.PlayRequest, game *models.Game) (*models.Game, *dto.PlayResponse, error)
}
