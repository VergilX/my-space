package main

import (
	"net/http"
)

func (app *application) cookieNotFoundError(w http.ResponseWriter, r *http.Request, err error, msg string) {
	app.logger.Error(err.Error(), "desc", msg, "url", r.URL)
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func (app *application) jsonParseError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error(err.Error(), "desc", "could not parse input JSON", "url", r.URL)
	http.Error(w, "Internal JSON conversion error", http.StatusInternalServerError)
}

func (app *application) formValueError(w http.ResponseWriter, r *http.Request, msg string) {
	app.logger.Error("Bad Request", "msg", msg, "url", r.URL)
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func (app *application) authenticationError(w http.ResponseWriter, r *http.Request, err error, msg string) {
	app.logger.Error(err.Error(), "msg", msg, "url", r.URL)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
