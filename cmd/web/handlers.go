package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

// Define a home handler function which will be where a judge
// can login to the judging platform.

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

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
		app.serverError(w, r, err)
		return
	}

	// Write the base template content as the response body. The last parameter
	// to Execute() represents dynamic data we want to pass in, for now is nil.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// Define a class rank handler function which will be where
// judges rank competitors in a given class.

func (app *application) classRank(w http.ResponseWriter, r *http.Request) {
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

func (app *application) createCompetitor(w http.ResponseWriter, r *http.Request) {
	// Create dummy data.
	competitorName := "Ashley Kaltwasser"
	competitorLocation := "Las Vegas, NV"

	// Pass data into ContestEntryModel.Insert() method, receiving
	// the ID of the new record back.
	_, err := app.competitors.Insert(competitorName, competitorLocation)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Redirect the user to the relevant page for the contest entry.
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) classRankPost(w http.ResponseWriter, r *http.Request) {
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

func (app *application) classResults(w http.ResponseWriter, r *http.Request) {
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
