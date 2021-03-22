package datasource

import (
	"database/sql"
	"errors"
	"github.com/obarra-dev/minesweeper"
)

func (ds *Datasource) GetGame(id int) (*minesweeper.Game, error) {
	var game minesweeper.Game
	var err = ds.db.Get(&game, `SELECT * FROM games WHERE id = ?`, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return &game, errors.New("game not found")
		} else {
			return nil, errors.New("error")
		}
	}

	return &game, nil
}

func (ds *Datasource) SaveGame(g *minesweeper.Game) (int, error) {
	t, err := ds.db.Begin()
	if err != nil {
		return 0, err
	}

	_, err = t.Exec(`INSERT INTO games (id, status, status_detail, date_created, date_updated) 
								VALUES (?,?,?,?,?,?)`, g.Columns, g.Columns)
	if err != nil {
		_ = t.Rollback()
		return 0, err
	}

	err = t.Commit()
	if err != nil {
		return 0, err
	}

	return 1, nil
}
