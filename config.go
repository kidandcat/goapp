package main

import "net/http"

func routes() {
	http.HandleFunc("/", home)
	http.HandleFunc("/tenis", tenis)
	http.HandleFunc("/padel", padel)
	http.HandleFunc("/reserve", reserve)
}
