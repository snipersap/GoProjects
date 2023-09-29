package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/config"
	handler "github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/handlers"
)

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

	mux = routes(mux)
	return mux
}

// chiRoutesWithNoSurf registers routes with middleware as NoSurf
func chiRoutesWithNoSurf(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Use middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux = routes(mux)
	return mux

}

// registeredRoutes returns routes of the app
func routes(mux *chi.Mux) *chi.Mux {
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	return mux
}
