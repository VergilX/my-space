package main

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/VergilX/my-space/internal/dblayer"
	"github.com/VergilX/my-space/internal/request"
	"github.com/VergilX/my-space/internal/response"
	"github.com/VergilX/my-space/internal/validator"
)

// time field should follow RFC 3339
// db currently uses UTC time
func (app *application) addPaste(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(userIDKey).(int64)
	if !ok {
		app.serverError(w, r, errors.New("type assertion error: user-id"))
		return
	}

	var input struct {
		Text   string
		Expiry time.Time // should follow RFC 3339
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	v.Check(input.Text != "", "text", "should not be empty")
	v.Check(time.Until(input.Expiry).Seconds() > 0, "expiry", "should be after current time")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v)
		return
	}

	err = app.querier.CreatePaste(app.ctx, dblayer.CreatePasteParams{
		Userid:  userId,
		Text:    input.Text,
		Expires: input.Expiry.Format(SQLITE_FORMAT_STRING),
	})
	if err != nil {
		app.serverError(w, r, err)
	}

	err = response.JSON(w, http.StatusOK, envelope{"message": "paste added!"})
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getAllPaste(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(userIDKey).(int64)
	if !ok {
		app.serverError(w, r, errors.New("type assertion error: user-id"))
		return
	}

	pastes, err := app.querier.GetAllPastes(app.ctx, userId)
	if err != nil {
		app.serverError(w, r, err)
	}

	err = response.JSON(w, http.StatusOK, envelope{"pastes": pastes})
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) editPaste(w http.ResponseWriter, r *http.Request) {
	var input struct {
		PasteID int64
		Text    string
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	v := validator.New()

	v.Check(input.PasteID > 0, "pasteid", "should be valid")
	v.Check(input.Text != "", "text", "should not be empty")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v)
		return
	}

	_, err = app.querier.UpdatePaste(app.ctx, dblayer.UpdatePasteParams{
		Text: input.Text,
		ID:   input.PasteID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			app.logError(r, err)
			err = response.JSON(w, http.StatusBadRequest, envelope{"error": "invalid pasteid"})
			if err != nil {
				app.serverError(w, r, err)
				return
			}
		}

		app.serverError(w, r, err)
	}

	err = response.JSON(w, http.StatusOK, envelope{"message": "paste edited!"})
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) deletePaste(w http.ResponseWriter, r *http.Request) {
	var input struct {
		PasteID int64
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	v := validator.New()

	v.Check(input.PasteID > 0, "pasteid", "should be valid")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v)
		return
	}

	_, err = app.querier.DeletePaste(app.ctx, input.PasteID)
	if err != nil {
		if err == sql.ErrNoRows {
			// means no item deleted, invalid id
			app.logError(r, err)
			err = response.JSON(w, http.StatusNotFound, envelope{"error": "invalid pasteid"})
			if err != nil {
				app.serverError(w, r, err)
				return
			}

			return
		}

		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, envelope{"message": "paste deleted!"})
	if err != nil {
		app.serverError(w, r, err)
	}
}
