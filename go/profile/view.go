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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/profile/view.html", "../templates/base.html")

	if err != nil {
		shared.ErrorResponse(w, err)
		return
	}

	tmpl.ExecuteTemplate(w, "view.html", Response{
		Title: "Профиль",
		User: &User{
			ID: "1",
			Username: "Olegarhchik",
			FullName: "Атабаев Олег",
			Initials: "АО",
			Email: "atabaev.o.k.04@gmail.com",
			Bio: "Будущий backend-разработчик",
		},
	})
}