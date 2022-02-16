package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type HomeData struct {
	Message string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, HomeData{Message: "Hello world!@"})
	if err != nil {
		panic(err)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./static/home.js")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/javascript")
	w.Write(data)
}

func main() {
	// db := utils.PrepareDB()
	// contacts := models.GetContacts(db)
	// fmt.Println(contacts[0].FirstName)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/static/home.js", staticHandler)
	fmt.Println("Starting server at :3000...")
	http.ListenAndServe(":3000", nil)
}
