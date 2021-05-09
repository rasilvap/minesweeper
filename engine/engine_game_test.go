package engine

import (
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
	game := NewGame(mockRepo)

	minedPointTile := [][2]int{{1, 1}}
	game := minesweeper.NewMinesweeper(3, 3, minedPointTile)
	id := 1
	mockRepo.On("Get", id).Return(game)

	result, _ := game.Get(id)

	//Mock assertion: Behavioral
	mockRepo.AssertExpectations(t)

	expected := models.GameResponse{3, 3, 1}
	assert.Equal(t, expected, *result)
}

func TestCreateGame(t *testing.T) {
	mockRepo := new(mockGameRepository)
	game := NewGame(mockRepo)

	minedPointTile := [][2]int{{1, 1}}
	game := minesweeper.NewMinesweeper(3, 3, minedPointTile)
	id := 1
	mockRepo.On("Save", game).Return(id)

	result, _ := game.Create(3, 3, 1)

	//Mock assertion: Behavioral
	mockRepo.AssertExpectations(t)

	assert.Equal(t, id, result)
}


*/
