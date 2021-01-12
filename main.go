package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	mux.HandleFunc("/login", login)   //ログイン画面
	mux.HandleFunc("/signup", signup) //新規登録画面

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
