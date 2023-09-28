package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/form", form)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}

func form(writer http.ResponseWriter, request *http.Request) {
	tmplt := template.Must(template.ParseFiles("gorm/template", "form.html"))
	if err := tmplt.Execute(writer, nil); err != nil {
		panic(err)
	}
}
