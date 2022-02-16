package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type HomeData struct {
	Message    string
	CurrentDoc string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, HomeData{
		Message:    "Hello world!@",
		CurrentDoc: "Welcome to golang generated variables in javascript",
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	// db := utils.PrepareDB()
	// contacts := models.GetContacts(db)
	// fmt.Println(contacts[0].FirstName)
	http.HandleFunc("/", homeHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Starting server at :3000...")
	http.ListenAndServe(":3000", nil)
}
