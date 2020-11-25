package commands

import (
	"os"

	"github.com/spf13/cobra"
	homedir "github.com/mitchellh/go-homedir"
	daemon "github.com/sevlyar/go-daemon"
)

func init() {
	rootCommand.AddCommand(stopCommand)
}

var stopCommand = &cobra.Command{
	Use: "stop",
	Short: "Game stop command",
	RunE: func(cmd *cobra.Command, args []string) error {
		home, err := homedir.Dir()
		if err != nil {
			return err
		}
		if err := os.Chdir(home+"/.termworld"); err != nil {
			return err
		}

		ctx := &daemon.Context{
			PidFileName: "termworld.pid",
			PidFilePerm: 0644,
			LogFileName: "termworld.log",
			LogFilePerm: 0640,
			WorkDir:     "./",
			Umask:       027,
		}
		ctx.Release()
		return nil
	},
}
