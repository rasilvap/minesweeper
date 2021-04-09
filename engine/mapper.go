package engine

import (
	"minesweeper-API/models"

	"github.com/obarra-dev/minesweeper"
)

func mapTypeMove(typeMove models.TypeMove) minesweeper.TypeMove {
	var move minesweeper.TypeMove
	switch typeMove {
	case models.TypeMoveFlag:
		move = minesweeper.TypeMoveFlag
	case models.TypeMoveQuestion:
		move = minesweeper.TypeMoveQuestion
	case models.TypeMoveClean:
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
		tileStateDTO = "EXPLOITED"
	default:
		tileStateDTO = ""
	}
	return tileStateDTO
}
