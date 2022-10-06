package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.go.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.go.html")
}

func renderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("err parsing template:", err)
		return
	}
}
