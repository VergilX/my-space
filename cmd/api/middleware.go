package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/tomasen/realip"
)

// idk why you need this but ok
type StringContextKey string

var userIDKey StringContextKey = "user-id"

func (app *application) requestLog(handler http.Handler) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		// information logged: ip, port, endpoint, method, protocol
		var (
			ip       = realip.FromRequest(r)
			url      = r.URL
			method   = r.Method
			protocol = r.Proto
		)

		app.logger.Info("log", "ip", ip, "url", url.RequestURI(), "method", method, "protocol", protocol)

		handler.ServeHTTP(w, r)
	}))
}

func (app *application) protected(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get csrf_token from header
		csrfToken := r.Header.Get("X-CSRF-Token")
		if csrfToken == "" {
			app.failedAuthentication(w, r, errors.New("CSRF token missing"))
			return
		}

		// get session token cookie
		cookie, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				app.failedAuthentication(w, r, err)
			} else {
				app.serverError(w, r, err)
			}

			return
		}

		id, err := app.querier.GetUserIDFromSessionToken(app.ctx, cookie.Value)
		if err == sql.ErrNoRows {
			app.failedAuthentication(w, r, err)
			return
		} else if err != nil {
			app.serverError(w, r, err)
			return
		}

		// verify CSRF
		exists, err := app.querier.VerifyCSRFToken(app.ctx, csrfToken)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		if exists == 0 {
			app.failedAuthentication(w, r, errors.New("invalid csrf token"))
			return
		}

		// session and CSRF are verified, store id in context
		ctx := context.WithValue(r.Context(), userIDKey, id)
		r = r.WithContext(ctx)

		handler.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")

				app.logger.Error(fmt.Errorf("%s", err).Error(), "msg", "server panic")
			}
		}()

		handler.ServeHTTP(w, r)
	})
}
