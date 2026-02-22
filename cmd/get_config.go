package cmd

import (
	"fmt"
	"rd-cli/config"

	"github.com/spf13/cobra"
)

var getConfigCmd = &cobra.Command{
	Use:   "get-config",
	Short: "Display current configuration",
	Long:  "Reads and displays the current rd-cli configuration from the user config directory.",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.GetConfig()

		if err != nil {
			cobra.CheckErr(err)
		}

		fmt.Printf("%+v\n", config)
	},
}

func init() {
	rootCmd.AddCommand(getConfigCmd)
}
