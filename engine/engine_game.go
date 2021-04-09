package engine

import (
	"log"

	"minesweeper-API/datasource"

	"minesweeper-API/models"
)

type Game interface {
	Create(rows, columns, mineAmount int) (int, error)
	Play(id int, playRequest models.PlayRequest) (*models.PlayResponse, error)
	Get(id int) (*models.GameResponse, error)
}

type game struct {
	gameDS      datasource.Spec
	minesWeeper MinesWeeper
}

func NewGame(gameDSImp datasource.Spec, minesWeeperService MinesWeeper) Game {
	return &game{
		gameDS:      gameDSImp,
		minesWeeper: minesWeeperService,
	}
}

func (s game) Create(rows, columns, mineAmount int) (int, error) {
	game, err := s.minesWeeper.BuildGame(rows, columns, mineAmount)
	if err != nil {
		log.Printf("Error building game, err: %v", err)
		return 0, err
	}

	id, err := s.gameDS.InsertGame(game)
	if err != nil {
		log.Printf("Error creating game, err: %v", err)
		return 0, err
	}

	return id, nil
}

func (s game) Get(id int) (*models.GameResponse, error) {
	g, err := s.gameDS.FindGame(id)
	if err != nil {
		log.Printf("Error finding game: %d, err: %v", id, err)
		return nil, err
	}

	if g.GameId == 0 {
		return nil, nil
	}

	return &models.GameResponse{
			Rows:       g.Rows,
			Columns:    g.Columns,
			MineAmount: g.MineAmount,
		},
		nil
}

func (s *game) Play(id int, playRequest models.PlayRequest) (*models.PlayResponse, error) {
	log.Println("Playing game", playRequest)
	game, err := s.gameDS.FindGame(id)
	if err != nil {
		log.Printf("Error finding g: %d, err: %v", id, err)
		return nil, err
	}

	gameDS, playResponse, err := s.minesWeeper.Play(playRequest, game)
	if err != nil {
		log.Printf("Error playing game: %d, err: %v", id, err)
		return nil, err
	}

	s.gameDS.UpdateGame(gameDS)

	return playResponse, nil
}
