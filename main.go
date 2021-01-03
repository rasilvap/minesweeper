package main

import (
	"log"
	"minesweeper-API/minesweeper-service/controller"
	"minesweeper-API/minesweeper-service/repository"
	"minesweeper-API/minesweeper-service/service"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/pdrum/swagger-automation/docs" // This line is necessary for go-swagger to find your docs!
)

var (
	gameRespository repository.GameRepository = repository.NewMemoryRepository()
	gameService     service.GameService       = service.NewGameService(gameRespository)
	gameController  controller.GameController = controller.NewGameController(gameService)
)

func main() {
	router := setupRoutes()
	port := getPort()

	log.Fatal(http.ListenAndServe(port, router))
}

// SetupRoutes: ..
func setupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/v1/games/{id:[0-9]+}", gameController.GetOne).Methods("GET")
	router.HandleFunc("/v1/games", gameController.Create).Methods("POST")
	router.HandleFunc("/v1/games/{id:[0-9]+}/play", gameController.Play).Methods("POST")

	return router
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Using default port :5000")
		port = ":5000"
	} else {
		port = ":" + port
		log.Println("Using port ", port)
	}

	return port
}
