package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/config"
	handler "github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/handlers"
)

// routes creates a new router, registers the middlewares and loads the session
func routes(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(LogEachApiHit)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux = endPoints(mux)
	return mux
}

// patRoutes registers the route handlers
func patRoutes(a *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))
	return mux
}

// chiRoutes registers the route handlers
func chiRoutes(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Use middlewares
	mux.Use(middleware.Recoverer) //Use chi middleware
	mux.Use(LogEachApiHit)        // Use custom middle ware

	mux = endPoints(mux)
	return mux
}

// chiRoutesWithNoSurf registers routes with middleware as NoSurf
func chiRoutesWithNoSurf(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Use middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux = endPoints(mux)
	return mux

}

// endPoints returns routes of the app
func endPoints(mux *chi.Mux) *chi.Mux {
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	return mux
}
