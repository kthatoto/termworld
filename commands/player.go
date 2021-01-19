package commands

import (
	"fmt"
	"strings"
	"net/rpc"

	"github.com/spf13/cobra"
)

type PlayerProcedureArgs struct {
	PlayerName string
	Options    []string
}

func init() {
	rootCommand.AddCommand(playerCommand)
}

var playerCommand = &cobra.Command{
	Use: "player",
	Short: "Player command",
	RunE: func(cmd *cobra.Command, args []string) error {
		if (len(args) == 0) {
			cmd.Usage()
			return nil
		}

		playerName := args[0]
		command := args[1]
		options := args[2:]
		client, err := rpc.DialHTTP("tcp", "localhost:8128")
		if err != nil {
			return err
		}
		proceduresCommand := fmt.Sprintf("PlayerProcedures.%s", strings.Title(command))

		playerProcedureArgs := PlayerProcedureArgs{ playerName, options }
		var result bool
		result = false
		err = client.Call(proceduresCommand, playerProcedureArgs, &result)
		if err != nil || !result {
			fmt.Println(err)
			return err
		}
		return nil
	},
}
