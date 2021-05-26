package engine

import (
	"minesweeper-API/models/dto"

	"github.com/obarra-dev/minesweeper"
)

func mapTypeMove(typeMove dto.MoveType) (move minesweeper.TypeMove) {
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

func mapToStateGameDTO(stateGame minesweeper.StateGame) (state dto.GameState) {
	switch stateGame {
	case minesweeper.StateGameRunning:
		state = dto.GameStateRunning
	case minesweeper.StateGameLost:
		state = dto.GameStateLost
	case minesweeper.StateGameNew:
		state = dto.GameStateNew
	case minesweeper.StateGameWon:
		state = dto.GameStateWon
	}
	return
}

func mapToTileStateDTO(tileState minesweeper.StateTile) (state dto.TileState) {
	switch tileState {
	case minesweeper.StateTileCovered:
		state = dto.TileStateCovered
	case minesweeper.StateTileClear:
		state = dto.TileStateClear
	case minesweeper.StateTileFlagged:
		state = dto.TileStateFlagged
	case minesweeper.StateTileNumbered:
		state = dto.TileStateNumbered
	case minesweeper.StateTileExploited:
		state = dto.TileStateExploited
	}
	return state
}
