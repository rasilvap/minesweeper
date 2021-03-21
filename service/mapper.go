package service

import (
	"minesweeper-API/minesweeper-service/model"

	"github.com/obarra-dev/minesweeper"
)

func mapTypeMove(typeMove model.TypeMove) minesweeper.TypeMove {
	var move minesweeper.TypeMove
	switch typeMove {
	case model.TypeMoveFlag:
		move = minesweeper.TypeMoveFlag
	case model.TypeMoveQuestion:
		move = minesweeper.TypeMoveQuestion
	case model.TypeMoveClean:
		move = minesweeper.TypeMoveClean
	}

	return move
}

func mapStateGame(stateGame minesweeper.StateGame) string {
	var gameStateDTO string
	switch stateGame {
	case minesweeper.StateGameRunning:
		gameStateDTO = "RUNNING"
	case minesweeper.StateGameLost:
		gameStateDTO = "LOST"
	case minesweeper.StateGameNew:
		gameStateDTO = "NEW"
	case minesweeper.StateGameWon:
		gameStateDTO = "WON"
	default:
		gameStateDTO = ""
	}
	return gameStateDTO
}

func mapTileState(tileState minesweeper.StateTile) string {
	var tileStateDTO string
	switch tileState {
	case minesweeper.StateTileCovered:
		tileStateDTO = "COVERED"
	case minesweeper.StateTileClear:
		tileStateDTO = "CLEAR"
	case minesweeper.StateTileFlagged:
		tileStateDTO = "FLAGGED"
	case minesweeper.StateTileNumbered:
		tileStateDTO = "NUMBERED"
	case minesweeper.StateTileExploited:
		tileStateDTO = "EXPLOTED"
	default:
		tileStateDTO = ""
	}
	return tileStateDTO
}
