package main

import (
	"testing"
	"time"
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

func TestReserve(t *testing.T) {
	g := setup(t)
	p := g.page("http://localhost:3000/reserve")

	p.MustElement(`input[name="name"]`).MustInput("Test")
	p.MustElement(`input[name="address"]`).MustInput("Test address")
	p.MustElement("input[type='submit']").MustClick()
	p.MustWaitStable()
	g.Eq(p.MustInfo().URL, "http://localhost:3000/")

	p.MustElementR("h3", "Padel").MustParent().MustClick()
	p.MustWaitStable()

	d := time.Now().Format("02/01/2006")
	g.Eq(p.MustElement("div.reserve").MustText(), "Test Test address "+d)
}
