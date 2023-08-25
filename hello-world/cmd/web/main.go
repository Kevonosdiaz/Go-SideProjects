package main

import (
	"fmt"
	"net/http"

	"github.com/Kevonosdiaz/Go-SideProjects/hello-world/pkg/handlers"
)

const PORT_NO = ":8080"

// main is the main app fn
func main() {
	// Handler function using an anonymous fn within it
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PORT_NO))
	// Start the web server, with port being 8080
	_ = http.ListenAndServe(PORT_NO, nil)

}
