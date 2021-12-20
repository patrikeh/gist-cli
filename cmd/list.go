package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listGistCmd = &cobra.Command{
	Args:    cobra.MaximumNArgs(1),
	Use:     "list [username]",
	Aliases: []string{"l"},
	Short:   "Lists gists for a user.",
	Long: `Lists gists for a user. 
	
In case no username is specified, lists gists for the current authenticated user.`,
}

func init() {
	listGistCmd.RunE = func(cmd *cobra.Command, args []string) error {
		username := ""
		if len(args) > 0 {
			username = args[0]
		}

		client, err := getGithubClient()
		if err != nil {
			return fmt.Errorf("error initializing github client: %w", err)
		}

		gists, err := client.ListGists(cmd.Context(), username)
		if err != nil {
			return fmt.Errorf("error listing gists: %w", err)
		}

		for _, gist := range gists {
			fmt.Printf("%s\n%s\n", *gist.ID, *gist.HTMLURL)
			for filename := range gist.Files {
				fmt.Printf(" * %s\n", filename)
			}
		}

		return nil
	}
}
