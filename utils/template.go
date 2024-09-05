package utils

import (
	"html/template"
	"net/http"
)

func RenderHome(w http.ResponseWriter, data interface{}) error {
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(w, tmpl, data)
}
