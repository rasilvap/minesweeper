package main

import (
	"log"
	"minesweeper-API/minesweeper-service/controller"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/pdrum/swagger-automation/docs" // This line is necessary for go-swagger to find your docs!
)

func main() {
	router := mux.NewRouter()

	controller.SetupRoutes(router)

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
