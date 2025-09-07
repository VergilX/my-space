package main

import (
	"net/http"
	"slices"

	"github.com/tomasen/realip"
)

func (app *application) requestLog(handler http.Handler) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		// information logged: ip, port, endpoint, method, protocol
		var (
			ip       = realip.FromRequest(r)
			url      = r.URL
			method   = r.Method
			protocol = r.Proto
		)

		app.logger.Info("request log:", "ip", ip, "url", url.String(), "method", method, "protocol", protocol)

		handler.ServeHTTP(w, r)
	}))
}

func (app *application) userAuthCheck(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authRequiredURL := []string{
			"status",
		}

		app.logger.Info(r.URL.Path)

		// check if url is in authRequiredURLS
		if !slices.Contains(authRequiredURL, "helo") {
			app.logger.Info("yaboi")
		}

		// if not, do nothing

		// else check session token authentication

		// if not successful, return to home or register

		var _, err = r.Cookie("session_token")
		if err == http.ErrNoCookie {
			app.cookieNotFoundError(w, r, err, "session_token")
		}
		handler.ServeHTTP(w, r)
	})
}
