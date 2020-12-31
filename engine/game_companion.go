package engine

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func BuildNewGame(rows, columns, mineAmount int) *Game {
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

	return &Game{board, rows, columns, mineAmount, 0}
}

func GenerateMinedPoints(amountPoints, maxRowIncluded, maxColumnIncluded int) [][2]int {
	tileMinePoints := make([][2]int, amountPoints)

	setPoints := make(map[string]bool)

	for len(setPoints) < amountPoints {
		concatenated := fmt.Sprint(rand.Intn(maxRowIncluded), "-", rand.Intn(maxColumnIncluded))
		setPoints[concatenated] = true
	}

	i := 0
	for key := range setPoints {
		point := strings.Split(key, "-")
		tileMinePoints[i][0], _ = strconv.Atoi(point[0])
		tileMinePoints[i][1], _ = strconv.Atoi(point[1])
		i++
	}

	return tileMinePoints
}
