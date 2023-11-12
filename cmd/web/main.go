package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/opensourcerror/go_webserv_04_mysql/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

// go run	 ./cmd/web -addr=":40000"

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	// :4000 will be default value if no flag
	addr := flag.String("addr", ":4000", "HTTP net addr")

	// Define a new command-line flag for the MySQL DSN string.
	dsn := flag.String("dsn", "web:notarealpass@/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse() //flags must be defined before this line

	// nil when u want to use the default settings
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// To keep the main() function tidy I've put the code for creating a connection
	// pool into the separate openDB() function below. We pass openDB() the DSN
	// from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}
	//folosesti asa:
	// app.snippets.Insert(...)

	logger.Info("srv up on ", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
