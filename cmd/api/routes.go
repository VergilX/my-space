package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /v1/status", app.checkStatus)

	// chain middleware
	baseChain := alice.New(app.userAuthCheck)

	// user stuff
	router.HandleFunc("POST /v1/register", app.registerUser)
	router.Handle("POST /v1/login", baseChain.ThenFunc(app.loginUser))
	router.Handle("POST /v1/logout", baseChain.ThenFunc(app.logoutUser))

	// protected routes

	return app.recoverPanic(app.requestLog(router))
}
