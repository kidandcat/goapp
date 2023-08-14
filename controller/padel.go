package controller

import (
	"log"
	"net/http"
	model "reservas/models"
)

type PadelPageData struct {
	Reservas []model.Reserva
}

func (c *Controller) Padel(w http.ResponseWriter, r *http.Request) {
	var data PadelPageData
	c.DB.Find(&data.Reservas)
	if err := c.Templates.ExecuteTemplate(w, "padel.html", data); err != nil {
		log.Fatal(err)
	}
}
