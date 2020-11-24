package commands

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCommand.AddCommand(logoutCommand)
}

var logoutCommand = &cobra.Command{
	Use: "logout",
	Short: "Logout command",
	RunE: func(cmd *cobra.Command, args []string) error {
		token := viper.Get("token").(string)
		if len(token) == 0 {
			fmt.Println("Not logged in")
			return nil
		}

		u := "http://localhost:8080/logout"
		req, err := http.NewRequest("DELETE", u, nil)
		if err != nil {
			return err
		}
		req.Header.Set("X-Termworld-Token", token)
		client := new(http.Client)
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode == 401 {
			fmt.Println("Not logged in")
			return nil
		}
		fmt.Println("Logout success!")
		viper.Set("token", "")
		viper.WriteConfig()

		return nil
	},
}
