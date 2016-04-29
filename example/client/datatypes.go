//************************************************************************//
// User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"io"
)

// Token authorization response
type Authorize struct {
	// access token
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// Time to expiration in seconds
	ExpiresIn *int `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// type of token
	TokenType *string `json:"token_type,omitempty" xml:"token_type,omitempty"`
}

// DecodeAuthorize decodes the Authorize instance encoded in r.
func DecodeAuthorize(r io.Reader, decoderFn goa.DecoderFunc) (*Authorize, error) {
	var decoded Authorize
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}

// ProposalCollection media type is a collection of Proposal.
type ProposalCollection []*Proposal

// DecodeProposalCollection decodes the ProposalCollection instance encoded in r.
func DecodeProposalCollection(r io.Reader, decoderFn goa.DecoderFunc) (ProposalCollection, error) {
	var decoded ProposalCollection
	err := decoderFn(r).Decode(&decoded)
	return decoded, err
}

// A response to a CFP
type Proposal struct {
	// Response abstract
	Abstract *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	// Response detail
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// API href of user
	Href *string `json:"href,omitempty" xml:"href,omitempty"`
	// ID of user
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// Reviews
	Reviews ReviewCollection `json:"reviews,omitempty" xml:"reviews,omitempty"`
	// Response title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

// DecodeProposal decodes the Proposal instance encoded in r.
func DecodeProposal(r io.Reader, decoderFn goa.DecoderFunc) (*Proposal, error) {
	var decoded Proposal
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}

// ReviewCollection media type is a collection of Review.
type ReviewCollection []*Review

// DecodeReviewCollection decodes the ReviewCollection instance encoded in r.
func DecodeReviewCollection(r io.Reader, decoderFn goa.DecoderFunc) (ReviewCollection, error) {
	var decoded ReviewCollection
	err := decoderFn(r).Decode(&decoded)
	return decoded, err
}

// A review is submitted by a reviewer
type Review struct {
	// Review comments
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// API href of user
	Href *string `json:"href,omitempty" xml:"href,omitempty"`
	// ID of user
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// Rating of proposal, from 1-5
	Rating *int `json:"rating,omitempty" xml:"rating,omitempty"`
}

// DecodeReview decodes the Review instance encoded in r.
func DecodeReview(r io.Reader, decoderFn goa.DecoderFunc) (*Review, error) {
	var decoded Review
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}

// UserCollection media type is a collection of User.
type UserCollection []*User

// DecodeUserCollection decodes the UserCollection instance encoded in r.
func DecodeUserCollection(r io.Reader, decoderFn goa.DecoderFunc) (UserCollection, error) {
	var decoded UserCollection
	err := decoderFn(r).Decode(&decoded)
	return decoded, err
}

// A user belonging to a tenant account
type User struct {
	// Biography of user
	Bio *string `json:"bio,omitempty" xml:"bio,omitempty"`
	// City of residence
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// Country of residence
	Country *string `json:"country,omitempty" xml:"country,omitempty"`
	// Email address of user
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// First name of user
	Firstname *string `json:"firstname,omitempty" xml:"firstname,omitempty"`
	// API href of user
	Href *string `json:"href,omitempty" xml:"href,omitempty"`
	// ID of user
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// Last name of user
	Lastname *string `json:"lastname,omitempty" xml:"lastname,omitempty"`
	// Role of user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// State of residence
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

// DecodeUser decodes the User instance encoded in r.
func DecodeUser(r io.Reader, decoderFn goa.DecoderFunc) (*User, error) {
	var decoded User
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}
