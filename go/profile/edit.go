package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"shared"
	"text/template"

	"github.com/golang-jwt/jwt/v5"
)

type (
	Form struct {
		FullName string `json:"fullName"`
		Email string `json:"email"`
		Bio string `json:"bio"`
	}
)

func update(form Form, user User) error {
	db, err := shared.Connect()

	if err != nil {
		return err
	}

	defer db.Close()

	if form.FullName != *user.FullName {
		err := shared.UpdateUser(db, "FullName", form.FullName, *user.ID)

		if err != nil {
			return err
		}
	}

	if form.Email != user.Email {
		err := shared.UpdateUser(db, "Email", form.Email, *user.ID)

		if err != nil {
        	return err
		}
	}

	if form.Bio != *user.Bio {
		err := shared.UpdateUser(db, "Bio", form.Bio, *user.ID)

		if err != nil {
			return err
		}
	}

	return nil
}

func unauthorized(tokenString string, username string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("access-token-secret-key"), nil
	})

	if err != nil || !token.Valid {
		return true
	}

	payload, _ := token.Claims.(*jwt.RegisteredClaims)

	return payload.Subject != username
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	username, err := shared.GetCookie(r, "username")

	if err != nil {
		shared.ErrorResponse(w, err)
		return
	}

	if r.Method == http.MethodPost {
		token := r.Header.Get("X-Auth-Token")
		form := Form{}

		err := json.NewDecoder(r.Body).Decode(&form)

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		user, err := getUser(username)

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		if unauthorized(token, user.Username) {
		 	shared.ErrorResponse(w, errors.New("для выполнения этого действия необходимо авторизоваться"))
			return
		}

		err = update(form, *user)

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		shared.SuccessResponse(w)
	}

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("../templates/profile/edit.html", "../templates/base.html")

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		user, err := getUser(username)

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		tmpl.ExecuteTemplate(w, "edit.html", Response{
			Title: "Редактирование профиля",
			User: user,
		})
	}
}