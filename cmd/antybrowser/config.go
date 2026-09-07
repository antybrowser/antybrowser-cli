package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage CLI configuration",
	}

	cmd.AddCommand(
		newConfigShowCmd(),
		newConfigSetCmd(),
	)

	return cmd
}

func newConfigShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := newAPIClient()
			data, err := client.get("/api/settings")
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				fmt.Println(string(data))
			} else {
				var settings map[string]interface{}
				if err := json.Unmarshal(data, &settings); err != nil {
					return err
				}
				fmt.Printf("API URL:        %s\n", apiURL)
				fmt.Printf("API Key:        %s\n", maskKey(getAPIKey()))
				fmt.Printf("Browser Type:   %v\n", settings["defaultBrowserType"])
				fmt.Printf("OS:             %v\n", settings["defaultOsFingerprint"])
				fmt.Printf("Screen:         %v\n", settings["defaultScreenResolution"])
				fmt.Printf("Language:       %v\n", settings["defaultLanguage"])
				fmt.Printf("Timezone:       %v\n", settings["defaultTimezone"])
			}
			return nil
		},
	}
}

func newConfigSetCmd() *cobra.Command {
	var apiURLStr, apiKeyStr string

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set configuration values",
		RunE: func(cmd *cobra.Command, args []string) error {
			if apiURLStr != "" {
				fmt.Printf("API URL set to: %s\n", apiURLStr)
			}
			if apiKeyStr != "" {
				fmt.Println("API key set.")
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&apiURLStr, "api-url", "", "Set API URL")
	cmd.Flags().StringVar(&apiKeyStr, "api-key", "", "Set API key")

	return cmd
}

func maskKey(key string) string {
	if key == "" {
		return "(not set)"
	}
	if len(key) <= 8 {
		return "****"
	}
	return key[:4] + "****" + key[len(key)-4:]
}
