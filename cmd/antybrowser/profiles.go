package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

type Profile struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	BrowserType string `json:"browserType"`
	OS          string `json:"osFingerprint"`
	Status      string `json:"status"`
	ProxyID     *int   `json:"proxyId"`
	GroupID     *int   `json:"groupId"`
	Notes       string `json:"notes"`
	CreatedAt   string `json:"createdAt"`
}

func newProfilesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profiles",
		Short: "Manage browser profiles",
	}

	cmd.AddCommand(
		newProfilesListCmd(),
		newProfilesGetCmd(),
		newProfilesCreateCmd(),
		newProfilesUpdateCmd(),
		newProfilesDeleteCmd(),
		newProfilesStartCmd(),
		newProfilesStopCmd(),
		newProfilesDuplicateCmd(),
	)

	return cmd
}

func newProfilesListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := newAPIClient()
			data, err := client.get("/api/profiles")
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				fmt.Println(string(data))
				return nil
			}

			var profiles []Profile
			if err := json.Unmarshal(data, &profiles); err != nil {
				return err
			}

			if len(profiles) == 0 {
				fmt.Println("No profiles found.")
				return nil
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tNAME\tBROWSER\tOS\tSTATUS")
			for _, p := range profiles {
				status := p.Status
				if status == "" {
					status = "Ready"
				}
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", p.ID, p.Name, p.BrowserType, p.OS, status)
			}
			w.Flush()
			return nil
		},
	}
}

func newProfilesGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <id>",
		Short: "Get profile details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := newAPIClient()
			data, err := client.get("/api/profiles/" + args[0])
			if err != nil {
				return err
			}
			if outputFormat == "json" {
				fmt.Println(string(data))
			} else {
				var p Profile
				if err := json.Unmarshal(data, &p); err != nil {
					return err
				}
				fmt.Printf("ID:        %d\n", p.ID)
				fmt.Printf("Name:      %s\n", p.Name)
				fmt.Printf("Browser:   %s\n", p.BrowserType)
				fmt.Printf("OS:        %s\n", p.OS)
				fmt.Printf("Status:    %s\n", p.Status)
				fmt.Printf("Notes:     %s\n", p.Notes)
				fmt.Printf("Created:   %s\n", p.CreatedAt)
			}
			return nil
		},
	}
}

func newProfilesCreateCmd() *cobra.Command {
	var name, browserType, osFingerprint, notes string
	var proxyID, groupID int

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			if name == "" {
				return fmt.Errorf("--name is required")
			}

			body := map[string]interface{}{
				"name": name,
			}
			if browserType != "" {
				body["browserType"] = browserType
			}
			if osFingerprint != "" {
				body["osFingerprint"] = osFingerprint
			}
			if notes != "" {
				body["notes"] = notes
			}
			if proxyID > 0 {
				body["proxyId"] = proxyID
			}
			if groupID > 0 {
				body["groupId"] = groupID
			}

			client := newAPIClient()
			data, err := client.post("/api/profiles", body)
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				fmt.Println(string(data))
			} else {
				var p Profile
				if err := json.Unmarshal(data, &p); err != nil {
					return err
				}
				fmt.Printf("Profile created: %s (ID: %d)\n", p.Name, p.ID)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Profile name (required)")
	cmd.Flags().StringVar(&browserType, "browser", "Chrome", "Browser type (Chrome, Edge, Chromium)")
	cmd.Flags().StringVar(&osFingerprint, "os", "Windows", "OS fingerprint (Windows, MacOS, Linux)")
	cmd.Flags().StringVar(&notes, "notes", "", "Profile notes")
	cmd.Flags().IntVar(&proxyID, "proxy-id", 0, "Proxy ID to assign")
	cmd.Flags().IntVar(&groupID, "group-id", 0, "Group ID to assign")

	return cmd
}

func newProfilesUpdateCmd() *cobra.Command {
	var name, browserType, osFingerprint, notes string

	cmd := &cobra.Command{
		Use:   "update <id>",
		Short: "Update a profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			body := map[string]interface{}{}
			if name != "" {
				body["name"] = name
			}
			if browserType != "" {
				body["browserType"] = browserType
			}
			if osFingerprint != "" {
				body["osFingerprint"] = osFingerprint
			}
			if notes != "" {
				body["notes"] = notes
			}

			if len(body) == 0 {
				return fmt.Errorf("at least one flag is required")
			}

			client := newAPIClient()
			data, err := client.put("/api/profiles/"+args[0], body)
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				fmt.Println(string(data))
			} else {
				fmt.Println("Profile updated.")
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "New name")
	cmd.Flags().StringVar(&browserType, "browser", "", "New browser type")
	cmd.Flags().StringVar(&osFingerprint, "os", "", "New OS fingerprint")
	cmd.Flags().StringVar(&notes, "notes", "", "New notes")

	return cmd
}

func newProfilesDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete a profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := newAPIClient()
			_, err := client.delete("/api/profiles/" + args[0])
			if err != nil {
				return err
			}
			fmt.Println("Profile deleted.")
			return nil
		},
	}
}

func newProfilesStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start <id>",
		Short: "Start a browser profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid profile ID: %s", args[0])
			}
			client := newAPIClient()
			_, err = client.post(fmt.Sprintf("/api/profiles/%d/start", id), nil)
			if err != nil {
				return err
			}
			fmt.Printf("Profile %d started.\n", id)
			return nil
		},
	}
}

func newProfilesStopCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stop <id>",
		Short: "Stop a running profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid profile ID: %s", args[0])
			}
			client := newAPIClient()
			_, err = client.post(fmt.Sprintf("/api/profiles/%d/stop", id), nil)
			if err != nil {
				return err
			}
			fmt.Printf("Profile %d stopped.\n", id)
			return nil
		},
	}
}

func newProfilesDuplicateCmd() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "duplicate <id>",
		Short: "Duplicate a profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			body := map[string]interface{}{}
			if name != "" {
				body["name"] = name
			}

			client := newAPIClient()
			data, err := client.post("/api/profiles/"+args[0]+"/duplicate", body)
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				fmt.Println(string(data))
			} else {
				var p Profile
				if err := json.Unmarshal(data, &p); err != nil {
					return err
				}
				fmt.Printf("Profile duplicated: %s (ID: %d)\n", p.Name, p.ID)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Name for the duplicate")

	return cmd
}
