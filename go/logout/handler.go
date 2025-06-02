package main

import (
	"net/http"
	"shared"
	"text/template"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	shared.DeleteCookie(w, "username")

	tmpl, err := template.ParseFiles("../templates/base.html", "../templates/home.html")

	if err != nil {
		shared.ErrorResponse(w, err)
		return
	}

	tmpl.ExecuteTemplate(w, "home.html", nil)
}