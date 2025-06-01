package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cgi"
	"text/template"
)

func main() {
	cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("../templates/base.html", "../template/home.html")

		if err != nil {
			fmt.Fprintf(w, "Template error")
			log.Print(err)
			return
		}

		tmpl.ExecuteTemplate(w, "home.html", nil)
	}))
}