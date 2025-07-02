package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("rraagg")
	http.Handle("/", templ.Handler(component))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
