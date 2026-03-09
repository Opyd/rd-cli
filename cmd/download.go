package cmd

import (
	"rd-cli/api"
	"rd-cli/config"
	"rd-cli/downloader"

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

		name, _ := cmd.Flags().GetString("name")

		config, err := config.GetConfig()

		if err != nil {
			cobra.CheckErr(err)
		}

		apiClient := api.NewClient(config.AccessToken)

		downloadLink, err := apiClient.UnrestrictLink(link)

		if err != nil {
			cobra.CheckErr(err)
		}

		if name == "" {
			name = downloadLink.Filename
		}

		path, _ := cmd.Flags().GetString("path")

		downloader := downloader.NewDownloader()

		err = downloader.Download(downloadLink.Download, path, name)

		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringP("name", "n", "", "Custom filename")
	downloadCmd.Flags().StringP("path", "p", ".", "Custom download path")
}
