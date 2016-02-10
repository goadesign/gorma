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
		if e.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
				err = goa.InvalidFormatError(`response[*].email`, *e.Email, goa.FormatEmail, err2, err)
			}
		}
	}
	return
}
