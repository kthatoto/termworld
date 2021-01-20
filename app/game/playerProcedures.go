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

func (p *PlayerProcedures) Start(args PlayerProcedureArgs, result *bool) error {
	requestId := utils.RandomString(12)
	message := fmt.Sprintf(
		"{\"playerName\":\"%s\",\"command\":\"start\",\"options\":null,\"requestId\":\"%s\"}",
		args.PlayerName,
		requestId,
	)
	WSConn.WriteMessage(websocket.TextMessage, []byte(message))
	*result = true
	return nil
}
