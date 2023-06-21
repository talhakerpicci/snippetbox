package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	/*
		After applying http.StripPrefix("/static", fileServer), the prefix "/static" is stripped from the URL path. The resulting modified URL path becomes "css/styles.css".
		Now, the modified URL path "css/styles.css" is passed to the fileServer, which serves files from the "./ui/static/" directory.
		It correctly resolves the file path as "./ui/static/css/styles.css" and serves the corresponding file.
	*/
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return app.logRequest(secureHeaders(mux))
}
