package main

import (
	"net/http"
	"shared"
	"text/template"
)

type (
	Response struct {
		Title string
		User *User
	}

	User = shared.User
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	shared.DeleteCookie(w, "username")

	response := Response{
		Title: "Главная",
		User: nil,
	}

	tmpl, err := template.ParseFiles("../templates/base.html", "../templates/home.html")

	if err != nil {
		shared.ErrorResponse(w, err)
		return
	}

	tmpl.ExecuteTemplate(w, "home.html", response)
}