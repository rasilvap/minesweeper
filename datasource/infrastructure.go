package datasource

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"minesweeper-API/models"
	"sync"
)

type datasourceSQL struct {
	*sqlx.DB
}

// New datasourceSQL creation
func NewDatasourceSQL(config models.DbConfig) (*datasourceSQL, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Server, config.Port, config.User, config.Password, config.Database)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxOpenConn)
	db.SetMaxIdleConns(config.MaxIdleConn)
	db.SetConnMaxLifetime(config.ConnMaxLifeTime)

	return &datasourceSQL{db}, nil
}


type datasourceMemory struct {
	sync.RWMutex
	cache map[int]interface{}
}

func NewDatasourceMemory() datasourceMemory {
	return datasourceMemory{
		cache: make(map[int]interface{}),
	}
}