package data

import (
	"strconv"
)

type Message struct {
	Id      int
	Name    string
	RoomId  string // メッセージ送信者のルームid
	Text    string
	Date    string
	Time    string
	ReplyId int
	Type    string
}

type Thread struct {
	Id        int
	Name      string
	RoomId    string
	Text      string
	Date      string
	Time      string
	ReplyId   int
	ReplyMsgs []Message
}

/* DBにメッセージを保存 */
func (msg *Message) CreateMessage() {
	db := DbInit()
	defer db.Close()
	//DBにメッセージを追加
	statement := "INSERT INTO messages (name, room_id, text, date, time, reply_id) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, _ := db.Prepare(statement)
	stmt.QueryRow(msg.Name, msg.RoomId, msg.Text, msg.Date, msg.Time, msg.ReplyId)
}

/* DBからルーム内の全メッセージを取得 */
func GetMessages(room_id int) (threads []Thread) {
	db := DbInit()
	defer db.Close()
	//DBからメッセージを取得
	statement := "SELECT id, name, room_id, text, date, time, reply_id FROM messages WHERE room_id = ?"
	stmt, _ := db.Prepare(statement)
	rows, _ := stmt.Query(room_id)

	var replyMsgs []Message
	for rows.Next() {
		var msg Message
		// var replyMsgs []Message
		rows.Scan(&msg.Id, &msg.Name, &msg.RoomId, &msg.Text, &msg.Date, &msg.Time, &msg.ReplyId)
		/* スレッドとリプライを区別 */
		if msg.ReplyId == 0 {
			thread := Thread{
				Id:      msg.Id,
				Name:    msg.Name,
				RoomId:  msg.RoomId,
				Text:    msg.Text,
				Date:    msg.Date,
				Time:    msg.Time,
				ReplyId: msg.ReplyId,
			}
			// msgs = append(msgs, msg)
			threads = append(threads, thread)
		} else {
			replyMsgs = append(replyMsgs, msg)
		}
	}
	/* 対応するスレッドにリプライを付与 */
	for i := range threads {
		for j := range replyMsgs {
			if threads[i].Id == replyMsgs[j].ReplyId {
				threads[i].ReplyMsgs = append(threads[i].ReplyMsgs, replyMsgs[j])
			}
		}
	}
	return
}

/* DBから最新のメッセージを取得 */
func (msg *Message) GetMessageId() {
	db := DbInit()
	defer db.Close()
	//同一ルーム内で最新のメッセージidを取得
	statement := "SELECT id FROM messages WHERE id = (SELECT max(id) FROM messages WHERE room_id = ?)"
	stmt, _ := db.Prepare(statement)
	defer stmt.Close()
	room_id, _ := strconv.Atoi(msg.RoomId)
	stmt.QueryRow(room_id).Scan(&msg.Id)
	return
}

func RemoveMsg(msgId int) (err error) {
	db := DbInit()
	defer db.Close()
	//DBからidと一致するmessageを削除
	statement := "DELETE FROM messages WHERE id = ?"
	stmt, _ := db.Prepare(statement)
	_, err = stmt.Exec(msgId)
	return
}
