//************************************************************************//
// API "congo": Application User Types
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

// ProposalPayload user type.
type ProposalPayload struct {
	Abstract  *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	Detail    *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	Withdrawn *bool   `json:"withdrawn,omitempty" xml:"withdrawn,omitempty"`
}

// Validate validates the ProposalPayload type instance.
func (ut *ProposalPayload) Validate() (err error) {
	if ut.Abstract != nil {
		if len(*ut.Abstract) < 50 {
			err = goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 50, true, err)
		}
	}
	if ut.Abstract != nil {
		if len(*ut.Abstract) > 500 {
			err = goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 500, false, err)
		}
	}
	if ut.Detail != nil {
		if len(*ut.Detail) < 100 {
			err = goa.InvalidLengthError(`response.detail`, *ut.Detail, len(*ut.Detail), 100, true, err)
		}
	}
	if ut.Detail != nil {
		if len(*ut.Detail) > 2000 {
			err = goa.InvalidLengthError(`response.detail`, *ut.Detail, len(*ut.Detail), 2000, false, err)
		}
	}
	if ut.Title != nil {
		if len(*ut.Title) < 10 {
			err = goa.InvalidLengthError(`response.title`, *ut.Title, len(*ut.Title), 10, true, err)
		}
	}
	if ut.Title != nil {
		if len(*ut.Title) > 200 {
			err = goa.InvalidLengthError(`response.title`, *ut.Title, len(*ut.Title), 200, false, err)
		}
	}
	return
}

// ReviewPayload user type.
type ReviewPayload struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	Rating  *int    `json:"rating,omitempty" xml:"rating,omitempty"`
}

// Validate validates the ReviewPayload type instance.
func (ut *ReviewPayload) Validate() (err error) {
	if ut.Comment != nil {
		if len(*ut.Comment) < 10 {
			err = goa.InvalidLengthError(`response.comment`, *ut.Comment, len(*ut.Comment), 10, true, err)
		}
	}
	if ut.Comment != nil {
		if len(*ut.Comment) > 200 {
			err = goa.InvalidLengthError(`response.comment`, *ut.Comment, len(*ut.Comment), 200, false, err)
		}
	}
	if ut.Rating != nil {
		if *ut.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *ut.Rating, 1, true, err)
		}
	}
	if ut.Rating != nil {
		if *ut.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *ut.Rating, 5, false, err)
		}
	}
	return
}

// UserPayload user type.
type UserPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty"`
	City      *string `json:"city,omitempty" xml:"city,omitempty"`
	Country   *string `json:"country,omitempty" xml:"country,omitempty"`
	Email     *string `json:"email,omitempty" xml:"email,omitempty"`
	Firstname *string `json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname  *string `json:"lastname,omitempty" xml:"lastname,omitempty"`
	State     *string `json:"state,omitempty" xml:"state,omitempty"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.Bio != nil {
		if len(*ut.Bio) > 500 {
			err = goa.InvalidLengthError(`response.bio`, *ut.Bio, len(*ut.Bio), 500, false, err)
		}
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2, err)
		}
	}
	return
}
