package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/amit1502/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creating a new movie ...")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParam(r)

	if err != nil {
		app.logger.Printf("Movie with id %d not found\n", id)
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Kaho na pyaar hai",
		// Year:      2000,
		Runtime: 120,
		Genres:  []string{"Romantic", "Action", "Comedy"},
		Version: 1,
	}

	err = app.writeJson(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logger.Println("Error while parsing data:", err)
		http.Error(w, "Error while parsing data", http.StatusInternalServerError)
		return
	}

}
