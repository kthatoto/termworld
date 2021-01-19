package commands

import (
	"fmt"
	"errors"
	"encoding/json"
	"io/ioutil"
	"strings"

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
		for i, player := range responseBody.Players {
			first := i == 0
			last := i == len(responseBody.Players) - 1
			displayPlayerInfo(player, first, last)
		}
		return nil
	},
}

func displayPlayerInfo(player *models.Player, first, last bool) {
	width := 30
	const horizontalLine := strings.Repeat("━", width)
	if (first) {
		fmt.Println("┏" + horizontalLine + "┓")
	}

	drawLine(fmt.Sprintf(" Name: %s", player.Name))
	drawLine(fmt.Sprintf(" Live: %s", player.Live ? "true" : "false"))
	drawLine(fmt.Sprintf(" Status:"))
	drawLine(fmt.Sprintf("   HP: 10 / 10"))

	if (last) {
		fmt.Println("┗" + horizontalLine + "┛")
	} else {
		fmt.Println("┣" + horizontalLine + "┫")
	}
}

func drawLine(content string, width: int) {
	fmt.Print("┃")
	fmt.Print(content)
	fmt.Print(strings.Repeat("━", width - len(content)))
	fmt.Println("┃")
}
