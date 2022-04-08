package main

import (
	"net/http"

	"github.com/SzymonSkursrki/goUdemyBookings/pkg/config"
	"github.com/SzymonSkursrki/goUdemyBookings/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//example using "github.com/bmizerany/pat"
// func routes(app *config.AppConfig) http.Handler {
// 	//mux multiplexer
// 	mux := pat.New()
// 	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
// 	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

// 	return mux
// }

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter() //multiplexer
	//midleware
	mux.Use(middleware.Recoverer) //recover after panic
	mux.Use(NoSurf)               //CSRF protection
	mux.Use(SessioLoad)           //load session on each request

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
