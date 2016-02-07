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
	APIVersion string
	Provider   string
}

// NewCallbackAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller callback action.
func NewCallbackAuthContext(c *goa.Context) (*CallbackAuthContext, error) {
	var err error
	ctx := CallbackAuthContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
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
	APIVersion string
	Provider   string
}

// NewOauthAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller oauth action.
func NewOauthAuthContext(c *goa.Context) (*OauthAuthContext, error) {
	var err error
	ctx := OauthAuthContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
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
	APIVersion string
	Payload    *RefreshAuthPayload
}

// NewRefreshAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller refresh action.
func NewRefreshAuthContext(c *goa.Context) (*RefreshAuthContext, error) {
	var err error
	ctx := RefreshAuthContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
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
	APIVersion string
	Payload    *TokenAuthPayload
}

// NewTokenAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller token action.
func NewTokenAuthContext(c *goa.Context) (*TokenAuthContext, error) {
	var err error
	ctx := TokenAuthContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
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
	APIVersion string
	Payload    *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(c *goa.Context) (*CreateUserContext, error) {
	var err error
	ctx := CreateUserContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
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
	APIVersion string
	UserID     int
}

// NewDeleteUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller delete action.
func NewDeleteUserContext(c *goa.Context) (*DeleteUserContext, error) {
	var err error
	ctx := DeleteUserContext{Context: c}
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
	APIVersion string
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller list action.
func NewListUserContext(c *goa.Context) (*ListUserContext, error) {
	var err error
	ctx := ListUserContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
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
	APIVersion string
	UserID     int
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(c *goa.Context) (*ShowUserContext, error) {
	var err error
	ctx := ShowUserContext{Context: c}
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

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(resp *User) error {
	ctx.Header().Set("Content-Type", "application/vnd.user")
	return ctx.Respond(200, resp)
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	*goa.Context
	APIVersion string
	UserID     int
	Payload    *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(c *goa.Context) (*UpdateUserContext, error) {
	var err error
	ctx := UpdateUserContext{Context: c}
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
