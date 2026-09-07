package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

type Automation struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	LastRun     string `json:"lastRun"`
	CreatedAt   string `json:"createdAt"`
}

func newAutomationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "automations",
		Short: "Manage automations",
	}

	cmd.AddCommand(
		newAutomationsListCmd(),
		newAutomationsRunCmd(),
	)

	return cmd
}

func newAutomationsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all automations",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := newAPIClient()
			data, err := client.get("/api/automations")
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				fmt.Println(string(data))
				return nil
			}

			var automations []Automation
			if err := json.Unmarshal(data, &automations); err != nil {
				return err
			}

			if len(automations) == 0 {
				fmt.Println("No automations found.")
				return nil
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tNAME\tSTATUS\tLAST RUN")
			for _, a := range automations {
				lastRun := a.LastRun
				if lastRun == "" {
					lastRun = "never"
				}
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", a.ID, a.Name, a.Status, lastRun)
			}
			w.Flush()
			return nil
		},
	}
}

func newAutomationsRunCmd() *cobra.Command {
	var profileID int
	var deleteCookies bool

	cmd := &cobra.Command{
		Use:   "run <automation-id>",
		Short: "Run an automation",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if profileID == 0 {
				return fmt.Errorf("--profile-id is required")
			}

			body := map[string]interface{}{
				"profileId": profileID,
			}
			if deleteCookies {
				body["deleteCookies"] = true
			}

			client := newAPIClient()
			data, err := client.post("/api/automations/"+args[0]+"/run", body)
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				fmt.Println(string(data))
			} else {
				fmt.Println("Automation started.")
			}
			return nil
		},
	}

	cmd.Flags().IntVar(&profileID, "profile-id", 0, "Profile ID to run automation on (required)")
	cmd.Flags().BoolVar(&deleteCookies, "delete-cookies", false, "Delete cookies before running")

	return cmd
}

func init() {
	_ = strconv.Itoa
}
