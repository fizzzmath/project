package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/cgi"
)

func main() {
	cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("../templates/auth/login.html", "../templates/base.html")

		if err != nil {
			fmt.Fprintf(w, "Template error")
			log.Print(err)
			return
		}

		tmpl.ExecuteTemplate(w, "login.html", nil)
	}))
}