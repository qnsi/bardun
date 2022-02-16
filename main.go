package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type HomeData struct {
	Message    string
	CurrentDoc string
	NoteMap    map[string]string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		panic(err)
	}

	keys, ok := r.URL.Query()["note"]
	var key string
	if !ok || len(keys[0]) < 1 {
		log.Println("Url param is missing")
		key = "Home"
	} else {
		key = keys[0]
	}

	noteMap := map[string]string{
		"Home":  "Welcome to golang generated variables in javascript\n[[Books]]",
		"Books": "Books. \n Go back to [[Home]]",
	}

	homeData := HomeData{
		NoteMap:    noteMap,
		Message:    "Hello world!@",
		CurrentDoc: noteMap[key],
	}

	err = tpl.Execute(w, homeData)
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
