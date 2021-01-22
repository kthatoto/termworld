package game

import (
	"fmt"
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Request struct {
	PlayerName string   `json:"playerName"`
	Command    string   `json:"command"`
	Options    []string `json:"options"`
	RequestId  string   `json:"requestId"`
}
type Response struct {
	RequestId string `json:"requestId"`
	Success   bool   `json:"success"`
	Message   string `json:"message"`
}
var responseMap map[string](chan Response)

func ReadMessages(conn *websocket.Conn, done chan bool) {
	responseMap = make(map[string](chan Response))
	for {
		_, jsonMessage, err := conn.ReadMessage()
		if err != nil {
			done <- true
			break
		}

		var response Response
		if err = json.Unmarshal(jsonMessage, &response); err != nil {
			continue
		}
		if len(response.RequestId) > 0 {
			if channel, ok := responseMap[response.RequestId]; ok {
				channel <- response
			}
		}

		fmt.Printf("read message: %+v\n", response)
	}
}
