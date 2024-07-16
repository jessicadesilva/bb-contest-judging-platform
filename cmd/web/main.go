package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux.
	mux := http.NewServeMux()
	// Register each function as the handler for the corresponding URL pattern.
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /{contest}/{year}/{division}/{class}/rank", classRank)
	mux.HandleFunc("POST /{contest}/{year}/{division}/{class}/rank", classRankPost)
	mux.HandleFunc("GET /{contest}/{year}/{division}/{class}/results", classResults)

	// Print a log message to say that the server is starting.
	log.Print("starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server.
	// If http.ListenAndServe() returns an error we use the log.Fatal()
	// function to log the error message and exit.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
