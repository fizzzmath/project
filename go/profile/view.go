package main

import (
	"errors"
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

func getUser(username string) (*User, error) {
	db, err := shared.Connect()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	user := User{}

	sel, err := db.Query(`
		SELECT *
		FROM User
		WHERE Username = ?
	`, username)

	if err != nil {
		return nil, err
	}

	if sel.Next() {
		err = sel.Scan(
			&user.ID, 
			&user.FullName, 
			&user.Email, 
			&user.Bio,
			&user.Username, 
			&user.Initials, 
		)

		if err != nil {
			return nil, err
		}

		return &user, nil
	}

	return nil, errors.New("пользователь не найден")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/profile/view.html", "../templates/base.html")

	if err != nil {
		shared.ErrorResponse(w, err)
		return
	}

	username, err := shared.GetCookie(r, "username")

	if err != nil {
		shared.ErrorResponse(w, err)
		return
	}

	user, err := getUser(username)

	if err != nil {
		shared.ErrorResponse(w, err)
		return
	}

	tmpl.ExecuteTemplate(w, "view.html", Response{
		Title: "Профиль",
		User: user,
	})
}