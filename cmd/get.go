package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// -o --output filePath?

var getGistCmd = &cobra.Command{
	Args:    cobra.ExactArgs(1),
	Use:     "get ID",
	Aliases: []string{"g"},
	Short:   "Gets a gist by ID.",
	Long:    "Gets a gist by ID.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 || args[0] == "" {
			return fmt.Errorf("Expect single argument ID")
		}

		client, err := getGithubClient()
		if err != nil {
			return fmt.Errorf("error initializing github client: %w", err)
		}
		gist, err := client.GetGist(cmd.Context(), args[0])
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", *gist.ID)
		for name, file := range gist.GetFiles() {
			fmt.Printf("--- %s ---\n%s\n", name, *file.Content)
		}
		return nil
	},
}
