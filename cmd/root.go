package cmd

import (
	"fmt"
	"os"

	"github.com/applandinc/appland-cli/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "appland",
	Short: "Manage AppLand resources",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func fail(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func Execute() {
	config.LoadConfig()

	if err := rootCmd.Execute(); err != nil {
		fail(err)
	}

	if err := config.WriteConfig(); err != nil {
		fail(err)
	}
}