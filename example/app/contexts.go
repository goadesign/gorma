//************************************************************************//
// API "congo": Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/bketelsen/gorma/example
// --design=github.com/bketelsen/gorma/example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"fmt"

	"github.com/raphael/goa"
)

// CallbackAuthContext provides the auth callback action context.
type CallbackAuthContext struct {
	*goa.Context
	Provider string
	Version  string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CallbackAuthContext) OK(resp []byte) error {
	return ctx.Respond(200, resp)
}

// OauthAuthContext provides the auth oauth action context.
type OauthAuthContext struct {
	*goa.Context
	Provider string
	Version  string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *OauthAuthContext) OK(resp *Authorize) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.authorize+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// RefreshAuthContext provides the auth refresh action context.
type RefreshAuthContext struct {
	*goa.Context
	Version string
	Payload *RefreshAuthPayload
}

// NewRefreshAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller refresh action.
func NewRefreshAuthContext(c *goa.Context) (*RefreshAuthContext, error) {
	var err error
	ctx := RefreshAuthContext{Context: c}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewRefreshAuthPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// RefreshAuthPayload is the auth refresh action payload.
type RefreshAuthPayload struct {
	// UUID of requesting application
	Application *string
	// email
	Email *string
	// password
	Password *string
}

// NewRefreshAuthPayload instantiates a RefreshAuthPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewRefreshAuthPayload(raw interface{}) (p *RefreshAuthPayload, err error) {
	p, err = UnmarshalRefreshAuthPayload(raw, err)
	return
}

// UnmarshalRefreshAuthPayload unmarshals and validates a raw interface{} into an instance of RefreshAuthPayload
func UnmarshalRefreshAuthPayload(source interface{}, inErr error) (target *RefreshAuthPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(RefreshAuthPayload)
		if v, ok := val["application"]; ok {
			var tmp1 string
			if val, ok := v.(string); ok {
				tmp1 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Application`, v, "string", err)
			}
			target.Application = &tmp1
		}
		if v, ok := val["email"]; ok {
			var tmp2 string
			if val, ok := v.(string); ok {
				tmp2 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			target.Email = &tmp2
		}
		if v, ok := val["password"]; ok {
			var tmp3 string
			if val, ok := v.(string); ok {
				tmp3 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Password`, v, "string", err)
			}
			target.Password = &tmp3
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *RefreshAuthContext) Created(resp *Authorize) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.authorize+json; charset=utf-8")
	return ctx.JSON(201, r)
}

// TokenAuthContext provides the auth token action context.
type TokenAuthContext struct {
	*goa.Context
	Version string
	Payload *TokenAuthPayload
}

// NewTokenAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller token action.
func NewTokenAuthContext(c *goa.Context) (*TokenAuthContext, error) {
	var err error
	ctx := TokenAuthContext{Context: c}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewTokenAuthPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// TokenAuthPayload is the auth token action payload.
type TokenAuthPayload struct {
	// UUID of requesting application
	Application *string
	// email
	Email *string
	// password
	Password *string
}

// NewTokenAuthPayload instantiates a TokenAuthPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewTokenAuthPayload(raw interface{}) (p *TokenAuthPayload, err error) {
	p, err = UnmarshalTokenAuthPayload(raw, err)
	return
}

// UnmarshalTokenAuthPayload unmarshals and validates a raw interface{} into an instance of TokenAuthPayload
func UnmarshalTokenAuthPayload(source interface{}, inErr error) (target *TokenAuthPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(TokenAuthPayload)
		if v, ok := val["application"]; ok {
			var tmp4 string
			if val, ok := v.(string); ok {
				tmp4 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Application`, v, "string", err)
			}
			target.Application = &tmp4
		}
		if v, ok := val["email"]; ok {
			var tmp5 string
			if val, ok := v.(string); ok {
				tmp5 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			target.Email = &tmp5
		}
		if v, ok := val["password"]; ok {
			var tmp6 string
			if val, ok := v.(string); ok {
				tmp6 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Password`, v, "string", err)
			}
			target.Password = &tmp6
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *TokenAuthContext) Created(resp *Authorize) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.authorize+json; charset=utf-8")
	return ctx.JSON(201, r)
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
	return ctx.Respond(200, resp)
}
