package data

import (
	"database/sql"
	"os"
	"time"
)

type User struct {
	Id       int
	Name     string
	Password string
}

type Session struct {
	Id   int
	Uuid string
	Name string
}

type Room struct {
	Id         int
	RoomName   string
	UpdateTime string
}

/* DB接続 */
func DbInit() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		panic(err)
	}
	return db
}

/* 新規ユーザーをDBに保存 */
func (user *User) Create() (err error) {
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

/* DBからユーザー名が一致するユーザーを取得 */
func UserByName(name string) (user User) {
	db := DbInit()
	defer db.Close()

	db.QueryRow("SELECT name, password FROM users WHERE name = ?", name).Scan(&user.Name, &user.Password)
	return
}

func (user *User) CreateSession() (session Session) {
	db := DbInit()
	defer db.Close()
	//セッション作成
	statement1 := "INSERT INTO sessions (uuid, name) VALUES(?, ?)"
	stmt1, _ := db.Prepare(statement1)
	defer stmt1.Close()
	stmt1.QueryRow(CreateUUID(), user.Name)
	//セッション取得
	statement2 := "SELECT uuid FROM sessions WHERE name = ?"
	stmt2, _ := db.Prepare(statement2)
	defer stmt2.Close()
	stmt2.QueryRow(user.Name).Scan(&session.Uuid)
	return
}

/* DBにroomを追加 */
func (room *Room) CreateRoom() {
	db := DbInit()
	defer db.Close()
	//DBにroomを追加
	statement := "INSERT INTO rooms(room_name, update_time) VALUES (?, ?)"
	t := time.Now()
	const layout = "2006-01-02 15:04:05"
	stmt, _ := db.Prepare(statement)
	defer stmt.Close()
	stmt.QueryRow(room.RoomName, t.Format(layout))
}

/* roomのupdate_timeを更新 */
func UpdateRoomTime(roomId int) {
	db := DbInit()
	defer db.Close()
	//DBにroomを追加
	statement := "UPDATE rooms SET update_time = ? WHERE id = ?"
	t := time.Now()
	const layout = "2006-01-02 15:04:05"
	stmt, _ := db.Prepare(statement)
	defer stmt.Close()
	stmt.QueryRow(t.Format(layout), roomId)
}

/* DBからroomを取得 */
func GetRooms() (rooms []Room) {
	db := DbInit()
	defer db.Close()
	// var msg Message
	//DBからroomを取得
	rows, _ := db.Query("SELECT id, room_name, update_time FROM rooms")
	for rows.Next() {
		room := Room{}
		rows.Scan(&room.Id, &room.RoomName, &room.UpdateTime)
		room.UpdateTime = room.UpdateTime[:10] + " " + room.UpdateTime[11:16]
		rooms = append(rooms, room)
	}
	rows.Close()
	return
}

func GetRoom(room_id int) (room Room) {
	db := DbInit()
	defer db.Close()
	//DBからidと一致するroomを取得
	db.QueryRow("SELECT id, room_name FROM rooms WHERE id = ?", room_id).
		Scan(&room.Id, &room.RoomName)
	return
}
