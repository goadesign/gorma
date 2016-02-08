//************************************************************************//
// API "congo" version v1: Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package v1

import "github.com/goadesign/goa"

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
