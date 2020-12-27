package config

import (
	"fmt"
	"github.com/graphikDB/graphikctl/version"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	Config.AddCommand(getCmd, openCmd)
}

var Config = &cobra.Command{
	Use:     "config",
	Short:   "configuration subcommands (get)",
	Version: version.Version,
}

var getCmd = &cobra.Command{
	Use:     "get",
	Short:   fmt.Sprintf("get a value from the registered config: %s", viper.ConfigFileUsed()),
	Version: version.Version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString(args[len(args)-1]))
	},
}

var openCmd = &cobra.Command{
	Use:     "open",
	Short:   fmt.Sprintf("open registered config: %s", viper.ConfigFileUsed()),
	Version: version.Version,
	Run: func(cmd *cobra.Command, args []string) {
		file := viper.ConfigFileUsed()
		if err := open.Start(file); err != nil {
			fmt.Printf("failed to open %s: %s", file, err)
		}
		return
	},
}
