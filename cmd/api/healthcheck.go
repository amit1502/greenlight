package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":  "available",
		"env":     app.config.env,
		"version": version,
	}

	err := app.writeJson(w, http.StatusOK, data, nil)

	if err != nil {
		app.logger.Println("Error while parsing data:", err)
		http.Error(w, "Error while parsing data", http.StatusInternalServerError)
		return
	}
}
