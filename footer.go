package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type footer struct {
	app.Compo
}

func (h *footer) Render() app.UI {
	return app.Footer().Class("footer").
		Body(
			app.Div().Class("content has-text-centered").
				Body(
					app.P().Body(
						app.Text("Made with "),
						app.A().Href("https://go-app.dev").Text("go-app"),
						app.Text(" by Jairo"),
					),
				),
		)
}
