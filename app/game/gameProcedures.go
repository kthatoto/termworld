package game

import (
	"fmt"
	"encoding/json"

	"github.com/gorilla/websocket"

	"github.com/kthatoto/termworld/app/utils"
)

type GameProcedures int
type GameProcedureArgs struct {
	Command    string
	PlayerName string
	Options    []string
}

type Command struct {
	Command    string   `json:"command"`
	PlayerName string   `json:"playerName"`
	Options    []string `json:"options"`
	RequestId  string   `json:"requestId"`
}

func (p *GameProcedures) Execute(args GameProcedureArgs, resp *interface{}) error {
	requestId := utils.RandomString(12)
	responseMap[requestId] = make(chan Response)

	command := Command{
		Command: args.Command,
		PlayerName: args.PlayerName,
		Options: args.Options,
		RequestId: requestId,
	}
	requestJson, _ := json.Marshal(command)
	WSConn.WriteMessage(websocket.TextMessage, []byte(requestJson))

	response := <-responseMap[requestId]
	fmt.Println(response.Message)
	close(responseMap[requestId])
	delete(responseMap, requestId)

	*resp = response.Message

	return nil
}
