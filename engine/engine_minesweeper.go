package engine

import (
	"encoding/json"
	"minesweeper-API/minesweeper-service/model"

	"github.com/obarra-dev/minesweeper"
)

type MinesWeeperService interface {
	BuildGame(rows, columns, mineAmount int) (*model.Game, error)
	Play(playRequest model.PlayRequest, game *model.Game) (*model.Game, *model.PlayResponse, error)
}

type minesWeeperService struct{}

func NewMinesweeper() MinesWeeperService {
	return &minesWeeperService{}
}

func (*minesWeeperService) BuildGame(rows, columns, mineAmount int) (*model.Game, error) {
	mines := minesweeper.GenerateMinedPoints(rows, columns, mineAmount)
	minesweeper := minesweeper.NewMinesweeper(rows, columns, mines)
	gameDS, err := buildGameDS(minesweeper)
	if err != nil {
		return nil, err
	}
	return gameDS, nil
}

func (*minesWeeperService) Play(playRequest model.PlayRequest, game *model.Game) (*model.Game, *model.PlayResponse, error) {
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

func buildGameDS(g *minesweeper.Game) (*model.Game, error) {
	j, err := json.Marshal(g.Board)
	if err != nil {
		return nil, err
	}

	return &model.Game{
			State:      int(g.State),
			Columns:    g.Columns,
			Rows:       g.Rows,
			MineAmount: g.MineAmount,
			FlagAmount: g.FlagAmount,
			Board:      string(j),
		},
		nil
}

func buildPlayResponse(game minesweeper.Game) model.PlayResponse {
	gameStateDTO := mapStateGame(game.State)
	row := len(game.Board)
	if row == 0 {
		return model.PlayResponse{
			StateGame: gameStateDTO,
			Game: model.GameDTO{Board: [][]model.TileDTO{},
				Rows:       game.Rows,
				Columns:    game.Columns,
				FlagAmount: game.FlagAmount,
			},
		}
	}

	boardDTO := make([][]model.TileDTO, row)
	for i := 0; i < row; i++ {
		column := len(game.Board[i])
		boardDTO[i] = make([]model.TileDTO, column)

		for j := 0; j < column; j++ {
			board := game.Board[i][j]
			tileStateDTO := mapTileState(board.State)
			boardDTO[i][j] = model.TileDTO{
				State:                tileStateDTO,
				Row:                  board.Row,
				Column:               board.Column,
				SurroundingMineCount: board.SurroundingMineCount,
				Mine:                 board.IsMine,

				ValueTest: -1}
		}
	}

	return model.PlayResponse{
		StateGame: gameStateDTO,
		Game: model.GameDTO{
			Board:      boardDTO,
			Rows:       game.Rows,
			Columns:    game.Columns,
			FlagAmount: game.FlagAmount,
		},
	}
}
