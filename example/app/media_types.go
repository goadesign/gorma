//************************************************************************//
// API "congo": Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// Token authorization response
// Identifier: application/vnd.authorize+json
type Authorize struct {
	// access token
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// Time to expiration in seconds
	ExpiresIn *int `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// type of token
	TokenType *string `json:"token_type,omitempty" xml:"token_type,omitempty"`
}

// Login media type
// Identifier: application/vnd.login+json
type Login struct {
	// UUID of requesting application
	Application *string `json:"application,omitempty" xml:"application,omitempty"`
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

// A response to a CFP
// Identifier: application/vnd.proposal+json
type Proposal struct {
	// Response abstract
	Abstract *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	// Response detail
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// API href of user
	Href *string `json:"href,omitempty" xml:"href,omitempty"`
	// ID of user
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// Links to related resources
	Links *ProposalLinks `json:"links,omitempty" xml:"links,omitempty"`
	// Reviews
	Reviews ReviewCollection `json:"reviews,omitempty" xml:"reviews,omitempty"`
	// Response title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the media type instance.
func (mt *Proposal) Validate() (err error) {
	if mt.Abstract != nil {
		if len(*mt.Abstract) < 50 {
			err = goa.InvalidLengthError(`response.abstract`, *mt.Abstract, len(*mt.Abstract), 50, true, err)
		}
	}
	if mt.Abstract != nil {
		if len(*mt.Abstract) > 500 {
			err = goa.InvalidLengthError(`response.abstract`, *mt.Abstract, len(*mt.Abstract), 500, false, err)
		}
	}
	if mt.Detail != nil {
		if len(*mt.Detail) < 100 {
			err = goa.InvalidLengthError(`response.detail`, *mt.Detail, len(*mt.Detail), 100, true, err)
		}
	}
	if mt.Detail != nil {
		if len(*mt.Detail) > 2000 {
			err = goa.InvalidLengthError(`response.detail`, *mt.Detail, len(*mt.Detail), 2000, false, err)
		}
	}
	for _, e := range mt.Reviews {
		if e.Comment != nil {
			if len(*e.Comment) < 10 {
				err = goa.InvalidLengthError(`response.reviews[*].comment`, *e.Comment, len(*e.Comment), 10, true, err)
			}
		}
		if e.Comment != nil {
			if len(*e.Comment) > 200 {
				err = goa.InvalidLengthError(`response.reviews[*].comment`, *e.Comment, len(*e.Comment), 200, false, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`response.reviews[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`response.reviews[*].rating`, *e.Rating, 5, false, err)
			}
		}
	}
	if mt.Title != nil {
		if len(*mt.Title) < 10 {
			err = goa.InvalidLengthError(`response.title`, *mt.Title, len(*mt.Title), 10, true, err)
		}
	}
	if mt.Title != nil {
		if len(*mt.Title) > 200 {
			err = goa.InvalidLengthError(`response.title`, *mt.Title, len(*mt.Title), 200, false, err)
		}
	}
	return
}

// A response to a CFP, link view
// Identifier: application/vnd.proposal+json
type ProposalLink struct {
	// API href of user
	Href *string `json:"href,omitempty" xml:"href,omitempty"`
	// ID of user
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// Response title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the media type instance.
func (mt *ProposalLink) Validate() (err error) {
	if mt.Title != nil {
		if len(*mt.Title) < 10 {
			err = goa.InvalidLengthError(`response.title`, *mt.Title, len(*mt.Title), 10, true, err)
		}
	}
	if mt.Title != nil {
		if len(*mt.Title) > 200 {
			err = goa.InvalidLengthError(`response.title`, *mt.Title, len(*mt.Title), 200, false, err)
		}
	}
	return
}

// ProposalLinks contains links to related resources of Proposal.
type ProposalLinks struct {
	Reviews ReviewLinkCollection `json:"reviews,omitempty" xml:"reviews,omitempty"`
}

// , default view
// Identifier: application/vnd.proposal+json; type=collection
type ProposalCollection []*Proposal

// Validate validates the media type instance.
func (mt ProposalCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Abstract != nil {
			if len(*e.Abstract) < 50 {
				err = goa.InvalidLengthError(`response[*].abstract`, *e.Abstract, len(*e.Abstract), 50, true, err)
			}
		}
		if e.Abstract != nil {
			if len(*e.Abstract) > 500 {
				err = goa.InvalidLengthError(`response[*].abstract`, *e.Abstract, len(*e.Abstract), 500, false, err)
			}
		}
		if e.Detail != nil {
			if len(*e.Detail) < 100 {
				err = goa.InvalidLengthError(`response[*].detail`, *e.Detail, len(*e.Detail), 100, true, err)
			}
		}
		if e.Detail != nil {
			if len(*e.Detail) > 2000 {
				err = goa.InvalidLengthError(`response[*].detail`, *e.Detail, len(*e.Detail), 2000, false, err)
			}
		}
		for _, e := range e.Reviews {
			if e.Comment != nil {
				if len(*e.Comment) < 10 {
					err = goa.InvalidLengthError(`response[*].reviews[*].comment`, *e.Comment, len(*e.Comment), 10, true, err)
				}
			}
			if e.Comment != nil {
				if len(*e.Comment) > 200 {
					err = goa.InvalidLengthError(`response[*].reviews[*].comment`, *e.Comment, len(*e.Comment), 200, false, err)
				}
			}
			if e.Rating != nil {
				if *e.Rating < 1 {
					err = goa.InvalidRangeError(`response[*].reviews[*].rating`, *e.Rating, 1, true, err)
				}
			}
			if e.Rating != nil {
				if *e.Rating > 5 {
					err = goa.InvalidRangeError(`response[*].reviews[*].rating`, *e.Rating, 5, false, err)
				}
			}
		}
		if e.Title != nil {
			if len(*e.Title) < 10 {
				err = goa.InvalidLengthError(`response[*].title`, *e.Title, len(*e.Title), 10, true, err)
			}
		}
		if e.Title != nil {
			if len(*e.Title) > 200 {
				err = goa.InvalidLengthError(`response[*].title`, *e.Title, len(*e.Title), 200, false, err)
			}
		}
	}
	return
}

// ProposalLinksArray contains links to related resources of ProposalCollection.
type ProposalLinksArray []*ProposalLinks

// A review is submitted by a reviewer
// Identifier: application/vnd.review+json
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

// Validate validates the media type instance.
func (mt *Review) Validate() (err error) {
	if mt.Comment != nil {
		if len(*mt.Comment) < 10 {
			err = goa.InvalidLengthError(`response.comment`, *mt.Comment, len(*mt.Comment), 10, true, err)
		}
	}
	if mt.Comment != nil {
		if len(*mt.Comment) > 200 {
			err = goa.InvalidLengthError(`response.comment`, *mt.Comment, len(*mt.Comment), 200, false, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false, err)
		}
	}
	return
}

// A review is submitted by a reviewer, link view
// Identifier: application/vnd.review+json
type ReviewLink struct {
	// API href of user
	Href *string `json:"href,omitempty" xml:"href,omitempty"`
	// ID of user
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
}

// , default view
// Identifier: application/vnd.review+json; type=collection
type ReviewCollection []*Review

// Validate validates the media type instance.
func (mt ReviewCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Comment != nil {
			if len(*e.Comment) < 10 {
				err = goa.InvalidLengthError(`response[*].comment`, *e.Comment, len(*e.Comment), 10, true, err)
			}
		}
		if e.Comment != nil {
			if len(*e.Comment) > 200 {
				err = goa.InvalidLengthError(`response[*].comment`, *e.Comment, len(*e.Comment), 200, false, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 5, false, err)
			}
		}
	}
	return
}

// , link view
// Identifier: application/vnd.review+json; type=collection
type ReviewLinkCollection []*ReviewLink

// A user belonging to a tenant account
// Identifier: application/vnd.user+json
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

// Validate validates the media type instance.
func (mt *User) Validate() (err error) {
	if mt.Bio != nil {
		if len(*mt.Bio) > 500 {
			err = goa.InvalidLengthError(`response.bio`, *mt.Bio, len(*mt.Bio), 500, false, err)
		}
	}
	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2, err)
		}
	}
	return
}

// A user belonging to a tenant account, link view
// Identifier: application/vnd.user+json
type UserLink struct {
	// Email address of user
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// API href of user
	Href *string `json:"href,omitempty" xml:"href,omitempty"`
	// ID of user
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
}

// Validate validates the media type instance.
func (mt *UserLink) Validate() (err error) {
	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2, err)
		}
	}
	return
}

// , default view
// Identifier: application/vnd.user+json; type=collection
type UserCollection []*User

// Validate validates the media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Bio != nil {
			if len(*e.Bio) > 500 {
				err = goa.InvalidLengthError(`response[*].bio`, *e.Bio, len(*e.Bio), 500, false, err)
			}
		}
		if e.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
				err = goa.InvalidFormatError(`response[*].email`, *e.Email, goa.FormatEmail, err2, err)
			}
		}
	}
	return
}
