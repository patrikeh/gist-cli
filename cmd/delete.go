package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteGistCmd = &cobra.Command{
	Args:    cobra.MinimumNArgs(1),
	Use:     "delete ID...",
	Aliases: []string{"d"},
	Short:   "Deletes gists by ID.",
	Long: `Deletes gists by ID.

Expects at least one ID. If multiple are present, deletes all specified gists.`,
}

func init() {
	deleteGistCmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := getGithubClient()
		if err != nil {
			return fmt.Errorf("error initializing github client: %w", err)
		}
		for _, gistId := range args {
			if err := client.DeleteGist(cmd.Context(), gistId); err != nil {
				return fmt.Errorf("error deleting gist %s: %w", gistId, err)
			}
			fmt.Printf("Deleted gist %s.\n", gistId)
		}
		return nil
	}
}
