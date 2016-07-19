package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// BootstrapUIPath computes a request path to the bootstrap action of ui.
func BootstrapUIPath() string {
	return fmt.Sprintf("/")
}

// Render single page app HTML
func (c *Client) BootstrapUI(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewBootstrapUIRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewBootstrapUIRequest create the request corresponding to the bootstrap action endpoint of the ui resource.
func (c *Client) NewBootstrapUIRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
