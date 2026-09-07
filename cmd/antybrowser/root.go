package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var cfgFile string
var apiURL string
var apiKey string
var outputFormat string

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "antybrowser",
		Short:   "CLI for Antybrowser",
		Long:    "Command-line interface for managing Antybrowser profiles, proxies, and automations.",
		Version: version,
		SilenceUsage: true,
		SilenceErrors: true,
	}

	root.PersistentFlags().StringVar(&apiURL, "api-url", "", "Antybrowser local API URL (overrides config)")
	root.PersistentFlags().StringVar(&apiKey, "api-key", "", "API key (overrides config)")
	root.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "Output format: table, json, yaml")

	root.AddCommand(
		newStatusCmd(),
		newProfilesCmd(),
		newProxiesCmd(),
		newExtensionsCmd(),
		newGroupsCmd(),
		newAutomationsCmd(),
		newConfigCmd(),
		newVersionCmd(),
	)

	return root
}

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version info",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("antybrowser-cli %s %s/%s\n", version, runtime.GOOS, runtime.GOARCH)
			fmt.Printf("Config: %s\n", getConfigFilePath())
			return nil
		},
	}
}
