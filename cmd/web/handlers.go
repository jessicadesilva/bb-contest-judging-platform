package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Define a home handler function which will be where a judge
// can login to the judging platform.

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	// Initialize a new structured logger that writes to the standard
	// out stream and includes the file source.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	// Initialize a slice containing the paths to the two HTML files.
	// The base template must be the first file in the slice.
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	// Read the template files into a template set.
	// We use ... to pass the contents of the files slice as variadic arguments.
	// If there's an error, log the detailed error message and return.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, "Internal Service Error", http.StatusInternalServerError)
		return
	}

	// Write the base template content as the response body. The last parameter
	// to Execute() represents dynamic data we want to pass in, for now is nil.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Define a class rank handler function which will be where
// judges rank competitors in a given class.

func classRank(w http.ResponseWriter, r *http.Request) {
	contest := r.PathValue("contest")
	year, err := strconv.Atoi(r.PathValue("year"))
	if err != nil || year != time.Now().Year() {
		http.NotFound(w, r)
		return
	}
	division := r.PathValue("division")
	class := r.PathValue("class")
	fmt.Fprintf(w, "Display a form to the judge to rank %s %s competitors for the %d %s contest.", division, class, year, contest)
}

func classRankPost(w http.ResponseWriter, r *http.Request) {
	contest := r.PathValue("contest")
	year, err := strconv.Atoi(r.PathValue("year"))
	if err != nil || year != time.Now().Year() {
		http.NotFound(w, r)
		return
	}
	division := r.PathValue("division")
	class := r.PathValue("class")

	// Use the w.WriteHeader() method to send a 201 status code.
	w.WriteHeader(http.StatusAccepted)

	fmt.Fprintf(w, "Save a ranking of %s %s competitors for the %d %s contest.", division, class, year, contest)
}

// Define a class results handler function which will be where
// compiled judge rankings are displayed.

func classResults(w http.ResponseWriter, r *http.Request) {
	contest := r.PathValue("contest")
	year, err := strconv.Atoi(r.PathValue("year"))
	if err != nil || year != time.Now().Year() {
		http.NotFound(w, r)
		return
	}
	division := r.PathValue("division")
	class := r.PathValue("class")
	message := fmt.Sprintf("See compiled rankings for the %d %s contest, Division: %s, Class: %s.", year, contest, division, class)
	w.Write([]byte(message))
}
