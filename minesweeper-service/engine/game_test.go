package engine

import (
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

func TestGetAdjecentTilesShouldBe8(t *testing.T) {
	game := BuildNewGame(3, 3)
	result := game.getAdjecentTiles(1, 1)

	if len(result) != 8 {
		t.Error("Error", len(result))
	}
}
