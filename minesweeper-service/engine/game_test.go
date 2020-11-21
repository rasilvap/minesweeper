package engine

import (
	"fmt"
	"testing"
)

func TestBuildGame(t *testing.T) {
	game := BuildNewGame(3, 3)
	game.ShowBoard()

	if len(game.Board) != 3 || game.Rows != 3 &&
		len(game.Board[0]) != 3 || game.Columns != 3 {
		t.Error("Error", len(game.Board), len(game.Board[0]))
	}
}

func TestGetAdjacentTilesShouldBe8(t *testing.T) {
	game := BuildNewGame(3, 3)
	result := game.getAdjacentTiles(1, 1)

	for d := 0; d < len(result); d++ {
		fmt.Println(result[d].valueTest)
	}

	if len(result) != 8 {
		t.Error("Error", len(result))
	}
}

func TestGetAdjacentTilesShouldBe3(t *testing.T) {
	game := BuildNewGame(3, 3)
	result := game.getAdjacentTiles(2, 2)
	if len(result) != 3 {
		t.Error("Error", len(result))
	}
}

func TestGetAdjacentTilesShouldBe5(t *testing.T) {
	game := BuildNewGame(3, 3)
	result := game.getAdjacentTiles(0, 1)
	if len(result) != 5 {
		t.Error("Error", len(result))
	}
}

func TestGetAdjacentTilesShouldBe0(t *testing.T) {
	game := BuildNewGame(1, 1)
	result := game.getAdjacentTiles(0, 0)
	if len(result) != 0 {
		t.Error("Error", len(result))
	}
}
