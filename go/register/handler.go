package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"shared"

	_ "github.com/go-sql-driver/mysql"
)

type (
	User = shared.User

	Response struct {
		Title string `json:"title,omitempty"`
		User *User `json:"user,omitempty"`
	}

	Form struct {
		Username string `json:"username"`
		Email string `json:"email"`
		Password string `json:"password"`
	}
)

func validate(form Form) error {
	db, err := shared.Connect()

	if err != nil {
		return err
	}

	defer db.Close();

	sel, err := db.Query(`
		SELECT *
		FROM User
		WHERE Email = ?
	`, form.Email)

	if err != nil {
		return err
	}

	defer sel.Close()

	for sel.Next() {
		return errors.New("пользователь с таким email уже существует")
	}

	sel, err = db.Query(`
		SELECT *
		FROM User
		WHERE Username = ?
	`, form.Username)

	if err != nil {
		return err
	}

	defer sel.Close()

	for sel.Next() {
		return errors.New("пользователь с таким логином уже существует")
	}

	return nil
}

func addUser(form Form) error {
	db, err := shared.Connect()

	if err != nil {
		return err
	}

	defer db.Close();

	_, err = db.Exec(`
		INSERT INTO Login
		VALUES (?, ?)
	`, form.Username, fmt.Sprintf("%x", sha256.Sum256([]byte(form.Password))))

	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO User(Email, Username)
		VALUES (?, ?)
	`, form.Email, form.Username)

	return err
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		form := Form{}

		err := json.NewDecoder(r.Body).Decode(&form)

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		err = validate(form)

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		err = addUser(form)

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		shared.SuccessResponse(w)
	}

	if r.Method == http.MethodGet {
		response := Response{
			Title: "Регистрация",
		}

		tmpl, err := template.ParseFiles("../templates/auth/register.html", "../templates/base.html")
	
		if err != nil {
			fmt.Fprintf(w, "Template error")
			log.Print(err)
			return
		}
	
		tmpl.ExecuteTemplate(w, "register.html", response)
	}
}