package socket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/* websocket用の変数 */
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/* websocketの開設 */
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln("websocketの開設に失敗しました。:", err)
	}
	/* クライアントの生成 */
	client := &client{
		socket: socket,
		send:   make(chan *Message),
	}

	go client.write() // messageの書き出し
	client.read()     // messageの読み込み
}
