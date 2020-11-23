package main

import (
	"net/url"
	"net/http"
	"fmt"
	"os"
	"bufio"
	"strings"

	"github.com/spf13/cobra"
	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/gaming"}
	httpHeader := http.Header{}
	httpHeader.Set("X-Termworld-Token", "")
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), httpHeader)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("read:", err)
				return
			}
			for _, line := range strings.Split(string(message), "\n") {
				fmt.Printf("read: %s\n", line)
			}
		}
	}()

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		fmt.Printf("wrote [%s]\n", text)
		if err != nil {
			fmt.Println("write:", err)
			return
		}
	}
}
