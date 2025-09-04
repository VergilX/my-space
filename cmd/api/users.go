package main

import (
	"net/http"
)

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	// get data from form and validate

	// use internals module auth to add user to db
	// while performing required functions

	// report errors as required

}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	// get data from POST request and validate

	// login using internals

	// report errors

}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	// get data from POST request and validate

	// logout using internals

	// report errors

}
