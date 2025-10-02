package main

import (
	"net/http"
	"time"

	"github.com/VergilX/my-space/internal/auth"
)

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	// get data from form and validate
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		app.formValueError(w, r, "username or password is missing")
	}

	// if username already exists give error
	exists, err := app.db.IfUserExists(username)
	if err != nil {
		app.logger.Error("Error accessing user database")
	}
	if exists {
		app.logger.Error("User already exists")
	}

}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	// get data from POST request and validate
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		app.formValueError(w, r, "username or password is missing")
	}

	// login using internals
	session_token, csrf_token, err := auth.Login(username, password)
	if err != nil {
		app.authenticationError(w, r, err, "login error")
	}

	// set session and csrf cookie for response
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session_token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	// set session and csrf cookie for response
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrf_token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	// get data from POST request and validate
	username := r.FormValue(("username"))

	if username == "" {
		app.formValueError(w, r, "username missing")
	}

	// logout using internals
	err := auth.Logout(username)
	if err != nil {
		app.authenticationError(w, r, err, "logout error")
	}

	// clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})
}
