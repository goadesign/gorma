package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty" form:"bio,omitempty"`
	City      *string `json:"city,omitempty" xml:"city,omitempty" form:"city,omitempty"`
	Country   *string `json:"country,omitempty" xml:"country,omitempty" form:"country,omitempty"`
	Email     string  `json:"email" xml:"email" form:"email"`
	Firstname string  `json:"firstname" xml:"firstname" form:"firstname"`
	Lastname  string  `json:"lastname" xml:"lastname" form:"lastname"`
	State     *string `json:"state,omitempty" xml:"state,omitempty" form:"state,omitempty"`
}

// CreateUserPath computes a request path to the create action of user.
func CreateUserPath() string {
	return fmt.Sprintf("/api/users")
}

// Record new user
func (c *Client) CreateUser(ctx context.Context, path string, payload *CreateUserPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateUserRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateUserRequest create the request corresponding to the create action endpoint of the user resource.
func (c *Client) NewCreateUserRequest(ctx context.Context, path string, payload *CreateUserPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
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
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteUserPath computes a request path to the delete action of user.
func DeleteUserPath(userID int) string {
	return fmt.Sprintf("/api/users/%v", userID)
}

// DeleteUser makes a request to the delete action endpoint of the user resource
func (c *Client) DeleteUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteUserRequest create the request corresponding to the delete action endpoint of the user resource.
func (c *Client) NewDeleteUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListUserPath computes a request path to the list action of user.
func ListUserPath() string {
	return fmt.Sprintf("/api/users")
}

// List all users in account
func (c *Client) ListUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListUserRequest create the request corresponding to the list action endpoint of the user resource.
func (c *Client) NewListUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowUserPath computes a request path to the show action of user.
func ShowUserPath(userID int) string {
	return fmt.Sprintf("/api/users/%v", userID)
}

// Retrieve user with given id
func (c *Client) ShowUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowUserRequest create the request corresponding to the show action endpoint of the user resource.
func (c *Client) NewShowUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty" form:"bio,omitempty"`
	City      *string `json:"city,omitempty" xml:"city,omitempty" form:"city,omitempty"`
	Country   *string `json:"country,omitempty" xml:"country,omitempty" form:"country,omitempty"`
	Email     string  `json:"email" xml:"email" form:"email"`
	Firstname *string `json:"firstname,omitempty" xml:"firstname,omitempty" form:"firstname,omitempty"`
	Lastname  *string `json:"lastname,omitempty" xml:"lastname,omitempty" form:"lastname,omitempty"`
	State     *string `json:"state,omitempty" xml:"state,omitempty" form:"state,omitempty"`
}

// UpdateUserPath computes a request path to the update action of user.
func UpdateUserPath(userID int) string {
	return fmt.Sprintf("/api/users/%v", userID)
}

// UpdateUser makes a request to the update action endpoint of the user resource
func (c *Client) UpdateUser(ctx context.Context, path string, payload *UpdateUserPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateUserRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateUserRequest create the request corresponding to the update action endpoint of the user resource.
func (c *Client) NewUpdateUserRequest(ctx context.Context, path string, payload *UpdateUserPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
