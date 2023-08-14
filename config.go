package main

import (
	"net/http"
	"reservas/controller"

	"github.com/gorilla/mux"
)

func routes(c *controller.Controller) {
	r := mux.NewRouter()
	r.HandleFunc("/", c.Home)
	r.HandleFunc("/tenis", c.Tenis)
	r.HandleFunc("/padel", c.Padel)
	r.HandleFunc("/reserve", c.Reserve)
	http.Handle("/", r)
}
