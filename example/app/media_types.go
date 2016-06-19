//************************************************************************//
// API "congo": Application Media Types
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/goadesign/gorma/example/design
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// Authorize media type.
//
// Identifier: application/vnd.authorize+json
type Authorize struct {
	// access token
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty" form:"access_token,omitempty"`
	// Time to expiration in seconds
	ExpiresIn *int `json:"expires_in,omitempty" xml:"expires_in,omitempty" form:"expires_in,omitempty"`
	// type of token
	TokenType *string `json:"token_type,omitempty" xml:"token_type,omitempty" form:"token_type,omitempty"`
}

// Login media type.
//
// Identifier: application/vnd.login+json
type Login struct {
	// UUID of requesting application
	Application *string `json:"application,omitempty" xml:"application,omitempty" form:"application,omitempty"`
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty" form:"email,omitempty"`
	// password
	Password *string `json:"password,omitempty" xml:"password,omitempty" form:"password,omitempty"`
}

// Proposal media type.
//
// Identifier: application/vnd.proposal+json
type Proposal struct {
	// Response abstract
	Abstract *string `json:"abstract,omitempty" xml:"abstract,omitempty" form:"abstract,omitempty"`
	// Response detail
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty" form:"detail,omitempty"`
	// API href of user
	Href *string `json:"href,omitempty" xml:"href,omitempty" form:"href,omitempty"`
	// ID of user
	ID *int `json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	// Links to related resources
	Links *ProposalLinks `json:"links,omitempty" xml:"links,omitempty" form:"links,omitempty"`
	// Reviews
	Reviews ReviewCollection `json:"reviews,omitempty" xml:"reviews,omitempty" form:"reviews,omitempty"`
	// Response title
	Title *string `json:"title,omitempty" xml:"title,omitempty" form:"title,omitempty"`
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

// ProposalLink media type.
//
// Identifier: application/vnd.proposal+json
type ProposalLink struct {
	// API href of user
	Href *string `json:"href,omitempty" xml:"href,omitempty" form:"href,omitempty"`
	// ID of user
	ID *int `json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	// Response title
	Title *string `json:"title,omitempty" xml:"title,omitempty" form:"title,omitempty"`
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
	Reviews ReviewLinkCollection `json:"reviews,omitempty" xml:"reviews,omitempty" form:"reviews,omitempty"`
}

// ProposalCollection media type is a collection of Proposal.
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

// Review media type.
//
// Identifier: application/vnd.review+json
type Review struct {
	// Review comments
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty" form:"comment,omitempty"`
	// API href of user
	Href *string `json:"href,omitempty" xml:"href,omitempty" form:"href,omitempty"`
	// ID of user
	ID *int `json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	// Rating of proposal, from 1-5
	Rating *int `json:"rating,omitempty" xml:"rating,omitempty" form:"rating,omitempty"`
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

// ReviewLink media type.
//
// Identifier: application/vnd.review+json
type ReviewLink struct {
	// API href of user
	Href *string `json:"href,omitempty" xml:"href,omitempty" form:"href,omitempty"`
	// ID of user
	ID *int `json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
}

// ReviewCollection media type is a collection of Review.
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

// ReviewLinkCollection media type is a collection of ReviewLink.
//
// Identifier: application/vnd.review+json; type=collection
type ReviewLinkCollection []*ReviewLink

// User media type.
//
// Identifier: application/vnd.user+json
type User struct {
	// Biography of user
	Bio *string `json:"bio,omitempty" xml:"bio,omitempty" form:"bio,omitempty"`
	// City of residence
	City *string `json:"city,omitempty" xml:"city,omitempty" form:"city,omitempty"`
	// Country of residence
	Country *string `json:"country,omitempty" xml:"country,omitempty" form:"country,omitempty"`
	// Email address of user
	Email *string `json:"email,omitempty" xml:"email,omitempty" form:"email,omitempty"`
	// First name of user
	Firstname *string `json:"firstname,omitempty" xml:"firstname,omitempty" form:"firstname,omitempty"`
	// API href of user
	Href *string `json:"href,omitempty" xml:"href,omitempty" form:"href,omitempty"`
	// ID of user
	ID *int `json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
	// Last name of user
	Lastname *string `json:"lastname,omitempty" xml:"lastname,omitempty" form:"lastname,omitempty"`
	// Role of user
	Role *string `json:"role,omitempty" xml:"role,omitempty" form:"role,omitempty"`
	// State of residence
	State *string `json:"state,omitempty" xml:"state,omitempty" form:"state,omitempty"`
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

// UserLink media type.
//
// Identifier: application/vnd.user+json
type UserLink struct {
	// Email address of user
	Email *string `json:"email,omitempty" xml:"email,omitempty" form:"email,omitempty"`
	// API href of user
	Href *string `json:"href,omitempty" xml:"href,omitempty" form:"href,omitempty"`
	// ID of user
	ID *int `json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
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

// UserCollection media type is a collection of User.
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
