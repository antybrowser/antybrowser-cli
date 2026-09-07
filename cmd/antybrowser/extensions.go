package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

type Extension struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

func newExtensionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extensions",
		Short: "Manage extensions",
	}

	cmd.AddCommand(
		newExtensionsListCmd(),
		newExtensionsDeleteCmd(),
	)

	return cmd
}

func newExtensionsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all extensions",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := newAPIClient()
			data, err := client.get("/api/extensions")
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				fmt.Println(string(data))
				return nil
			}

			var extensions []Extension
			if err := json.Unmarshal(data, &extensions); err != nil {
				return err
			}

			if len(extensions) == 0 {
				fmt.Println("No extensions found.")
				return nil
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tNAME\tTYPE\tVERSION")
			for _, e := range extensions {
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", e.ID, e.Name, e.Type, e.Version)
			}
			w.Flush()
			return nil
		},
	}
}

func newExtensionsDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete an extension",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := newAPIClient()
			_, err := client.delete("/api/extensions/" + args[0])
			if err != nil {
				return err
			}
			fmt.Println("Extension deleted.")
			return nil
		},
	}
}
