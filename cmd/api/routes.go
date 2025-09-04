package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/status", app.checkStatus)

	// user stuff
	router.HandlerFunc(http.MethodPost, "/v1/register", app.registerUser)
	router.HandlerFunc(http.MethodPost, "/v1/login", app.loginUser)
	router.HandlerFunc(http.MethodPost, "/v1/logout", app.logoutUser)

	return router
}
