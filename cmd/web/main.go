package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := application{
		logger: logger,
	}

	logger.Info("starting server", slog.String("addr", *addr))

	if err := http.ListenAndServe(*addr, app.routes()); err != nil {
		log.Fatal(err)
	}
}
