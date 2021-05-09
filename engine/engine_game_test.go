package engine

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"minesweeper-API/models"
	"testing"
)

type gameDSMock struct {
	mock.Mock
}

func (mock *gameDSMock) Find(id int) (*models.Game, error) {
	args := mock.Called(id)

	if r := args.Error(1); r != nil {
		return nil, r
	}

	return args.Get(0).(*models.Game), nil
}

func (mock *gameDSMock) Insert(g *models.Game) (int, error) {
	args := mock.Called(g)
	return args.Int(0), args.Error(1)
}

func (mock *gameDSMock) Update(g *models.Game) error {
	args := mock.Called(g)
	return args.Error(0)
}

type minesWeeperMock struct {
	mock.Mock
}

func (mock *minesWeeperMock) BuildGame(rows, columns, mineAmount int) (*models.Game, error) {
	args := mock.Called(rows, columns, mineAmount)
	if r := args.Error(1); r != nil {
		return nil, r
	}

	return args.Get(0).(*models.Game), nil
}

func (mock *minesWeeperMock) Play(playRequest models.PlayRequest, game *models.Game) (*models.Game, *models.PlayResponse, error) {
	args := mock.Called(playRequest, game)
	resultGame := args.Get(0).(*models.Game)
	resultPlayResponse := args.Get(1).(*models.PlayResponse)
	return resultGame, resultPlayResponse, args.Error(2)
}

func Test_CreateGame(t *testing.T) {
	t.Run("Crete ok", func(t *testing.T) {
		minesWeeper := new(minesWeeperMock)
		gameDS := new(gameDSMock)
		game := NewGame(gameDS, minesWeeper)
		g := models.Game{
			GameId:     0,
			State:      1,
			Columns:    3,
			Rows:       3,
			MineAmount: 1,
			FlagAmount: 0,
			Board:      "",
		}
		minesWeeper.On("BuildGame", 3, 3, 1).Return(&g, nil)

		gameDS.On("Insert", &g).Return(123, nil)

		//act
		got, err := game.Create(3, 3, 1)

		gameDS.AssertExpectations(t)
		minesWeeper.AssertExpectations(t)
		assert.Equal(t, 123, got)
		assert.Nil(t, err)

	})

	t.Run("when build with error", func(t *testing.T) {
		minesWeeper := new(minesWeeperMock)
		gameDS := new(gameDSMock)
		game := NewGame(gameDS, minesWeeper)

		minesWeeper.On("BuildGame", 3, 3, 1).Return(nil, errors.New("some error"))

		//act
		got, err := game.Create(3, 3, 1)

		gameDS.AssertExpectations(t)
		minesWeeper.AssertExpectations(t)
		assert.Zero(t, got)
		assert.Error(t, err)
	})

	t.Run("when insert with error", func(t *testing.T) {
		minesWeeper := new(minesWeeperMock)
		gameDS := new(gameDSMock)
		game := NewGame(gameDS, minesWeeper)

		g := models.Game{
			GameId:     0,
			State:      1,
			Columns:    3,
			Rows:       3,
			MineAmount: 1,
			FlagAmount: 0,
			Board:      "",
		}
		minesWeeper.On("BuildGame", 3, 3, 1).Return(&g, nil)
		gameDS.On("Insert", &g).Return(0, errors.New("some error"))

		//act
		got, err := game.Create(3, 3, 1)

		gameDS.AssertExpectations(t)
		minesWeeper.AssertExpectations(t)
		assert.Zero(t, got)
		assert.Error(t, err)
	})

}
