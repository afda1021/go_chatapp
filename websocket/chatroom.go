package socket

import (
	"chat/data"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/* websocket用の変数 */
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

/* 全てのクライアントを管理 */
type chatroom struct {
	forward chan *data.Message
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

/* チャットルームを作成 */
func NewChatroom() *chatroom {
	fmt.Println("chatroom が生成されました。")
	return &chatroom{
		forward: make(chan *data.Message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (c *chatroom) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/* websocketの開設 */
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln("websocketの開設に失敗しました。:", err)
	}
	query := r.URL.Query()
	id := query.Get("id")
	name := query.Get("name")

	/* クライアントの生成 */
	client := &client{
		name:     name,
		roomId:   id, //ルームid(クエリ)
		socket:   socket,
		send:     make(chan *data.Message),
		sendUser: make(chan []*clientUser),
		room:     c,
	}
	// クライアントを入室させる。最後には必ず退室させる。
	c.join <- client
	defer func() {
		c.leave <- client
	}()

	go client.write()     // messageの書き出し
	go client.writeUser() // userの書き出し
	client.read()         // messageの読み込み
}

/* チャットルームを起動する */
func (c *chatroom) Run() {
	/* forwardチャネルに動きがあった場合(メッセージの受信) */
	var clientUsers []*clientUser
	for {
		select {
		/* joinチャネルに動きがあった場合(クライアントの入室) */
		case clt := <-c.join:
			c.clients[clt] = true
			// 入室したユーザー
			clientUser := &clientUser{
				Name:   clt.name,
				RoomId: clt.roomId,
			}
			fmt.Println(clt.name, "ok")
			clientUsers = append(clientUsers, clientUser)
			// 存在するクライアント全てに対してユーザー一覧を送信
			for target := range c.clients {
				select {
				case target.sendUser <- clientUsers:
					fmt.Println("ユーザー送信")
				default:
					fmt.Println("ユーザー送信に失敗")
					delete(c.clients, target)
				}
			}
			fmt.Printf("クライアント入室。現在 %x 人。\n", len(c.clients))

		/* leaveチャネルに動きがあった場合(クライアントの退室) */
		case clt := <-c.leave:
			delete(c.clients, clt)
			// 退室したユーザーを一覧から削除
			var result []*clientUser
			for _, clientUser := range clientUsers {
				if clientUser.Name != clt.name {
					result = append(result, clientUser)
				}
			}
			clientUsers = result
			// 存在するクライアント全てに対してユーザー一覧を送信
			for target := range c.clients {
				select {
				case target.sendUser <- clientUsers:
					fmt.Println("ユーザー削除")
				default:
					fmt.Println("ユーザー削除に失敗")
					delete(c.clients, target)
				}
			}
			fmt.Printf("クライアント退室。現在 %x 人。\n", len(c.clients))

		/* forwardチャネルに動きがあった場合(メッセージの受信) */
		case msg := <-c.forward:
			fmt.Println("メッセージ受信")
			// 存在するクライアント全てに対してメッセージを送信する
			for target := range c.clients {
				if target.roomId == msg.RoomId { //同じルームのクライアントのみに送信
					select {
					case target.send <- msg:
						fmt.Println("メッセージ送信")
					default:
						fmt.Println("メッセージ送信に失敗")
						delete(c.clients, target)
					}
				}
			}
		}
	}
}
