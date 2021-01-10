package service

import (
	"minesweeper-API/minesweeper-service/engine"
	"minesweeper-API/minesweeper-service/model"
)

func mapTypeMove(typeMove model.TypeMove) engine.TypeMove {
	var move engine.TypeMove
	switch typeMove {
	case model.TypeMoveFlag:
		move = engine.TypeMoveFlag
	case model.TypeMoveQuestion:
		move = engine.TypeMoveQuestion
	case model.TypeMoveClean:
		move = engine.TypeMoveClean
	}

	return move
}

func mapStateGame(stateGame engine.StateGame) string {
	var gameStateDTO string
	switch stateGame {
	case engine.StateGameRunning:
		gameStateDTO = "RUNNING"
	case engine.StateGameLost:
		gameStateDTO = "LOST"
	case engine.StateGameNew:
		gameStateDTO = "NEW"
	case engine.StateGameWon:
		gameStateDTO = "WON"
	default:
		gameStateDTO = ""
	}
	return gameStateDTO
}

func mapTileState(tileState engine.StateTile) string {
	var tileStateDTO string
	switch tileState {
	case engine.StateTileCovered:
		tileStateDTO = "COVERED"
	case engine.StateTileClear:
		tileStateDTO = "CLEAR"
	case engine.StateTileFlagged:
		tileStateDTO = "FLAGGED"
	case engine.StateTileNumberd:
		tileStateDTO = "NUMBERED"
	case engine.StateTileExploted:
		tileStateDTO = "EXPLOTED"
	default:
		tileStateDTO = ""
	}
	return tileStateDTO
}
