package service

import (
	"minesweeper-API/minesweeper-service/engine"
	"testing"
)

func TestMarkPlayWhenLost3x3(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := engine.BuildNewGame(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(1, 1, engine.TypeMoveClean)
	res := buildPlayResponse(gameCopy)

	//assert
	if res.StateGame != "LOST" || res.Game.Columns != 3 || res.Game.Rows != 3 ||
		len(res.Game.Board) != res.Game.Rows || len(res.Game.Board[0]) != res.Game.Columns {
		t.Error("Error", res, len(res.Game.Board), len(res.Game.Board[0]))
	}
}

func TestMarkPlayWhenWon3x3(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := engine.BuildNewGame(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(0, 0, engine.TypeMoveClean)
	res := buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(0, 1, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res)
	}

	//execute
	gameCopy = game.Play(0, 2, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res)
	}

	//execute
	gameCopy = game.Play(1, 0, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(1, 2, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(2, 0, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(2, 1, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(2, 2, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "WON" {
		t.Error("Error", res, len(res.Game.Board))
	}
}

func TestMarkPlayWhenRunning3X8(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := engine.BuildNewGame(3, 8, minedPointTile)

	//execute
	gameCopy := game.Play(0, 5, engine.TypeMoveClean)
	game.ShowBoard()
	res := buildPlayResponse(gameCopy)

	//assert
	if res.StateGame != "RUNNING" || res.Game.Rows != 3 || res.Game.Columns != 8 ||
		len(res.Game.Board) != 3 || len(res.Game.Board[0]) != 6 {
		t.Error("Error", res, len(res.Game.Board), len(res.Game.Board[0]))
	}
}
