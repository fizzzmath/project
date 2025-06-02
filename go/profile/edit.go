package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"shared"
	"text/template"
)

type (
	Form struct {
		FullName string `json:"fullName"`
		Email string `json:"email"`
		Bio string `json:"bio"`
	}
)

func update(id string, form Form) {
	// ...
}

func unauthorized(token string) bool {
	return true
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.URL.Query().Get("user_id")
		token := r.Header.Get("X-Auth-Token")
		form := Form{}

		err := json.NewDecoder(r.Body).Decode(&form)

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		if unauthorized(token) {
		 	shared.ErrorResponse(w, errors.New("для выполнения этого действия необходимо авторизоваться"))
			return
		}

		update(id, form)

		shared.SuccessResponse(w)
	}

	if r.Method == http.MethodGet {
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
}