package commands

import (
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
		cmd.Usage()
		return nil
	},
}
