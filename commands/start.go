package commands

import (
	"fmt"
	"time"
	// "net/url"
	// "net/http"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
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
		// u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/gaming"}
		// httpHeader := http.Header{}
		// token := viper.Get("token").(string)
		// conn, _, err := websocket.DefaultDialer.Dial(u.String(), httpHeader)
		// if err != nil {
		// 	return err
		// }
		// defer conn.Close()

		home, err := homedir.Dir()
		if err != nil {
			return err
		}
		ctx := &daemon.Context{
			PidFileName: "termworld.pid",
			PidFilePerm: 0644,
			LogFileName: "termworld.log",
			LogFilePerm: 0640,
			WorkDir:     fmt.Sprintf("%s/.termworld", home),
			Umask:       027,
		}
		child, err := ctx.Rebort()
		if err != nil {
			return err
		}
		if child != nil {
			return nil
		}
		defer ctx.Release()
		fmt.Println("-----------------------")
		fmt.Println("Start daemon")

		for {
			time.Sleep(3 * time.Second)
			fmt.Println(time.Now())
		}
	},
}
