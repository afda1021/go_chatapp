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

//新規登録画面
func signup(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/signup.html"))
	t.ExecuteTemplate(w, "signup.html", nil)
}

//ログイン画面
func login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/login.html"))
	t.ExecuteTemplate(w, "login.html", nil)
}

//新規登録処理
func signupAccount(w http.ResponseWriter, r *http.Request) {
	user := &data.User{
		Name:     r.FormValue("name"),
		Password: r.FormValue("password"),
	}

	err := user.Create()
	if err == nil { //ユーザー名が既に存在する場合
		http.Redirect(w, r, "/signup", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

//ログイン処理
func authenticate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")
	user := data.UserByName(name) //DBから一致するユーザーを取得

	if user.Password != data.Encrypt(password) {
		http.Redirect(w, r, "/login", 302)
	} else {
		http.Redirect(w, r, "/room_top", 302)
	}
}

func roomTop(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/room_top.html"))
	t.ExecuteTemplate(w, "room_top.html", nil)
}
