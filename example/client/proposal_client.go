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

// CreateProposalPayload is the proposal create action payload.
type CreateProposalPayload struct {
	Abstract  string `json:"abstract" xml:"abstract"`
	Detail    string `json:"detail" xml:"detail"`
	Title     string `json:"title" xml:"title"`
	Withdrawn *bool  `json:"withdrawn,omitempty" xml:"withdrawn,omitempty"`
}

// CreateProposalPath computes a request path to the create action of proposal.
func CreateProposalPath(userID string) string {
	return fmt.Sprintf("/api/users/%v/proposals", userID)
}

// Create a new proposal
func (c *Client) CreateProposal(ctx context.Context, path string, payload *CreateProposalPayload) (*http.Response, error) {
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

// DeleteProposalPath computes a request path to the delete action of proposal.
func DeleteProposalPath(proposalID int, userID string) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID)
}

// DeleteProposal makes a request to the delete action endpoint of the proposal resource
func (c *Client) DeleteProposal(ctx context.Context, path string) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}

// ListProposalPath computes a request path to the list action of proposal.
func ListProposalPath(userID string) string {
	return fmt.Sprintf("/api/users/%v/proposals", userID)
}

// List all proposals for a user
func (c *Client) ListProposal(ctx context.Context, path string) (*http.Response, error) {
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

// ShowProposalPath computes a request path to the show action of proposal.
func ShowProposalPath(proposalID int, userID string) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID)
}

// Retrieve proposal with given id
func (c *Client) ShowProposal(ctx context.Context, path string) (*http.Response, error) {
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

// UpdateProposalPayload is the proposal update action payload.
type UpdateProposalPayload struct {
	Abstract  *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	Detail    *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	Withdrawn *bool   `json:"withdrawn,omitempty" xml:"withdrawn,omitempty"`
}

// UpdateProposalPath computes a request path to the update action of proposal.
func UpdateProposalPath(proposalID int, userID string) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v", userID, proposalID)
}

// UpdateProposal makes a request to the update action endpoint of the proposal resource
func (c *Client) UpdateProposal(ctx context.Context, path string, payload *UpdateProposalPayload) (*http.Response, error) {
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
	req, err := http.NewRequest("PATCH", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}
