package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/VergilX/my-space/internal/response"
	"github.com/VergilX/my-space/internal/validator"
)

func (app *application) logError(r *http.Request, err error) {
	var (
		message = err.Error()
		method  = r.Method
		url     = r.URL.String()
	)

	requestAttrs := slog.Group("request", "method", method, "url", url)

	if app.config.traceEnabled {
		trace := string(debug.Stack())
		app.logger.Error(message, requestAttrs, "trace", trace)
	} else {
		app.logger.Error(message, requestAttrs)
	}
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := response.JSONWithHeaders(w, status, envelope{"error": message}, headers)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "It's civil war in here!"
	app.errorResponse(w, r, http.StatusInternalServerError, message, nil)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "your request stinks!"
	app.errorResponse(w, r, http.StatusBadRequest, message, nil)
}

func (app *application) methodNotAllowedError(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message, nil)
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, v *validator.Validator) {
	app.logError(r, errors.New("failed validation"))

	// you need to give information to the frontend so give validator
	err := response.JSON(w, http.StatusUnprocessableEntity, v)
	if err != nil {
		app.logError(r, err)
		app.serverError(w, r, err)
	}
}
