package main

import (
	"net/http"

	"github.com/Kevonosdiaz/Go-SideProjects/hello-world/pkg/config"
	"github.com/Kevonosdiaz/Go-SideProjects/hello-world/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// routes uses the chi external package to process routing
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Use middleware from chi pkg
	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}

// // routes uses the pat external package to process routing
// func routes(app *config.AppConfig) http.Handler {
// 	mux := pat.New()
// 	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
// 	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
// 	return mux
// }
