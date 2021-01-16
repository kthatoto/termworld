package commands

import (
	"fmt"
	"strings"
	"net/rpc"

	"github.com/spf13/cobra"
)

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
		// options := args[2:]
		client, err := rpc.DialHTTP("tcp", "localhost:8128")
		if err != nil {
			return err
		}
		proceduresCommand := fmt.Sprintf("Procedures.%s", strings.Title(command))
		var result *bool
		err = client.Call(proceduresCommand, playerName, result)
		if (!*result) {
			fmt.Println(err)
			return err
		}
		return nil
	},
}
