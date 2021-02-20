package socket

import (
	"chat/data"
	"fmt"
	"strconv"

	"github.com/gorilla/websocket"
)

type client struct {
	name     string
	roomId   string
	socket   *websocket.Conn // websocketへのコネクション
	send     chan *data.Message
	sendUser chan []*clientUser
	room     *chatroom
}
type clientUser struct {
	Name   string
	RoomId string
}

/* websocketに書き出されたメッセージを読み込む。*/
func (c *client) read() {
	for {
		var msg *data.Message
		if err := c.socket.ReadJSON(&msg); err == nil {
			if msg.Type == "publish" || msg.Type == "reply" { //新規メッセージ受信
				msg.CreateMessage() //DBにメッセージを保存
				msg.GetMessage()    //保存したメッセージのidとdatetimeを取得
				roomId, _ := strconv.Atoi(msg.RoomId)
				data.UpdateRoomTime(roomId) //roomのupdate_timeを更新
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

/* 入室中のユーザーをwebsocketに書き込む */
func (c *client) writeUser() {
	for clients := range c.sendUser {
		fmt.Println(clients)
		if err := c.socket.WriteJSON(clients); err != nil {
			break
		}
	}
	c.socket.Close()
}
