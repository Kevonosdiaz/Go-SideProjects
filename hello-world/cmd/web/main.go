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

	// Handler function using an anonymous fn within it
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PORT_NO))
	// Start the web server, with port being 8080
	_ = http.ListenAndServe(PORT_NO, nil)

}
