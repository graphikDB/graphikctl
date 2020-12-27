package auth

import (
	"fmt"
	"github.com/graphikDB/graphikctl/version"
	"github.com/spf13/cobra"
)

func init() {
	Auth.AddCommand(login)
}

var Auth = &cobra.Command{
	Use: "auth",
	Short: "authentication/authorization subcommands",
	Version: version.Version,
}

var login = &cobra.Command{
	Use: "login",
	Short: "launch a login flow to an identity provider",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unimplemented")
	},
	Version: version.Version,
}