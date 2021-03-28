package main

import (
	"encoding/json"
	"log"
	"minesweeper-API/minesweeper-service/datasource"
	"minesweeper-API/minesweeper-service/engine"
	"net/http"
	"strconv"

	"minesweeper-API/minesweeper-service/model"

	"github.com/gorilla/mux"
)

//TODO avoid global variables, and avoid interface segregation

var (
	ds, _ = datasource.NewDataSource(model.DbConfig{
		Server:          "localhost",
		Port:            5432,
		User:            "postgres",
		Password:        "postgres",
		Database:        "postgres",
		MaxOpenConn:     100,
		MaxIdleConn:     50,
		ConnMaxLifeTime: 0,
	})

	gameService = engine.NewGameService(ds,
		engine.NewMinesWeeperService(),
	)
)

func getOne(w http.ResponseWriter, r *http.Request) {
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

func create(w http.ResponseWriter, r *http.Request) {
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

	j, err := json.Marshal(model.GameSimpleResponse{GameId: id})
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

	playResponse, err := gameService.PlayMove(id, playRequest)
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
