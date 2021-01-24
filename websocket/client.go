package socket

import (
	"chat/data"
	"fmt"

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
			fmt.Println(msg.Date)
			msg.CreateMessage() //DBにメッセージを保存
			c.room.forward <- msg
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
