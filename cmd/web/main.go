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
	// set address at runtime
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	// structured logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	// start server
	logger.Info("starting server", "addr", *addr)

	// call new app.routes() method to get servemux containing our routes
	// pass to http.ListenAndServe
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
