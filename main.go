package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "HOST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// When the user sends a request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // declare and define the variable fileServer. This line tells GoLang to look at the static directory tree, and to look at the ./static file
	http.Handle("/", fileServer)                        // this will direct the function to the fileserver location at /
	http.HandleFunc("/form", formHandler)               // will show the form in html
	http.HandleFunc("/hello", helloHandler)             // prints 'hello' to the screen

	fmt.Printf("Starting server at Port localhost:3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err) // error handling for GoLang
	}
}
