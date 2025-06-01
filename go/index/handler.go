package main

import (
	"net/http"
	"shared"
	"text/template"
)

type (
	Response struct {
		Title string
		User User
	}

	User struct {
		ID string
		FullName string
		Email string
		Bio string
	}
)

func getUser(username string) User {
	return User{
		ID: "1", 
		FullName: "Атабаев Олег Константинович", 
		Email: "atabaev.o.k.04@gmail.com", 
		Bio: "Будущий backend-разработчик",
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/base.html", "../templates/home.html")

	if err != nil {
		shared.ErrorResponse(w, err)
		return
	}

	response := Response{
		Title: "Главная",
	}
	
	username, err := shared.GetCookie(r, "username")

	if err != nil {
		tmpl.ExecuteTemplate(w, "home.html", response)
		return
	}

	response.User = getUser(username)

	tmpl.ExecuteTemplate(w, "home.html", response)
}