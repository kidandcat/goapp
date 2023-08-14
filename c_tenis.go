package main

import (
	"log"
	"net/http"
)

func tenis(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "tenis.html", nil); err != nil {
		log.Fatal(err)
	}
}
