package cmd

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rd-cli",
	Short: "Download files via Real-Debrid",
	Long: `rd-cli is a command-line tool for downloading files through the Real-Debrid service.

It unrestricts links from supported file hosters (e.g. Rapidgator, 1Fichier)
and downloads them at full speed using your Real-Debrid account.

Examples:
  rd-cli set-token                               # configure your API token interactively
  rd-cli download <url>                          # unrestrict and download a file
  rd-cli download <url> -p ~/Downloads           # download to specific directory
  rd-cli download <url> -n movie.mkv             # custom filename
  rd-cli download -f links.txt -p ~/Downloads    # batch download from file`,
}

func Execute() {
	err := fang.Execute(context.Background(), rootCmd)
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
