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
	var board []dto.TileDTO
	for r := 0; r < game.Rows; r++ {
		for c := 0; c < game.Columns; c++ {
			tile := game.Board[r][c]
			tileStateDTO := mapTileState(tile.State)
			board = append(board, dto.TileDTO{
				State:                tileStateDTO,
				Row:                  tile.Row,
				Column:               tile.Column,
				SurroundingMineCount: tile.SurroundingMineCount,
				Mine:                 tile.IsMine,
			})
		}
	}

	gameStateDTO := mapStateGame(game.State)
	return dto.PlayResponse{
		Game: dto.GameDTO{
			StateGame: gameStateDTO,
			Rows:      game.Rows,
			Columns:   game.Columns,
			Board:     board,
		},
	}
}
