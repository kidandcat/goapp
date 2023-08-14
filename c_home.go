package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "home.html", nil); err != nil {
		log.Fatal(err)
	}
}
