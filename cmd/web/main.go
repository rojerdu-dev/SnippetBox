package main

import (
	"log"
	"net/http"
)

func main() {
	// ServeMux
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	const port = "8000"
	log.Printf("starting server on port: %s", port)

	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
