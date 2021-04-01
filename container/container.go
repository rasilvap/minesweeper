package container

import (
	"fmt"
	"minesweeper-API/minesweeper-service/datasource"
	"minesweeper-API/minesweeper-service/engine"
	"minesweeper-API/minesweeper-service/model"
)

func CreateEngine() engine.GameService {
	var ds = createDataSourceSQL()
	return engine.NewGame(ds, engine.NewMinesweeper())
}

func createDataSourceSQL() datasource.Spec {
	ds, err := datasource.NewDatasourceSQL(model.DbConfig{
		Server:          "localhost",
		Port:            5432,
		User:            "postgres",
		Password:        "postgres",
		Database:        "postgres",
		MaxOpenConn:     100,
		MaxIdleConn:     50,
		ConnMaxLifeTime: 0,
	})

	if err != nil {
		panic(fmt.Sprintf("can't start DatasourceSQL: %v", err))
	}

	return ds
}

func createDataSourceMemory() datasource.Spec {
	return datasource.NewDatasourceMemory()
}
