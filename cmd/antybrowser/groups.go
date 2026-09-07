package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

type Group struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

func newGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "groups",
		Short: "Manage profile groups",
	}

	cmd.AddCommand(
		newGroupsListCmd(),
		newGroupsCreateCmd(),
		newGroupsUpdateCmd(),
		newGroupsDeleteCmd(),
	)

	return cmd
}

func newGroupsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all groups",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := newAPIClient()
			data, err := client.get("/api/groups")
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				fmt.Println(string(data))
				return nil
			}

			var groups []Group
			if err := json.Unmarshal(data, &groups); err != nil {
				return err
			}

			if len(groups) == 0 {
				fmt.Println("No groups found.")
				return nil
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tNAME\tCOLOR\tDESCRIPTION")
			for _, g := range groups {
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", g.ID, g.Name, g.Color, g.Description)
			}
			w.Flush()
			return nil
		},
	}
}

func newGroupsCreateCmd() *cobra.Command {
	var name, description, color string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if name == "" {
				return fmt.Errorf("--name is required")
			}

			body := map[string]interface{}{
				"name": name,
			}
			if description != "" {
				body["description"] = description
			}
			if color != "" {
				body["color"] = color
			}

			client := newAPIClient()
			data, err := client.post("/api/groups", body)
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				fmt.Println(string(data))
			} else {
				var g Group
				if err := json.Unmarshal(data, &g); err != nil {
					return err
				}
				fmt.Printf("Group created: %s (ID: %d)\n", g.Name, g.ID)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Group name (required)")
	cmd.Flags().StringVar(&description, "description", "", "Group description")
	cmd.Flags().StringVar(&color, "color", "", "Group color")

	return cmd
}

func newGroupsUpdateCmd() *cobra.Command {
	var name, description, color string

	cmd := &cobra.Command{
		Use:   "update <id>",
		Short: "Update a group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			body := map[string]interface{}{}
			if name != "" {
				body["name"] = name
			}
			if description != "" {
				body["description"] = description
			}
			if color != "" {
				body["color"] = color
			}

			if len(body) == 0 {
				return fmt.Errorf("at least one flag is required")
			}

			client := newAPIClient()
			_, err := client.put("/api/groups/"+args[0], body)
			if err != nil {
				return err
			}
			fmt.Println("Group updated.")
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "New name")
	cmd.Flags().StringVar(&description, "description", "", "New description")
	cmd.Flags().StringVar(&color, "color", "", "New color")

	return cmd
}

func newGroupsDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete a group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := newAPIClient()
			_, err := client.delete("/api/groups/" + args[0])
			if err != nil {
				return err
			}
			fmt.Println("Group deleted.")
			return nil
		},
	}
}

func init() {
	_ = strconv.Itoa
}
