package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
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
	dsn := flag.String(
		"dsn",
		"root:mysql@tcp(127.0.0.1:3306)/snippetbox?parseTime=true",
		"MySQL data source name")
	flag.Parse()

	// structured logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// connection pool into separate openDB() function below
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger: logger,
	}

	// start server
	logger.Info("starting server", "addr", *addr)

	// call new app.routes() method to get servemux containing our routes
	// pass to http.ListenAndServe
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

// openDB() function wraps sql.open() and
// returns a sql.DB connection pool for a given DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
