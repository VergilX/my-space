package main

import (
	"net/http"
)

func (app *application) checkStatus(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeToJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Server could not convert the data into JSON", http.StatusInternalServerError)
	}
}
