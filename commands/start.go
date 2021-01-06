package commands

import (
	"fmt"
	"os"
	"time"
	"errors"
	"net/url"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/gorilla/websocket"
	homedir "github.com/mitchellh/go-homedir"
	daemon "github.com/sevlyar/go-daemon"
)

func init() {
	rootCommand.AddCommand(startCommand)
}

var startCommand = &cobra.Command{
	Use: "start",
	Short: "Game start command",
	RunE: func(cmd *cobra.Command, args []string) error {
		u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/gaming"}
		httpHeader := http.Header{}
		token := viper.Get("token").(string)
		httpHeader.Set("X-Termworld-Token", token)
		conn, _, err := websocket.DefaultDialer.Dial(u.String(), httpHeader)
		if err != nil {
			return err
		}
		defer conn.Close()

		home, err := homedir.Dir()
		if err != nil {
			return err
		}
		if err := os.Chdir(home+"/.termworld"); err != nil {
			return err
		}

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
			return errors.New("Already started!")
		}
		child, err := ctx.Reborn()
		if err != nil {
			return err
		}
		if child != nil {
			return nil
		}
		defer ctx.Release()
		fmt.Println("-----------------------")
		fmt.Println("Start daemon")

		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println(i)
		}
		return nil
	},
}
