package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	mux.HandleFunc("/signup", signup)                //新規登録画面
	mux.HandleFunc("/login", login)                  //ログイン画面
	mux.HandleFunc("/signup_account", signupAccount) //新規登録処理
	mux.HandleFunc("/authenticate", authenticate)    //ログイン処理

	mux.HandleFunc("/room_top", roomTop)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
