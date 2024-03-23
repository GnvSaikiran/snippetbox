package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog, errorLog *log.Logger
	staticDir         *string
}

func main() {
	// Using command-line flags for configuration
	addr := flag.String("addr", ":4000", "HTTP network address")
	staticDir := flag.String("staticdir", "./ui/static", "Static files directory path")
	flag.Parse()

	// Creating custom loggers for leveled logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Creating an application struct for dependency injection
	app := application{
		infoLog:   infoLog,
		errorLog:  errorLog,
		staticDir: staticDir,
	}

	// Configuring the server
	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.Handler(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
