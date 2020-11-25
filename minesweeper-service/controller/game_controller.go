package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"minesweeper-API/minesweeper-service/model"
	"minesweeper-API/minesweeper-service/service"

	"github.com/gorilla/mux"
)

// SetupRoutes: ..
func SetupRoutes(router *mux.Router) {

	router.HandleFunc("/v1/games/{id:[0-9]+}", getOne).Methods("GET")

	router.HandleFunc("/v1/games", create).Methods("POST")

	router.HandleFunc("/v1/games/{id:[0-9]+}/play", play).Methods("POST")

	router.HandleFunc("/v1/games/{id:[0-9]+}/mark", mark).Methods("POST")
}

func getOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	game, err := service.GetOneGame(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if game == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	j, err := json.Marshal(game)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, err = w.Write(j)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	var gameRequest model.GameRequest
	err := json.NewDecoder(r.Body).Decode(&gameRequest)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := service.CreateGame(gameRequest.Rows, gameRequest.Columns, gameRequest.MineAmount)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(model.GameSimpleResponse{id})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(j)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	log.Println("Playing")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var playRequest model.PlayRequest
	err = json.NewDecoder(r.Body).Decode(&playRequest)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	playRespose, err := service.PlayMovement(id, playRequest.Row, playRequest.Column)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if playRespose == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Println(playRespose)
	j, err := json.Marshal(playRespose)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	_, err = w.Write(j)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func mark(w http.ResponseWriter, r *http.Request) {
	log.Println("Marking")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var markRequest model.MarkRequest
	err = json.NewDecoder(r.Body).Decode(&markRequest)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = service.MarkTile(id, markRequest.Row, markRequest.Column, markRequest.Mark)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
}
