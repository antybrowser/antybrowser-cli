package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

type CLIConfig struct {
	APIURL string `json:"api_url"`
	APIKey string `json:"api_key"`
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

func getConfigFilePath() string {
	return filepath.Join(getConfigPath(), "config.json")
}

func loadConfig() CLIConfig {
	var cfg CLIConfig
	data, err := os.ReadFile(getConfigFilePath())
	if err != nil {
		return cfg
	}
	_ = json.Unmarshal(data, &cfg)
	return cfg
}

func saveConfig(cfg CLIConfig) error {
	dir := getConfigPath()
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("create config directory: %w", err)
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}
	return os.WriteFile(getConfigFilePath(), data, 0600)
}

func getAPIKey() string {
	if apiKey != "" {
		return apiKey
	}
	if v := os.Getenv("ANTYBROWSER_API_KEY"); v != "" {
		return v
	}
	cfg := loadConfig()
	return cfg.APIKey
}

func getDefaultAPIURL() string {
	if v := os.Getenv("ANTYBROWSER_API_URL"); v != "" {
		return v
	}
	cfg := loadConfig()
	if cfg.APIURL != "" {
		return cfg.APIURL
	}
	return "http://127.0.0.1:5173"
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

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage CLI configuration",
	}

	cmd.AddCommand(
		newConfigShowCmd(),
		newConfigSetCmd(),
		newConfigResetCmd(),
	)

	return cmd
}

func newConfigShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("CLI Configuration")
			fmt.Println("=================")
			fmt.Printf("Config file:   %s\n", getConfigFilePath())
			fmt.Printf("API URL:       %s\n", getResolvedAPIURL())
			fmt.Printf("API Key:       %s\n", maskKey(getResolvedAPIKey()))
			fmt.Println()

			client := newAPIClient()
			data, err := client.get("/api/settings")
			if err != nil {
				fmt.Printf("Cannot connect to Antybrowser: %v\n", err)
				fmt.Println("Start the desktop app to see server settings.")
				return nil
			}

			var settings map[string]interface{}
			if err := json.Unmarshal(data, &settings); err != nil {
				return err
			}

			fmt.Println("Antybrowser Settings")
			fmt.Println("====================")
			fmt.Printf("Browser Type:   %v\n", settings["defaultBrowserType"])
			fmt.Printf("OS:             %v\n", settings["defaultOsFingerprint"])
			fmt.Printf("Screen:         %v\n", settings["defaultScreenResolution"])
			fmt.Printf("Language:       %v\n", settings["defaultLanguage"])
			fmt.Printf("Timezone:       %v\n", settings["defaultTimezone"])
			return nil
		},
	}
}

func newConfigSetCmd() *cobra.Command {
	var urlStr, keyStr string

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set and save configuration values",
		Example: `  antybrowser config set --api-url http://192.168.1.100:5173
  antybrowser config set --api-key abc123
  antybrowser config set --api-url http://custom:9999 --api-key mykey`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := loadConfig()
			changed := false

			if urlStr != "" {
				cfg.APIURL = urlStr
				changed = true
				fmt.Printf("API URL set to: %s\n", urlStr)
			}
			if keyStr != "" {
				cfg.APIKey = keyStr
				changed = true
				fmt.Printf("API key set to: %s\n", maskKey(keyStr))
			}

			if !changed {
				return fmt.Errorf("specify --api-url or --api-key")
			}

			if err := saveConfig(cfg); err != nil {
				return fmt.Errorf("save config: %w", err)
			}
			fmt.Printf("Config saved to: %s\n", getConfigFilePath())
			return nil
		},
	}

	cmd.Flags().StringVar(&urlStr, "api-url", "", "API URL (e.g. http://192.168.1.100:5173)")
	cmd.Flags().StringVar(&keyStr, "api-key", "", "API key")

	return cmd
}

func newConfigResetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "Reset configuration to defaults",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := CLIConfig{}
			if err := saveConfig(cfg); err != nil {
				return fmt.Errorf("save config: %w", err)
			}
			fmt.Println("Configuration reset to defaults.")
			return nil
		},
	}
}

func getResolvedAPIURL() string {
	if apiURL != "" && apiURL != "http://127.0.0.1:5173" {
		return apiURL
	}
	return getDefaultAPIURL()
}

func getResolvedAPIKey() string {
	return getAPIKey()
}
