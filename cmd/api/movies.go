package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/amit1502/greenlight/internal/data"
	"github.com/amit1502/greenlight/internal/validation"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		app.logger.Println("Invalid movie input:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// add some validation
	v := validation.Validator{
		Error: make(map[string]string),
	}

	v.Check(len(input.Title) < 10, "title", "title len should be less then 10")
	v.Check(input.Runtime < 180, "runtime", "movie length should be less than 3 hrs")

	if len(v.Error) > 0 {
		app.logger.Println("validation failed")
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
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
