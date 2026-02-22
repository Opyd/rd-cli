package cmd

import (
	"bufio"
	"fmt"
	"os"
	"rd-cli/config"
	"strings"

	"github.com/spf13/cobra"
)

var setTokenCmd = &cobra.Command{
	Use:   "set-token [token]",
	Short: "Save your Real-Debrid API token",
	Long: `Save your Real-Debrid API token to the local configuration file.

If a token is provided as an argument it will be saved directly.
If no argument is given, you will be prompted to enter the token interactively.

Examples:
  rd-cli set-token                  # interactive prompt
  rd-cli set-token YOUR_TOKEN_HERE  # pass token directly`,
	Run: func(cmd *cobra.Command, args []string) {
		token := ""
		var err error

		if len(args) > 0 {
			token = args[0]
		}

		if token == "" {
			token, err = promptForToken()

			if err != nil {
				cobra.CheckErr(err)
			}
		}

		err = config.SetToken(token)
		if err != nil {
			cobra.CheckErr(err)
		}

		fmt.Println("Successfully set up the access token")

	},
}

func promptForToken() (string, error) {
	fmt.Println("Enter RealDebrid token:")

	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	line = strings.TrimSpace(line)

	return line, nil
}

func init() {
	rootCmd.AddCommand(setTokenCmd)
}
