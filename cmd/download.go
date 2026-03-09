package cmd

import (
	"bufio"
	"fmt"
	"os"
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
  rd-cli download https://1fichier.com/?abc123 -n movie.mkv
  rd-cli download https://rapidgator.net/file/example -p ~/Downloads
  rd-cli download -f links.txt -p ~/Downloads`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var hadErrors bool

		path, _ := cmd.Flags().GetString("path")

		file, _ := cmd.Flags().GetString("file")

		if file != "" {
			hadErrors, err = processFile(file, path)

			if hadErrors {
				os.Exit(1)
			}

			return err
		}

		if len(args) == 0 {
			cobra.CheckErr("No link passed")
		}

		link := args[0]

		name, _ := cmd.Flags().GetString("name")

		err = processLink(link, path, name)

		return err
	},
}

func processFile(inputFilePath string, path string) (bool, error) {
	inputFile, err := os.Open(inputFilePath)

	hasErrors := false

	if err != nil {
		return false, err
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		rowLink := scanner.Text()
		err := processLink(rowLink, path, "")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to download: %s %v\n", rowLink, err)
			hasErrors = true
		}
	}

	return hasErrors, nil
}

func processLink(link string, path string, name string) error {
	config, err := config.GetConfig()

	if err != nil {
		return err
	}

	apiClient := api.NewClient(config.AccessToken)

	downloadLink, err := apiClient.UnrestrictLink(link)

	if err != nil {
		return err
	}

	if name == "" {
		name = downloadLink.Filename
	}

	downloader := downloader.NewDownloader()

	err = downloader.Download(downloadLink.Download, path, name)

	return err
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringP("name", "n", "", "Custom filename")
	downloadCmd.Flags().StringP("path", "p", ".", "Custom download path")
	downloadCmd.Flags().StringP("file", "f", "", "File with links to download")
}
