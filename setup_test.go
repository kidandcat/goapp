package main

import (
	"os"
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/ysmood/got"
)

type G struct {
	got.G
	browser *rod.Browser
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Remove(DB_NAME)
	os.Exit(code)
}

var setup = func() func(t *testing.T) G {
	browser := rod.New().MustConnect()

	DB_NAME = "test.db"
	PORT = "3000"
	go main()

	return func(t *testing.T) G {
		t.Parallel()
		return G{got.New(t), browser}
	}
}()

func (g G) page(url string) *rod.Page {
	page := g.browser.MustIncognito().MustPage(url).Timeout(10 * time.Second)
	g.Cleanup(page.MustClose)
	return page
}
