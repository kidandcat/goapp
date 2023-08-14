package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routes() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/tenis", tenis)
	r.HandleFunc("/padel", padel)
	r.HandleFunc("/reserve", reserve)
	http.Handle("/", r)
}
