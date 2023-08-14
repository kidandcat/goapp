package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:embed views/*.html static/*
var tmplFS embed.FS
var templates = template.Must(template.New("").ParseFS(tmplFS, "views/*.html"))

var db *gorm.DB

var store = sessions.NewCookieStore([]byte("key"))

func main() {
	var err error

	// initialize DB
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Reserva{})
	if err != nil {
		panic("failed to migrate database")
	}

	// initialize static files
	http.Handle("/static/", cacheMiddleware(http.FileServer(http.FS(tmplFS))))

	routes()

	fmt.Println("Listening on 127.0.0.1:8000")
	fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}

func cacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Type of files to cache
		if r.URL.Path[len(r.URL.Path)-4:] == ".jpg" ||
			r.URL.Path[len(r.URL.Path)-5:] == ".jpeg" ||
			r.URL.Path[len(r.URL.Path)-4:] == ".png" ||
			r.URL.Path[len(r.URL.Path)-4:] == ".gif" {
			w.Header().Set("Cache-Control", "max-age=2630000") // 1 month
		}
		next.ServeHTTP(w, r)
	})
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
