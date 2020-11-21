package engine

import "fmt"

type StateTile int

const (
	StateTileCovered  StateTile = 1
	StateTileClear              = 2
	StateTileFlagged            = 3
	StateTileNumberd            = 4
	StateTileExploted           = 4
)

type Tile struct {
	state                StateTile
	r                    int
	c                    int
	surroundingMineCount int
	isMine               bool
	valueTest            int
}

type Mine struct {
	r      int
	c      int
	active bool
}

type Game struct {
	Board      [][]Tile
	Rows       int
	Columns    int
	FlagAmount int
}

func BuildNewGame(rows int, columns int) Game {
	//create the rows
	board := make([][]Tile, rows)

	//create the columns
	for r := range board {
		board[r] = make([]Tile, columns)
	}

	cont := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			cont++
			board[i][j] = Tile{StateTileCovered, i, j, 0, false, cont}
		}
	}

	return Game{board, rows, columns, 0}
}

func (g Game) ShowBoard() {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			fmt.Print(g.Board[i][j], " ")
		}
		fmt.Println()
	}
}
