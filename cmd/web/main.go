package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	/*
		After applying http.StripPrefix("/static", fileServer), the prefix "/static" is stripped from the URL path. The resulting modified URL path becomes "css/styles.css".
		Now, the modified URL path "css/styles.css" is passed to the fileServer, which serves files from the "./ui/static/" directory.
		It correctly resolves the file path as "./ui/static/css/styles.css" and serves the corresponding file.
	*/
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
