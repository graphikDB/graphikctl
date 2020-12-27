package cmd

import (
	"fmt"
	"github.com/graphikDB/graphikctl/cmd/auth"
	"github.com/graphikDB/graphikctl/cmd/config"
	"github.com/graphikDB/graphikctl/version"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "graphikctl",
	Long: `
A command line utility for graphikDB

---
env-prefix: GRAPHIKCTL
`,
	Version: version.Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.graphikctl.yaml)")
	rootCmd.AddCommand(auth.Auth, config.Config)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.SetEnvPrefix("GRAPHIKCTL")
		// Search config in home directory with name ".graphikctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".graphikctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	viper.SetDefault("auth.scopes", []string{"openid", "email", "profile"})
	viper.SetDefault("auth.redirect", "http://localhost:8080/")
	viper.SetDefault("server.port", ":8080")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
