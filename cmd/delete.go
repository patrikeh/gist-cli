package cmd

import (
	"fmt"

	githubclient "github.com/patrikeh/gist/githubClient"
	"github.com/spf13/cobra"
)

var deleteGistCmd = &cobra.Command{
	Args:    cobra.ExactArgs(1),
	Use:     "delete ID",
	Aliases: []string{"d"},
	Short:   "Deletes a gist by ID.",
	Long:    "Deletes a gist by ID.",
}

func init() {
	deleteGistCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 || args[0] == "" {
			return fmt.Errorf("Expect single argument [ID]")
		}

		token, err := getToken()
		if err != nil {
			return err
		}

		client := githubclient.New(token)

		if err := client.DeleteGist(cmd.Context(), args[0]); err != nil {
			return fmt.Errorf("error deleting gist %s: %w", args[0], err)
		}
		fmt.Printf("Deleted gist %s.\n", args[0])
		return nil
	}
}
