package main

import (
	"fmt"
	"net/http"
)

const PORT_NO = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 = %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 = %d", sum))
}

// addValues adds two integers and returns the sum (private)
func addValues(x, y int) int {
	return x + y
}

func

// main is the main app fn
func main() {
	// Handler function using an anonymous fn within it
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PORT_NO))
	// Start the web server, with port being 8080
	_ = http.ListenAndServe(PORT_NO, nil)
}
