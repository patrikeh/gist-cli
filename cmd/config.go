package cmd

import (
	"fmt"
	"os"
	"path"

	githubclient "github.com/patrikeh/gist/githubClient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config [set-url|set-token]",
	Short: "Set persistent configuration.",
	Long:  "Set persistent configuration.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s\n", viper.ConfigFileUsed())
		for k, v := range viper.AllSettings() {
			fmt.Printf("%s: %s\n", k, v)
		}
		return nil
	},
}

var setUrlCmd = &cobra.Command{
	Args:     cobra.ExactArgs(1),
	Use:      "set-url github-url",
	Short:    "Sets github URL in config.",
	Long:     "Sets github URL in config.",
	PostRunE: writeConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Expected a single argument 'url'")
		}
		viper.Set("url", args[0])
		return nil
	},
}

var setTokenCmd = &cobra.Command{
	Args:     cobra.ExactArgs(1),
	Use:      "set-token access-token",
	Short:    "Sets github token.",
	Long:     "Sets github token. Can be generated at https://github.com/settings/tokens with gist privileges.",
	PostRunE: writeConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Expected a single argument 'access-token'")
		}
		viper.Set("access-token", args[0])
		return nil
	},
}

func init() {
	configCmd.AddCommand(setTokenCmd, setUrlCmd)
}

func writeConfig(cmd *cobra.Command, args []string) error {
	if _, err := os.Stat(viper.ConfigFileUsed()); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(viper.ConfigFileUsed()), os.ModePerm)
	}

	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}
	return nil
}

func getToken() (string, error) {
	apiToken := viper.GetString("access-token")
	if len(apiToken) == 0 {
		return "", fmt.Errorf("Missing GitHub access token. Use gist set-token [access-token] in order to add a token.")
	}
	return apiToken, nil
}

func getUrl() (string, error) {
	url := viper.GetString("url")
	if len(url) == 0 {
		return "https://api.github.com/", nil
	}
	return url, nil
}

func getGithubClient() (*githubclient.Client, error) {
	token, err := getToken()
	if err != nil {
		return nil, err
	}
	url, err := getUrl()
	if err != nil {
		return nil, err
	}

	return githubclient.New(token, url)
}
