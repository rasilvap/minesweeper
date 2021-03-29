package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/pdrum/swagger-automation/docs" // This line is necessary for go-swagger to find your docs!
)

func main() {
	port := getPort()

	router := setupRoutes()

	log.Fatal(http.ListenAndServe(port, router))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Using default port :5000")
		port = ":5000"
	} else {
		port = ":" + port
		log.Println("Using port ", port)
	}

	return port
}
