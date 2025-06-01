package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"shared"
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

func addUser(form Form) {
	// ...
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		form := Form{}

		err := json.NewDecoder(r.Body).Decode(&form)

		if err != nil {
			shared.ErrorResponse(w, err)
			return
		}

		addUser(form)

		err = json.NewEncoder(w).Encode(`"status": "success"`)

		if err != nil {
			shared.ErrorResponse(w, err)
		}
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