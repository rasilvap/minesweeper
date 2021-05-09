package main

import (
	"flag"
	"fmt"
	"log"
	"minesweeper-API/config"
	"minesweeper-API/container"
	"minesweeper-API/models"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "github.com/pdrum/swagger-automation/docs" // This line is necessary for go-swagger to find your docs!
)

var (
	 LoggerInfo  *log.Logger
	 LoggerError *log.Logger
)

func main() {
	file, _ := os.OpenFile("log.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)

	LoggerInfo = log.New(file, "INFO: ", log.LUTC|log.Lmicroseconds|log.Lshortfile )
	LoggerError = log.New(file, "Error: ", log.LUTC|log.Lmicroseconds|log.Lshortfile )
	env := flag.String("env", "dev", "Execution environment")
	flag.Parse()
	LoggerInfo.Printf("Starting application server - %s\n", *env)

	cfg := config.BuildConfig(*env)
	c := container.New(cfg)

	r := createServer()
	setupRoutes(r, c.GameHandler)
	log.Fatal(http.ListenAndServe(getPort(cfg.Server), r))
}

func createServer() *mux.Router {
	return mux.NewRouter()
}

func getPort(c models.ServerConfig) string {
	log.Println("Using port ", c.Port)
	return fmt.Sprintf(":%d", c.Port)
}
