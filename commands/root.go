package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	homedir "github.com/mitchellh/go-homedir"
)

var rootCommand = &cobra.Command{
	Use: "termworld",
	Short: "Welcome to termworld!",
	Long: "Welcome to termworld! Long ver",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Usage()
		return nil
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		er(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		er(err)
	}

	viper.AddConfigPath(fmt.Sprintf("%s/.termworld", home))
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.SetDefault("token", "")
	if err := viper.ReadInConfig(); err != nil {
		os.Mkdir(fmt.Sprintf("%s/.termworld", home), 0755)
		if err = viper.SafeWriteConfigAs(fmt.Sprintf("%s/.termworld/config", home)); err != nil {
			er(err)
		}
		if err = viper.ReadInConfig(); err != nil {
			er(err)
		}
	}
}

func er(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
