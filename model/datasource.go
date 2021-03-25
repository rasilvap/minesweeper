package model

type Game struct {
	GameId     int `db:"game_id"`
	State      string
	Columns    int
	Rows       int
	MineAmount int `db:"mine_amount"`
	FlagAmount int `db:"flag_amount"`
	Board      string
}
