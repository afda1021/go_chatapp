package main

import (
	"chat/data"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

func index(w http.ResponseWriter, r *http.Request) {
	session := session(w, r) //クッキーと一致するセッションをDBから取得

	if session.Id == 0 {
		t := template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))
		t.ExecuteTemplate(w, "layout", nil)
	} else {
		user := data.UserByUuid(w, r)
		rooms := data.GetRooms() //DBからルームを取得
		type Data struct {
			User  data.User
			Rooms []data.Room
		}
		data := Data{User: user, Rooms: rooms} //構造体dataにuserとroonsを格納

		t := template.Must(template.ParseFiles("templates/layout.html", "templates/room_top.html"))
		t.ExecuteTemplate(w, "layout", data)
	}
}

/* 新規登録画面 */
func signup(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/signup.html"))
	t.ExecuteTemplate(w, "layout", nil)
}

/* ログイン画面 */
func login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/login.html"))
	t.ExecuteTemplate(w, "layout", nil)
}

/* ゲストログイン画面 */
func loginGuest(w http.ResponseWriter, r *http.Request) {
	cfg, _ := ini.Load("config.ini")
	type Data struct {
		Name     string
		Password string
	}
	data := Data{
		Name:     cfg.Section("guest_account").Key("username").String(),
		Password: cfg.Section("guest_account").Key("password").String(),
	}

	t := template.Must(template.ParseFiles("templates/layout.html", "templates/login_guest.html"))
	t.ExecuteTemplate(w, "layout", data)
}

/* 新規登録処理 */
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

/* ログイン処理 */
func authenticate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")
	user := data.UserByName(name) //DBから一致するユーザーを取得

	if user.Password != data.Encrypt(password) {
		http.Redirect(w, r, "/login", 302) //ログイン失敗
	} else {
		session := user.CreateSession() //セッション作成
		//クッキー作成
		cookie := http.Cookie{
			Name:   "_cookie",
			Value:  session.Uuid,
			MaxAge: 60 * 60 * 6, //有効期限(秒)
			// Expires:  time.Now().Add(1 * time.Minute),
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302) //ログイン成功
	}
}

/* ログアウト処理 */
func logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("_cookie") //ブラウザのクッキー取得
	session := data.Session{Uuid: cookie.Value}
	session.DeleteByUUID() //DBのセッションを削除

	http.Redirect(w, r, "/", 302)
}

/* ルーム作成画面 */
func newRoom(w http.ResponseWriter, r *http.Request) {
	session := session(w, r) //クッキーと一致するセッションを取得
	if session.Id == 0 {
		http.Redirect(w, r, "/", 302)
	} else {
		t := template.Must(template.ParseFiles("templates/layout.html", "templates/room_new.html"))
		t.ExecuteTemplate(w, "layout", nil)
	}
}

/* ルーム作成 */
func createRoom(w http.ResponseWriter, r *http.Request) {
	session := session(w, r) //クッキーと一致するセッションを取得

	if session.Id == 0 {
		http.Redirect(w, r, "/", 302)
	} else {
		room := data.Room{
			RoomName: r.PostFormValue("room_name"), //入力したルーム名
		}
		room.CreateRoom() //DBにroomを追加
		http.Redirect(w, r, "/", 302)
	}
}

/* ルーム画面 */
func room(w http.ResponseWriter, r *http.Request) {
	session := session(w, r) //クッキーと一致するセッションを取得

	if session.Id == 0 {
		http.Redirect(w, r, "/", 302)
	} else {
		query := r.URL.Query()
		id := query.Get("id") //roomのIdを取得
		intID, _ := strconv.Atoi(id)

		room := data.GetRoom(intID) //ルームidと一致するroomを取得
		var user data.User
		user = data.UserByUuid(w, r)       //ユーザー名を取得
		threads := data.GetMessages(intID) //ルームidと一致するmessageを全て取得

		type Data struct {
			Room    data.Room
			Name    string
			Threads []data.Thread
		}
		data := Data{Room: room, Name: user.Name, Threads: threads}
		t := template.Must(template.ParseFiles("templates/layout.html", "templates/room.html"))
		t.ExecuteTemplate(w, "layout", data)
	}
}

/* メッセージ送信取消 */
// func deleteMsg(w http.ResponseWriter, r *http.Request) {
// 	session := session(w, r) //クッキーと一致するセッションを取得

// 	if session.Id == 0 {
// 		http.Redirect(w, r, "/", 302)
// 	} else {
// 		query := r.URL.Query()
// 		id := query.Get("id")        //roomのIdを取得
// 		msgId := query.Get("msg_id") //msgのIdを取得
// 		// intId, _ := strconv.Atoi(id)
// 		intMsgId, _ := strconv.Atoi(msgId)

// 		data.RemoveMsg(intMsgId) //idと一致するmessageを削除
// 		http.Redirect(w, r, "/room?id="+id, 302)
// 	}
// }
