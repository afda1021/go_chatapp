package data

type Message struct {
	Id     int
	Name   string
	RoomId string // メッセージ送信者のルームid
	Text   string
	Date   string
	Time   string
}

/* DBにメッセージを保存 */
func (msg *Message) CreateMessage() {
	db := DbInit()
	defer db.Close()
	//DBにメッセージを追加
	statement := "INSERT INTO messages (name, room_id, text, date, time) VALUES (?, ?, ?, ?, ?)"
	stmt, _ := db.Prepare(statement)
	stmt.QueryRow(msg.Name, msg.RoomId, msg.Text, msg.Date, msg.Time)
}

/* DBからメッセージを取得 */
func GetMessage(room_id int) (msgs []Message) {
	db := DbInit()
	defer db.Close()
	//DBからメッセージを取得
	statement := "SELECT id, name, room_id, text, date, time FROM messages WHERE room_id = ?"
	stmt, _ := db.Prepare(statement)
	rows, _ := stmt.Query(room_id)

	for rows.Next() {
		var msg Message
		rows.Scan(&msg.Id, &msg.Name, &msg.RoomId, &msg.Text, &msg.Date, &msg.Time)
		msgs = append(msgs, msg)
	}
	return
}

func DeleteMsg(msgId int) (err error) {
	db := DbInit()
	defer db.Close()
	//DBからidと一致するmessageを削除
	statement := "DELETE FROM messages WHERE id = ?"
	stmt, _ := db.Prepare(statement)
	_, err = stmt.Exec(msgId)
	return
}
