package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// HomeHandler is a handler for the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

// AboutHandler is a handler for the about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	sum := addValue(1, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("About Page, sum: %d", sum))
}

// addValue adds two ints a and b
func addValue(a, b int) int {
	return a + b
}

// main is the entry point for the program
func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)

	fmt.Println(fmt.Sprintf("Server is listening on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
