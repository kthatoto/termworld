package game

import (
	"fmt"
	"encoding/json"

	"github.com/gorilla/websocket"

	"github.com/kthatoto/termworld/utils"
)

type GameProcedures int
type GameProcedureArgs struct {
	PlayerName string
	Options []string
}

type Command struct {
	PlayerName string   `json:"playerName"`
	Command    string   `json:"command"`
	Options    []string `json:"options"`
	RequestId  string   `json:"requestId"`
}

func (p *GameProcedures) Start(args GameProcedureArgs, resp *interface{}) error {
	requestId := utils.RandomString(12)
	responseMap[requestId] = make(chan Response)

	command := Command{
		PlayerName: args.PlayerName,
		Command: "start",
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

func (p *GameProcedures) Stop(args GameProcedureArgs, resp *interface{}) error {
	requestId := utils.RandomString(12)
	responseMap[requestId] = make(chan Response)

	command := Command{
		PlayerName: args.PlayerName,
		Command: "stop",
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

func (p *GameProcedures) Move(args GameProcedureArgs, resp *interface{}) error {
	requestId := utils.RandomString(12)
	responseMap[requestId] = make(chan Response)

	command := Command{
		PlayerName: args.PlayerName,
		Command: "move",
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
