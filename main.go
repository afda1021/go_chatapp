package main

import (
	socket "chat/websocket"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files)) ///static/をhttp.FileServer()が捜索するURLから取り除

	mux.HandleFunc("/", index)

	mux.HandleFunc("/signup", signup)                //新規登録画面
	mux.HandleFunc("/login", login)                  //ログイン画面
	mux.HandleFunc("/login_guest", loginGuest)       //ゲストログイン画面
	mux.HandleFunc("/signup_account", signupAccount) //新規登録処理
	mux.HandleFunc("/authenticate", authenticate)    //ログイン処理
	mux.HandleFunc("/logout", logout)                //ログアウト処理

	mux.HandleFunc("/room/new", newRoom)       //ルーム作成画面
	mux.HandleFunc("/room/create", createRoom) //ルーム作成
	mux.HandleFunc("/room", room)              //ルーム画面

	// mux.HandleFunc("/delete_msg", deleteMsg) //メッセージ送信取消

	chatroom := socket.NewChatroom() // チャットルームを作成
	mux.HandleFunc("/ws", chatroom.ServeHTTP)
	go chatroom.Run() // チャットルームを起動する

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
