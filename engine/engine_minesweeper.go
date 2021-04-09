package engine

import (
	"encoding/json"
	"minesweeper-API/models"

	"github.com/obarra-dev/minesweeper"
)

type MinesWeeper interface {
	BuildGame(rows, columns, mineAmount int) (*models.Game, error)
	Play(playRequest models.PlayRequest, game *models.Game) (*models.Game, *models.PlayResponse, error)
}

type minesWeeper struct{}

func NewMinesweeper() MinesWeeper {
	return minesWeeper{}
}

func (minesWeeper) BuildGame(rows, columns, mineAmount int) (*models.Game, error) {
	mines := minesweeper.GenerateMinedPoints(rows, columns, mineAmount)
	minesweeper := minesweeper.NewMinesweeper(rows, columns, mines)
	gameDS, err := buildGameDS(minesweeper)
	if err != nil {
		return nil, err
	}
	return gameDS, nil
}

func (minesWeeper) Play(playRequest models.PlayRequest, game *models.Game) (*models.Game, *models.PlayResponse, error) {
	var board [][]minesweeper.Tile
	err := json.Unmarshal([]byte(game.Board), &board)
	if err != nil {
		return nil, nil, err
	}

	minesweeper := minesweeper.Game{
		State:      minesweeper.StateGame(game.State),
		Rows:       game.Rows,
		Columns:    game.Columns,
		MineAmount: game.MineAmount,
		FlagAmount: game.FlagAmount,
		Board:      board,
	}

	gamedMinesweeper := minesweeper.Play(playRequest.Row, playRequest.Column, mapTypeMove(playRequest.Move))

	gamedDS, err := buildGameDS(&gamedMinesweeper)
	if err != nil {
		return nil, nil, err
	}

	playResponse := buildPlayResponse(gamedMinesweeper)

	return gamedDS, &playResponse, nil
}

func buildGameDS(g *minesweeper.Game) (*models.Game, error) {
	j, err := json.Marshal(g.Board)
	if err != nil {
		return nil, err
	}

	return &models.Game{
			State:      int(g.State),
			Columns:    g.Columns,
			Rows:       g.Rows,
			MineAmount: g.MineAmount,
			FlagAmount: g.FlagAmount,
			Board:      string(j),
		},
		nil
}

func buildPlayResponse(game minesweeper.Game) models.PlayResponse {
	gameStateDTO := mapStateGame(game.State)
	row := len(game.Board)
	if row == 0 {
		return models.PlayResponse{
			StateGame: gameStateDTO,
			Game: models.GameDTO{Board: [][]models.TileDTO{},
				Rows:       game.Rows,
				Columns:    game.Columns,
				FlagAmount: game.FlagAmount,
			},
		}
	}

	boardDTO := make([][]models.TileDTO, row)
	for i := 0; i < row; i++ {
		column := len(game.Board[i])
		boardDTO[i] = make([]models.TileDTO, column)

		for j := 0; j < column; j++ {
			board := game.Board[i][j]
			tileStateDTO := mapTileState(board.State)
			boardDTO[i][j] = models.TileDTO{
				State:                tileStateDTO,
				Row:                  board.Row,
				Column:               board.Column,
				SurroundingMineCount: board.SurroundingMineCount,
				Mine:                 board.IsMine,

				ValueTest: -1}
		}
	}

	return models.PlayResponse{
		StateGame: gameStateDTO,
		Game: models.GameDTO{
			Board:      boardDTO,
			Rows:       game.Rows,
			Columns:    game.Columns,
			FlagAmount: game.FlagAmount,
		},
	}
}
