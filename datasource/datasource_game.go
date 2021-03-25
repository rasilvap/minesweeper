package datasource

import (
	"database/sql"
	"minesweeper-API/minesweeper-service/model"
)

func (ds *Datasource) GetGame(id int) (*model.Game, error) {
	var game model.Game
	switch err := ds.db.Get(&game, `SELECT * FROM minesweeper.games WHERE game_id = $1`, id); err {
	case nil, sql.ErrNoRows:
		return &game, nil
	default:
		return nil, err
	}
}

func (ds *Datasource) SaveGame(g *model.Game) (int, error) {
	res, err := ds.db.NamedQuery(
		`INSERT INTO minesweeper.games (state, columns, rows, mine_amount, flag_amount)
 		VALUES (:state, :columns, :rows, :mine_amount, :flag_amount) returning game_id`,
		&g)

	res.Next()
	var id int
	err = res.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
