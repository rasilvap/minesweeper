package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, h Game) {
	router.HandleFunc("/", HandleHealth).Methods(http.MethodGet)
	setupRoutesGameHandler(router, h)
}

func setupRoutesGameHandler(router *mux.Router, h Game) {
	router.HandleFunc("/v1/games/{id:[0-9]+}", h.Get).Methods(http.MethodGet)
	router.HandleFunc("/v1/games", h.Create).Methods(http.MethodPost)
	router.HandleFunc("/v1/games/{id:[0-9]+}/play", h.Play).Methods(http.MethodPost)
}
