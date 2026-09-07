package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type apiClient struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func newAPIClient() *apiClient {
	return &apiClient{
		baseURL: apiURL,
		apiKey:  getAPIKey(),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *apiClient) do(method, path string, body interface{}) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, c.baseURL+path, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.apiKey != "" {
		req.Header.Set("x-api-key", c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, string(data))
	}

	return data, nil
}

func (c *apiClient) get(path string) ([]byte, error) {
	return c.do("GET", path, nil)
}

func (c *apiClient) post(path string, body interface{}) ([]byte, error) {
	return c.do("POST", path, body)
}

func (c *apiClient) put(path string, body interface{}) ([]byte, error) {
	return c.do("PUT", path, body)
}

func (c *apiClient) delete(path string) ([]byte, error) {
	return c.do("DELETE", path, nil)
}

func printJSON(data []byte) error {
	var pretty bytes.Buffer
	if err := json.Indent(&pretty, data, "", "  "); err != nil {
		fmt.Println(string(data))
		return nil
	}
	fmt.Println(pretty.String())
	return nil
}
