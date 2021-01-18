package data

import (
	"database/sql"
)

type User struct {
	Name     string
	Password string
}

//DB接続
func DbInit() *sql.DB {
	db, err := sql.Open("mysql", "user1:0000@/chat")
	if err != nil {
		panic(err)
	}
	return db
}

//新規ユーザーをDBに保存
func (user User) Create() (err error) {
	db := DbInit()
	defer db.Close()
	//既にユーザーが存在しないか確認
	var username string
	err = db.QueryRow("SELECT name FROM users WHERE name = ?", user.Name).Scan(&username)

	if err == nil {
		return
	} else {
		stmt, _ := db.Prepare("INSERT INTO users (name, password) VALUES (?, ?)")
		stmt.QueryRow(user.Name, Encrypt(user.Password))
		return
	}
}

//DBからユーザー名が一致するユーザーを取得
func UserByName(name string) (user User) {
	db := DbInit()
	defer db.Close()

	db.QueryRow("SELECT name, password FROM users WHERE name = ?", name).Scan(&user.Name, &user.Password)
	return
}
