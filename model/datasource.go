package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Game struct {
	GameId     int `db:"game_id"`
	State      string
	Columns    int
	Rows       int
	MineAmount int `db:"mine_amount"`
	FlagAmount int `db:"flag_amount"`
	Board      Board
}

type Board struct {
	B [][]Tile `json:"b,omitempty"`
}

type Tile struct {
	State                string `json:"state,omitempty"`
	Row                  int    `json:"row,omitempty"`
	Column               int    `json:"column,omitempty"`
	SurroundingMineCount int    `json:"surrounding_mine_count,omitempty"`
	IsMine               bool   `json:"is_mine,omitempty"`
}

func (a Board) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Board) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
