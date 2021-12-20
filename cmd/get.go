package cmd

import "github.com/spf13/cobra"

var getGistCmd = &cobra.Command{
	Args:    cobra.ExactArgs(1),
	Use:     "get ID",
	Aliases: []string{"g"},
	Short:   "Gets a gist by ID.",
	Long:    "Gets a gist by ID.",
}
