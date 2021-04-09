package models

import "github.com/obarra-dev/minesweeper"

// GameRequest : ...
type GameRequest struct {
	Rows       int `json:"rows"`
	Columns    int `json:"columns"`
	MineAmount int `json:"mineAmount"`
}

// GameSimpleResponse : ...
type GameSimpleResponse struct {
	GameId int `json:"gameId"`
}

// GameResponse : ...
type GameResponse struct {
	Rows       int `json:"rows"`
	Columns    int `json:"columns"`
	MineAmount int `json:"mineAmount"`
}

type TypeMove string

const (
	TypeMoveClean          TypeMove = "CLEAN"
	TypeMoveFlag           TypeMove = "FLAG"
	TypeMoveQuestion       TypeMove = "QUESTION"
	TypeMoveRevertQuestion TypeMove = "REVERT_QUESTION"
	TypeMoveRevertFlag     TypeMove = "REVERT_FLAG"
)

// PlayRequest : ...
type PlayRequest struct {
	Row    int      `json:"row"`
	Column int      `json:"column"`
	Move   TypeMove `json:"move"`
}

// PlayResponse : ...
type PlayResponse struct {
	StateGame string  `json:"stateGame"`
	Game      GameDTO `json:"game"`
}

// GameCompleteResponse : ...
type GameCompleteResponse struct {
	Game minesweeper.Game
}

type TileDTO struct {
	State                string `json:"state"`
	Row                  int    `json:"row"`
	Column               int    `json:"column"`
	SurroundingMineCount int    `json:"surroundingMineCount"`
	Mine                 bool   `json:"mine"`
	ValueTest            int    `json:"valueTest"`
}

type GameDTO struct {
	Board      [][]TileDTO `json:"board"`
	Rows       int         `json:"rows"`
	Columns    int         `json:"columns"`
	FlagAmount int         `json:"flagAmount"`
}
