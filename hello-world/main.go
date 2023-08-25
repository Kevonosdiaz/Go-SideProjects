package main

import (
	"fmt"
	"net/http"
)

const PORT_NO = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {

}

// main is the main app fn
func main() {
	// Handler function using an anonymous fn within it
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PORT_NO))
	// Start the web server, with port being 8080
	_ = http.ListenAndServe(PORT_NO, nil)
}
