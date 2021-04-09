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

type GameHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Play(w http.ResponseWriter, r *http.Request)
}

type gameHandler struct {
	gameEngine engine.Game
}

func NewGameHandler(gameEngine engine.Game) GameHandler {
	return gameHandler{gameEngine: gameEngine}
}

func (h gameHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	game, err := h.gameEngine.Get(id)
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

func (h gameHandler) Create(w http.ResponseWriter, r *http.Request) {
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

func (h gameHandler) Play(w http.ResponseWriter, r *http.Request) {
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
