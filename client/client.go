package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	deviceURL  *url.URL
	httpClient *http.Client
}

func NewClient(awtrix3Url string) (*Client, error) {
	deviceURL, err := url.Parse(awtrix3Url)
	if err != nil {
		return nil, err
	}
	return &Client{
		deviceURL: deviceURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil
}

func (c *Client) doRequest(method, path string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	fullUrl := c.deviceURL.JoinPath("api", path).String()
	req, err := http.NewRequest(method, fullUrl, reqBody)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}
