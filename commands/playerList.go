package commands

import (
	"fmt"
	"errors"
	"encoding/json"
	"io/ioutil"

	"github.com/spf13/cobra"

	"github.com/kthatoto/termworld/utils"
	"github.com/kthatoto/termworld/app/models"
)

func init() {
	playerCommand.AddCommand(playerListCommand)
}

var playerListCommand = &cobra.Command{
	Use: "list",
	Short: "Player list command",
	RunE: func(cmd *cobra.Command, args []string) error {
		httpClient := utils.HttpClient{WithToken: true}
		resp, err := httpClient.Call("GET", "/players", nil)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			return errors.New("Request failed")
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var responseBody models.PlayersResponseBody
		if err := json.Unmarshal(bytes, &responseBody); err != nil {
			return err
		}
		for _, player := range responseBody.Players {
			fmt.Printf("ID:%s Name:%s\n", player.ID, player.Name)
		}
		return nil
	},
}
