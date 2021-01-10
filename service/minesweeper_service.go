package service

import (
	"minesweeper-API/minesweeper-service/engine"
	"minesweeper-API/minesweeper-service/model"
)

type MinesWeeperService interface {
	BuildGame(rows, columns, mineAmount int) *engine.Game
	Play(playRequest model.PlayRequest, game *engine.Game) engine.Game
}

type minesWeeperService struct{}

func NewMinesWeeperService() MinesWeeperService {
	return &minesWeeperService{}
}

func (*minesWeeperService) BuildGame(rows, columns, mineAmount int) *engine.Game {
	minedPointTile := engine.GenerateMinedPoints(rows, columns, mineAmount)
	return engine.BuildNewGame(rows, columns, minedPointTile)
}

func (*minesWeeperService) Play(playRequest model.PlayRequest, game *engine.Game) engine.Game {
	return game.Play(playRequest.Row, playRequest.Column, mapTypeMove(playRequest.Move))
}
