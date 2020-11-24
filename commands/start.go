package commands

import (
	"net/url"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		conn, _, err := websocket.DefaultDialer.Dial(u.String(), httpHeader)
		if err != nil {
			return err
		}
		defer conn.Close()
	},
}
