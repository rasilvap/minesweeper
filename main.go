package main

import (
	"log"
	"minesweeper-API/minesweeper-service/controller"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/pdrum/swagger-automation/docs" // This line is necessary for go-swagger to find your docs!
)

var (
	gameController controller.GameController = controller.NewGameController()
)

func main() {
	router := mux.NewRouter()

	setupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Using default port :5000")
		port = ":5000"
	} else {
		port = ":" + port
		log.Println("Using port ", port)
	}

	log.Fatal(http.ListenAndServe(port, router))
}

// SetupRoutes: ..
func setupRoutes(router *mux.Router) {

	router.HandleFunc("/v1/games/{id:[0-9]+}", gameController.GetOne).Methods("GET")

	router.HandleFunc("/v1/games", gameController.Create).Methods("POST")

	router.HandleFunc("/v1/games/{id:[0-9]+}/play", gameController.Play).Methods("POST")
}
