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
			fmt.Println("Not working")
			return nil
		}

		var result bool
		result = false
		err = client.Call("Procedures.Stop", 0, &result)
		if err != nil {
			return err
		}

		fmt.Println("Stopped")
		return nil
	},
}
