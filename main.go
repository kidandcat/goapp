package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed views/*.html static/*
var tmplFS embed.FS

var templates = template.Must(template.New("").ParseFS(tmplFS, "views/*.html"))

func main() {
	http.Handle("/static/", http.FileServer(http.FS(tmplFS)))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "home.html", nil); err != nil {
			log.Fatal(err)
		}
	})

	if err := http.ListenAndServe("127.0.0.1:8000", nil); err != nil {
		log.Fatal(err)
	}
}
