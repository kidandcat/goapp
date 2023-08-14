package main

import (
	"log"
	"net/http"
)

type PadelPageData struct {
	Reservas []Reserva
}

func padel(w http.ResponseWriter, r *http.Request) {
	var data PadelPageData
	db.Find(&data.Reservas)
	if err := templates.ExecuteTemplate(w, "padel.html", data); err != nil {
		log.Fatal(err)
	}
}
