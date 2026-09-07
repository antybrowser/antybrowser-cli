package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Check connection to Antybrowser",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := newAPIClient()
			data, err := client.get("/api/status")
			if err != nil {
				return fmt.Errorf("cannot connect to Antybrowser: %w\n\nMake sure the Antybrowser desktop app is running.", err)
			}
			if outputFormat == "json" {
				fmt.Println(string(data))
			} else {
				fmt.Println("Connected to Antybrowser")
			}
			return nil
		},
	}
}
