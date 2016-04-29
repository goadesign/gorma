package client

import (
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
)

// BootstrapUIPath computes a request path to the bootstrap action of ui.
func BootstrapUIPath() string {
	return fmt.Sprintf("/")
}

// Render single page app HTML
func (c *Client) BootstrapUI(ctx context.Context, path string) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}
