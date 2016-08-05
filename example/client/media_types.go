//************************************************************************//
// API "congo": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/goadesign/gorma/example/design
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// Token authorization response (default view)
//
// Identifier: application/vnd.authorize+json
type Authorize struct {
	// access token
	AccessToken *string `form:"access_token,omitempty" json:"access_token,omitempty" xml:"access_token,omitempty"`
	// Time to expiration in seconds
	ExpiresIn *int `form:"expires_in,omitempty" json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// type of token
	TokenType *string `form:"token_type,omitempty" json:"token_type,omitempty" xml:"token_type,omitempty"`
}

// DecodeAuthorize decodes the Authorize instance encoded in resp body.
func (c *Client) DecodeAuthorize(resp *http.Response) (*Authorize, error) {
	var decoded Authorize
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Login media type (default view)
//
// Identifier: application/vnd.login+json
type Login struct {
	// UUID of requesting application
	Application *string `form:"application,omitempty" json:"application,omitempty" xml:"application,omitempty"`
	// email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// DecodeLogin decodes the Login instance encoded in resp body.
func (c *Client) DecodeLogin(resp *http.Response) (*Login, error) {
	var decoded Login
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A response to a CFP (default view)
//
// Identifier: application/vnd.proposal+json
type Proposal struct {
	// Response abstract
	Abstract *string `form:"abstract,omitempty" json:"abstract,omitempty" xml:"abstract,omitempty"`
	// Response detail
	Detail *string `form:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
	// API href of user
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// ID of user
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Links to related resources
	Links *ProposalLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// Reviews
	Reviews ReviewCollection `form:"reviews,omitempty" json:"reviews,omitempty" xml:"reviews,omitempty"`
	// Response title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the Proposal media type instance.
func (mt *Proposal) Validate() (err error) {
	if mt.Abstract != nil {
		if len(*mt.Abstract) < 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.abstract`, *mt.Abstract, len(*mt.Abstract), 50, true))
		}
	}
	if mt.Abstract != nil {
		if len(*mt.Abstract) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.abstract`, *mt.Abstract, len(*mt.Abstract), 500, false))
		}
	}
	if mt.Detail != nil {
		if len(*mt.Detail) < 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.detail`, *mt.Detail, len(*mt.Detail), 100, true))
		}
	}
	if mt.Detail != nil {
		if len(*mt.Detail) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.detail`, *mt.Detail, len(*mt.Detail), 2000, false))
		}
	}
	if err2 := mt.Reviews.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if mt.Title != nil {
		if len(*mt.Title) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *mt.Title, len(*mt.Title), 10, true))
		}
	}
	if mt.Title != nil {
		if len(*mt.Title) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *mt.Title, len(*mt.Title), 200, false))
		}
	}
	return
}

// A response to a CFP (link view)
//
// Identifier: application/vnd.proposal+json
type ProposalLink struct {
	// API href of user
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// ID of user
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Response title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the ProposalLink media type instance.
func (mt *ProposalLink) Validate() (err error) {
	if mt.Title != nil {
		if len(*mt.Title) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *mt.Title, len(*mt.Title), 10, true))
		}
	}
	if mt.Title != nil {
		if len(*mt.Title) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *mt.Title, len(*mt.Title), 200, false))
		}
	}
	return
}

// ProposalLinks contains links to related resources of Proposal.
type ProposalLinks struct {
}

// DecodeProposal decodes the Proposal instance encoded in resp body.
func (c *Client) DecodeProposal(resp *http.Response) (*Proposal, error) {
	var decoded Proposal
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeProposalLink decodes the ProposalLink instance encoded in resp body.
func (c *Client) DecodeProposalLink(resp *http.Response) (*ProposalLink, error) {
	var decoded ProposalLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ProposalCollection is the media type for an array of Proposal (default view)
//
// Identifier: application/vnd.proposal+json; type=collection
type ProposalCollection []*Proposal

// Validate validates the ProposalCollection media type instance.
func (mt ProposalCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Abstract != nil {
			if len(*e.Abstract) < 50 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].abstract`, *e.Abstract, len(*e.Abstract), 50, true))
			}
		}
		if e.Abstract != nil {
			if len(*e.Abstract) > 500 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].abstract`, *e.Abstract, len(*e.Abstract), 500, false))
			}
		}
		if e.Detail != nil {
			if len(*e.Detail) < 100 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].detail`, *e.Detail, len(*e.Detail), 100, true))
			}
		}
		if e.Detail != nil {
			if len(*e.Detail) > 2000 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].detail`, *e.Detail, len(*e.Detail), 2000, false))
			}
		}
		if err2 := e.Reviews.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
		if e.Title != nil {
			if len(*e.Title) < 10 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].title`, *e.Title, len(*e.Title), 10, true))
			}
		}
		if e.Title != nil {
			if len(*e.Title) > 200 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].title`, *e.Title, len(*e.Title), 200, false))
			}
		}
	}
	return
}

// ProposalLinksArray contains links to related resources of ProposalCollection.
type ProposalLinksArray []*ProposalLinks

// DecodeProposalCollection decodes the ProposalCollection instance encoded in resp body.
func (c *Client) DecodeProposalCollection(resp *http.Response) (ProposalCollection, error) {
	var decoded ProposalCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A review is submitted by a reviewer (default view)
//
// Identifier: application/vnd.review+json
type Review struct {
	// Review comments
	Comment *string `form:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
	// API href of user
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// ID of user
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Rating of proposal, from 1-5
	Rating *int `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// Validate validates the Review media type instance.
func (mt *Review) Validate() (err error) {
	if mt.Comment != nil {
		if len(*mt.Comment) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.comment`, *mt.Comment, len(*mt.Comment), 10, true))
		}
	}
	if mt.Comment != nil {
		if len(*mt.Comment) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.comment`, *mt.Comment, len(*mt.Comment), 200, false))
		}
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true))
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false))
		}
	}
	return
}

// A review is submitted by a reviewer (link view)
//
// Identifier: application/vnd.review+json
type ReviewLink struct {
	// API href of user
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// ID of user
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// DecodeReview decodes the Review instance encoded in resp body.
func (c *Client) DecodeReview(resp *http.Response) (*Review, error) {
	var decoded Review
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeReviewLink decodes the ReviewLink instance encoded in resp body.
func (c *Client) DecodeReviewLink(resp *http.Response) (*ReviewLink, error) {
	var decoded ReviewLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ReviewCollection is the media type for an array of Review (default view)
//
// Identifier: application/vnd.review+json; type=collection
type ReviewCollection []*Review

// Validate validates the ReviewCollection media type instance.
func (mt ReviewCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Comment != nil {
			if len(*e.Comment) < 10 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].comment`, *e.Comment, len(*e.Comment), 10, true))
			}
		}
		if e.Comment != nil {
			if len(*e.Comment) > 200 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].comment`, *e.Comment, len(*e.Comment), 200, false))
			}
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, true))
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].rating`, *e.Rating, 5, false))
			}
		}
	}
	return
}

// ReviewCollection is the media type for an array of Review (link view)
//
// Identifier: application/vnd.review+json; type=collection
type ReviewLinkCollection []*ReviewLink

// DecodeReviewCollection decodes the ReviewCollection instance encoded in resp body.
func (c *Client) DecodeReviewCollection(resp *http.Response) (ReviewCollection, error) {
	var decoded ReviewCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeReviewLinkCollection decodes the ReviewLinkCollection instance encoded in resp body.
func (c *Client) DecodeReviewLinkCollection(resp *http.Response) (ReviewLinkCollection, error) {
	var decoded ReviewLinkCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A user belonging to a tenant account (default view)
//
// Identifier: application/vnd.user+json
type User struct {
	// Biography of user
	Bio *string `form:"bio,omitempty" json:"bio,omitempty" xml:"bio,omitempty"`
	// City of residence
	City *string `form:"city,omitempty" json:"city,omitempty" xml:"city,omitempty"`
	// Country of residence
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Email address of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// First name of user
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" xml:"firstname,omitempty"`
	// API href of user
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// ID of user
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Last name of user
	Lastname *string `form:"lastname,omitempty" json:"lastname,omitempty" xml:"lastname,omitempty"`
	// Role of user
	Role *string `form:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// State of residence
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {
	if mt.Bio != nil {
		if len(*mt.Bio) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.bio`, *mt.Bio, len(*mt.Bio), 500, false))
		}
	}
	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// A user belonging to a tenant account (link view)
//
// Identifier: application/vnd.user+json
type UserLink struct {
	// Email address of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// API href of user
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// ID of user
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// Validate validates the UserLink media type instance.
func (mt *UserLink) Validate() (err error) {
	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// DecodeUser decodes the User instance encoded in resp body.
func (c *Client) DecodeUser(resp *http.Response) (*User, error) {
	var decoded User
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeUserLink decodes the UserLink instance encoded in resp body.
func (c *Client) DecodeUserLink(resp *http.Response) (*UserLink, error) {
	var decoded UserLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.user+json; type=collection
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Bio != nil {
			if len(*e.Bio) > 500 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].bio`, *e.Bio, len(*e.Bio), 500, false))
			}
		}
		if e.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFormatError(`response[*].email`, *e.Email, goa.FormatEmail, err2))
			}
		}
	}
	return
}

// DecodeUserCollection decodes the UserCollection instance encoded in resp body.
func (c *Client) DecodeUserCollection(resp *http.Response) (UserCollection, error) {
	var decoded UserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
