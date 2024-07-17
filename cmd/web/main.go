package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux.
	mux := http.NewServeMux()

	// Create a file server out of th "./ui/static" directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// Register the file server as the handler for all URL paths
	// that start with "/static/".
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register each function as the handler for the corresponding URL pattern.
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /rank/{contest}/{year}/{division}/{class}", classRank)
	mux.HandleFunc("POST /rank/{contest}/{year}/{division}/{class}", classRankPost)
	mux.HandleFunc("GET /results/{contest}/{year}/{division}/{class}", classResults)

	// Print a log message to say that the server is starting.
	log.Print("starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server.
	// If http.ListenAndServe() returns an error we use the log.Fatal()
	// function to log the error message and exit.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
