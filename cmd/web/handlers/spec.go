package handlers

import "net/http"

type Game interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Play(w http.ResponseWriter, r *http.Request)
}