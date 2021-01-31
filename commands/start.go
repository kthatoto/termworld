package commands

import (
	"os"
	"fmt"
	"net/url"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/gorilla/websocket"
	homedir "github.com/mitchellh/go-homedir"
	daemon "github.com/sevlyar/go-daemon"

	"github.com/kthatoto/termworld/app/game"
)

func init() {
	rootCommand.AddCommand(startCommand)
}

var startCommand = &cobra.Command{
	Use: "start",
	Short: "Game start command",
	RunE: func(cmd *cobra.Command, args []string) error {
		chdirHome()

		ctx := &daemon.Context{
			PidFileName: "termworld.pid",
			PidFilePerm: 0644,
			LogFileName: "termworld.log",
			LogFilePerm: 0640,
			WorkDir:     "./",
			Umask:       027,
		}

		already, _ := ctx.Search()
		if already != nil {
			fmt.Println("Already started")
			return nil
		}
		child, err := ctx.Reborn()
		if err != nil {
			return err
		}
		if child != nil {
			fmt.Println("Started")
			return nil
		}
		defer ctx.Release()

		u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/gaming"}
		httpHeader := http.Header{}
		token := viper.Get("token").(string)
		httpHeader.Set("X-Termworld-Token", token)
		conn, _, err := websocket.DefaultDialer.Dial(u.String(), httpHeader)
		if err != nil {
			return err
		}

		daemonWork(conn)
		return nil
	},
}

func chdirHome() {
	home, err := homedir.Dir()
	if err != nil {
		return
	}
	if err := os.Chdir(home+"/.termworld"); err != nil {
		return
	}
}

func daemonWork(conn *websocket.Conn) {
	defer conn.Close()
	done := make(chan bool)
	go game.HandleProcedures(conn, done)
	go game.ReadMessages(conn, done)
	<-done
	return
}
