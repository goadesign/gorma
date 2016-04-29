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

// CreateReviewPayload is the review create action payload.
type CreateReviewPayload struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	Rating  int     `json:"rating" xml:"rating"`
}

// CreateReviewPath computes a request path to the create action of review.
func CreateReviewPath(proposalID string, userID string) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v/review", userID, proposalID)
}

// Create a new review
func (c *Client) CreateReview(ctx context.Context, path string, payload *CreateReviewPayload) (*http.Response, error) {
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

// DeleteReviewPath computes a request path to the delete action of review.
func DeleteReviewPath(proposalID string, reviewID int, userID string) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID)
}

// DeleteReview makes a request to the delete action endpoint of the review resource
func (c *Client) DeleteReview(ctx context.Context, path string) (*http.Response, error) {
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

// ListReviewPath computes a request path to the list action of review.
func ListReviewPath(proposalID string, userID string) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v/review", userID, proposalID)
}

// List all reviews for a proposal
func (c *Client) ListReview(ctx context.Context, path string) (*http.Response, error) {
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

// ShowReviewPath computes a request path to the show action of review.
func ShowReviewPath(proposalID string, reviewID int, userID string) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID)
}

// Retrieve review with given id
func (c *Client) ShowReview(ctx context.Context, path string) (*http.Response, error) {
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

// UpdateReviewPayload is the review update action payload.
type UpdateReviewPayload struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	Rating  *int    `json:"rating,omitempty" xml:"rating,omitempty"`
}

// UpdateReviewPath computes a request path to the update action of review.
func UpdateReviewPath(proposalID string, reviewID int, userID string) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID)
}

// UpdateReview makes a request to the update action endpoint of the review resource
func (c *Client) UpdateReview(ctx context.Context, path string, payload *UpdateReviewPayload) (*http.Response, error) {
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
