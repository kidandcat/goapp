package controller

import (
	"log"
	"net/http"
)

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	session, _ := c.Store.Get(r, "session-name")
	session.Values["foo"] = "bar"
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := c.Templates.ExecuteTemplate(w, "home.html", nil); err != nil {
		log.Fatal(err)
	}
}
