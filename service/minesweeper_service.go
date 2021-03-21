package service

import (
	"minesweeper-API/minesweeper-service/model"

	"github.com/obarra-dev/minesweeper"
)

type MinesWeeperService interface {
	BuildGame(rows, columns, mineAmount int) *minesweeper.Game
	Play(playRequest model.PlayRequest, game *minesweeper.Game) minesweeper.Game
}

type minesWeeperService struct{}

func NewMinesWeeperService() MinesWeeperService {
	return &minesWeeperService{}
}

func (*minesWeeperService) BuildGame(rows, columns, mineAmount int) *minesweeper.Game {
	minedPointTile := minesweeper.GenerateMinedPoints(rows, columns, mineAmount)
	return minesweeper.NewMinesweeper(rows, columns, minedPointTile)
}

func (*minesWeeperService) Play(playRequest model.PlayRequest, game *minesweeper.Game) minesweeper.Game {
	return game.Play(playRequest.Row, playRequest.Column, mapTypeMove(playRequest.Move))
}
