package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	t, _ = t.ParseFiles("index.html")
	s := User{Name: "Inti"}
	t.Execute(w, s)
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	t := template.New("hi")
	t, _ = t.Parse("Hi, {{ .Name }}!")
	s := User{Name: "Inti"}
	t.Execute(w, s)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
