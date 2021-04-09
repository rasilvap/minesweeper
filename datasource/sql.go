package datasource

import (
	"database/sql"
	"fmt"
	"minesweeper-API/models"

	"github.com/jmoiron/sqlx"
)

type datasourceSQL struct {
	db *sqlx.DB
}

// New datasourceSQL creation
func NewDatasourceSQL(config models.DbConfig) (Spec, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Server, config.Port, config.User, config.Password, config.Database)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxOpenConn)
	db.SetMaxIdleConns(config.MaxIdleConn)
	db.SetConnMaxLifetime(config.ConnMaxLifeTime)

	return &datasourceSQL{
		db: db,
	}, nil
}

func (ds *datasourceSQL) FindGame(id int) (*models.Game, error) {
	var game models.Game
	switch err := ds.db.Get(&game, `SELECT * FROM minesweeper.games WHERE game_id = $1`, id); err {
	case nil, sql.ErrNoRows:
		return &game, nil
	default:
		return nil, err
	}
}

func (ds *datasourceSQL) InsertGame(g *models.Game) (int, error) {
	res, err := ds.db.NamedQuery(
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

func (ds *datasourceSQL) UpdateGame(g *models.Game) error {
	_, err := ds.db.NamedQuery(
		`UPDATE  minesweeper.games  SET state = :state, 
                               columns = :columns, rows = :rows, mine_amount = :mine_amount, flag_amount = :flag_amount, 
                               board = :board WHERE game_id = :game_id`,
		&g)

	return err
}
