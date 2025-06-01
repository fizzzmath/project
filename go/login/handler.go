package main

import (
	"encoding/json"
	"errors"
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
		Token string `json:"token"`
	}
)

func invalid(form Form) bool {
	return false
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
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("../templates/auth/login.html", "../templates/base.html")

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		tmpl.ExecuteTemplate(w, "login.html", nil)
	}

	if r.Method == http.MethodPost {
		form := Form{}

		err := json.Unmarshal([]byte(r.Header.Get("Body")), &form)

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

		json.NewEncoder(w).Encode(Response{
			Token: token,
		})
	}
}