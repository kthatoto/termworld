package commands

import (
	"fmt"
	"net/rpc"

	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(stopCommand)
}

var stopCommand = &cobra.Command{
	Use: "stop",
	Short: "Game stop command",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := rpc.DialHTTP("tcp", "localhost:8128")
		if err != nil {
			return err
		}

		var result *bool
		err = client.Call("Procedures.Stop", 0, result)
		if err != nil {
			fmt.Printf("ERROR: %+v\n", err)
		}
		return nil
	},
}
