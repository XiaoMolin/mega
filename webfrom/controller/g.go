package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home
	templates      map[string]*template.Template
)

func init() {
	templates = PopulateTemplates()
}

// Startup func
func Startup() {
	homeController.registerRoutes()
	_ = http.ListenAndServe(":8888", nil)
}
