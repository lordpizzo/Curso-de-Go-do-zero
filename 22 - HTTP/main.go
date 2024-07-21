package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{"João", "XXXXXXXXXXXXX"}
	templates.ExecuteTemplate(w, "home.html", u)
}
func main() {
	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("/home", home)
	log.Fatal(http.ListenAndServe(":50001", nil))
}
