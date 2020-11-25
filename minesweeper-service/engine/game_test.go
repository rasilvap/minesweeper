package engine

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateMinedPoints(t *testing.T) {
	minedPoints := generateMinedPoints(24, 3, 8)
	if len(minedPoints) != 24 {
		t.Error("Error", minedPoints)
	}
}
func TestGenerateMinedPointsTwoMines(t *testing.T) {
	minedPoints := generateMinedPoints(2, 3, 3)
	if len(minedPoints) != 2 {
		t.Error("Error", minedPoints)
	}
}

func TestSetUpMines(t *testing.T) {
	minedPointTile := [][2]int{{0, 1}, {1, 1}, {1, 0}}
	game := BuildNewGame(3, 3, len(minedPointTile))
	game.SetUpMines(minedPointTile)
	game.ShowBoard()

	expected := [][]Tile{
		{Tile{StateTileCovered, 0, 0, 3, false, 1}, Tile{StateTileCovered, 0, 1, 2, true, 2}, Tile{StateTileCovered, 0, 2, 2, false, 3}},
		{Tile{StateTileCovered, 1, 0, 2, true, 4}, Tile{StateTileCovered, 1, 1, 2, true, 5}, Tile{StateTileCovered, 1, 2, 2, false, 6}},
		{Tile{StateTileCovered, 2, 0, 2, false, 7}, Tile{StateTileCovered, 2, 1, 2, false, 8}, Tile{StateTileCovered, 2, 2, 1, false, 9}}}

	if !reflect.DeepEqual(expected, game.Board) {
		t.Error("Error", game.Board)
	}
}

func TestMarkPlayMovementWhenRunning(t *testing.T) {
	//setup
	game := BuildNewGame(3, 3, 1)
	minedPointTile := [][2]int{{1, 1}}
	game.SetUpMines(minedPointTile)

	//execute
	stateGame, gameCopy := game.PlayMovement(0, 0)

	//assert
	if stateGame != StateGameRunning {
		t.Error("Error", stateGame, gameCopy)
	}
}

func TestMarkPlayMovementWhenRunningAndShowNumber(t *testing.T) {
	//setup
	game := BuildNewGame(3, 3, 1)
	minedPointTile := [][2]int{{1, 1}}
	game.SetUpMines(minedPointTile)

	//execute
	stateGame, gameCopy := game.PlayMovement(0, 0)
	fmt.Println(gameCopy)

	//assert
	if stateGame != StateGameRunning || len(gameCopy.Board) != 1 {
		t.Error("Error", stateGame, gameCopy, len(gameCopy.Board))
	}
}

func TestMarkPlayMovementWhenGameLost(t *testing.T) {
	//setup
	game := BuildNewGame(3, 3, 1)
	minedPointTile := [][2]int{{1, 1}}
	game.SetUpMines(minedPointTile)

	//execute
	stateGame, gameCopy := game.PlayMovement(1, 1)

	//assert
	if stateGame != StateGameLost {
		t.Error("Error", stateGame, gameCopy)
	}
}

func TestMarkPlayMovementWhenGameWon(t *testing.T) {
	//setup
	game := BuildNewGame(3, 3, 1)
	minedPointTile := [][2]int{{1, 1}}
	game.SetUpMines(minedPointTile)

	//execute
	stateGame, gameCopy := game.PlayMovement(0, 0)
	//assert
	if stateGame != StateGameRunning {
		t.Error("Error", stateGame, gameCopy)
	}

	stateGame, gameCopy = game.PlayMovement(0, 1)
	//assert
	if stateGame != StateGameRunning {
		t.Error("Error", stateGame, gameCopy)
	}

	stateGame, gameCopy = game.PlayMovement(0, 2)
	//assert
	if stateGame != StateGameRunning {
		t.Error("Error", stateGame, gameCopy)
	}

	stateGame, gameCopy = game.PlayMovement(1, 0)
	//assert
	if stateGame != StateGameRunning {
		t.Error("Error", stateGame, gameCopy)
	}

	stateGame, gameCopy = game.PlayMovement(1, 2)
	//assert
	if stateGame != StateGameRunning {
		t.Error("Error", stateGame, gameCopy)
	}

	stateGame, gameCopy = game.PlayMovement(2, 0)
	//assert
	if stateGame != StateGameRunning {
		t.Error("Error", stateGame, gameCopy)
	}

	stateGame, gameCopy = game.PlayMovement(2, 1)
	//assert
	if stateGame != StateGameRunning {
		t.Error("Error", stateGame, gameCopy)
	}

	stateGame, gameCopy = game.PlayMovement(2, 2)
	//assert
	if stateGame != StateGameWon {
		t.Error("Error", stateGame, gameCopy)
	}
}

func TestRevealEmptyAdjecentTiles3x3(t *testing.T) {
	game := BuildNewGame(3, 3, 1)
	minedPointTile := [][2]int{{1, 1}}
	game.SetUpMines(minedPointTile)

	game.RevealEmptyAdjacentTiles(0, 0)
	game.ShowBoard()

	expected := [][]StateTile{
		{StateTileCovered, StateTileCovered, StateTileCovered},
		{StateTileCovered, StateTileCovered, StateTileCovered},
		{StateTileCovered, StateTileCovered, StateTileCovered}}

	if !reflect.DeepEqual(expected, game.GetStates()) {
		t.Error("Error", game.GetStates())
	}
}

func TestRevealEmptyAdjecentTiles3x8(t *testing.T) {
	minedPointTile := [][2]int{{1, 1}}
	game := BuildNewGame(3, 8, len(minedPointTile))
	game.SetUpMines(minedPointTile)

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
	game := BuildNewGame(3, 3, 0)

	flagAmount := game.MarkFlag(1, 1)
	if flagAmount != 1 || game.FlagAmount != 1 {
		t.Error("Error", flagAmount, game.FlagAmount)
	}
}

func TestMarkFlagWhenRevert(t *testing.T) {
	game := BuildNewGame(3, 3, 0)

	flagAmount := game.MarkFlag(1, 1)
	flagAmount = game.MarkFlag(1, 1)

	if flagAmount != 0 || game.FlagAmount != 0 {
		t.Error("Error", flagAmount, game.FlagAmount)
	}
}

func TestGetAdjacentTilesShouldBe8(t *testing.T) {
	game := BuildNewGame(3, 3, 0)
	result := game.getAdjacentTiles(1, 1)

	for d := 0; d < len(result); d++ {
		fmt.Println(result[d].ValueTest)
	}

	if len(result) != 8 {
		t.Error("Error", len(result))
	}
}

func TestGetAdjacentTilesShouldBe3(t *testing.T) {
	game := BuildNewGame(3, 3, 0)
	result := game.getAdjacentTiles(2, 2)
	if len(result) != 3 {
		t.Error("Error", len(result))
	}
}

func TestGetAdjacentTilesShouldBe5(t *testing.T) {
	game := BuildNewGame(3, 3, 0)
	result := game.getAdjacentTiles(0, 1)
	if len(result) != 5 {
		t.Error("Error", len(result))
	}
}

func TestGetAdjacentTilesShouldBe0(t *testing.T) {
	game := BuildNewGame(1, 1, 0)
	result := game.getAdjacentTiles(0, 0)
	if len(result) != 0 {
		t.Error("Error", len(result))
	}
}

func TestBuildGame(t *testing.T) {
	game := BuildNewGame(3, 3, 0)
	game.ShowBoard()

	if len(game.Board) != 3 || game.Rows != 3 &&
		len(game.Board[0]) != 3 || game.Columns != 3 {
		t.Error("Error", len(game.Board), len(game.Board[0]))
	}
}
