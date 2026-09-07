package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

type Proxy struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	IP        string `json:"ip"`
	Country   string `json:"country"`
	City      string `json:"city"`
	ISP       string `json:"isp"`
	CreatedAt string `json:"createdAt"`
}

func newProxiesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proxies",
		Short: "Manage proxies",
	}

	cmd.AddCommand(
		newProxiesListCmd(),
		newProxiesCreateCmd(),
		newProxiesDeleteCmd(),
		newProxiesCheckCmd(),
	)

	return cmd
}

func newProxiesListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all proxies",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := newAPIClient()
			data, err := client.get("/api/proxies")
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				fmt.Println(string(data))
				return nil
			}

			var proxies []Proxy
			if err := json.Unmarshal(data, &proxies); err != nil {
				return err
			}

			if len(proxies) == 0 {
				fmt.Println("No proxies found.")
				return nil
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tNAME\tHOST:PORT\tTYPE\tSTATUS\tCOUNTRY")
			for _, p := range proxies {
				fmt.Fprintf(w, "%d\t%s\t%s:%d\t%s\t%s\t%s\n",
					p.ID, p.Name, p.Host, p.Port, p.Type, p.Status, p.Country)
			}
			w.Flush()
			return nil
		},
	}
}

func newProxiesCreateCmd() *cobra.Command {
	var name, host, username, password, proxyType string
	var port int

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new proxy",
		RunE: func(cmd *cobra.Command, args []string) error {
			if host == "" || port == 0 {
				return fmt.Errorf("--host and --port are required")
			}

			body := map[string]interface{}{
				"host": host,
				"port": port,
			}
			if name != "" {
				body["name"] = name
			}
			if proxyType != "" {
				body["type"] = proxyType
			}
			if username != "" {
				body["username"] = username
			}
			if password != "" {
				body["password"] = password
			}

			client := newAPIClient()
			data, err := client.post("/api/proxies", body)
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				fmt.Println(string(data))
			} else {
				var p Proxy
				if err := json.Unmarshal(data, &p); err != nil {
					return err
				}
				fmt.Printf("Proxy created: %s (ID: %d)\n", p.Name, p.ID)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Proxy name")
	cmd.Flags().StringVar(&host, "host", "", "Proxy host")
	cmd.Flags().IntVar(&port, "port", 0, "Proxy port")
	cmd.Flags().StringVar(&proxyType, "type", "http", "Proxy type (http, socks4, socks5)")
	cmd.Flags().StringVar(&username, "username", "", "Proxy username")
	cmd.Flags().StringVar(&password, "password", "", "Proxy password")

	return cmd
}

func newProxiesDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete a proxy",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := newAPIClient()
			_, err := client.delete("/api/proxies/" + args[0])
			if err != nil {
				return err
			}
			fmt.Println("Proxy deleted.")
			return nil
		},
	}
}

func newProxiesCheckCmd() *cobra.Command {
	var host, username, password, proxyType string
	var port int

	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check a proxy connection",
		RunE: func(cmd *cobra.Command, args []string) error {
			if host == "" || port == 0 {
				return fmt.Errorf("--host and --port are required")
			}

			body := map[string]interface{}{
				"host": host,
				"port": port,
			}
			if proxyType != "" {
				body["type"] = proxyType
			}
			if username != "" {
				body["username"] = username
			}
			if password != "" {
				body["password"] = password
			}

			client := newAPIClient()
			data, err := client.post("/api/proxies/check", body)
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				fmt.Println(string(data))
			} else {
				var result map[string]interface{}
				if err := json.Unmarshal(data, &result); err != nil {
					return err
				}
				success, _ := result["success"].(bool)
				if success {
					fmt.Println("Proxy is working.")
				} else {
					fmt.Println("Proxy check failed.")
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&host, "host", "", "Proxy host")
	cmd.Flags().IntVar(&port, "port", 0, "Proxy port")
	cmd.Flags().StringVar(&proxyType, "type", "http", "Proxy type")
	cmd.Flags().StringVar(&username, "username", "", "Proxy username")
	cmd.Flags().StringVar(&password, "password", "", "Proxy password")

	return cmd
}

func init() {
	_ = strconv.Itoa // ensure import
}
