package data

import "database/sql"

//DB接続
func DbInit() *sql.DB {
	db, err := sql.Open("mysql", "user1:0000@/chat")
	if err != nil {
		panic(err)
	}
	return db
}

func Create(name, password string) (err error) {
	db := DbInit()
	defer db.Close()
	//既にユーザーが存在しないか確認
	var username string
	err = db.QueryRow("SELECT name FROM users WHERE name = ?", name).Scan(&username)

	if err == nil {
		return
	} else {
		stmt, _ := db.Prepare("INSERT INTO users (name, password) VALUES (?, ?)")
		stmt.QueryRow(name, password)
		return
	}
}
