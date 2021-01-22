package game

import (
	"fmt"

	"github.com/gorilla/websocket"

	"github.com/kthatoto/termworld/utils"
)

type PlayerProcedures int
type PlayerProcedureArgs struct {
	PlayerName string
	Options []string
}

func (p *PlayerProcedures) Start(args PlayerProcedureArgs, resp *interface{}) error {
	requestId := utils.RandomString(12)
	responseMap[requestId] = make(chan Response)

	request := fmt.Sprintf(
		"{\"playerName\":\"%s\",\"command\":\"start\",\"options\":null,\"requestId\":\"%s\"}",
		args.PlayerName,
		requestId,
	)
	WSConn.WriteMessage(websocket.TextMessage, []byte(request))

	response := <-responseMap[requestId]
	fmt.Println(response.Message)
	close(responseMap[requestId])
	delete(responseMap, requestId)

	*resp = response.Message

	return nil
}
