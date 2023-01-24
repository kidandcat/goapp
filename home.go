package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type home struct {
	app.Compo
	name string
}

func (h *home) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Hello "+h.name),
		app.P().Body(
			app.Input().
				Type("text").
				Value(h.name).
				Class("input").
				Placeholder("What is your name?").
				AutoFocus(true).
				OnInput(h.ValueTo(&h.name)),
		),
		&footer{},
	)
}
