package main

import (
	"fmt"
	"net/http"
	"qnsi/bardun/models"
	"qnsi/bardun/utils"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World!</h1>")
}

func main() {
	db := utils.PrepareDB()
	contacts := models.GetContacts(db)
	fmt.Println(contacts[0].FirstName)
	// http.HandleFunc("/", handleFunc)
	// fmt.Println("Starting server at :3000...")
	// http.ListenAndServe(":3000", nil)
}
