package dto

import (
	"github.com/obarra-dev/minesweeper"
)

// CreateGameRequest : ...
type CreateGameRequest struct {
	Rows       int `json:"rows"`
	Columns    int `json:"columns"`
	MineAmount int `json:"mineAmount"`
}

// CreateGameResponse : ...
type CreateGameResponse struct {
	GameID int `json:"id"`
}

// GetGameResponse : ...
type GetGameResponse struct {
	Rows       int `json:"rows"`
	Columns    int `json:"columns"`
	MineAmount int `json:"mineAmount"`
}

// PlayRequest : ...
type PlayRequest struct {
	Row    int      `json:"row"`
	Column int      `json:"column"`
	Move   MoveType `json:"move"`
}

// PlayResponse : ...
type PlayResponse struct {
	Game GameDTO `json:"game"`
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
}

type GameDTO struct {
	StateGame string    `json:"stateGame"`
	Rows      int       `json:"rows"`
	Columns   int       `json:"columns"`
	Board     []TileDTO `json:"board"`
}
