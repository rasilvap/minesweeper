package model

import (
	"minesweeper-API/minesweeper-service/engine"
)

// GameRequest : ...
type GameRequest struct {
	Rows       int `json:"rows"`
	Columns    int `json:"columns"`
	MineAmount int `json:"mineAmount"`
}

// GameResponse : ...
type GameResponse struct {
	Rows       int `json:"rows"`
	Columns    int `json:"columns"`
	MineAmount int `json:"mineAmount"`
}

// MarkRequest : ...
type MarkRequest struct {
	Row    int    `json:"row"`
	Column int    `json:"column"`
	Mark   string `json:"mark"`
}

// PlayRequest : ...
type PlayRequest struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

// PlayResponse : ...
type PlayResponse struct {
	StateGame string  `json:"stateGame"`
	Game      GameDTO `json:"game"`
}

// GameCompleteResponse : ...
type GameCompleteResponse struct {
	Game engine.Game
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
