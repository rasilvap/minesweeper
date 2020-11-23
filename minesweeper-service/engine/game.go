package engine

import (
	"fmt"
	"log"
	"math/rand"
)

type StateTile int

const (
	StateTileCovered  StateTile = 1
	StateTileClear    StateTile = 2
	StateTileFlagged  StateTile = 3
	StateTileNumberd  StateTile = 4
	StateTileExploted StateTile = 5
)

type StateGame int

const (
	StateGameNew     StateGame = 1
	StateGameRunning StateGame = 2
	StateGameWon     StateGame = 3
	StateGameLost    StateGame = 4
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
	MineAmount int

	FlagAmount int
}

func (g Game) PlayMovement(r, c int) (StateGame, Game) {
	tile := &g.Board[r][c]

	//TODO unify with MarkFlag
	if tile.State != StateTileCovered {
		if tile.State == StateTileFlagged {
			tile.State = StateTileCovered
		}

		log.Println("Tile has already been played")
		return StateGameRunning, g.buildGameWithVisibleTiles()
	}

	//game over, so show all tiles
	if tile.IsMine {
		log.Println("Game Over")
		tile.State = StateTileExploted
		return StateGameLost, g.copyGame()
	}

	//simple case, only mark
	if tile.SurroundingMineCount == 0 {
		log.Println("Tile was Cleaned")
		tile.State = StateTileClear
	} else {
		log.Println("Tile was Numbered")
		tile.State = StateTileNumberd
	}

	g.RevealEmptyAdjacentTiles(r, c)

	// game won, clear all tiles
	if g.isFlawlessVictory() {
		log.Println("Flawless Victory")
		return StateGameWon, g.copyGame()
	}

	log.Println("The Game is Running")
	//return showable tiles
	return StateGameRunning, g.buildGameWithVisibleTiles()
}

func (g Game) isFlawlessVictory() bool {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			if board := g.Board[i][j]; !board.IsMine &&
				(board.State == StateTileCovered || board.State == StateTileFlagged) {
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

func (g Game) SetUpMines(minedPointTiles [][2]int) {
	mines := make([]Mine, len(minedPointTiles))
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

//TODO return points adjacent
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

func BuildNewGame(rows, columns, mineAmount int) Game {
	board := make([][]Tile, rows)

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

	return Game{board, rows, columns, mineAmount, 0}
}

func (g Game) copyGame() Game {
	board := make([][]Tile, g.Rows)

	for r := range board {
		board[r] = make([]Tile, g.Columns)
	}

	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			privateBoard := g.Board[i][j]
			board[i][j] = Tile{privateBoard.State, i, j, privateBoard.SurroundingMineCount, privateBoard.IsMine, -1}
		}
	}

	return Game{board, g.Rows, g.Columns, g.MineAmount, g.FlagAmount}
}

func (g Game) buildGameWithVisibleTiles() Game {
	var board [][]Tile
	for i := 0; i < g.Rows; i++ {
		var column []Tile
		for j := 0; j < g.Columns; j++ {
			if board := g.Board[i][j]; !board.IsMine &&
				(board.State == StateTileClear || board.State == StateTileNumberd) {
				column = append(column, g.Board[i][j])
			}
		}
		if column != nil && len(column) > 0 {
			board = append(board, column)
		}
	}

	if board == nil {
		board = [][]Tile{}
	}
	return Game{board, g.Rows, g.Columns, g.MineAmount, g.FlagAmount}
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
