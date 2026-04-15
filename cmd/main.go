package main

import (
	"net/http"
	"html/template"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		templ, err := template.ParseFiles("../tmpl/index.tmpl")
		if err != nil {
			panic(err)
		}

		err = templ.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
