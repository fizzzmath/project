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

	tmpl.ExecuteTemplate(w, "base", Response{
		Title: "Редактирование профиля",
	})
}