package main

import (
	"testing"
)

func TestNavigate(t *testing.T) {
	g := setup(t)
	p := g.page("http://localhost:3000")

	p.MustElementR("h3", "Padel").MustParent().MustClick()
	p.MustWaitStable()
	g.Eq(p.MustInfo().URL, "http://localhost:3000/padel")

	p.MustElementR("h3", "Tenis").MustParent().MustClick()
	p.MustWaitStable()
	g.Eq(p.MustInfo().URL, "http://localhost:3000/tenis")
}
