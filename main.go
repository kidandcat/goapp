package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"reservas/controller"
	"reservas/model"

	"github.com/gorilla/sessions"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	PORT    = os.Getenv("PORT")
	DB_NAME = os.Getenv("DB_NAME")
)

//go:embed views/*.html static/*
var tmplFS embed.FS

func main() {
	if DB_NAME == "" {
		DB_NAME = "reservas.db"
	}

	// initialize DB
	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(
		&model.Reserva{},
	)
	if err != nil {
		panic("failed to migrate database")
	}

	// initialize static files
	http.Handle("/static/", cacheMiddleware(http.FileServer(http.FS(tmplFS))))

	// initialize routes
	routes(&controller.Controller{
		DB:        db,
		Templates: template.Must(template.New("").ParseFS(tmplFS, "views/*.html")),
		Store:     sessions.NewCookieStore([]byte("key")),
	})

	if PORT == "" {
		PORT = "3000"
	}
	fmt.Println("Listening on 127.0.0.1:" + PORT)
	fatal(http.ListenAndServe("127.0.0.1:"+PORT, nil))
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
