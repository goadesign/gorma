//************************************************************************//
// API "congo": Application Contexts
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/goadesign/gorma/example/design
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --version=v1.0.0
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
func NewCallbackAuthContext(ctx context.Context, service *goa.Service) (*CallbackAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CallbackAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProvider := req.Params["provider"]
	if len(paramProvider) > 0 {
		rawProvider := paramProvider[0]
		rctx.Provider = rawProvider
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CallbackAuthContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/html")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
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
func NewOauthAuthContext(ctx context.Context, service *goa.Service) (*OauthAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := OauthAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProvider := req.Params["provider"]
	if len(paramProvider) > 0 {
		rawProvider := paramProvider[0]
		rctx.Provider = rawProvider
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *OauthAuthContext) OK(r *Authorize) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.authorize+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
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
func NewRefreshAuthContext(ctx context.Context, service *goa.Service) (*RefreshAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := RefreshAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// refreshAuthPayload is the auth refresh action payload.
type refreshAuthPayload struct {
	// UUID of requesting application
	Application *string `form:"application,omitempty" json:"application,omitempty" xml:"application,omitempty"`
	// email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Publicize creates RefreshAuthPayload from refreshAuthPayload
func (payload *refreshAuthPayload) Publicize() *RefreshAuthPayload {
	var pub RefreshAuthPayload
	if payload.Application != nil {
		pub.Application = payload.Application
	}
	if payload.Email != nil {
		pub.Email = payload.Email
	}
	if payload.Password != nil {
		pub.Password = payload.Password
	}
	return &pub
}

// RefreshAuthPayload is the auth refresh action payload.
type RefreshAuthPayload struct {
	// UUID of requesting application
	Application *string `form:"application,omitempty" json:"application,omitempty" xml:"application,omitempty"`
	// email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Created sends a HTTP response with status code 201.
func (ctx *RefreshAuthContext) Created(r *Authorize) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.authorize+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
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
func NewTokenAuthContext(ctx context.Context, service *goa.Service) (*TokenAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := TokenAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// tokenAuthPayload is the auth token action payload.
type tokenAuthPayload struct {
	// UUID of requesting application
	Application *string `form:"application,omitempty" json:"application,omitempty" xml:"application,omitempty"`
	// email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Publicize creates TokenAuthPayload from tokenAuthPayload
func (payload *tokenAuthPayload) Publicize() *TokenAuthPayload {
	var pub TokenAuthPayload
	if payload.Application != nil {
		pub.Application = payload.Application
	}
	if payload.Email != nil {
		pub.Email = payload.Email
	}
	if payload.Password != nil {
		pub.Password = payload.Password
	}
	return &pub
}

// TokenAuthPayload is the auth token action payload.
type TokenAuthPayload struct {
	// UUID of requesting application
	Application *string `form:"application,omitempty" json:"application,omitempty" xml:"application,omitempty"`
	// email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Created sends a HTTP response with status code 201.
func (ctx *TokenAuthContext) Created(r *Authorize) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.authorize+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
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
func NewCreateProposalContext(ctx context.Context, service *goa.Service) (*CreateProposalContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateProposalContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// createProposalPayload is the proposal create action payload.
type createProposalPayload struct {
	Abstract  *string `form:"abstract,omitempty" json:"abstract,omitempty" xml:"abstract,omitempty"`
	Detail    *string `form:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
	Title     *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	Withdrawn *bool   `form:"withdrawn,omitempty" json:"withdrawn,omitempty" xml:"withdrawn,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createProposalPayload) Validate() (err error) {
	if payload.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "title"))
	}
	if payload.Abstract == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "abstract"))
	}
	if payload.Detail == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "detail"))
	}

	if payload.Abstract != nil {
		if len(*payload.Abstract) < 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.abstract`, *payload.Abstract, len(*payload.Abstract), 50, true))
		}
	}
	if payload.Abstract != nil {
		if len(*payload.Abstract) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.abstract`, *payload.Abstract, len(*payload.Abstract), 500, false))
		}
	}
	if payload.Detail != nil {
		if len(*payload.Detail) < 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.detail`, *payload.Detail, len(*payload.Detail), 100, true))
		}
	}
	if payload.Detail != nil {
		if len(*payload.Detail) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.detail`, *payload.Detail, len(*payload.Detail), 2000, false))
		}
	}
	if payload.Title != nil {
		if len(*payload.Title) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.title`, *payload.Title, len(*payload.Title), 10, true))
		}
	}
	if payload.Title != nil {
		if len(*payload.Title) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.title`, *payload.Title, len(*payload.Title), 200, false))
		}
	}
	return
}

// Publicize creates CreateProposalPayload from createProposalPayload
func (payload *createProposalPayload) Publicize() *CreateProposalPayload {
	var pub CreateProposalPayload
	if payload.Abstract != nil {
		pub.Abstract = *payload.Abstract
	}
	if payload.Detail != nil {
		pub.Detail = *payload.Detail
	}
	if payload.Title != nil {
		pub.Title = *payload.Title
	}
	if payload.Withdrawn != nil {
		pub.Withdrawn = payload.Withdrawn
	}
	return &pub
}

// CreateProposalPayload is the proposal create action payload.
type CreateProposalPayload struct {
	Abstract  string `form:"abstract" json:"abstract" xml:"abstract"`
	Detail    string `form:"detail" json:"detail" xml:"detail"`
	Title     string `form:"title" json:"title" xml:"title"`
	Withdrawn *bool  `form:"withdrawn,omitempty" json:"withdrawn,omitempty" xml:"withdrawn,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateProposalPayload) Validate() (err error) {
	if payload.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "title"))
	}
	if payload.Abstract == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "abstract"))
	}
	if payload.Detail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "detail"))
	}

	if len(payload.Abstract) < 50 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.abstract`, payload.Abstract, len(payload.Abstract), 50, true))
	}
	if len(payload.Abstract) > 500 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.abstract`, payload.Abstract, len(payload.Abstract), 500, false))
	}
	if len(payload.Detail) < 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.detail`, payload.Detail, len(payload.Detail), 100, true))
	}
	if len(payload.Detail) > 2000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.detail`, payload.Detail, len(payload.Detail), 2000, false))
	}
	if len(payload.Title) < 10 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.title`, payload.Title, len(payload.Title), 10, true))
	}
	if len(payload.Title) > 200 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.title`, payload.Title, len(payload.Title), 200, false))
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
func NewDeleteProposalContext(ctx context.Context, service *goa.Service) (*DeleteProposalContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeleteProposalContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProposalID := req.Params["proposalID"]
	if len(paramProposalID) > 0 {
		rawProposalID := paramProposalID[0]
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = proposalID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("proposalID", rawProposalID, "integer"))
		}
	}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
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
func NewListProposalContext(ctx context.Context, service *goa.Service) (*ListProposalContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListProposalContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListProposalContext) OK(r ProposalCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.proposal+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
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
func NewShowProposalContext(ctx context.Context, service *goa.Service) (*ShowProposalContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowProposalContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProposalID := req.Params["proposalID"]
	if len(paramProposalID) > 0 {
		rawProposalID := paramProposalID[0]
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = proposalID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("proposalID", rawProposalID, "integer"))
		}
	}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowProposalContext) OK(r *Proposal) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.proposal+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowProposalContext) OKLink(r *ProposalLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.proposal+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
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
	Payload    *ProposalPayload
}

// NewUpdateProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller update action.
func NewUpdateProposalContext(ctx context.Context, service *goa.Service) (*UpdateProposalContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpdateProposalContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProposalID := req.Params["proposalID"]
	if len(paramProposalID) > 0 {
		rawProposalID := paramProposalID[0]
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = proposalID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("proposalID", rawProposalID, "integer"))
		}
	}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
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
func NewCreateReviewContext(ctx context.Context, service *goa.Service) (*CreateReviewContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateReviewContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProposalID := req.Params["proposalID"]
	if len(paramProposalID) > 0 {
		rawProposalID := paramProposalID[0]
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = proposalID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("proposalID", rawProposalID, "integer"))
		}
	}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// createReviewPayload is the review create action payload.
type createReviewPayload struct {
	Comment *string `form:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
	Rating  *int    `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createReviewPayload) Validate() (err error) {
	if payload.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "rating"))
	}

	if payload.Comment != nil {
		if len(*payload.Comment) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.comment`, *payload.Comment, len(*payload.Comment), 10, true))
		}
	}
	if payload.Comment != nil {
		if len(*payload.Comment) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.comment`, *payload.Comment, len(*payload.Comment), 200, false))
		}
	}
	if payload.Rating != nil {
		if *payload.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 1, true))
		}
	}
	if payload.Rating != nil {
		if *payload.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 5, false))
		}
	}
	return
}

// Publicize creates CreateReviewPayload from createReviewPayload
func (payload *createReviewPayload) Publicize() *CreateReviewPayload {
	var pub CreateReviewPayload
	if payload.Comment != nil {
		pub.Comment = payload.Comment
	}
	if payload.Rating != nil {
		pub.Rating = *payload.Rating
	}
	return &pub
}

// CreateReviewPayload is the review create action payload.
type CreateReviewPayload struct {
	Comment *string `form:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
	Rating  int     `form:"rating" json:"rating" xml:"rating"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateReviewPayload) Validate() (err error) {
	if payload.Comment != nil {
		if len(*payload.Comment) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.comment`, *payload.Comment, len(*payload.Comment), 10, true))
		}
	}
	if payload.Comment != nil {
		if len(*payload.Comment) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.comment`, *payload.Comment, len(*payload.Comment), 200, false))
		}
	}
	if payload.Rating < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, payload.Rating, 1, true))
	}
	if payload.Rating > 5 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, payload.Rating, 5, false))
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
func NewDeleteReviewContext(ctx context.Context, service *goa.Service) (*DeleteReviewContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeleteReviewContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProposalID := req.Params["proposalID"]
	if len(paramProposalID) > 0 {
		rawProposalID := paramProposalID[0]
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = proposalID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("proposalID", rawProposalID, "integer"))
		}
	}
	paramReviewID := req.Params["reviewID"]
	if len(paramReviewID) > 0 {
		rawReviewID := paramReviewID[0]
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			rctx.ReviewID = reviewID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("reviewID", rawReviewID, "integer"))
		}
	}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
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
func NewListReviewContext(ctx context.Context, service *goa.Service) (*ListReviewContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListReviewContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProposalID := req.Params["proposalID"]
	if len(paramProposalID) > 0 {
		rawProposalID := paramProposalID[0]
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = proposalID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("proposalID", rawProposalID, "integer"))
		}
	}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListReviewContext) OK(r ReviewCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.review+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListReviewContext) OKLink(r ReviewLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.review+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
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
func NewShowReviewContext(ctx context.Context, service *goa.Service) (*ShowReviewContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowReviewContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProposalID := req.Params["proposalID"]
	if len(paramProposalID) > 0 {
		rawProposalID := paramProposalID[0]
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = proposalID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("proposalID", rawProposalID, "integer"))
		}
	}
	paramReviewID := req.Params["reviewID"]
	if len(paramReviewID) > 0 {
		rawReviewID := paramReviewID[0]
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			rctx.ReviewID = reviewID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("reviewID", rawReviewID, "integer"))
		}
	}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowReviewContext) OK(r *Review) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.review+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowReviewContext) OKLink(r *ReviewLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.review+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
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
	Payload    *ReviewPayload
}

// NewUpdateReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller update action.
func NewUpdateReviewContext(ctx context.Context, service *goa.Service) (*UpdateReviewContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpdateReviewContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProposalID := req.Params["proposalID"]
	if len(paramProposalID) > 0 {
		rawProposalID := paramProposalID[0]
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			rctx.ProposalID = proposalID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("proposalID", rawProposalID, "integer"))
		}
	}
	paramReviewID := req.Params["reviewID"]
	if len(paramReviewID) > 0 {
		rawReviewID := paramReviewID[0]
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			rctx.ReviewID = reviewID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("reviewID", rawReviewID, "integer"))
		}
	}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
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
func NewBootstrapUIContext(ctx context.Context, service *goa.Service) (*BootstrapUIContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := BootstrapUIContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *BootstrapUIContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/html")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
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
func NewCreateUserContext(ctx context.Context, service *goa.Service) (*CreateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createUserPayload is the user create action payload.
type createUserPayload struct {
	Bio       *string `form:"bio,omitempty" json:"bio,omitempty" xml:"bio,omitempty"`
	City      *string `form:"city,omitempty" json:"city,omitempty" xml:"city,omitempty"`
	Country   *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	Email     *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname  *string `form:"lastname,omitempty" json:"lastname,omitempty" xml:"lastname,omitempty"`
	State     *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createUserPayload) Validate() (err error) {
	if payload.Firstname == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "firstname"))
	}
	if payload.Lastname == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "lastname"))
	}
	if payload.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}

	if payload.Bio != nil {
		if len(*payload.Bio) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.bio`, *payload.Bio, len(*payload.Bio), 500, false))
		}
	}
	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// Publicize creates CreateUserPayload from createUserPayload
func (payload *createUserPayload) Publicize() *CreateUserPayload {
	var pub CreateUserPayload
	if payload.Bio != nil {
		pub.Bio = payload.Bio
	}
	if payload.City != nil {
		pub.City = payload.City
	}
	if payload.Country != nil {
		pub.Country = payload.Country
	}
	if payload.Email != nil {
		pub.Email = *payload.Email
	}
	if payload.Firstname != nil {
		pub.Firstname = *payload.Firstname
	}
	if payload.Lastname != nil {
		pub.Lastname = *payload.Lastname
	}
	if payload.State != nil {
		pub.State = payload.State
	}
	return &pub
}

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	Bio       *string `form:"bio,omitempty" json:"bio,omitempty" xml:"bio,omitempty"`
	City      *string `form:"city,omitempty" json:"city,omitempty" xml:"city,omitempty"`
	Country   *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	Email     string  `form:"email" json:"email" xml:"email"`
	Firstname string  `form:"firstname" json:"firstname" xml:"firstname"`
	Lastname  string  `form:"lastname" json:"lastname" xml:"lastname"`
	State     *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateUserPayload) Validate() (err error) {
	if payload.Firstname == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "firstname"))
	}
	if payload.Lastname == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "lastname"))
	}
	if payload.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}

	if payload.Bio != nil {
		if len(*payload.Bio) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.bio`, *payload.Bio, len(*payload.Bio), 500, false))
		}
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, payload.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, payload.Email, goa.FormatEmail, err2))
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
func NewDeleteUserContext(ctx context.Context, service *goa.Service) (*DeleteUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeleteUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
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
func NewListUserContext(ctx context.Context, service *goa.Service) (*ListUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUserContext) OK(r UserCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
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
func NewShowUserContext(ctx context.Context, service *goa.Service) (*ShowUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OKLink(r *UserLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
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
func NewUpdateUserContext(ctx context.Context, service *goa.Service) (*UpdateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpdateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// updateUserPayload is the user update action payload.
type updateUserPayload struct {
	Bio       *string `form:"bio,omitempty" json:"bio,omitempty" xml:"bio,omitempty"`
	City      *string `form:"city,omitempty" json:"city,omitempty" xml:"city,omitempty"`
	Country   *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	Email     *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname  *string `form:"lastname,omitempty" json:"lastname,omitempty" xml:"lastname,omitempty"`
	State     *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateUserPayload) Validate() (err error) {
	if payload.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}

	if payload.Bio != nil {
		if len(*payload.Bio) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.bio`, *payload.Bio, len(*payload.Bio), 500, false))
		}
	}
	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// Publicize creates UpdateUserPayload from updateUserPayload
func (payload *updateUserPayload) Publicize() *UpdateUserPayload {
	var pub UpdateUserPayload
	if payload.Bio != nil {
		pub.Bio = payload.Bio
	}
	if payload.City != nil {
		pub.City = payload.City
	}
	if payload.Country != nil {
		pub.Country = payload.Country
	}
	if payload.Email != nil {
		pub.Email = *payload.Email
	}
	if payload.Firstname != nil {
		pub.Firstname = payload.Firstname
	}
	if payload.Lastname != nil {
		pub.Lastname = payload.Lastname
	}
	if payload.State != nil {
		pub.State = payload.State
	}
	return &pub
}

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
	Bio       *string `form:"bio,omitempty" json:"bio,omitempty" xml:"bio,omitempty"`
	City      *string `form:"city,omitempty" json:"city,omitempty" xml:"city,omitempty"`
	Country   *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	Email     string  `form:"email" json:"email" xml:"email"`
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname  *string `form:"lastname,omitempty" json:"lastname,omitempty" xml:"lastname,omitempty"`
	State     *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateUserPayload) Validate() (err error) {
	if payload.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}

	if payload.Bio != nil {
		if len(*payload.Bio) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.bio`, *payload.Bio, len(*payload.Bio), 500, false))
		}
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, payload.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, payload.Email, goa.FormatEmail, err2))
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
