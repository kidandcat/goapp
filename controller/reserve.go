package controller

import (
	"log"
	"net/http"
	"reservas/model"
	"time"
)

func (c *Controller) Reserve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		if err := c.Templates.ExecuteTemplate(w, "reserve.html", nil); err != nil {
			log.Fatal(err)
		}
		return
	}

	// read form data
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	// create reservation
	re := &model.Reserva{
		Name:         r.FormValue("name"),
		Address:      r.FormValue("address"),
		Date:         time.Now(),
		Duration:     time.Minute * 90,
		Participants: "jairo",
	}

	c.DB.Save(re)

	// redirect to home
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
