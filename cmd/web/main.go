package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	/*
		After applying http.StripPrefix("/static", fileServer), the prefix "/static" is stripped from the URL path. The resulting modified URL path becomes "css/styles.css".
		Now, the modified URL path "css/styles.css" is passed to the fileServer, which serves files from the "./ui/static/" directory.
		It correctly resolves the file path as "./ui/static/css/styles.css" and serves the corresponding file.
	*/
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	/* mux.HandleFunc("/", http.HandlerFunc(home)) */
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
