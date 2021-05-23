package engine

import (
	"encoding/json"
	"minesweeper-API/models"
	"minesweeper-API/models/dto"

	"github.com/obarra-dev/minesweeper"
)

type minesWeeper struct{}

func NewMinesweeper() MinesWeeper {
	return minesWeeper{}
}

func (minesWeeper) BuildGame(rows, columns, mineAmount int) (*models.Game, error) {
	mines := minesweeper.GenerateMines(rows, columns, mineAmount)
	minesweeper := minesweeper.New(rows, columns, mines)
	gameDS, err := buildGameDS(minesweeper)
	if err != nil {
		return nil, err
	}
	return gameDS, nil
}

func (minesWeeper) Play(playRequest dto.PlayRequest, game *models.Game) (*models.Game, *dto.PlayResponse, error) {
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
			Board:      string(j),
		},
		nil
}

func buildPlayResponse(game minesweeper.Game) dto.PlayResponse {
	gameStateDTO := mapStateGame(game.State)
	row := len(game.Board)
	if row == 0 {
		return dto.PlayResponse{
			StateGame: gameStateDTO,
			Game: dto.GameDTO{Board: [][]dto.TileDTO{},
				Rows:       game.Rows,
				Columns:    game.Columns,
			},
		}
	}

	boardDTO := make([][]dto.TileDTO, row)
	for i := 0; i < row; i++ {
		column := len(game.Board[i])
		boardDTO[i] = make([]dto.TileDTO, column)

		for j := 0; j < column; j++ {
			board := game.Board[i][j]
			tileStateDTO := mapTileState(board.State)
			boardDTO[i][j] = dto.TileDTO{
				State:                tileStateDTO,
				Row:                  board.Row,
				Column:               board.Column,
				SurroundingMineCount: board.SurroundingMineCount,
				Mine:                 board.IsMine,

				ValueTest: -1}
		}
	}

	return dto.PlayResponse{
		StateGame: gameStateDTO,
		Game: dto.GameDTO{
			Board:      boardDTO,
			Rows:       game.Rows,
			Columns:    game.Columns,
		},
	}
}
