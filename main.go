package main

import (
	"log"
	"net/http"
)

func main() {
	const port = ":8000"

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Printf("starting server on port%s ...", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home page for project"))
}
