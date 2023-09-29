package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// LogEachApiHit is a custom middleware that logs the path of the pages requested
func LogEachApiHit(next http.Handler) http.Handler {

	// Return a function which is of type http.Handler, therefore typecast
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Page was hit!:", r.URL)
		next.ServeHTTP(w, r)
	})
}

// NoSurf sets the base cookie for http requests, we will use it in forms
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Secure:   false,
	})
	return csrfHandler
}
