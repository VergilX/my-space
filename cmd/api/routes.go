package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /v1/status", app.checkStatus)

	// chain middleware
	baseChain := alice.New(app.protected)

	// user stuff
	router.HandleFunc("POST /v1/register", app.registerUser)
	router.HandleFunc("POST /v1/login", app.loginUser)
	// router.HandleFunc("POST /v1/logout", app.logoutUser)

	// protected routes
	router.Handle("POST /v1/logout", baseChain.ThenFunc(app.logoutUser))

	router.Handle("GET /v1/clipboard", baseChain.ThenFunc(app.getClipContent))
	router.Handle("POST /v1/setclip", baseChain.ThenFunc(app.setClipText))

	router.Handle("GET /v1/viewallpaste", baseChain.ThenFunc(app.getAllPaste))
	router.Handle("POST /v1/newpaste", baseChain.ThenFunc(app.addPaste))
	router.Handle("POST /v1/editpaste", baseChain.ThenFunc(app.editPaste))
	router.Handle("POST /v1/deletepaste", baseChain.ThenFunc(app.deletePaste))

	return app.recoverPanic(app.requestLog(router))
}
