package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// address flag
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	//logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	//logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelInfo,
		ReplaceAttr: nil,
	}))

	// ServeMux
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	//const port = "8080"
	//log.Printf("starting server on port%s\n", *addr)
	logger.Info("starting server", slog.Any("addr", ":4000"))

	err := http.ListenAndServe(*addr, mux)
	//log.Fatal(err)
	logger.Error(err.Error())
	os.Exit(1)
}
