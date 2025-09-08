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
	router.HandleFunc("POST /v1/login", app.loginUser)

	// protected routes
	router.Handle("POST /v1/logout", baseChain.ThenFunc(app.logoutUser))

	router.Handle("GET /v1/clipboard", baseChain.ThenFunc(app.getClipContent))

	return app.recoverPanic(app.requestLog(router))
}
