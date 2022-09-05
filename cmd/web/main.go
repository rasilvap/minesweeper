package main

import (
	"flag"
	"fmt"
	"log"
	"minesweeper-API/cmd/web/handlers"
	"minesweeper-API/config"
	"minesweeper-API/container"
	"minesweeper-API/models"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "github.com/pdrum/swagger-automation/docs" // This line is necessary for go-swagger to find your docs!
)

/*I'm adding this project to my whitelist, guys */
func main() {
	env := flag.String("env", "dev", "Execution environment")
	flag.Parse()
	log.Printf("Starting application server - %s", *env)

	cfg := config.BuildConfig(*env)
	c := container.New(cfg)

	r := createServer()
	handlers.SetupRoutes(r, c.GameHandler)
	log.Fatal(http.ListenAndServe(getPort(cfg.Server), r))
}

func createServer() *mux.Router {
	return mux.NewRouter()
}

func getPort(c models.ServerConfig) string {
	log.Println("Using port ", c.Port)
	return fmt.Sprintf(":%d", c.Port)
}
