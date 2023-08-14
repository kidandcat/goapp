package main

import (
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/ysmood/got"
)

// test context.
type G struct {
	got.G

	browser *rod.Browser
}

// setup for tests.
var setup = func() func(t *testing.T) G {
	browser := rod.New().MustConnect()

	DB_NAME = "test.db"
	PORT = "3000"
	go main()

	return func(t *testing.T) G {
		t.Parallel() // run each test concurrently

		return G{got.New(t), browser}
	}
}()

// a helper function to create an incognito page.
func (g G) page(url string) *rod.Page {
	page := g.browser.MustIncognito().MustPage(url).Timeout(10 * time.Second)
	g.Cleanup(page.MustClose)
	return page
}
