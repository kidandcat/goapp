package controller

import (
	"log"
	"net/http"
)

func (c *Controller) Tenis(w http.ResponseWriter, r *http.Request) {
	if err := c.Templates.ExecuteTemplate(w, "tenis.html", nil); err != nil {
		log.Fatal(err)
	}
}
