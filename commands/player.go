package commands

import (
	"errors"
	"encoding/json"

	"github.com/spf13/cobra"

	"github.com/kthatoto/termworld/utils"
	"github.com/kthatoto/termworld/app/models"
)

func init() {
	rootCommand.AddCommand(playerCommand)
}

var playerCommand = &cobra.Command{
	Use: "player",
	Short: "Player command",
	RunE: func(cmd *cobra.Command, args []string) error {
		httpClient := utils.HttpClient{WithToken: true}
		resp, err := httpClient.Call("GET", "/players", nil)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			return errors.New("Request failed")
		}

		var responseBody models.PlayersResponseBody
		if err := json.Unmarshal(responseBody); err != nil {
			return err
		}
		for player := range responseBody {
			fmt.Printf("ID:%s Name:%s\n", player.ID, player.Name)
		}
		return nil
	},
}
