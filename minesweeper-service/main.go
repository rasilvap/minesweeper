package main

import (
	"log"
	"minesweeper-API/minesweeper-service/controller"
	"net/http"

	"github.com/gorilla/mux"
)

const basePathAPI = "/api"

func main() {
	router := mux.NewRouter()

	controller.SetupRoutes(basePathAPI, router)

	log.Fatal(http.ListenAndServe(":5000", router))
}
