package main

import (
	"errors"
	"net/http"

	_ "github.com/VergilX/my-space/internal/db"
	"github.com/VergilX/my-space/internal/dblayer"
	"github.com/VergilX/my-space/internal/request"
	"github.com/VergilX/my-space/internal/response"
)

func (app *application) getClipContent(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(userIDKey).(int64)
	if !ok {
		app.serverError(w, r, errors.New("type assertion error: user-id"))
		return
	}

	clipText, err := app.querier.GetClipContent(app.ctx, userId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, envelope{"clip-text": clipText})
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) setClipText(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(userIDKey).(int64)
	if !ok {
		app.serverError(w, r, errors.New("type assertion error: user-id"))
		return
	}

	var input struct {
		ClipText string
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.querier.UpdateClip(app.ctx, dblayer.UpdateClipParams{
		Text:   input.ClipText,
		Userid: userId,
	})
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, envelope{"message": "clipboard text updated!"})
	if err != nil {
		app.serverError(w, r, err)
	}

}
