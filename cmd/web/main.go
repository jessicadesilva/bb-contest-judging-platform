package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// Struct type for storing configuration settings.
	type config struct {
		addr      string
		staticDir string
	}

	var cfg config

	// Command line flag for network address with default value :4000.
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	// Command line flag for directory of static assets with default value ./ui/static.
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")

	// Parse command line flag values and assign to variables.
	flag.Parse()

	// Initialize a new structured logger that writes to the standard
	// out stream and includes the file source.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	// Use the http.NewServeMux() function to initialize a new servemux.
	mux := http.NewServeMux()

	// Create a file server out of the static assets directory.
	fileServer := http.FileServer(http.Dir(cfg.staticDir))
	// Register the file server as the handler for all URL paths
	// that start with "/static/".
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register each function as the handler for the corresponding URL pattern.
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /rank/{contest}/{year}/{division}/{class}", classRank)
	mux.HandleFunc("POST /rank/{contest}/{year}/{division}/{class}", classRankPost)
	mux.HandleFunc("GET /results/{contest}/{year}/{division}/{class}", classResults)

	// Print a log message to say that the server is starting.
	logger.Info("starting server", slog.String("addr", cfg.addr))

	// Use the http.ListenAndServe() function to start a new web server.
	// If http.ListenAndServe() returns an error, we log any error message
	// returned at Error severity and then terminate the application with
	// exit code 1.
	err := http.ListenAndServe(cfg.addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
