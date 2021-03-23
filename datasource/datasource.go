package datasource

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"minesweeper-API/minesweeper-service/model"
)

type Datasource struct {
	db *sqlx.DB
}

// New Datasource creation
func NewDataSource(config model.DbConfig) (Spec, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Server, config.Port, config.User, config.Password, config.Database)

	db, err := sqlx.Connect("postgres", psqlInfo)

	if err != nil {

		panic("error")
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxOpenConn)
	db.SetMaxIdleConns(config.MaxIdleConn)
	db.SetConnMaxLifetime(config.ConnMaxLifeTime)

	return &Datasource{
		db: db,
	}, err
}
