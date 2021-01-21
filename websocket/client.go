package socket

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn // websocketへのコネクション
	send   chan *Message
	room   *chatroom
}

type Message struct {
	Text string
}

/* websocketに書き出されたメッセージを読み込む。*/
func (c *client) read() {
	for {
		var msg *Message
		if err := c.socket.ReadJSON(&msg); err == nil {
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
