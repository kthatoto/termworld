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
		var response interface{}
		response = nil
		err = client.Call(proceduresCommand, playerProcedureArgs, &response)

		if err != nil {
			if strings.Contains(err.Error(), "can't find method") {
				fmt.Printf("error: can't find method [%s]\n", command)
				return nil
			}
			fmt.Printf("error: %s\n", err)
			return nil
		}
		fmt.Println(response)
		return nil
	},
}
