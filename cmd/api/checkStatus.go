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
		app.jsonParseError(w, r, err)
	}
}
