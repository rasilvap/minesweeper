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
	Game engine.Game
}
