package nyt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	apiKey string

	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,

		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	q := req.URL.Query()
	q.Set("api-key", c.apiKey)
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}

func (c *Client) Get(target string, dest any) error {
	req, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, body)
	}
	return json.NewDecoder(resp.Body).Decode(dest)
}
