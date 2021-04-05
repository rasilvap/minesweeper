package container

import (
	"fmt"
	"minesweeper-API/datasource"
	"minesweeper-API/engine"
	"minesweeper-API/model"
)

func CreateEngine(c model.Config) engine.Game {
	var ds = createDataSourceSQL(c.Database)
	return engine.NewGame(ds, engine.NewMinesweeper())
}

func createDataSourceSQL(c model.DbConfig) datasource.Spec {
	ds, err := datasource.NewDatasourceSQL(c)
	if err != nil {
		panic(fmt.Sprintf("can't start DatasourceSQL: %v", err))
	}

	return ds
}

func createDataSourceMemory() datasource.Spec {
	return datasource.NewDatasourceMemory()
}
