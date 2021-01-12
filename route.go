package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.ExecuteTemplate(w, "index.html", nil)
}

//ログイン画面
func login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/login.html"))
	t.ExecuteTemplate(w, "login.html", nil)
}

//新規登録画面
func signup(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/signup.html"))
	t.ExecuteTemplate(w, "signup.html", nil)
}
