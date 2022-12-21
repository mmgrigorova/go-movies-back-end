package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// every var app (the receiver) of type application will have access to that function
// no params, returns http handler
func (app *application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	//middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Get("/", app.Home)

	mux.Get("/movies", app.AllMovies)

	return mux
}
