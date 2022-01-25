/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var RootCmd = &cobra.Command{
	Version: "1.0.0",
	Use:     "gist [fileNames]",
	Short:   "CLI for interacting with GitHub gists.",
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		log.Fatalf("error executing command: %s", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/gist/config.yaml)")

	RootCmd.AddCommand(configCmd,
		createGistCmd,
		deleteGistCmd,
		getGistCmd,
		listGistCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.SetConfigFile(path.Join(home, ".config", "gist", "config.yaml"))
	}
	viper.BindEnv("host", "GH_HOST")
	viper.BindEnv("access-token", "GH_TOKEN")
	viper.ReadInConfig()
}
