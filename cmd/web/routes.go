package main

import (
	"net/http"

	"github.com/Goodmorningpeople/booking_go/pkg/config"
	"github.com/Goodmorningpeople/booking_go/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes func contains all the routes, which are served to an http server in main.go (using chi router)
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	mux.Get("/", handlers.Repo.Home)
	return mux
}
