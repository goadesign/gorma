package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
)

// CallbackAuthPath computes a request path to the callback action of auth.
func CallbackAuthPath(provider string) string {
	return fmt.Sprintf("/api/auth/%v/callback", provider)
}

// OAUTH2 callback endpoint
func (c *Client) CallbackAuth(ctx context.Context, path string) (*http.Response, error) {
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

// OauthAuthPath computes a request path to the oauth action of auth.
func OauthAuthPath(provider string) string {
	return fmt.Sprintf("/api/auth/%v", provider)
}

// OAUTH2 login endpoint
func (c *Client) OauthAuth(ctx context.Context, path string) (*http.Response, error) {
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

// RefreshAuthPayload is the auth refresh action payload.
type RefreshAuthPayload struct {
	// UUID of requesting application
	Application *string `json:"application,omitempty" xml:"application,omitempty"`
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

// RefreshAuthPath computes a request path to the refresh action of auth.
func RefreshAuthPath() string {
	return fmt.Sprintf("/api/auth/refresh")
}

// Obtain a refreshed access token
func (c *Client) RefreshAuth(ctx context.Context, path string, payload *RefreshAuthPayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}

// TokenAuthPayload is the auth token action payload.
type TokenAuthPayload struct {
	// UUID of requesting application
	Application *string `json:"application,omitempty" xml:"application,omitempty"`
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

// TokenAuthPath computes a request path to the token action of auth.
func TokenAuthPath() string {
	return fmt.Sprintf("/api/auth/token")
}

// Obtain an access token
func (c *Client) TokenAuth(ctx context.Context, path string, payload *TokenAuthPayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}
