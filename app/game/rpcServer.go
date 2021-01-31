package game

import (
	"net"
	"net/http"
	"net/rpc"

	"github.com/gorilla/websocket"
)

type Procedures int
var proceduresDone chan bool
var WSConn *websocket.Conn

func (p *Procedures) Stop(_ int, result *bool) error {
	proceduresDone <- true
	*result = true
	return nil
}

func HandleProcedures(conn *websocket.Conn, done chan bool) {
	WSConn = conn
	proceduresDone = make(chan bool)
	go func() {
		done <- <-proceduresDone
	}()

	procedures := new(Procedures)
	gameProcedures := new(GameProcedures)
	rpc.Register(procedures)
	rpc.Register(gameProcedures)

	rpc.HandleHTTP()
	var listener net.Listener
	listener, err := net.Listen("tcp", ":8128")
	if err != nil {
		done <- true
		return
	}
	err = http.Serve(listener, nil)
	if err != nil {
		done <- true
		return
	}
}
