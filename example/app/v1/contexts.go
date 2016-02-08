//************************************************************************//
// API "congo" version v1: Application Contexts
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

import (
	"github.com/goadesign/goa"
	"strconv"
)

// CreateProposalContext provides the proposal create action context.
type CreateProposalContext struct {
	*goa.Context
	APIVersion string
	UserID     int
	Payload    *CreateProposalPayload
}

// NewCreateProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller create action.
func NewCreateProposalContext(c *goa.Context) (*CreateProposalContext, error) {
	var err error
	ctx := CreateProposalContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &ctx, err
}

// CreateProposalPayload is the proposal create action payload.
type CreateProposalPayload struct {
	Abstract  string `json:"abstract" xml:"abstract"`
	Detail    string `json:"detail" xml:"detail"`
	Title     string `json:"title" xml:"title"`
	Withdrawn *bool  `json:"withdrawn,omitempty" xml:"withdrawn,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateProposalPayload) Validate() (err error) {
	if payload.Title == "" {
		err = goa.MissingAttributeError(`raw`, "title", err)
	}

	if payload.Abstract == "" {
		err = goa.MissingAttributeError(`raw`, "abstract", err)
	}

	if payload.Detail == "" {
		err = goa.MissingAttributeError(`raw`, "detail", err)
	}

	if len(payload.Abstract) < 50 {
		err = goa.InvalidLengthError(`raw.abstract`, payload.Abstract, len(payload.Abstract), 50, true, err)
	}
	if len(payload.Abstract) > 500 {
		err = goa.InvalidLengthError(`raw.abstract`, payload.Abstract, len(payload.Abstract), 500, false, err)
	}
	if len(payload.Detail) < 100 {
		err = goa.InvalidLengthError(`raw.detail`, payload.Detail, len(payload.Detail), 100, true, err)
	}
	if len(payload.Detail) > 2000 {
		err = goa.InvalidLengthError(`raw.detail`, payload.Detail, len(payload.Detail), 2000, false, err)
	}
	if len(payload.Title) < 10 {
		err = goa.InvalidLengthError(`raw.title`, payload.Title, len(payload.Title), 10, true, err)
	}
	if len(payload.Title) > 200 {
		err = goa.InvalidLengthError(`raw.title`, payload.Title, len(payload.Title), 200, false, err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateProposalContext) Created() error {
	return ctx.RespondBytes(201, nil)
}

// DeleteProposalContext provides the proposal delete action context.
type DeleteProposalContext struct {
	*goa.Context
	APIVersion string
	ProposalID int
	UserID     int
}

// NewDeleteProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller delete action.
func NewDeleteProposalContext(c *goa.Context) (*DeleteProposalContext, error) {
	var err error
	ctx := DeleteProposalContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteProposalContext) NoContent() error {
	return ctx.RespondBytes(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteProposalContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// ListProposalContext provides the proposal list action context.
type ListProposalContext struct {
	*goa.Context
	APIVersion string
	UserID     int
}

// NewListProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller list action.
func NewListProposalContext(c *goa.Context) (*ListProposalContext, error) {
	var err error
	ctx := ListProposalContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListProposalContext) OK(resp ProposalCollection) error {
	ctx.Header().Set("Content-Type", "application/vnd.proposal+json; type=collection")
	return ctx.Respond(200, resp)
}

// ShowProposalContext provides the proposal show action context.
type ShowProposalContext struct {
	*goa.Context
	APIVersion string
	ProposalID int
	UserID     int
}

// NewShowProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller show action.
func NewShowProposalContext(c *goa.Context) (*ShowProposalContext, error) {
	var err error
	ctx := ShowProposalContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowProposalContext) OK(resp *Proposal) error {
	ctx.Header().Set("Content-Type", "application/vnd.proposal")
	return ctx.Respond(200, resp)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowProposalContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// UpdateProposalContext provides the proposal update action context.
type UpdateProposalContext struct {
	*goa.Context
	APIVersion string
	ProposalID int
	UserID     int
	Payload    *UpdateProposalPayload
}

// NewUpdateProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller update action.
func NewUpdateProposalContext(c *goa.Context) (*UpdateProposalContext, error) {
	var err error
	ctx := UpdateProposalContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &ctx, err
}

// UpdateProposalPayload is the proposal update action payload.
type UpdateProposalPayload struct {
	Abstract  *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	Detail    *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	Withdrawn *bool   `json:"withdrawn,omitempty" xml:"withdrawn,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateProposalPayload) Validate() (err error) {
	if payload.Abstract != nil {
		if len(*payload.Abstract) < 50 {
			err = goa.InvalidLengthError(`raw.abstract`, *payload.Abstract, len(*payload.Abstract), 50, true, err)
		}
	}
	if payload.Abstract != nil {
		if len(*payload.Abstract) > 500 {
			err = goa.InvalidLengthError(`raw.abstract`, *payload.Abstract, len(*payload.Abstract), 500, false, err)
		}
	}
	if payload.Detail != nil {
		if len(*payload.Detail) < 100 {
			err = goa.InvalidLengthError(`raw.detail`, *payload.Detail, len(*payload.Detail), 100, true, err)
		}
	}
	if payload.Detail != nil {
		if len(*payload.Detail) > 2000 {
			err = goa.InvalidLengthError(`raw.detail`, *payload.Detail, len(*payload.Detail), 2000, false, err)
		}
	}
	if payload.Title != nil {
		if len(*payload.Title) < 10 {
			err = goa.InvalidLengthError(`raw.title`, *payload.Title, len(*payload.Title), 10, true, err)
		}
	}
	if payload.Title != nil {
		if len(*payload.Title) > 200 {
			err = goa.InvalidLengthError(`raw.title`, *payload.Title, len(*payload.Title), 200, false, err)
		}
	}
	return
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateProposalContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateProposalContext) NoContent() error {
	return ctx.RespondBytes(204, nil)
}

// CreateReviewContext provides the review create action context.
type CreateReviewContext struct {
	*goa.Context
	APIVersion string
	ProposalID int
	UserID     int
	Payload    *CreateReviewPayload
}

// NewCreateReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller create action.
func NewCreateReviewContext(c *goa.Context) (*CreateReviewContext, error) {
	var err error
	ctx := CreateReviewContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &ctx, err
}

// CreateReviewPayload is the review create action payload.
type CreateReviewPayload struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	Rating  int     `json:"rating" xml:"rating"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateReviewPayload) Validate() (err error) {

	if payload.Comment != nil {
		if len(*payload.Comment) < 10 {
			err = goa.InvalidLengthError(`raw.comment`, *payload.Comment, len(*payload.Comment), 10, true, err)
		}
	}
	if payload.Comment != nil {
		if len(*payload.Comment) > 200 {
			err = goa.InvalidLengthError(`raw.comment`, *payload.Comment, len(*payload.Comment), 200, false, err)
		}
	}
	if payload.Rating < 1 {
		err = goa.InvalidRangeError(`raw.rating`, payload.Rating, 1, true, err)
	}
	if payload.Rating > 5 {
		err = goa.InvalidRangeError(`raw.rating`, payload.Rating, 5, false, err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateReviewContext) Created() error {
	return ctx.RespondBytes(201, nil)
}

// DeleteReviewContext provides the review delete action context.
type DeleteReviewContext struct {
	*goa.Context
	APIVersion string
	ProposalID int
	ReviewID   int
	UserID     int
}

// NewDeleteReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller delete action.
func NewDeleteReviewContext(c *goa.Context) (*DeleteReviewContext, error) {
	var err error
	ctx := DeleteReviewContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawReviewID := c.Get("reviewID")
	if rawReviewID != "" {
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			ctx.ReviewID = int(reviewID)
		} else {
			err = goa.InvalidParamTypeError("reviewID", rawReviewID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteReviewContext) NoContent() error {
	return ctx.RespondBytes(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteReviewContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// ListReviewContext provides the review list action context.
type ListReviewContext struct {
	*goa.Context
	APIVersion string
	ProposalID int
	UserID     int
}

// NewListReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller list action.
func NewListReviewContext(c *goa.Context) (*ListReviewContext, error) {
	var err error
	ctx := ListReviewContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListReviewContext) OK(resp ReviewCollection) error {
	ctx.Header().Set("Content-Type", "application/vnd.review+json; type=collection")
	return ctx.Respond(200, resp)
}

// ShowReviewContext provides the review show action context.
type ShowReviewContext struct {
	*goa.Context
	APIVersion string
	ProposalID int
	ReviewID   int
	UserID     int
}

// NewShowReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller show action.
func NewShowReviewContext(c *goa.Context) (*ShowReviewContext, error) {
	var err error
	ctx := ShowReviewContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawReviewID := c.Get("reviewID")
	if rawReviewID != "" {
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			ctx.ReviewID = int(reviewID)
		} else {
			err = goa.InvalidParamTypeError("reviewID", rawReviewID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowReviewContext) OK(resp *Review) error {
	ctx.Header().Set("Content-Type", "application/vnd.review")
	return ctx.Respond(200, resp)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowReviewContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// UpdateReviewContext provides the review update action context.
type UpdateReviewContext struct {
	*goa.Context
	APIVersion string
	ProposalID int
	ReviewID   int
	UserID     int
	Payload    *UpdateReviewPayload
}

// NewUpdateReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller update action.
func NewUpdateReviewContext(c *goa.Context) (*UpdateReviewContext, error) {
	var err error
	ctx := UpdateReviewContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawReviewID := c.Get("reviewID")
	if rawReviewID != "" {
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			ctx.ReviewID = int(reviewID)
		} else {
			err = goa.InvalidParamTypeError("reviewID", rawReviewID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &ctx, err
}

// UpdateReviewPayload is the review update action payload.
type UpdateReviewPayload struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	Rating  *int    `json:"rating,omitempty" xml:"rating,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateReviewPayload) Validate() (err error) {
	if payload.Comment != nil {
		if len(*payload.Comment) < 10 {
			err = goa.InvalidLengthError(`raw.comment`, *payload.Comment, len(*payload.Comment), 10, true, err)
		}
	}
	if payload.Comment != nil {
		if len(*payload.Comment) > 200 {
			err = goa.InvalidLengthError(`raw.comment`, *payload.Comment, len(*payload.Comment), 200, false, err)
		}
	}
	if payload.Rating != nil {
		if *payload.Rating < 1 {
			err = goa.InvalidRangeError(`raw.rating`, *payload.Rating, 1, true, err)
		}
	}
	if payload.Rating != nil {
		if *payload.Rating > 5 {
			err = goa.InvalidRangeError(`raw.rating`, *payload.Rating, 5, false, err)
		}
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateReviewContext) NoContent() error {
	return ctx.RespondBytes(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateReviewContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}
