package main

import (
	"net/http"
	"net/http/cgi"
)

func main() {
	cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action := r.URL.Query().Get("action")

		if action == "edit" {
			editHandler(w, r)
		} else {
			viewHandler(w, r)
		}
	}))
}