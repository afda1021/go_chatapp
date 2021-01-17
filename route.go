package main

import (
	"chat/data"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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

//新規登録処理
func signupAccount(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")

	err := data.Create(name, password)
	if err == nil { //ユーザー名が既に存在する場合
		http.Redirect(w, r, "/signup", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
