package service

import (
	"minesweeper-API/minesweeper-service/engine"
	"minesweeper-API/minesweeper-service/model"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGameRepository struct {
	mock.Mock
}

func (mock *mockGameRepository) Save(game *engine.Game) int {
	args := mock.Called(game)
	result := args.Get(0)
	return result.(int)
}

func (mock *mockGameRepository) Get(id int) *engine.Game {
	args := mock.Called(id)
	result := args.Get(0)
	return result.(*engine.Game)
}

func TestGetOneGame(t *testing.T) {
	mockRepo := new(mockGameRepository)
	service := NewGameService(mockRepo)

	minedPointTile := [][2]int{{1, 1}}
	game := engine.BuildNewGame(3, 3, minedPointTile)
	id := 1
	mockRepo.On("Get", id).Return(game)

	result, _ := service.GetOneGame(id)

	//Mock assertion: Behavioral
	mockRepo.AssertExpectations(t)

	expected := model.GameResponse{3, 3, 1}
	assert.Equal(t, expected, *result)
}

func TestCreateGame(t *testing.T) {
	mockRepo := new(mockGameRepository)
	service := NewGameService(mockRepo)

	minedPointTile := [][2]int{{1, 1}}
	game := engine.BuildNewGame(3, 3, minedPointTile)
	id := 1
	mockRepo.On("Save", game).Return(id)

	result, _ := service.CreateGame(3, 3, 1)

	//Mock assertion: Behavioral
	mockRepo.AssertExpectations(t)

	assert.Equal(t, id, result)
}

func TestMarkPlayWhenLost3x3(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := engine.BuildNewGame(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(1, 1, engine.TypeMoveClean)
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
	game := engine.BuildNewGame(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(0, 0, engine.TypeMoveClean)
	res := buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(0, 1, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res)
	}

	//execute
	gameCopy = game.Play(0, 2, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res)
	}

	//execute
	gameCopy = game.Play(1, 0, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(1, 2, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(2, 0, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(2, 1, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "RUNNING" {
		t.Error("Error", res, len(res.Game.Board))
	}

	//execute
	gameCopy = game.Play(2, 2, engine.TypeMoveClean)
	res = buildPlayResponse(gameCopy)
	//assert
	if res.StateGame != "WON" {
		t.Error("Error", res, len(res.Game.Board))
	}
}

func TestMarkPlayWhenRunning3X8(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := engine.BuildNewGame(3, 8, minedPointTile)

	//execute
	gameCopy := game.Play(0, 5, engine.TypeMoveClean)
	game.ShowBoard()
	res := buildPlayResponse(gameCopy)

	//assert
	if res.StateGame != "RUNNING" || res.Game.Rows != 3 || res.Game.Columns != 8 ||
		len(res.Game.Board) != 3 || len(res.Game.Board[0]) != 6 {
		t.Error("Error", res, len(res.Game.Board), len(res.Game.Board[0]))
	}
}
