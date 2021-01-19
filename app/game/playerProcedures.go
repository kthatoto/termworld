package game

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type PlayerProcedures int
type PlayerProcedureArgs struct {
	PlayerName string
	Options []string
}

func (p *PlayerProcedures) Start(args PlayerProcedureArgs, result *bool) error {
	message := fmt.Sprintf("{\"playerName\":\"%s\",\"command\":\"start\",\"options\":null}", args.PlayerName)
	WSConn.WriteMessage(websocket.TextMessage, []byte(message))
	*result = true
	return nil
}
