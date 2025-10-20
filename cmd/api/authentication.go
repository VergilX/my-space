package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/VergilX/my-space/internal/auth"
	"github.com/VergilX/my-space/internal/dblayer"
	"github.com/VergilX/my-space/internal/request"
	"github.com/VergilX/my-space/internal/response"
	"github.com/VergilX/my-space/internal/validator"
)

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := dblayer.CreateUserParams{
		Username: input.Username,
		Password: input.Password,
	}

	v := validator.New()

	if !validateUser(v, &user) {
		app.failedValidationResponse(w, r, v)
	} else {

		err = app.querier.CreateUser(r.Context(), user)
		if err != nil {
			app.logError(r, err)
			response.JSON(w, http.StatusConflict, envelope{"error": "Could not create user"})
		} else {
			response.JSON(w, http.StatusOK, envelope{"message": "User registered successfully!"})
		}
	}

}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	// get data from POST request and validate
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		app.badRequestResponse(w, r, errors.New("username or password not given"))
	}

	// login using internals
	session_token, csrf_token, err := auth.Login(username, password)
	if err != nil {
		app.badRequestResponse(w, r, err) // replace with authentication error
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
		app.badRequestResponse(w, r, errors.New("username not found"))
	}

	// logout using internals
	err := auth.Logout(username)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("username not found"))
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
