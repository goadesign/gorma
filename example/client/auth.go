package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CallbackAuthPath computes a request path to the callback action of auth.
func CallbackAuthPath(provider string) string {
	return fmt.Sprintf("/api/auth/%v/callback", provider)
}

// OAUTH2 callback endpoint
func (c *Client) CallbackAuth(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewCallbackAuthRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCallbackAuthRequest create the request corresponding to the callback action endpoint of the auth resource.
func (c *Client) NewCallbackAuthRequest(ctx context.Context, path string) (*http.Request, error) {
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

// OauthAuthPath computes a request path to the oauth action of auth.
func OauthAuthPath(provider string) string {
	return fmt.Sprintf("/api/auth/%v", provider)
}

// OAUTH2 login endpoint
func (c *Client) OauthAuth(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewOauthAuthRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewOauthAuthRequest create the request corresponding to the oauth action endpoint of the auth resource.
func (c *Client) NewOauthAuthRequest(ctx context.Context, path string) (*http.Request, error) {
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
	req, err := c.NewRefreshAuthRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRefreshAuthRequest create the request corresponding to the refresh action endpoint of the auth resource.
func (c *Client) NewRefreshAuthRequest(ctx context.Context, path string, payload *RefreshAuthPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*") // Use default encoder
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
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
	req, err := c.NewTokenAuthRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewTokenAuthRequest create the request corresponding to the token action endpoint of the auth resource.
func (c *Client) NewTokenAuthRequest(ctx context.Context, path string, payload *TokenAuthPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*") // Use default encoder
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}
