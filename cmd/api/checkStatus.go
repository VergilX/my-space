package main

import (
	"net/http"

	"github.com/VergilX/my-space/internal/response"
)

func (app *application) checkStatus(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}
}
