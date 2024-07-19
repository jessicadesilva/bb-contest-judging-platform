package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	// Use the http.NewServeMux() function to initialize a new servemux.
	mux := http.NewServeMux()

	// Create a file server out of the static assets directory.
	fileServer := http.FileServer(http.Dir(app.staticDir))
	// Register the file server as the handler for all URL paths
	// that start with "/static/".
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register each function as the handler for the corresponding URL pattern.
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /rank/{contest}/{year}/{division}/{class}", app.classRank)
	mux.HandleFunc("POST /rank/{contest}/{year}/{division}/{class}", app.classRankPost)
	mux.HandleFunc("POST /competitor/{location}/{competitor}", app.createCompetitor)
	mux.HandleFunc("GET /results/{contest}/{year}/{division}/{class}", app.classResults)

	return mux
}
