package controller

import (
	"github.com/gorilla/sessions"
	"html/template"
)

var (
	homeController home
	templates      map[string]*template.Template
	sessionName    string
	flashName      string
	store          *sessions.CookieStore
	pageLimit      int
)

func init() {
	templates = PopulateTemplates()
	store = sessions.NewCookieStore([]byte(".....___++++()RTYUJNBVFTYUIOigfrtyuj"))
	sessionName = "go-mega"
	flashName = "go-flash"
	pageLimit = 5

}

// Startup func
func Startup() {
	homeController.registerRoutes()
}
