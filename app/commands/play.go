package commands

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"net/rpc"

	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(playCommand)
}

var playCommand = &cobra.Command{
	Use: "play",
	Short: "Play command",
	RunE: func(cmd *cobra.Command, args []string) error {
		if (len(args) == 0) {
			fmt.Println("player name is required")
			return nil
		}

		playerName := args[0]
		options := args[1:]
		client, err := rpc.DialHTTP("tcp", "localhost:8128")
		if err != nil {
			return err
		}

		var response interface{}
		response = nil
		gameProcedureArgs := GameProcedureArgs{
			Command: "touch",
			PlayerName: playerName,
			Options: options,
		}
		err = client.Call("GameProcedures.Execute", gameProcedureArgs, &response)
		if err != nil {
			return err
		}
		if response != "live" {
			fmt.Println(response)
			return nil
		}

		stdin := bufio.NewScanner(os.Stdin)
		for {
			stdin.Scan()
			text := stdin.Text()
			input := strings.Split(text, " ")
			command := input[0]
			options := input[1:]
			if command == "q" {
				break
			}

			gameProcedureArgs = GameProcedureArgs{ command, playerName, options }
			err = client.Call("GameProcedures.Execute", gameProcedureArgs, &response)
			if err != nil {
				break
			}
			fmt.Println(response)
		}

		if err != nil {
			return err
		}
		return nil
	},
}
