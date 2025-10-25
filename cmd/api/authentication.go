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

var SQLITE_FORMAT_STRING = "2006-01-02T15:04:05.000Z"
var TOKEN_SIZE int = 128

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	// json request
	var input struct {
		Username string
		Password string
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	// validate user
	v.Check(input.Username != "", "username", "must not be empty")
	v.Check(input.Password != "", "password", "must not be empty")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v)
	} else {
		// if username exists give error
		exists, err := app.querier.DoesUserExist(app.ctx, input.Username)
		if err != nil {
			app.logger.Warn("Could not verify existence of user")
			app.serverError(w, r, err)
		}

		if exists == 0 {
			// generate hashpassword
			input.Password, err = auth.HashPassword(input.Password)
			if err != nil {
				app.failedAuthentication(w, r, err)
				return
			}

			userId, err := app.querier.CreateUser(app.ctx, input)
			if err != nil {
				app.logError(r, err)
				response.JSON(w, http.StatusConflict, envelope{"error": "Could not create user"})
			} else {
				// create entries in session and csrf tables
				app.querier.CreateSessionToken(app.ctx, dblayer.CreateSessionTokenParams{
					Userid: userId,
				})
				app.querier.CreateCSRFToken(app.ctx, dblayer.CreateCSRFTokenParams{
					Userid: userId,
				})
				response.JSON(w, http.StatusOK, envelope{"message": "User registered successfully!"})
			}
		} else {
			err = response.JSON(w, http.StatusConflict, envelope{"error": "User already exists!"})
			if err != nil {
				app.serverError(w, r, err)
			}
		}
	}

}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	// json request
	var input struct {
		Username string
		Password string
	}

	err := request.DecodeJSONStrict(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	v := validator.New()

	v.Check(input.Username != "", "username", "must not be empty")
	v.Check(input.Password != "", "password", "must not be empty")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v)
	} else {
		// check if username already exists
		exists, err := app.querier.DoesUserExist(app.ctx, input.Username)
		if err != nil {
			app.logger.Warn("could not verify existance of user")
			app.serverError(w, r, err)
		}

		if exists == 0 {
			app.failedAuthentication(w, r, errors.New("unregistered user"))
		} else {
			// get password from db and validate
			user, err := app.querier.GetUser(app.ctx, input.Username)
			if err != nil {
				app.serverError(w, r, err)
			}

			if !auth.ValidatePasswordWithHash(input.Password, user.Password) {
				app.failedAuthentication(w, r, errors.New("invalid password"))
				return
			}

			// generate tokens and send in header
			session_token, err := auth.GenerateToken(TOKEN_SIZE)
			if err != nil {
				app.serverError(w, r, err)
			}
			csrf_token, err := auth.GenerateToken(TOKEN_SIZE)
			if err != nil {
				app.serverError(w, r, err)
			}

			expiry := time.Now().Add(24 * time.Hour)

			// set in db
			err = app.querier.RenewSessionToken(app.ctx, dblayer.RenewSessionTokenParams{
				Token:  session_token,
				Expiry: expiry.Format(SQLITE_FORMAT_STRING),
				Userid: user.ID,
			},
			)
			if err != nil {
				app.serverError(w, r, err)
				return
			}

			err = app.querier.RenewCSRFToken(app.ctx, dblayer.RenewCSRFTokenParams{
				Token:  csrf_token,
				Expiry: expiry.Format(SQLITE_FORMAT_STRING),
				Userid: user.ID,
			})
			if err != nil {
				app.serverError(w, r, err)
				return
			}

			// set session and csrf cookie for response
			http.SetCookie(w, &http.Cookie{
				Name:     "session_token",
				Value:    session_token,
				Expires:  expiry,
				HttpOnly: true,
			})

			err = response.JSON(w, http.StatusOK, envelope{"message": "User logged in!", "csrf_token": csrf_token})
			if err != nil {
				app.serverError(w, r, err)
			}
		}
	}
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	// this is a protected route
	// so http cookies should have the username and stuff
	// this is checked in middleware

	// clear session and session and csrf from db
	expiry := time.Now().Add(-time.Hour)

	userId, ok := r.Context().Value(userIDKey).(int64)
	if !ok {
		app.serverError(w, r, errors.New("type assertion error: user-id"))
		return
	}

	err := app.querier.RenewSessionToken(app.ctx, dblayer.RenewSessionTokenParams{
		Token:  "",
		Expiry: expiry.Format(SQLITE_FORMAT_STRING),
		Userid: userId,
	})
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = app.querier.RenewCSRFToken(app.ctx, dblayer.RenewCSRFTokenParams{
		Token:  "",
		Expiry: expiry.Format(SQLITE_FORMAT_STRING),
		Userid: userId,
	})
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// clear cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  expiry,
		HttpOnly: true,
	})

	response.JSON(w, http.StatusOK, envelope{"message": "User logged out!"})
}
