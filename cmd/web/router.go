package main

import (
	"minesweeper-API/cmd/web/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(router *mux.Router, h handlers.GameHandler) {
	router.HandleFunc("/", handlers.HandleHealth).Methods(http.MethodGet)
	setupRoutesGameHandler(router, h)
}

func setupRoutesGameHandler(router *mux.Router, h handlers.GameHandler) {
	router.HandleFunc("/v1/games/{id:[0-9]+}", h.Get).Methods(http.MethodGet)
	router.HandleFunc("/v1/games", h.Create).Methods(http.MethodPost)
	router.HandleFunc("/v1/games/{id:[0-9]+}/play", h.Play).Methods(http.MethodPost)
}
