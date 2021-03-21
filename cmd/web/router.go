package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func setupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handleHealth).Methods(http.MethodGet)
	router.HandleFunc("/v1/games/{id:[0-9]+}", getOne).Methods(http.MethodGet)
	router.HandleFunc("/v1/games", create).Methods(http.MethodPost)
	router.HandleFunc("/v1/games/{id:[0-9]+}/play", play).Methods(http.MethodPost)

	return router
}
