package cmd

import (
	"fmt"
	"rd-cli/api"
	"rd-cli/config"

	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download <url>",
	Short: "Unrestrict and download a file via Real-Debrid",
	Long: `Unrestrict a link from a supported file hoster and download it at full speed using your Real-Debrid account.

Examples:
  rd-cli download https://rapidgator.net/file/example
  rd-cli download https://1fichier.com/?abc123`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cobra.CheckErr("No link passed")
		}

		link := args[0]

		config, err := config.GetConfig()

		if err != nil {
			cobra.CheckErr(err)
		}

		apiClient := api.NewClient(config.AccessToken)

		downloadLink, err := apiClient.UnrestrictLink(link)

		if err != nil {
			cobra.CheckErr(err)
		}

		fmt.Println(downloadLink)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
