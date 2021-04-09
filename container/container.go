package container

import (
	"fmt"
	"minesweeper-API/cmd/web/handlers"
	"minesweeper-API/datasource"
	"minesweeper-API/engine"
	"minesweeper-API/models"
)

func CreateHandler(c models.Config) handlers.GameHandler {
	var ds = createDataSourceSQL(c.Database)
	e := engine.NewGame(ds, engine.NewMinesweeper())
	return handlers.NewGameHandler(e)
}

func createDataSourceSQL(c models.DbConfig) datasource.Spec {
	ds, err := datasource.NewDatasourceSQL(c)
	if err != nil {
		panic(fmt.Sprintf("can't start DatasourceSQL: %v", err))
	}

	return ds
}

func createDataSourceMemory() datasource.Spec {
	return datasource.NewDatasourceMemory()
}
