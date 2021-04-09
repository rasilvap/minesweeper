package container

import (
	"fmt"
	"minesweeper-API/cmd/web/handlers"
	"minesweeper-API/datasource"
	"minesweeper-API/engine"
	"minesweeper-API/models"
)

//TODO return a generic container
func CreateHandler(c models.Config) handlers.Game {
	var ds = createDataSourceSQL(c.Database)
	e := engine.NewGame(ds, engine.NewMinesweeper())
	return handlers.NewGame(e)
}

func createDataSourceSQL(c models.DbConfig) datasource.Game {
	ds, err := datasource.NewDatasourceSQL(c)
	if err != nil {
		panic(fmt.Sprintf("can't start DatasourceSQL: %v", err))
	}

	return datasource.NewGameSQl(ds)
}

func createDataSourceMemory() datasource.Game {
	return datasource.NewDatasourceMemory()
}
