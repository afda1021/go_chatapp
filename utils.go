package main

import (
	"chat/data"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (session data.Session) {
	//ブラウザのクッキー取得
	cookie, err := r.Cookie("_cookie")

	if err != nil { //クッキーが存在しない場合
		return
	} else {
		db := data.DbInit()
		defer db.Close()
		//クッキーと一致するセッションを取得
		db.QueryRow("SELECT id FROM sessions WHERE uuid = ?", cookie.Value).
			Scan(&session.Id)
		return
	}
}
