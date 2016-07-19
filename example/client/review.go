package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateReviewPayload is the review create action payload.
type CreateReviewPayload struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty" form:"comment,omitempty"`
	Rating  int     `json:"rating" xml:"rating" form:"rating"`
}

// CreateReviewPath computes a request path to the create action of review.
func CreateReviewPath(userID string, proposalID string) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v/review", userID, proposalID)
}

// Create a new review
func (c *Client) CreateReview(ctx context.Context, path string, payload *CreateReviewPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateReviewRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateReviewRequest create the request corresponding to the create action endpoint of the review resource.
func (c *Client) NewCreateReviewRequest(ctx context.Context, path string, payload *CreateReviewPayload, contentType string) (*http.Request, error) {
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

// DeleteReviewPath computes a request path to the delete action of review.
func DeleteReviewPath(userID string, proposalID string, reviewID int) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID)
}

// DeleteReview makes a request to the delete action endpoint of the review resource
func (c *Client) DeleteReview(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteReviewRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteReviewRequest create the request corresponding to the delete action endpoint of the review resource.
func (c *Client) NewDeleteReviewRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListReviewPath computes a request path to the list action of review.
func ListReviewPath(userID string, proposalID string) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v/review", userID, proposalID)
}

// List all reviews for a proposal
func (c *Client) ListReview(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListReviewRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListReviewRequest create the request corresponding to the list action endpoint of the review resource.
func (c *Client) NewListReviewRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowReviewPath computes a request path to the show action of review.
func ShowReviewPath(userID string, proposalID string, reviewID int) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID)
}

// Retrieve review with given id
func (c *Client) ShowReview(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowReviewRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowReviewRequest create the request corresponding to the show action endpoint of the review resource.
func (c *Client) NewShowReviewRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateReviewPayload is the review update action payload.
type UpdateReviewPayload struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty" form:"comment,omitempty"`
	Rating  *int    `json:"rating,omitempty" xml:"rating,omitempty" form:"rating,omitempty"`
}

// UpdateReviewPath computes a request path to the update action of review.
func UpdateReviewPath(userID string, proposalID string, reviewID int) string {
	return fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", userID, proposalID, reviewID)
}

// UpdateReview makes a request to the update action endpoint of the review resource
func (c *Client) UpdateReview(ctx context.Context, path string, payload *UpdateReviewPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateReviewRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateReviewRequest create the request corresponding to the update action endpoint of the review resource.
func (c *Client) NewUpdateReviewRequest(ctx context.Context, path string, payload *UpdateReviewPayload, contentType string) (*http.Request, error) {
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
