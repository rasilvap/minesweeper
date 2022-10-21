package engine

import (
	"log"
	"minesweeper-API/models/dto"

	"minesweeper-API/datasource"
)

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

func (e game) Get(id int) (*dto.GetGameResponse, error) {
	e.gameDS.Find(id)

	return &dto.GetGameResponse{
			Rows: 1,
		},
		nil
}

func (e game) Play(id int, playRequest dto.PlayRequest) (*dto.PlayResponse, error) {
	log.Println("Playing game test", playRequest)
	g, err := e.gameDS.Find(id)
	if err != nil {
		log.Printf("Error finding g: %d, err: %v", id, err)
		return nil, err
	}

	if g == nil {
		return nil, nil
	}

	gameDS, playResponse, err := e.minesWeeperEngine.Play(playRequest, g)
	if err != nil {
		log.Printf("Error playing game: %d, err: %v", id, err)
		return nil, err
	}

	e.gameDS.Update(gameDS)

	return playResponse, nil
}
