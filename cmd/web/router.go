package main

import (
	"minesweeper-API/engine"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(router *mux.Router, e engine.Game) {
	router.HandleFunc("/", handleHealth).Methods(http.MethodGet)

	router.HandleFunc("/v1/games/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			getOne(e, w, r)
		}).Methods(http.MethodGet)

	router.HandleFunc("/v1/games",
		func(w http.ResponseWriter, r *http.Request) {
			create(e, w, r)
		}).Methods(http.MethodPost)

	router.HandleFunc("/v1/games/{id:[0-9]+}/play",
		func(w http.ResponseWriter, r *http.Request) {
			play(e, w, r)
		}).Methods(http.MethodPost)
}
