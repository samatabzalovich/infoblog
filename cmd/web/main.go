package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/infoBlog/about", about)
	mux.HandleFunc("/infoBlog/samplePost", samplePost)
	mux.HandleFunc("/infoBlog/contact", contact)
	mux.HandleFunc("/infoBlog/login", login)
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
