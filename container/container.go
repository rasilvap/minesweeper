package container

import (
	"fmt"
	"minesweeper-API/cmd/web/handlers"
	"minesweeper-API/datasource"
	"minesweeper-API/engine"
	"minesweeper-API/models"
)

type container struct {
	GameHandler handlers.Game
}

//TODO return a generic container
func New(c models.Config) container {
	var ds = createDataSourceSQL(c.Database)
	var mw = engine.NewMinesweeper()

	e := engine.NewGame(ds, mw)

	return container{
		GameHandler: handlers.NewGame(e),
	}
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
