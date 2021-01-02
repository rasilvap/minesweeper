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

type controller struct{}

type GameController interface {
	GetOne(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Play(w http.ResponseWriter, r *http.Request)
}

var (
	gameService service.GameService = service.NewGameService()
)

func NewGameController() GameController {
	return &controller{}
}

func (*controller) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	game, err := gameService.GetOneGame(id)
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

func (*controller) Create(w http.ResponseWriter, r *http.Request) {
	var gameRequest model.GameRequest
	err := json.NewDecoder(r.Body).Decode(&gameRequest)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := gameService.CreateGame(gameRequest.Rows, gameRequest.Columns, gameRequest.MineAmount)
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

func (*controller) Play(w http.ResponseWriter, r *http.Request) {
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

	playRespose, err := gameService.PlayMove(id, playRequest)
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
