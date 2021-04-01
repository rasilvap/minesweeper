package main

import (
	"github.com/gorilla/mux"
	"log"
	"minesweeper-API/minesweeper-service/container"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/pdrum/swagger-automation/docs" // This line is necessary for go-swagger to find your docs!
)

func main() {
	e := container.CreateEngine()
	r := createServer()
	setupRoutes(r, e)
	p := getPort()
	log.Fatal(http.ListenAndServe(p, r))
}

func createServer() *mux.Router {
	return mux.NewRouter()
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Using default port :5000")
		port = ":5001"
	} else {
		port = ":" + port
		log.Println("Using port ", port)
	}

	return port
}
