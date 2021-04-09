package datasource

import (
	"database/sql"
	"minesweeper-API/models"
)


type gameSql struct {
	*datasourceSQL
}

func NewGameSQl(ds *datasourceSQL) Game{
	return gameSql{ds}
}

func (ds gameSql) Find(id int) (*models.Game, error) {
	var game models.Game
	switch err := ds.Get(&game, `SELECT * FROM minesweeper.games WHERE game_id = $1`, id); err {
	case nil, sql.ErrNoRows:
		return &game, nil
	default:
		return nil, err
	}
}

func (ds gameSql) Insert(g *models.Game) (int, error) {
	res, err := ds.NamedQuery(
		`INSERT INTO minesweeper.games (state, columns, rows, mine_amount, flag_amount, board)
 		VALUES (:state, :columns, :rows, :mine_amount, :flag_amount, :board) returning game_id`,
		&g)

	if err != nil {
		return 0, err
	}

	res.Next()
	var id int
	err = res.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ds gameSql) Update(g *models.Game) error {
	_, err := ds.NamedQuery(
		`UPDATE  minesweeper.games  SET state = :state, 
                               columns = :columns, rows = :rows, mine_amount = :mine_amount, flag_amount = :flag_amount, 
                               board = :board WHERE game_id = :game_id`,
		&g)

	return err
}
