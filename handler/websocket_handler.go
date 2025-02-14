package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// true to allow the request
		fmt.Println("CheckOrigin called and return true value")
		return true
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	//simple echo server logic

	for {
		mType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error to read message : ", err)
			break
		}
		fmt.Println("Message received :", string(p))

		// Echo message back
		sendMessage := "from server : " + string(p)
		if err := conn.WriteMessage(mType, []byte(sendMessage)); err != nil {
			fmt.Println("Error to Write message : ", err)
			break
		}
	}

}
