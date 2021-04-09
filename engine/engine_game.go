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
	gameDS            datasource.Game
	minesWeeperEngine MinesWeeper
}

//TODO validate with pprof this with pointer
func NewGame(gameDS datasource.Game, minesWeeperEngine MinesWeeper) Game {
	return game{
		gameDS:            gameDS,
		minesWeeperEngine: minesWeeperEngine,
	}
}

func (e game) Create(rows, columns, mineAmount int) (int, error) {
	g, err := e.minesWeeperEngine.BuildGame(rows, columns, mineAmount)
	if err != nil {
		log.Printf("Error building game, err: %v", err)
		return 0, err
	}

	id, err := e.gameDS.Insert(g)
	if err != nil {
		log.Printf("Error creating game, err: %v", err)
		return 0, err
	}

	return id, nil
}

func (e game) Get(id int) (*models.GameResponse, error) {
	g, err := e.gameDS.Find(id)
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

func (e game) Play(id int, playRequest models.PlayRequest) (*models.PlayResponse, error) {
	log.Println("Playing game", playRequest)
	g, err := e.gameDS.Find(id)
	if err != nil {
		log.Printf("Error finding g: %d, err: %v", id, err)
		return nil, err
	}

	gameDS, playResponse, err := e.minesWeeperEngine.Play(playRequest, g)
	if err != nil {
		log.Printf("Error playing game: %d, err: %v", id, err)
		return nil, err
	}

	e.gameDS.Update(gameDS)

	return playResponse, nil
}
