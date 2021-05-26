package handlers

import (
	"encoding/json"
	"log"
	"minesweeper-API/engine"
	"minesweeper-API/models/dto"
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

func (g game) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := g.gameEngine.Get(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if response == nil {
		log.Println("game not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	j, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, err = w.Write(j)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (g game) Create(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if request.Columns <= 0 || request.Rows <= 0 || request.MineAmount < 0 {
		log.Println("request invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := g.gameEngine.Create(request.Rows, request.Columns, request.MineAmount)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(dto.CreateGameResponse{GameID: id})
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

func (g game) Play(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var request dto.PlayRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if request.Column <= 0 || request.Row <= 0 {
		log.Println("request invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := request.Move.IsValid(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := g.gameEngine.Play(id, request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if response == nil {
		log.Println("game not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	j, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, err = w.Write(j)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
