package main

import (
	"log"
	"net/http"
)

func main() {
	const port = ":8000"

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("starting server on port%s ...", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "GET" {
		_, err := w.Write([]byte("home page for project"))
		if err != nil {
			log.Printf("home http.ResponseWriter failed to write: %s", err)
			return
		}
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := w.Write([]byte("Display a specific snippet ..."))
		if err != nil {
			log.Printf("snippetCreate http.ResponseWriter failed to write: %s", err)
			return
		}
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := w.Write([]byte("Create a new snippet ..."))
		if err != nil {
			log.Printf("snippetView http.ResponseWriter failed to write: %s", err)
			return
		}
	}
}
