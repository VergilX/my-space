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
	router.Handle("/v1/register", baseChain.ThenFunc(app.registerUser))
	router.Handle("/v1/login", baseChain.ThenFunc(app.loginUser))
	router.Handle("/v1/logout", baseChain.ThenFunc(app.logoutUser))

	return app.requestLog(router)
}
