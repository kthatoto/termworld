package game

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func ReadMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Finished")
			break
		}
		fmt.Printf("read message: %s\n", message)
	}
}

func SendRequest(conn *websocket.Conn) {
	requestString := "requestMap"
	conn.WriteMessage(websocket.TextMessage, []byte(requestString))
}
