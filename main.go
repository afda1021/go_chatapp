package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	mux.HandleFunc("/signup", signup)                //新規登録画面
	mux.HandleFunc("/login", login)                  //ログイン画面
	mux.HandleFunc("/signup_account", signupAccount) //新規登録処理
	mux.HandleFunc("/authenticate", authenticate)    //ログイン処理
	mux.HandleFunc("/logout", logout)                //ログアウト処理

	mux.HandleFunc("/room/new", newRoom)       //ルーム作成画面
	mux.HandleFunc("/room/create", createRoom) //ルーム作成
	mux.HandleFunc("/room", room)              //ルーム画面

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
