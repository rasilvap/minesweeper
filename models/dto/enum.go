package dto

import "errors"

type MoveType string

const (
	TypeMoveClean          MoveType = "CLEAN"
	TypeMoveFlag                    = "FLAG"
	TypeMoveQuestion                = "QUESTION"
	TypeMoveRevertQuestion          = "REVERT_QUESTION"
	TypeMoveRevertFlag              = "REVERT_FLAG"
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
	GameStateLost              = "LOST"
	GameStateWon               = "WON"
	GameStateNew               = "NEW"
)

type TileState string

const (
	TileStateCovered   TileState = "COVERED"
	TileStateClear               = "CLEAR"
	TileStateFlagged             = "FLAGGED"
	TileStateNumbered            = "NUMBERED"
	TileStateExploited           = "EXPLOITED"
)
