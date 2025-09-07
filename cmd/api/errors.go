package main

import (
	"net/http"
)

func (app *application) cookieNotFoundError(w http.ResponseWriter, r *http.Request, err error, msg string) {
	app.logger.Info("Required cookie is not found", "errorMsg", err.Error())
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func (app *application) jsonParseError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Info(err.Error(), "msg", "could not parse input JSON")
	http.Error(w, "Internal JSON conversion error", http.StatusInternalServerError)
}
