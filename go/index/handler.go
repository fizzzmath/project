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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/base.html", "../templates/home.html")

	if err != nil {
		shared.ErrorResponse(w, err)
		return
	}

	response := Response{
		Title: "Главная",
		User: nil,
	}
	
	username, err := shared.GetCookie(r, "username")

	if err != nil {
		tmpl.ExecuteTemplate(w, "home.html", response)
		return
	}
	
	response.User = &User{
		Username: username,
	}

	tmpl.ExecuteTemplate(w, "home.html", response)
}