package commands

import (
	"fmt"
	"bufio"
	"os"

	"github.com/spf13/cobra"

	"github.com/kthatoto/termworld/utils"
)

var name string

func init() {
	playerCommand.AddCommand(playerCreateCommand)
	playerCreateCommand.PersistentFlags().StringVar(&name, "name", "", "Player name")
}

var playerCreateCommand = &cobra.Command{
	Use: "create",
	Short: "Player create command",
	RunE: func(cmd *cobra.Command, args []string) error {
		for {
			if len(name) > 0 {
				break
			}
			fmt.Print("Please enter player name: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			name = string(scanner.Text())
		}

		httpClient := utils.HttpClient{WithToken: true}
		type RequestBody struct {
			Name string `json:"name"`
		}
		param := RequestBody{name}
		resp, err := httpClient.Call("POST", "/players", param)
		if err != nil {
			return err
		}
		switch(resp.StatusCode) {
		case 201:
			fmt.Println("Player create success!")
			return nil
		case 403:
			fmt.Println("Your players count already reached max count")
			return nil
		case 409:
			fmt.Println("The name is already used")
			return nil
		default:
			fmt.Println("Request failed")
		}
		return nil
	},
}
