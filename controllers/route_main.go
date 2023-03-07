package controllers

import (
	"app/models"
	"html/template"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, "Hello")
}

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("views/templates/signup.html")
		if err != nil {
			log.Fatalln(err)
		}
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user := models.User{
			Name: r.PostFormValue("name"),
			Email: r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 302)
	}
}
