package engine

import (
	"minesweeper-API/models/dto"

	"github.com/obarra-dev/minesweeper"
)

func mapTypeMove(typeMove dto.TypeMove) (move minesweeper.TypeMove) {
	switch typeMove {
	case dto.TypeMoveFlag:
		move = minesweeper.TypeMoveFlag
	case dto.TypeMoveQuestion:
		move = minesweeper.TypeMoveQuestion
	case dto.TypeMoveClean:
		move = minesweeper.TypeMoveClean
	}
	return
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
