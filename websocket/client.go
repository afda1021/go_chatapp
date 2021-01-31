package socket

import (
	"chat/data"

	"github.com/gorilla/websocket"
)

type client struct {
	roomId string
	socket *websocket.Conn // websocketへのコネクション
	send   chan *data.Message
	room   *chatroom
}

/* websocketに書き出されたメッセージを読み込む。*/
func (c *client) read() {
	for {
		var msg *data.Message
		if err := c.socket.ReadJSON(&msg); err == nil {
			if msg.Type == "publish" || msg.Type == "reply" { //新規メッセージ受信
				msg.CreateMessage() //DBにメッセージを保存
				msg.GetMessageId()  //保存したメッセージのidを取得
				c.room.forward <- msg
			} else if msg.Type == "remove" { //送信取り消し
				data.RemoveMsg(msg.Id) //DBからmessageを削除
				c.room.forward <- msg
			}
		} else {
			break
		}
	}
	c.socket.Close()
}

/* 読み込んだメッセージをwebsocketに書き込む。*/
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
