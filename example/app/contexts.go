//************************************************************************//
// API "congo": Application Contexts
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

import (
	"github.com/goadesign/goa"
	"strconv"
)

// CallbackAuthContext provides the auth callback action context.
type CallbackAuthContext struct {
	*goa.Context
	Provider string
}

// NewCallbackAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller callback action.
func NewCallbackAuthContext(c *goa.Context) (*CallbackAuthContext, error) {
	var err error
	ctx := CallbackAuthContext{Context: c}
	rawProvider := c.Get("provider")
	if rawProvider != "" {
		ctx.Provider = rawProvider
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CallbackAuthContext) OK(resp []byte) error {
	ctx.Header().Set("Content-Type", "text/html")
	return ctx.RespondBytes(200, resp)
}

// OauthAuthContext provides the auth oauth action context.
type OauthAuthContext struct {
	*goa.Context
	Provider string
}

// NewOauthAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller oauth action.
func NewOauthAuthContext(c *goa.Context) (*OauthAuthContext, error) {
	var err error
	ctx := OauthAuthContext{Context: c}
	rawProvider := c.Get("provider")
	if rawProvider != "" {
		ctx.Provider = rawProvider
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *OauthAuthContext) OK(resp *Authorize) error {
	ctx.Header().Set("Content-Type", "application/vnd.authorize")
	return ctx.Respond(200, resp)
}

// RefreshAuthContext provides the auth refresh action context.
type RefreshAuthContext struct {
	*goa.Context
	Payload *RefreshAuthPayload
}

// NewRefreshAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller refresh action.
func NewRefreshAuthContext(c *goa.Context) (*RefreshAuthContext, error) {
	var err error
	ctx := RefreshAuthContext{Context: c}
	return &ctx, err
}

// RefreshAuthPayload is the auth refresh action payload.
type RefreshAuthPayload struct {
	// UUID of requesting application
	Application *string `json:"application,omitempty" xml:"application,omitempty"`
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

// Created sends a HTTP response with status code 201.
func (ctx *RefreshAuthContext) Created(resp *Authorize) error {
	ctx.Header().Set("Content-Type", "application/vnd.authorize+json")
	return ctx.Respond(201, resp)
}

// TokenAuthContext provides the auth token action context.
type TokenAuthContext struct {
	*goa.Context
	Payload *TokenAuthPayload
}

// NewTokenAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller token action.
func NewTokenAuthContext(c *goa.Context) (*TokenAuthContext, error) {
	var err error
	ctx := TokenAuthContext{Context: c}
	return &ctx, err
}

// TokenAuthPayload is the auth token action payload.
type TokenAuthPayload struct {
	// UUID of requesting application
	Application *string `json:"application,omitempty" xml:"application,omitempty"`
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

// Created sends a HTTP response with status code 201.
func (ctx *TokenAuthContext) Created(resp *Authorize) error {
	ctx.Header().Set("Content-Type", "application/vnd.authorize+json")
	return ctx.Respond(201, resp)
}

// CreateProposalContext provides the proposal create action context.
type CreateProposalContext struct {
	*goa.Context
	UserID  int
	Payload *CreateProposalPayload
}

// NewCreateProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller create action.
func NewCreateProposalContext(c *goa.Context) (*CreateProposalContext, error) {
	var err error
	ctx := CreateProposalContext{Context: c}
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
	ProposalID int
	UserID     int
}

// NewDeleteProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller delete action.
func NewDeleteProposalContext(c *goa.Context) (*DeleteProposalContext, error) {
	var err error
	ctx := DeleteProposalContext{Context: c}
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
	UserID int
}

// NewListProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller list action.
func NewListProposalContext(c *goa.Context) (*ListProposalContext, error) {
	var err error
	ctx := ListProposalContext{Context: c}
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
	ProposalID int
	UserID     int
}

// NewShowProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller show action.
func NewShowProposalContext(c *goa.Context) (*ShowProposalContext, error) {
	var err error
	ctx := ShowProposalContext{Context: c}
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
	ProposalID int
	UserID     int
	Payload    *UpdateProposalPayload
}

// NewUpdateProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller update action.
func NewUpdateProposalContext(c *goa.Context) (*UpdateProposalContext, error) {
	var err error
	ctx := UpdateProposalContext{Context: c}
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

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateProposalContext) NoContent() error {
	return ctx.RespondBytes(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateProposalContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// CreateReviewContext provides the review create action context.
type CreateReviewContext struct {
	*goa.Context
	ProposalID int
	UserID     int
	Payload    *CreateReviewPayload
}

// NewCreateReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller create action.
func NewCreateReviewContext(c *goa.Context) (*CreateReviewContext, error) {
	var err error
	ctx := CreateReviewContext{Context: c}
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
	ProposalID int
	ReviewID   int
	UserID     int
}

// NewDeleteReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller delete action.
func NewDeleteReviewContext(c *goa.Context) (*DeleteReviewContext, error) {
	var err error
	ctx := DeleteReviewContext{Context: c}
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
	ProposalID int
	UserID     int
}

// NewListReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller list action.
func NewListReviewContext(c *goa.Context) (*ListReviewContext, error) {
	var err error
	ctx := ListReviewContext{Context: c}
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
	ProposalID int
	ReviewID   int
	UserID     int
}

// NewShowReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller show action.
func NewShowReviewContext(c *goa.Context) (*ShowReviewContext, error) {
	var err error
	ctx := ShowReviewContext{Context: c}
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

// BootstrapUiContext provides the ui bootstrap action context.
type BootstrapUiContext struct {
	*goa.Context
}

// NewBootstrapUiContext parses the incoming request URL and body, performs validations and creates the
// context used by the ui controller bootstrap action.
func NewBootstrapUiContext(c *goa.Context) (*BootstrapUiContext, error) {
	var err error
	ctx := BootstrapUiContext{Context: c}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *BootstrapUiContext) OK(resp []byte) error {
	ctx.Header().Set("Content-Type", "text/html")
	return ctx.RespondBytes(200, resp)
}

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	*goa.Context
	Payload *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(c *goa.Context) (*CreateUserContext, error) {
	var err error
	ctx := CreateUserContext{Context: c}
	return &ctx, err
}

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty"`
	City      *string `json:"city,omitempty" xml:"city,omitempty"`
	Country   *string `json:"country,omitempty" xml:"country,omitempty"`
	Email     string  `json:"email" xml:"email"`
	Firstname string  `json:"firstname" xml:"firstname"`
	Lastname  string  `json:"lastname" xml:"lastname"`
	State     *string `json:"state,omitempty" xml:"state,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateUserPayload) Validate() (err error) {
	if payload.Firstname == "" {
		err = goa.MissingAttributeError(`raw`, "firstname", err)
	}

	if payload.Lastname == "" {
		err = goa.MissingAttributeError(`raw`, "lastname", err)
	}

	if payload.Email == "" {
		err = goa.MissingAttributeError(`raw`, "email", err)
	}

	if payload.Bio != nil {
		if len(*payload.Bio) > 500 {
			err = goa.InvalidLengthError(`raw.bio`, *payload.Bio, len(*payload.Bio), 500, false, err)
		}
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, payload.Email); err2 != nil {
		err = goa.InvalidFormatError(`raw.email`, payload.Email, goa.FormatEmail, err2, err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUserContext) Created() error {
	return ctx.RespondBytes(201, nil)
}

// DeleteUserContext provides the user delete action context.
type DeleteUserContext struct {
	*goa.Context
	UserID int
}

// NewDeleteUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller delete action.
func NewDeleteUserContext(c *goa.Context) (*DeleteUserContext, error) {
	var err error
	ctx := DeleteUserContext{Context: c}
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
func (ctx *DeleteUserContext) NoContent() error {
	return ctx.RespondBytes(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteUserContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// ListUserContext provides the user list action context.
type ListUserContext struct {
	*goa.Context
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller list action.
func NewListUserContext(c *goa.Context) (*ListUserContext, error) {
	var err error
	ctx := ListUserContext{Context: c}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUserContext) OK(resp UserCollection) error {
	ctx.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	return ctx.Respond(200, resp)
}

// ShowUserContext provides the user show action context.
type ShowUserContext struct {
	*goa.Context
	UserID int
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(c *goa.Context) (*ShowUserContext, error) {
	var err error
	ctx := ShowUserContext{Context: c}
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
func (ctx *ShowUserContext) OK(resp *User) error {
	ctx.Header().Set("Content-Type", "application/vnd.user")
	return ctx.Respond(200, resp)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	*goa.Context
	UserID  int
	Payload *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(c *goa.Context) (*UpdateUserContext, error) {
	var err error
	ctx := UpdateUserContext{Context: c}
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

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty"`
	City      *string `json:"city,omitempty" xml:"city,omitempty"`
	Country   *string `json:"country,omitempty" xml:"country,omitempty"`
	Email     string  `json:"email" xml:"email"`
	Firstname *string `json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname  *string `json:"lastname,omitempty" xml:"lastname,omitempty"`
	State     *string `json:"state,omitempty" xml:"state,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateUserPayload) Validate() (err error) {
	if payload.Email == "" {
		err = goa.MissingAttributeError(`raw`, "email", err)
	}

	if payload.Bio != nil {
		if len(*payload.Bio) > 500 {
			err = goa.InvalidLengthError(`raw.bio`, *payload.Bio, len(*payload.Bio), 500, false, err)
		}
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, payload.Email); err2 != nil {
		err = goa.InvalidFormatError(`raw.email`, payload.Email, goa.FormatEmail, err2, err)
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateUserContext) NoContent() error {
	return ctx.RespondBytes(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUserContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}
