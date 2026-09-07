package main

import (
	"fmt"
	"os"
	"path/filepath"
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

	root.PersistentFlags().StringVar(&apiURL, "api-url", getDefaultAPIURL(), "Antybrowser local API URL")
	root.PersistentFlags().StringVar(&apiKey, "api-key", "", "API key (or set ANTYBROWSER_API_KEY)")
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

func getDefaultAPIURL() string {
	if v := os.Getenv("ANTYBROWSER_API_URL"); v != "" {
		return v
	}
	return "http://127.0.0.1:5173"
}

func getAPIKey() string {
	if apiKey != "" {
		return apiKey
	}
	if v := os.Getenv("ANTYBROWSER_API_KEY"); v != "" {
		return v
	}
	cfgPath := getConfigPath()
	if cfgPath == "" {
		return ""
	}
	data, err := os.ReadFile(filepath.Join(cfgPath, "config.json"))
	if err != nil {
		return ""
	}
	_ = data
	return ""
}

func getConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "Antybrowser", "CLI")
	case "darwin":
		return filepath.Join(home, ".config", "antybrowser")
	default:
		return filepath.Join(home, ".config", "antybrowser")
	}
}

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version info",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("antybrowser-cli %s %s/%s\n", version, runtime.GOOS, runtime.GOARCH)
			return nil
		},
	}
}
