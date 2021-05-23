package models

type Game struct {
	GameID     int `db:"game_id"`
	State      int
	Columns    int
	Rows       int
	MineAmount int `db:"mine_amount"`
	FlagAmount int `db:"flag_amount"`
	Board      string
}
