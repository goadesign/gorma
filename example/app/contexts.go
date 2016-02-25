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
	"golang.org/x/net/context"
	"strconv"
)

// CallbackAuthContext provides the auth callback action context.
type CallbackAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Provider string
}

// NewCallbackAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller callback action.
func NewCallbackAuthContext(ctx context.Context) (*CallbackAuthContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := CallbackAuthContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawProvider := req.Params.Get("provider")
	if rawProvider != "" {
		rctx.Provider = rawProvider
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CallbackAuthContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/html")
	ctx.ResponseData.WriteHeader(200)
	ctx.ResponseData.Write(resp)
	return nil
}

// OauthAuthContext provides the auth oauth action context.
type OauthAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Provider string
}

// NewOauthAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller oauth action.
func NewOauthAuthContext(ctx context.Context) (*OauthAuthContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := OauthAuthContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawProvider := req.Params.Get("provider")
	if rawProvider != "" {
		rctx.Provider = rawProvider
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *OauthAuthContext) OK(r *Authorize) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.authorize")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// RefreshAuthContext provides the auth refresh action context.
type RefreshAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *RefreshAuthPayload
}

// NewRefreshAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller refresh action.
func NewRefreshAuthContext(ctx context.Context) (*RefreshAuthContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := RefreshAuthContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
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
func (ctx *RefreshAuthContext) Created(r *Authorize) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.authorize+json")
	return ctx.ResponseData.Send(ctx.Context, 201, r)
}

// TokenAuthContext provides the auth token action context.
type TokenAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *TokenAuthPayload
}

// NewTokenAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller token action.
func NewTokenAuthContext(ctx context.Context) (*TokenAuthContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := TokenAuthContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
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
func (ctx *TokenAuthContext) Created(r *Authorize) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.authorize+json")
	return ctx.ResponseData.Send(ctx.Context, 201, r)
}

// CreateProposalContext provides the proposal create action context.
type CreateProposalContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID  int
	Payload *CreateProposalPayload
}

// NewCreateProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller create action.
func NewCreateProposalContext(ctx context.Context) (*CreateProposalContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := CreateProposalContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawUserID := req.Params.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &rctx, err
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
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeleteProposalContext provides the proposal delete action context.
type DeleteProposalContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ProposalID int
	UserID     int
}

// NewDeleteProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller delete action.
func NewDeleteProposalContext(ctx context.Context) (*DeleteProposalContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := DeleteProposalContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawProposalID := req.Params.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := req.Params.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteProposalContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteProposalContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListProposalContext provides the proposal list action context.
type ListProposalContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID int
}

// NewListProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller list action.
func NewListProposalContext(ctx context.Context) (*ListProposalContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ListProposalContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawUserID := req.Params.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListProposalContext) OK(r ProposalCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.proposal+json; type=collection")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// ShowProposalContext provides the proposal show action context.
type ShowProposalContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ProposalID int
	UserID     int
}

// NewShowProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller show action.
func NewShowProposalContext(ctx context.Context) (*ShowProposalContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ShowProposalContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawProposalID := req.Params.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := req.Params.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowProposalContext) OK(r *Proposal) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.proposal")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowProposalContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateProposalContext provides the proposal update action context.
type UpdateProposalContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ProposalID int
	UserID     int
	Payload    *UpdateProposalPayload
}

// NewUpdateProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller update action.
func NewUpdateProposalContext(ctx context.Context) (*UpdateProposalContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := UpdateProposalContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawProposalID := req.Params.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := req.Params.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &rctx, err
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
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateProposalContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateReviewContext provides the review create action context.
type CreateReviewContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ProposalID int
	UserID     int
	Payload    *CreateReviewPayload
}

// NewCreateReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller create action.
func NewCreateReviewContext(ctx context.Context) (*CreateReviewContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := CreateReviewContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawProposalID := req.Params.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := req.Params.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &rctx, err
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
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeleteReviewContext provides the review delete action context.
type DeleteReviewContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ProposalID int
	ReviewID   int
	UserID     int
}

// NewDeleteReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller delete action.
func NewDeleteReviewContext(ctx context.Context) (*DeleteReviewContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := DeleteReviewContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawProposalID := req.Params.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawReviewID := req.Params.Get("reviewID")
	if rawReviewID != "" {
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			rctx.ReviewID = int(reviewID)
		} else {
			err = goa.InvalidParamTypeError("reviewID", rawReviewID, "integer", err)
		}
	}
	rawUserID := req.Params.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteReviewContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteReviewContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListReviewContext provides the review list action context.
type ListReviewContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ProposalID int
	UserID     int
}

// NewListReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller list action.
func NewListReviewContext(ctx context.Context) (*ListReviewContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ListReviewContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawProposalID := req.Params.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := req.Params.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListReviewContext) OK(r ReviewCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.review+json; type=collection")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// ShowReviewContext provides the review show action context.
type ShowReviewContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ProposalID int
	ReviewID   int
	UserID     int
}

// NewShowReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller show action.
func NewShowReviewContext(ctx context.Context) (*ShowReviewContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ShowReviewContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawProposalID := req.Params.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawReviewID := req.Params.Get("reviewID")
	if rawReviewID != "" {
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			rctx.ReviewID = int(reviewID)
		} else {
			err = goa.InvalidParamTypeError("reviewID", rawReviewID, "integer", err)
		}
	}
	rawUserID := req.Params.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowReviewContext) OK(r *Review) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.review")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowReviewContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateReviewContext provides the review update action context.
type UpdateReviewContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ProposalID int
	ReviewID   int
	UserID     int
	Payload    *UpdateReviewPayload
}

// NewUpdateReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller update action.
func NewUpdateReviewContext(ctx context.Context) (*UpdateReviewContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := UpdateReviewContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawProposalID := req.Params.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawReviewID := req.Params.Get("reviewID")
	if rawReviewID != "" {
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			rctx.ReviewID = int(reviewID)
		} else {
			err = goa.InvalidParamTypeError("reviewID", rawReviewID, "integer", err)
		}
	}
	rawUserID := req.Params.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &rctx, err
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
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateReviewContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// BootstrapUIContext provides the ui bootstrap action context.
type BootstrapUIContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewBootstrapUIContext parses the incoming request URL and body, performs validations and creates the
// context used by the ui controller bootstrap action.
func NewBootstrapUIContext(ctx context.Context) (*BootstrapUIContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := BootstrapUIContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *BootstrapUIContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/html")
	ctx.ResponseData.WriteHeader(200)
	ctx.ResponseData.Write(resp)
	return nil
}

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(ctx context.Context) (*CreateUserContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := CreateUserContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
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
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeleteUserContext provides the user delete action context.
type DeleteUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID int
}

// NewDeleteUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller delete action.
func NewDeleteUserContext(ctx context.Context) (*DeleteUserContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := DeleteUserContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawUserID := req.Params.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteUserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListUserContext provides the user list action context.
type ListUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller list action.
func NewListUserContext(ctx context.Context) (*ListUserContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ListUserContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUserContext) OK(r UserCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// ShowUserContext provides the user show action context.
type ShowUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID int
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(ctx context.Context) (*ShowUserContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ShowUserContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawUserID := req.Params.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID  int
	Payload *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(ctx context.Context) (*UpdateUserContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := UpdateUserContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawUserID := req.Params.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &rctx, err
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
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
