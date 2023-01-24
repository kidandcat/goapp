package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &home{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Reservas",
		Description: "Reserva de mesas",
		Styles: []string{
			"/web/global.css",
			"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css",
		},
	})

	if err := http.ListenAndServe("127.0.0.1:8000", nil); err != nil {
		log.Fatal(err)
	}
}
