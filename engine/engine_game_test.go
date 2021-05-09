package engine

import (
	"errors"
	"minesweeper-API/models/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"minesweeper-API/models"
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

func (mock *minesWeeperMock) Play(playRequest dto.PlayRequest, game *models.Game) (*models.Game, *dto.PlayResponse, error) {
	args := mock.Called(playRequest, game)

	if r := args.Error(2); r != nil {
		return nil, nil, r
	}

	resultGame := args.Get(0).(*models.Game)
	resultPlayResponse := args.Get(1).(*dto.PlayResponse)
	return resultGame, resultPlayResponse, nil
}

func Test_Create(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
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

	t.Run("error when build", func(t *testing.T) {
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

	t.Run("error when insert", func(t *testing.T) {
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

func Test_Get(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		minesWeeper := new(minesWeeperMock)
		gameDS := new(gameDSMock)
		game := NewGame(gameDS, minesWeeper)
		g := models.Game{
			GameId:     123,
			State:      1,
			Columns:    3,
			Rows:       3,
			MineAmount: 1,
			FlagAmount: 0,
			Board:      "",
		}

		gameDS.On("Find", 123).Return(&g, nil)

		//act
		got, err := game.Get(123)

		gameDS.AssertExpectations(t)
		minesWeeper.AssertExpectations(t)
		assert.Equal(t, &dto.GameResponse{
			Rows:       3,
			Columns:    3,
			MineAmount: 1,
		}, got)
		assert.Nil(t, err)
	})

	t.Run("Not found", func(t *testing.T) {
		minesWeeper := new(minesWeeperMock)
		gameDS := new(gameDSMock)
		game := NewGame(gameDS, minesWeeper)
		gameDS.On("Find", 123).Return(&models.Game{}, nil)

		//act
		got, err := game.Get(123)

		gameDS.AssertExpectations(t)
		minesWeeper.AssertExpectations(t)
		assert.Nil(t, got)
		assert.Nil(t, err)
	})

	t.Run("Error when find", func(t *testing.T) {
		minesWeeper := new(minesWeeperMock)
		gameDS := new(gameDSMock)
		game := NewGame(gameDS, minesWeeper)
		gameDS.On("Find", 123).Return(nil, errors.New("some error"))

		//act
		got, err := game.Get(123)

		gameDS.AssertExpectations(t)
		minesWeeper.AssertExpectations(t)
		assert.Nil(t, got)
		assert.Error(t, err)
	})
}

func Test_Play(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		minesWeeper := new(minesWeeperMock)
		gameDS := new(gameDSMock)
		game := NewGame(gameDS, minesWeeper)
		g := models.Game{
			GameId:     123,
			State:      1,
			Columns:    3,
			Rows:       3,
			MineAmount: 1,
			FlagAmount: 0,
			Board:      "",
		}
		req := dto.PlayRequest{
			Row:    0,
			Column: 0,
			Move:   "CLEAN",
		}

		res := dto.PlayResponse{
			StateGame: "ACTIVE",
			Game: dto.GameDTO{
				Board:      nil,
				Rows:       3,
				Columns:    3,
				FlagAmount: 0,
			},
		}

		gPlayed := g
		gPlayed.State = 2
		gameDS.On("Find", 123).Return(&g, nil)
		minesWeeper.On("Play", req, &g).Return(&gPlayed, &res, nil)
		gameDS.On("Update", &gPlayed).Return(nil)

		//act
		got, err := game.Play(123, req)

		gameDS.AssertExpectations(t)
		minesWeeper.AssertExpectations(t)
		assert.Equal(t, &res, got)
		assert.Nil(t, err)
	})

	t.Run("error when find", func(t *testing.T) {
		minesWeeper := new(minesWeeperMock)
		gameDS := new(gameDSMock)
		game := NewGame(gameDS, minesWeeper)

		req := dto.PlayRequest{
			Row:    0,
			Column: 0,
			Move:   "CLEAN",
		}
		gameDS.On("Find", 123).Return(nil, errors.New("some error"))

		//act
		got, err := game.Play(123, req)

		gameDS.AssertExpectations(t)
		minesWeeper.AssertExpectations(t)
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("error when play", func(t *testing.T) {
		minesWeeper := new(minesWeeperMock)
		gameDS := new(gameDSMock)
		game := NewGame(gameDS, minesWeeper)
		g := models.Game{
			GameId:     123,
			State:      1,
			Columns:    3,
			Rows:       3,
			MineAmount: 1,
			FlagAmount: 0,
			Board:      "",
		}
		req := dto.PlayRequest{
			Row:    0,
			Column: 0,
			Move:   "CLEAN",
		}

		gameDS.On("Find", 123).Return(&g, nil)
		minesWeeper.On("Play", req, &g).Return(nil, nil, errors.New("some error"))

		//act
		got, err := game.Play(123, req)

		gameDS.AssertExpectations(t)
		minesWeeper.AssertExpectations(t)
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("error when update", func(t *testing.T) {
		minesWeeper := new(minesWeeperMock)
		gameDS := new(gameDSMock)
		game := NewGame(gameDS, minesWeeper)
		g := models.Game{
			GameId:     123,
			State:      1,
			Columns:    3,
			Rows:       3,
			MineAmount: 1,
			FlagAmount: 0,
			Board:      "",
		}
		req := dto.PlayRequest{
			Row:    0,
			Column: 0,
			Move:   "CLEAN",
		}

		res := dto.PlayResponse{
			StateGame: "ACTIVE",
			Game: dto.GameDTO{
				Board:      nil,
				Rows:       3,
				Columns:    3,
				FlagAmount: 0,
			},
		}

		gPlayed := g
		gPlayed.State = 2
		gameDS.On("Find", 123).Return(&g, nil)
		minesWeeper.On("Play", req, &g).Return(&gPlayed, &res, nil)
		gameDS.On("Update", &gPlayed).Return(errors.New("some error"))

		//act
		got, err := game.Play(123, req)

		gameDS.AssertExpectations(t)
		minesWeeper.AssertExpectations(t)
		assert.Nil(t, got)
		assert.Error(t, err)
	})
}
