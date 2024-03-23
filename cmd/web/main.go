package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Using command-line flags for configuration
	addr := flag.String("addr", ":4000", "HTTP network address")
	staticDir := flag.String("staticdir", "./ui/static", "Static files directory path")
	flag.Parse()

	// Creating custom loggers for leveled loggin
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	fileServer := http.FileServer(http.Dir(*staticDir))

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Configuring the server
	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
