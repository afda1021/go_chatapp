package data

type Message struct {
	Id     int
	Name   string
	RoomId string // メッセージ送信者のルームid
	Text   string
}

/* DBにメッセージを保存 */
func (msg *Message) CreateMessage() {
	db := DbInit()
	defer db.Close()
	//DBにメッセージを追加
	statement := "INSERT INTO messages (name, room_id, text) VALUES (?, ?, ?)"
	stmt, _ := db.Prepare(statement)
	stmt.QueryRow(msg.Name, msg.RoomId, msg.Text)
}

/* DBからメッセージを取得 */
func GetMessage(room_id int) (msgs []Message) {
	db := DbInit()
	defer db.Close()
	//DBからメッセージを取得
	statement := "SELECT id, name, room_id, text FROM messages WHERE room_id = ?"
	stmt, _ := db.Prepare(statement)
	rows, _ := stmt.Query(room_id)

	for rows.Next() {
		var msg Message
		rows.Scan(&msg.Id, &msg.Name, &msg.RoomId, &msg.Text)
		msgs = append(msgs, msg)
	}
	return
}
