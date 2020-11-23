package service

import (
	"minesweeper-API/minesweeper-service/engine"
	"testing"
)

func TestMarkPlayMovementWhenRunning(t *testing.T) {
	//setup
	game := engine.BuildNewGame(3, 3, 1)
	minedPointTile := [][2]int{{1, 1}}
	game.SetUpMines(minedPointTile)

	//execute
	stateGame, gameCopy := game.PlayMovement(1, 1)
	res := buildPlayResponse(stateGame, gameCopy)

	//assert
	if res.StateGame != "LOST" || res.Game.Columns != 3 || res.Game.Rows != 3 ||
		len(res.Game.Board) != res.Game.Rows || len(res.Game.Board[0]) != res.Game.Columns {
		t.Error("Error", res, len(res.Game.Board), len(res.Game.Board[0]))
	}
}

func TestMarkPlayMovementWhenRunningXX(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := engine.BuildNewGame(3, 8, len(minedPointTile))

	game.SetUpMines(minedPointTile)

	//execute
	stateGame, gameCopy := game.PlayMovement(0, 5)
	game.ShowBoard()
	res := buildPlayResponse(stateGame, gameCopy)

	//assert
	if res.StateGame != "RUNNING" || res.Game.Rows != 3 || res.Game.Columns != 8 ||
		len(res.Game.Board) != 3 || len(res.Game.Board[0]) != 5 {
		t.Error("Error", res, len(res.Game.Board), len(res.Game.Board[0]))
	}
}
