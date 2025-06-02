package main

import (
	"net/http"
	"shared"
	"text/template"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/profile/edit.html", "../templates/base.html")

	if err != nil {
		shared.ErrorResponse(w, err)
		return
	}

	tmpl.ExecuteTemplate(w, "edit.html", Response{
		Title: "Редактирование профиля",
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