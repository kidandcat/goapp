package controller

import (
	"html/template"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

type Controller struct {
	DB        *gorm.DB
	Templates *template.Template
	Store     *sessions.CookieStore
}
