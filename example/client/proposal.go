package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateProposalPayload is the proposal create action payload.
type CreateProposalPayload struct {
	Abstract  string `form:"abstract" json:"abstract" xml:"abstract"`
	Detail    string `form:"detail" json:"detail" xml:"detail"`
	Title     string `form:"title" json:"title" xml:"title"`
	Withdrawn *bool  `form:"withdrawn,omitempty" json:"withdrawn,omitempty" xml:"withdrawn,omitempty"`
}

// CreateProposalPath computes a request path to the create action of proposal.
func CreateProposalPath(userID int) string {
	return fmt.Sprintf("/api/users/%v/proposals", userID)
}

// Create a new proposal
func (c *Client) CreateProposal(ctx context.Context, path string, payload *CreateProposalPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateProposalRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateProposalRequest create the request corresponding to the create action endpoint of the proposal resource.
func (c *Client) NewCreateProposalRequest(ctx context.Context, path string, payload *CreateProposalPayload, contentType string) (*http.Request, error) {
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

// DeleteProposalPath computes a request path to the delete action of proposal.
func DeleteProposalPath(userID int, proposalID int) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID)
}

// DeleteProposal makes a request to the delete action endpoint of the proposal resource
func (c *Client) DeleteProposal(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteProposalRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteProposalRequest create the request corresponding to the delete action endpoint of the proposal resource.
func (c *Client) NewDeleteProposalRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListProposalPath computes a request path to the list action of proposal.
func ListProposalPath(userID int) string {
	return fmt.Sprintf("/api/users/%v/proposals", userID)
}

// List all proposals for a user
func (c *Client) ListProposal(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListProposalRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListProposalRequest create the request corresponding to the list action endpoint of the proposal resource.
func (c *Client) NewListProposalRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowProposalPath computes a request path to the show action of proposal.
func ShowProposalPath(userID int, proposalID int) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID)
}

// Retrieve proposal with given id
func (c *Client) ShowProposal(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowProposalRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowProposalRequest create the request corresponding to the show action endpoint of the proposal resource.
func (c *Client) NewShowProposalRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateProposalPath computes a request path to the update action of proposal.
func UpdateProposalPath(userID int, proposalID int) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID)
}

// UpdateProposal makes a request to the update action endpoint of the proposal resource
func (c *Client) UpdateProposal(ctx context.Context, path string, payload *ProposalPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateProposalRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateProposalRequest create the request corresponding to the update action endpoint of the proposal resource.
func (c *Client) NewUpdateProposalRequest(ctx context.Context, path string, payload *ProposalPayload, contentType string) (*http.Request, error) {
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
