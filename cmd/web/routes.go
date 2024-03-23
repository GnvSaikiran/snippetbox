package main

import "net/http"

func (app *application) Handler() *http.ServeMux {
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	fileServer := http.FileServer(http.Dir(*app.staticDir))

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
