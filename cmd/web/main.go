package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/GnvSaikiran/snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	infoLog, errorLog *log.Logger
	staticDir         *string
	snippets          *models.SnippetModel
	templateCache     map[string]*template.Template
}

func main() {
	// Using command-line flags for configuration
	addr := flag.String("addr", ":4000", "HTTP network address")
	staticDir := flag.String("staticdir", "./ui/static", "Static files directory path")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Creating custom loggers for leveled logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Creating a sql.DB object
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	// Initializing template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	// Creating an application struct for dependency injection
	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		staticDir:     staticDir,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	// Configuring the server
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.Handler(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

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
