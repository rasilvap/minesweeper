package handlers

import (
	"encoding/json"
	"log"
	"minesweeper-API/engine"
	"minesweeper-API/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type game struct {
	gameEngine engine.Game
}

func NewGame(gameEngine engine.Game) Game {
	return game{gameEngine: gameEngine}
}

func (h game) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	g, err := h.gameEngine.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if g == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	j, err := json.Marshal(g)
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

func (h game) Create(w http.ResponseWriter, r *http.Request) {
	var gameRequest models.GameRequest
	err := json.NewDecoder(r.Body).Decode(&gameRequest)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := h.gameEngine.Create(gameRequest.Rows, gameRequest.Columns, gameRequest.MineAmount)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(models.GameSimpleResponse{GameId: id})
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

func (h game) Play(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var playRequest models.PlayRequest
	err = json.NewDecoder(r.Body).Decode(&playRequest)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	playResponse, err := h.gameEngine.Play(id, playRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if playResponse == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Println(playResponse)
	j, err := json.Marshal(playResponse)
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
