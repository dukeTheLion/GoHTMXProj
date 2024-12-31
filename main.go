package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Pet struct {
	Value string
	Name  string
}

func main() {

	pets := map[string][]Pet{
		"Pets": {
			{Value: string(byte(1)), Name: "I"},
			{Value: string(byte(2)), Name: "J"},
			{Value: string(byte(3)), Name: "K"},
		},
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseGlob("templates/*.html"))
		tmpl.ExecuteTemplate(w, "index.html", pets)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		slt := r.PostFormValue("pets")
		frm := r.PostFormValue("frm")

		fmt.Println([]byte(slt))
		fmt.Println([]byte(frm))

		h1(w, r)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/test/", h2)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
