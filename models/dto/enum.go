package dto

import "errors"

type MoveType string

const (
	TypeMoveClean          MoveType = "CLEAN"
	TypeMoveFlag           MoveType = "FLAG"
	TypeMoveQuestion       MoveType = "QUESTION"
	TypeMoveRevertQuestion MoveType = "REVERT_QUESTION"
	TypeMoveRevertFlag     MoveType = "REVERT_FLAG"
)

func (mt MoveType) IsValid() error {
	switch mt {
	case TypeMoveClean, TypeMoveFlag, TypeMoveQuestion, TypeMoveRevertQuestion, TypeMoveRevertFlag:
		return nil
	}
	return errors.New("invalid move type")
}

type GameState string

const (
	GameStateRunning GameState = "RUNNING"
	GameStateLost    GameState = "LOST"
	GameStateWon     GameState = "WON"
	GameStateNew     GameState = "NEW"
)

type TileState string

const (
	TileStateCovered   TileState = "COVERED"
	TileStateClear     TileState = "CLEAR"
	TileStateFlagged   TileState = "FLAGGED"
	TileStateNumbered  TileState = "NUMBERED"
	TileStateExploited TileState = "EXPLOITED"
)
