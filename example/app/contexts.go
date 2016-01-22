//************************************************************************//
// API "cellar": Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/bketelsen/gorma/example
// --design=github.com/bketelsen/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/goadesign/goa"
)

// CreateAccountContext provides the account create action context.
type CreateAccountContext struct {
	*goa.Context
	Payload *CreateAccountPayload
}

// NewCreateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller create action.
func NewCreateAccountContext(c *goa.Context) (*CreateAccountContext, error) {
	var err error
	ctx := CreateAccountContext{Context: c}
	return &ctx, err
}

// CreateAccountPayload is the account create action payload.
type CreateAccountPayload struct {
	// Name of account
	Name string
}

// Validate runs the validation rules defined in the design.
func (payload *CreateAccountPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MissingAttributeError(`raw`, "name", err)
	}

	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateAccountContext) Created() error {
	return ctx.RespondBytes(201, nil)
}

// DeleteAccountContext provides the account delete action context.
type DeleteAccountContext struct {
	*goa.Context
	AccountID int
}

// NewDeleteAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller delete action.
func NewDeleteAccountContext(c *goa.Context) (*DeleteAccountContext, error) {
	var err error
	ctx := DeleteAccountContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteAccountContext) NoContent() error {
	return ctx.RespondBytes(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteAccountContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// ShowAccountContext provides the account show action context.
type ShowAccountContext struct {
	*goa.Context
	AccountID int
}

// NewShowAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller show action.
func NewShowAccountContext(c *goa.Context) (*ShowAccountContext, error) {
	var err error
	ctx := ShowAccountContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowAccountContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAccountContext) OK(resp *Account, view AccountViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.account+json; charset=utf-8")
	return ctx.Respond(200, r)
}

// UpdateAccountContext provides the account update action context.
type UpdateAccountContext struct {
	*goa.Context
	AccountID int
	Payload   *UpdateAccountPayload
}

// NewUpdateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller update action.
func NewUpdateAccountContext(c *goa.Context) (*UpdateAccountContext, error) {
	var err error
	ctx := UpdateAccountContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &ctx, err
}

// UpdateAccountPayload is the account update action payload.
type UpdateAccountPayload struct {
	// Name of account
	Name string
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateAccountPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MissingAttributeError(`raw`, "name", err)
	}

	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateAccountContext) NoContent() error {
	return ctx.RespondBytes(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateAccountContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// CreateBottleContext provides the bottle create action context.
type CreateBottleContext struct {
	*goa.Context
	AccountID int
	Payload   *CreateBottlePayload
}

// NewCreateBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller create action.
func NewCreateBottleContext(c *goa.Context) (*CreateBottleContext, error) {
	var err error
	ctx := CreateBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &ctx, err
}

// CreateBottlePayload is the bottle create action payload.
type CreateBottlePayload struct {
	Color     string
	Country   *string
	Name      string
	Region    *string
	Review    *string
	Sweetness *int
	Varietal  string
	Vineyard  string
	Vintage   int
}

// Validate runs the validation rules defined in the design.
func (payload *CreateBottlePayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MissingAttributeError(`raw`, "name", err)
	}
	if payload.Vineyard == "" {
		err = goa.MissingAttributeError(`raw`, "vineyard", err)
	}
	if payload.Varietal == "" {
		err = goa.MissingAttributeError(`raw`, "varietal", err)
	}

	if payload.Color == "" {
		err = goa.MissingAttributeError(`raw`, "color", err)
	}

	if !(payload.Color == "red" || payload.Color == "white" || payload.Color == "rose" || payload.Color == "yellow" || payload.Color == "sparkling") {
		err = goa.InvalidEnumValueError(`raw.color`, payload.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
	}
	if payload.Country != nil {
		if len(*payload.Country) < 2 {
			err = goa.InvalidLengthError(`raw.country`, *payload.Country, len(*payload.Country), 2, true, err)
		}
	}
	if len(payload.Name) < 2 {
		err = goa.InvalidLengthError(`raw.name`, payload.Name, len(payload.Name), 2, true, err)
	}
	if payload.Review != nil {
		if len(*payload.Review) < 3 {
			err = goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 3, true, err)
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) > 300 {
			err = goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 300, false, err)
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness < 1 {
			err = goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 1, true, err)
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness > 5 {
			err = goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 5, false, err)
		}
	}
	if len(payload.Varietal) < 4 {
		err = goa.InvalidLengthError(`raw.varietal`, payload.Varietal, len(payload.Varietal), 4, true, err)
	}
	if len(payload.Vineyard) < 2 {
		err = goa.InvalidLengthError(`raw.vineyard`, payload.Vineyard, len(payload.Vineyard), 2, true, err)
	}
	if payload.Vintage < 1900 {
		err = goa.InvalidRangeError(`raw.vintage`, payload.Vintage, 1900, true, err)
	}
	if payload.Vintage > 2020 {
		err = goa.InvalidRangeError(`raw.vintage`, payload.Vintage, 2020, false, err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateBottleContext) Created() error {
	return ctx.RespondBytes(201, nil)
}

// DeleteBottleContext provides the bottle delete action context.
type DeleteBottleContext struct {
	*goa.Context
	AccountID int
	BottleID  int
}

// NewDeleteBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller delete action.
func NewDeleteBottleContext(c *goa.Context) (*DeleteBottleContext, error) {
	var err error
	ctx := DeleteBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := c.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			ctx.BottleID = int(bottleID)
		} else {
			err = goa.InvalidParamTypeError("bottleID", rawBottleID, "integer", err)
		}
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteBottleContext) NoContent() error {
	return ctx.RespondBytes(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteBottleContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// ListBottleContext provides the bottle list action context.
type ListBottleContext struct {
	*goa.Context
	AccountID int
	Years     []int
}

// NewListBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller list action.
func NewListBottleContext(c *goa.Context) (*ListBottleContext, error) {
	var err error
	ctx := ListBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawYears := c.Get("years")
	if rawYears != "" {
		elemsYears := strings.Split(rawYears, ",")
		elemsYears2 := make([]int, len(elemsYears))
		for i, rawElem := range elemsYears {
			if elem, err2 := strconv.Atoi(rawElem); err2 == nil {
				elemsYears2[i] = int(elem)
			} else {
				err = goa.InvalidParamTypeError("elem", rawElem, "integer", err)
			}
		}
		ctx.Years = elemsYears2
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListBottleContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ListBottleContext) OK(resp BottleCollection, view BottleCollectionViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.bottle+json; type=collection; charset=utf-8")
	return ctx.Respond(200, r)
}

// RateBottleContext provides the bottle rate action context.
type RateBottleContext struct {
	*goa.Context
	AccountID int
	BottleID  int
	Payload   *RateBottlePayload
}

// NewRateBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller rate action.
func NewRateBottleContext(c *goa.Context) (*RateBottleContext, error) {
	var err error
	ctx := RateBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := c.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			ctx.BottleID = int(bottleID)
		} else {
			err = goa.InvalidParamTypeError("bottleID", rawBottleID, "integer", err)
		}
	}
	return &ctx, err
}

// RateBottlePayload is the bottle rate action payload.
type RateBottlePayload struct {
	// Rating of bottle between 1 and 5
	Rating int
}

// Validate runs the validation rules defined in the design.
func (payload *RateBottlePayload) Validate() (err error) {

	if payload.Rating < 1 {
		err = goa.InvalidRangeError(`raw.rating`, payload.Rating, 1, true, err)
	}
	if payload.Rating > 5 {
		err = goa.InvalidRangeError(`raw.rating`, payload.Rating, 5, false, err)
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *RateBottleContext) NoContent() error {
	return ctx.RespondBytes(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *RateBottleContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// ShowBottleContext provides the bottle show action context.
type ShowBottleContext struct {
	*goa.Context
	AccountID int
	BottleID  int
}

// NewShowBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller show action.
func NewShowBottleContext(c *goa.Context) (*ShowBottleContext, error) {
	var err error
	ctx := ShowBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := c.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			ctx.BottleID = int(bottleID)
		} else {
			err = goa.InvalidParamTypeError("bottleID", rawBottleID, "integer", err)
		}
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowBottleContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OK(resp *Bottle, view BottleViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.bottle+json; charset=utf-8")
	return ctx.Respond(200, r)
}

// UpdateBottleContext provides the bottle update action context.
type UpdateBottleContext struct {
	*goa.Context
	AccountID int
	BottleID  int
	Payload   *UpdateBottlePayload
}

// NewUpdateBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller update action.
func NewUpdateBottleContext(c *goa.Context) (*UpdateBottleContext, error) {
	var err error
	ctx := UpdateBottleContext{Context: c}
	rawAccountID := c.Get("accountID")
	if rawAccountID != "" {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawBottleID := c.Get("bottleID")
	if rawBottleID != "" {
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			ctx.BottleID = int(bottleID)
		} else {
			err = goa.InvalidParamTypeError("bottleID", rawBottleID, "integer", err)
		}
	}
	return &ctx, err
}

// UpdateBottlePayload is the bottle update action payload.
type UpdateBottlePayload struct {
	Color     *string
	Country   *string
	Name      *string
	Region    *string
	Review    *string
	Sweetness *int
	Varietal  *string
	Vineyard  *string
	Vintage   *int
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateBottlePayload) Validate() (err error) {
	if payload.Color != nil {
		if !(*payload.Color == "red" || *payload.Color == "white" || *payload.Color == "rose" || *payload.Color == "yellow" || *payload.Color == "sparkling") {
			err = goa.InvalidEnumValueError(`raw.color`, *payload.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
		}
	}
	if payload.Country != nil {
		if len(*payload.Country) < 2 {
			err = goa.InvalidLengthError(`raw.country`, *payload.Country, len(*payload.Country), 2, true, err)
		}
	}
	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true, err)
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) < 3 {
			err = goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 3, true, err)
		}
	}
	if payload.Review != nil {
		if len(*payload.Review) > 300 {
			err = goa.InvalidLengthError(`raw.review`, *payload.Review, len(*payload.Review), 300, false, err)
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness < 1 {
			err = goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 1, true, err)
		}
	}
	if payload.Sweetness != nil {
		if *payload.Sweetness > 5 {
			err = goa.InvalidRangeError(`raw.sweetness`, *payload.Sweetness, 5, false, err)
		}
	}
	if payload.Varietal != nil {
		if len(*payload.Varietal) < 4 {
			err = goa.InvalidLengthError(`raw.varietal`, *payload.Varietal, len(*payload.Varietal), 4, true, err)
		}
	}
	if payload.Vineyard != nil {
		if len(*payload.Vineyard) < 2 {
			err = goa.InvalidLengthError(`raw.vineyard`, *payload.Vineyard, len(*payload.Vineyard), 2, true, err)
		}
	}
	if payload.Vintage != nil {
		if *payload.Vintage < 1900 {
			err = goa.InvalidRangeError(`raw.vintage`, *payload.Vintage, 1900, true, err)
		}
	}
	if payload.Vintage != nil {
		if *payload.Vintage > 2020 {
			err = goa.InvalidRangeError(`raw.vintage`, *payload.Vintage, 2020, false, err)
		}
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateBottleContext) NoContent() error {
	return ctx.RespondBytes(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateBottleContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}
