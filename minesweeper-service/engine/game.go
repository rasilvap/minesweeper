package engine

import (
	"fmt"
	"math/rand"
)

type StateTile int

const (
	StateTileCovered  StateTile = 1
	StateTileClear              = 2
	StateTileFlagged            = 3
	StateTileNumberd            = 4
	StateTileExploted           = 4
)

type StateGame int

const (
	StateGameNew     StateGame = 1
	StateGameRunning           = 2
	StateGameWon               = 3
	StateGameLost              = 4
)

type Tile struct {
	State                StateTile
	Row                  int
	Column               int
	SurroundingMineCount int
	IsMine               bool
	ValueTest            int
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

func (g Game) PlayMovement(r, c int) (StateGame, [][]Tile) {
	tile := &g.Board[r][c]

	//played tile, maybe revert
	if tile.State != StateTileCovered {
		if tile.State == StateTileFlagged {
			tile.State = StateTileCovered
		}

		return StateGameRunning, [][]Tile{{Tile{tile.State,
			tile.Row,
			tile.Column,
			tile.SurroundingMineCount,
			tile.IsMine,
			tile.ValueTest}}}
	}

	//game over, clear all tiles
	if tile.IsMine {
		tile.State = StateTileExploted
		return StateGameLost, nil
	}

	//simple case, only mark
	if tile.SurroundingMineCount == 0 {
		tile.State = StateTileClear
	} else {
		tile.State = StateTileNumberd
	}

	g.RevealEmptyAdjacentTiles(r, c)

	// game won, clear all tiles
	if g.isFlawlessVictory() {
		return StateGameWon, nil
	}

	//return showable tiles
	return StateGameRunning, make([][]Tile, 2)
}

func (g Game) isFlawlessVictory() bool {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			if board := g.Board[i][j]; !board.IsMine &&
				(board.State == StateTileClear || board.State == StateTileFlagged) {
				return false
			}
		}
	}

	return true
}

func (g *Game) MarkFlag(r, c int) int {
	tile := &g.Board[r][c]

	if tile.State == StateTileCovered {
		tile.State = StateTileFlagged
		g.FlagAmount++
	} else if tile.State == StateTileFlagged {
		tile.State = StateTileCovered
		g.FlagAmount--
	}

	return g.FlagAmount
}

func (g Game) SetUpMines(amountMines int, minedPointTiles [][2]int) {
	mines := make([]Mine, amountMines)
	for i := range mines {
		r := minedPointTiles[i][0]
		c := minedPointTiles[i][1]

		//TODO use to stats
		mines[i] = Mine{r, c, true}
		g.Board[r][c].IsMine = true

		adjacentTiles := g.getAdjacentTiles(r, c)
		for i := 0; i < len(adjacentTiles); i++ {
			g.Board[adjacentTiles[i].Row][adjacentTiles[i].Column].SurroundingMineCount++
		}
	}
}

func (g Game) RevealEmptyAdjacentTiles(r int, c int) {
	if g.Board[r][c].SurroundingMineCount == 0 {
		adjecentTiles := g.getAdjacentTiles(r, c)
		for i := 0; i < len(adjecentTiles); i++ {
			if adjecentTiles[i].IsMine != true &&
				(adjecentTiles[i].State == StateTileCovered || adjecentTiles[i].State == StateTileFlagged) {
				if adjecentTiles[i].SurroundingMineCount == 0 {
					g.Board[adjecentTiles[i].Row][adjecentTiles[i].Column].State = StateTileClear
					g.RevealEmptyAdjacentTiles(adjecentTiles[i].Row, adjecentTiles[i].Column)
				} else {
					g.Board[adjecentTiles[i].Row][adjecentTiles[i].Column].State = StateTileNumberd
				}
			}
		}
	}
}

func (g Game) getAdjacentTiles(f int, c int) []Tile {
	minF := -1
	if f == 0 {
		minF = 0
	}

	minC := -1
	if c == 0 {
		minC = 0
	}

	maxF := 1
	if f == (g.Rows - 1) {
		maxF = 0
	}

	maxC := 1
	if c == (g.Columns - 1) {
		maxC = 0
	}

	var adjecentTiles []Tile
	for cc := minC; cc <= maxC; cc++ {
		for ff := minF; ff <= maxF; ff++ {
			if cc == 0 && ff == 0 {
				continue
			}

			var resultF = ff + f
			var resultC = cc + c

			adjecentTiles = append(adjecentTiles, g.Board[resultF][resultC])
		}
	}

	return adjecentTiles
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

func (g Game) GetStates() [][]StateTile {
	states := make([][]StateTile, g.Rows)

	for i := 0; i < g.Rows; i++ {
		states[i] = make([]StateTile, g.Columns)
		for j := 0; j < g.Columns; j++ {
			states[i][j] = g.Board[i][j].State
		}
	}
	return states
}

//TODO improve random in order to dont repeat
func generateMinedPointTiles(amountPoints int, maxIncluded int) [][2]int {
	max := maxIncluded + 1
	tileMinePoints := make([][2]int, amountPoints)
	for i := 0; i < amountPoints; i++ {
		rand.Seed(int64(i))
		tileMinePoints[i][0] = rand.Intn(max)
		tileMinePoints[i][1] = rand.Intn(max)
	}

	return tileMinePoints
}
