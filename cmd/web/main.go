package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "github.com/pdrum/swagger-automation/docs" // This line is necessary for go-swagger to find your docs!
	"log"
	"minesweeper-API/minesweeper-service/config"
	"minesweeper-API/minesweeper-service/container"
	"minesweeper-API/minesweeper-service/model"
	"net/http"
)

func main() {
	c := config.BuildConfig("dev")
	e := container.CreateEngine(c)
	r := createServer()
	setupRoutes(r, e)

	log.Fatal(http.ListenAndServe(getPort(c.Server), r))
}

func createServer() *mux.Router {
	return mux.NewRouter()
}

func getPort(c model.ServerConfig) string {
	log.Println("Using port ", c.Port)
	return fmt.Sprintf(":%d", c.Port)
}
