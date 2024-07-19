package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/jessicadesilva/bb-contest-judging-platform/internal/models"

	// Alias unused package that we need to run its
	// init() function to register itself with the
	// database/sql package.
	_ "github.com/go-sql-driver/mysql"
)

// Struct type for application-wide dependencies.
type application struct {
	logger      *slog.Logger
	staticDir   string
	competitors *models.CompetitorModel
}

// Struct type for storing configuration settings.
type config struct {
	addr      string
	dsn       string
	staticDir string
}

func main() {

	var cfg config

	// Command line flag for network address with default value :4000.
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	// Command line flag for MySQL data source name (DSN) string.
	flag.StringVar(&cfg.dsn, "dsn", "web:pass@tcp(localhost:3306)/judgingplatform?parseTime=true", "MySQL data source name")
	// Command line flag for directory of static assets with default value ./ui/static.
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")

	// Parse command line flag values and assign to variables.
	flag.Parse()

	// Initialize a new structured logger that writes to the standard
	// out stream and includes the file source.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	// Open a connection pool and ping the database to check correct setup.
	db, err := openDB(cfg.dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Ensure connection pool is closed before main() exits.
	defer db.Close()

	app := &application{logger: logger,
		staticDir:   cfg.staticDir,
		competitors: &models.CompetitorModel{DB: db},
	}

	// Print a log message to say that the server is starting.
	logger.Info("starting server", slog.String("addr", cfg.addr))

	// Use the http.ListenAndServe() function to start a new web server.
	// If http.ListenAndServe() returns an error, we log any error message
	// returned at Error severity and then terminate the application with
	// exit code 1.
	err = http.ListenAndServe(cfg.addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

// The openDB() function wraps sql.Open() and returns a
// sql.DB connection pool for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
