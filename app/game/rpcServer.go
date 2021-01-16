package game

import (
	"net"
	"net/http"
	"net/rpc"

	"github.com/gorilla/websocket"
)

type Procedures int
var proceduresDone chan bool

func (p *Procedures) Stop() {
	proceduresDone <- true
}

func HandleProcedure(conn *websocket.Conn, done chan bool) {
	go func() {
		done <- <-proceduresDone
	}()
	procedures := new(Procedures)
	err := rpc.Register(procedures)
	if err != nil {
		done <- true
		return
	}
	rpc.HandleHTTP()
	var listener net.Listener
	listener, err = net.Listen("tcp", ":8128")
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
