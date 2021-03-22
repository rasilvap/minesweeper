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
	dbCs := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=True",
		config.User,
		config.Password,
		config.Server,
		config.Port,
		config.Database)

	db, err := sqlx.Connect("postgres", dbCs)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxOpenConn)
	db.SetMaxIdleConns(config.MaxIdleConn)
	db.SetConnMaxLifetime(config.ConnMaxLifeTime)

	return &Datasource{
		db: db,
	}, err
}
