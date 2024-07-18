package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Struct type for application-wide dependencies.
type application struct {
	logger    *slog.Logger
	staticDir string
}

// Struct type for storing configuration settings.
type config struct {
	addr      string
	staticDir string
}

func main() {

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

	app := &application{logger: logger, staticDir: cfg.staticDir}

	// Print a log message to say that the server is starting.
	logger.Info("starting server", slog.String("addr", cfg.addr))

	// Use the http.ListenAndServe() function to start a new web server.
	// If http.ListenAndServe() returns an error, we log any error message
	// returned at Error severity and then terminate the application with
	// exit code 1.
	err := http.ListenAndServe(cfg.addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
