package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	data := map[string]string{
		"Title": "Imersão Full Cycle - Desafio 1",
		"Titleh1" : "Imersão Full Cycle",
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Servidor iniciado em http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

