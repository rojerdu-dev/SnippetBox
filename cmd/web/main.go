package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// addr flag
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	// structured logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// init new instance of application struct containing the
	// dependencies (for now just structured logging)
	app := &application{
		logger: logger,
	}

	// mux
	mux := http.NewServeMux()

	// file serve & handlers
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// start server
	//log.Printf("starting server on %s\n", *addr)
	logger.Info("starting server", "address", *addr)

	// handle errors
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
