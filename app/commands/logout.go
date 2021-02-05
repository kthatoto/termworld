package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/kthatoto/termworld/app/utils"
)

func init() {
	rootCommand.AddCommand(logoutCommand)
}

var logoutCommand = &cobra.Command{
	Use: "logout",
	Short: "Logout command",
	RunE: func(cmd *cobra.Command, args []string) error {
		httpClient := utils.HttpClient{WithToken: true}
		resp, err := httpClient.Call("DELETE", "/logout", nil)
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
