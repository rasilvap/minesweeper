package container

import (
	"fmt"
	"minesweeper-API/cmd/web/handler"
	"minesweeper-API/datasource"
	"minesweeper-API/engine"
	"minesweeper-API/model"
)

func CreateHandler(c model.Config) handler.GameHandler {
	var ds = createDataSourceSQL(c.Database)
	e := engine.NewGame(ds, engine.NewMinesweeper())
	return handler.NewGameHandler(e)
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
