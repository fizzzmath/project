package main

import (
	"fmt"
	"net/http"
	"net/http/cgi"
	"os"
)

func main() {
	cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// tmpl, err := template.ParseFiles("templates/base.html", "template/home.html")

		// if err != nil {
		// 	fmt.Fprintf(w, "Template error")
		// 	log.Print(err)
		// 	return
		// }

		// tmpl.ExecuteTemplate(w, "home.html", nil)

		fmt.Fprintf(w, "%s", os.Getenv("DOCUMENT_ROOT"))
	}))
}