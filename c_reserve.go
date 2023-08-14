package main

import (
	"log"
	"net/http"
	"time"
)

func reserve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		if err := templates.ExecuteTemplate(w, "reserve.html", nil); err != nil {
			log.Fatal(err)
		}
		return
	}

	// read form data
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	// create reservation
	re := &Reserva{
		Name:         r.FormValue("name"),
		Address:      r.FormValue("address"),
		Date:         time.Now(),
		Duration:     time.Minute * 90,
		Participants: "jairo",
	}

	db.Save(re)

	// redirect to home
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
