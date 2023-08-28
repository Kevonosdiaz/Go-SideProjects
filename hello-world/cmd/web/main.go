package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kevonosdiaz/Go-SideProjects/hello-world/pkg/config"
	"github.com/Kevonosdiaz/Go-SideProjects/hello-world/pkg/handlers"
	"github.com/Kevonosdiaz/Go-SideProjects/hello-world/pkg/render"
)

const PORT_NO = ":8080"

// main is the main app fn
func main() {
	var app config.AppConfig
	cache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not create template cache:", err)
	}

	app.TemplateCache = cache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PORT_NO))

	// Start the web server, with port being 8080
	srv := *&http.Server{
		Addr:    PORT_NO,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
