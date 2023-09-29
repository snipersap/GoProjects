package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/config"
	handler "github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/handlers"
)

// routes registers the route handlers
func routes(a *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))

	return mux
}
