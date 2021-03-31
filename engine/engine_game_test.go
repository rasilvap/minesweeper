package engine

import (
	"testing"

	"github.com/obarra-dev/minesweeper"
	"github.com/stretchr/testify/mock"
)

type mockGameRepository struct {
	mock.Mock
}

func (mock *mockGameRepository) Save(game *minesweeper.Game) int {
	args := mock.Called(game)
	result := args.Get(0)
	return result.(int)
}

func (mock *mockGameRepository) Get(id int) *minesweeper.Game {
	args := mock.Called(id)
	result := args.Get(0)
	return result.(*minesweeper.Game)
}

/*
func TestGetOneGame(t *testing.T) {
	mockRepo := new(mockGameRepository)
	engine := NewGame(mockRepo)

	minedPointTile := [][2]int{{1, 1}}
	game := minesweeper.NewMinesweeper(3, 3, minedPointTile)
	id := 1
	mockRepo.On("Get", id).Return(game)

	result, _ := engine.Get(id)

	//Mock assertion: Behavioral
	mockRepo.AssertExpectations(t)

	expected := model.GameResponse{3, 3, 1}
	assert.Equal(t, expected, *result)
}

func TestCreateGame(t *testing.T) {
	mockRepo := new(mockGameRepository)
	engine := NewGame(mockRepo)

	minedPointTile := [][2]int{{1, 1}}
	game := minesweeper.NewMinesweeper(3, 3, minedPointTile)
	id := 1
	mockRepo.On("Save", game).Return(id)

	result, _ := engine.Create(3, 3, 1)

	//Mock assertion: Behavioral
	mockRepo.AssertExpectations(t)

	assert.Equal(t, id, result)
}


*/
func TestMarkPlayWhenLost3x3(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := minesweeper.NewMinesweeper(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(1, 1, minesweeper.TypeMoveClean)
	res := buildPlayResponse(gameCopy)

	//assert
	if res.StateGame != "LOST" || res.Game.Columns != 3 || res.Game.Rows != 3 ||
		len(res.Game.Board) != res.Game.Rows || len(res.Game.Board[0]) != res.Game.Columns {
		t.Error("Error", res, len(res.Game.Board), len(res.Game.Board[0]))
	}
}

func TestMarkPlayWhenWon3x3(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := minesweeper.NewMinesweeper(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(0, 0, minesweeper.TypeMoveClean)
	res := buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(0, 1, minesweeper.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res)
	}

	//execute
	gameCopy = game.Play(0, 2, minesweeper.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res)
	}

	//execute
	gameCopy = game.Play(1, 0, minesweeper.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(1, 2, minesweeper.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(2, 0, minesweeper.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(2, 1, minesweeper.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(2, 2, minesweeper.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "WON" {
		t.Error("Error", res, len(res.Game.Board))
	}
}

func TestMarkPlayWhenRunning3X8(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := minesweeper.NewMinesweeper(3, 8, minedPointTile)

	//execute
	gameCopy := game.Play(0, 5, minesweeper.TypeMoveClean)
	game.ShowBoard()
	res := buildPlayResponse(gameCopy)

	//assert
	if res.StateGame != "RUNNING" || res.Game.Rows != 3 || res.Game.Columns != 8 ||
		len(res.Game.Board) != 3 || len(res.Game.Board[0]) != 6 {
		t.Error("Error", res, len(res.Game.Board), len(res.Game.Board[0]))
	}
}
