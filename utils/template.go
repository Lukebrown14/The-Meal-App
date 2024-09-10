package utils

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templates *template.Template

func init() {
	// Parse all template files in the templates directory
	templatePath := filepath.Join("templates", "*.html")
	templates = template.Must(template.ParseGlob(templatePath))
}

func RenderHome(w http.ResponseWriter, data interface{}) error {
	return templates.ExecuteTemplate(w, "index.html", data)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	return templates.ExecuteTemplate(w, tmpl, data)
}
