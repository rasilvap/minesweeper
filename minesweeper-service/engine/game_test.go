package engine

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetUpMines(t *testing.T) {
	game := BuildNewGame(3, 3)
	minedPointTile := [][2]int{{0, 1}, {1, 1}, {1, 0}}
	game.SetUpMines(3, minedPointTile)
	game.ShowBoard()

	expected := [][]Tile{
		{Tile{StateTileCovered, 0, 0, 3, false, 1}, Tile{StateTileCovered, 0, 1, 2, true, 2}, Tile{StateTileCovered, 0, 2, 2, false, 3}},
		{Tile{StateTileCovered, 1, 0, 2, true, 4}, Tile{StateTileCovered, 1, 1, 2, true, 5}, Tile{StateTileCovered, 1, 2, 2, false, 6}},
		{Tile{StateTileCovered, 2, 0, 2, false, 7}, Tile{StateTileCovered, 2, 1, 2, false, 8}, Tile{StateTileCovered, 2, 2, 1, false, 9}}}

	if !reflect.DeepEqual(expected, game.Board) {
		t.Error("Error", game.Board)
	}
}

func TestRevealEmptyAdjecentTiles3x8(t *testing.T) {
	game := BuildNewGame(3, 8)
	minedPointTile := [][2]int{{1, 1}}
	game.SetUpMines(1, minedPointTile)

	game.RevealEmptyAdjacentTiles(0, 5)
	game.ShowBoard()

	expected := [][]StateTile{
		{StateTileCovered, StateTileCovered, StateTileNumberd, StateTileClear, StateTileClear, StateTileClear, StateTileClear, StateTileClear},
		{StateTileCovered, StateTileCovered, StateTileNumberd, StateTileClear, StateTileClear, StateTileClear, StateTileClear, StateTileClear},
		{StateTileCovered, StateTileCovered, StateTileNumberd, StateTileClear, StateTileClear, StateTileClear, StateTileClear, StateTileClear},
	}
	if !reflect.DeepEqual(expected, game.GetStates()) {
		t.Error("Error", game.GetStates())
	}
}

func TestMarkFlag(t *testing.T) {
	game := BuildNewGame(3, 3)

	flagAmount := game.MarkFlag(1, 1)
	if flagAmount != 1 || game.FlagAmount != 1 {
		t.Error("Error", flagAmount, game.FlagAmount)
	}
}

func TestMarkFlagWhenRevert(t *testing.T) {
	game := BuildNewGame(3, 3)

	flagAmount := game.MarkFlag(1, 1)
	flagAmount = game.MarkFlag(1, 1)

	if flagAmount != 0 || game.FlagAmount != 0 {
		t.Error("Error", flagAmount, game.FlagAmount)
	}
}

func TestMarkPlayMovement(t *testing.T) {
	game := BuildNewGame(2, 2)
	game.MarkFlag(0, 0)

	stateGame, pointTiles := game.PlayMovement(0, 0)

	expected := [][]StateTile{
		{StateTileCovered, StateTileCovered},
		{StateTileCovered, StateTileCovered},
	}

	if stateGame != StateGameRunning ||
		pointTiles[0][0].Row != 0 || pointTiles[0][0].Column != 0 ||
		!reflect.DeepEqual(expected, game.GetStates()) {
		t.Error("Error", stateGame, pointTiles[0][0].Row, pointTiles[0][0].Column, game.GetStates())
	}
}

func TestGetAdjacentTilesShouldBe8(t *testing.T) {
	game := BuildNewGame(3, 3)
	result := game.getAdjacentTiles(1, 1)

	for d := 0; d < len(result); d++ {
		fmt.Println(result[d].ValueTest)
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

func TestBuildGame(t *testing.T) {
	game := BuildNewGame(3, 3)
	game.ShowBoard()

	if len(game.Board) != 3 || game.Rows != 3 &&
		len(game.Board[0]) != 3 || game.Columns != 3 {
		t.Error("Error", len(game.Board), len(game.Board[0]))
	}
}
