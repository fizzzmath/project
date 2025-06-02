package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"shared"
	"text/template"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type (
	Form struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Response struct {
		Title string `json:"title,omitempty"`
		Token string `json:"token"`
		User *User `json:"user,omitempty"`
	}

	User = shared.User
)

func invalid(form Form) bool {
	db, err := shared.Connect()

	if err != nil {
		return true
	}

	defer db.Close()

	sel, err := db.Query(`
		SELECT *
		FROM Login
		WHERE Username = ? AND Password = ?
	`, form.Username, fmt.Sprintf("%x", sha256.Sum256([]byte(form.Password))))

	if err != nil {
		return true
	}

	defer sel.Close()

	for sel.Next() {
		return false
	}

	return true
}

func grantAccessToken(username string) (string, error) {
	payload := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Subject: username,
	}
	key := []byte("access-token-secret-key")

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := accessToken.SignedString(key)

	if err != nil {
		return "", err
	}

	return t, nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("../templates/auth/login.html", "../templates/base.html")

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		tmpl.ExecuteTemplate(w, "login.html", Response{
			Title: "Вход",
		})
	}

	if r.Method == http.MethodPost {
		form := Form{}
		
		err := json.NewDecoder(r.Body).Decode(&form)

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		if invalid(form) {
			shared.ErrorResponse(w, errors.New("неверные логин или пароль"))
			return
		}

		token, err := grantAccessToken(form.Username)

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		shared.SetCookie(w, "username", form.Username)

		err = json.NewEncoder(w).Encode(Response{
			Token: token,
		})

		if err != nil {
			shared.ErrorResponse(w, err)
		}
	}
}