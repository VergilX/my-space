package main

import (
	"fmt"
	"net/http"

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

		app.logger.Info("log", "ip", ip, "url", url.RequestURI(), "method", method, "protocol", protocol)

		handler.ServeHTTP(w, r)
	}))
}

func (app *application) userAuthCheck(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// first value to be used after DB implementation
		_, err := r.Cookie("session_token")

		if err != nil {
			if err == http.ErrNoCookie {
				app.badRequestResponse(w, r, err) // replace with auth error
			}
		}

		// check database for session_token value and authenticate
		// to be implmented after db implementation

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
