//************************************************************************//
// API "congo" version v1: Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/bketelsen/gorma/example
// --design=github.com/bketelsen/gorma/example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package v1

import (
	"fmt"
	"strconv"

	"github.com/raphael/goa"
)

// CreateProposalContext provides the proposal create action context.
type CreateProposalContext struct {
	*goa.Context
	UserID  int
	Version string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewCreateProposalPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateProposalPayload is the proposal create action payload.
type CreateProposalPayload struct {
	Abstract  string
	CreatedAt *string
	DeletedAt *string
	Detail    string
	FirstName *string
	ID        int
	M2reviews *string
	Reviews   *string
	Title     string
	UpdatedAt *string
	UserId    *int
	Withdrawn *bool
}

// NewCreateProposalPayload instantiates a CreateProposalPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateProposalPayload(raw interface{}) (p *CreateProposalPayload, err error) {
	p, err = UnmarshalCreateProposalPayload(raw, err)
	return
}

// UnmarshalCreateProposalPayload unmarshals and validates a raw interface{} into an instance of CreateProposalPayload
func UnmarshalCreateProposalPayload(source interface{}, inErr error) (target *CreateProposalPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateProposalPayload)
		if v, ok := val["abstract"]; ok {
			var tmp48 string
			if val, ok := v.(string); ok {
				tmp48 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Abstract`, v, "string", err)
			}
			target.Abstract = tmp48
		} else {
			err = goa.MissingAttributeError(`payload`, "abstract", err)
		}
		if v, ok := val["created_at"]; ok {
			var tmp49 string
			if val, ok := v.(string); ok {
				tmp49 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.CreatedAt`, v, "string", err)
			}
			target.CreatedAt = &tmp49
		}
		if v, ok := val["deleted_at"]; ok {
			var tmp50 string
			if val, ok := v.(string); ok {
				tmp50 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.DeletedAt`, v, "string", err)
			}
			target.DeletedAt = &tmp50
		}
		if v, ok := val["detail"]; ok {
			var tmp51 string
			if val, ok := v.(string); ok {
				tmp51 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Detail`, v, "string", err)
			}
			target.Detail = tmp51
		} else {
			err = goa.MissingAttributeError(`payload`, "detail", err)
		}
		if v, ok := val["first_name"]; ok {
			var tmp52 string
			if val, ok := v.(string); ok {
				tmp52 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.FirstName`, v, "string", err)
			}
			target.FirstName = &tmp52
		}
		if v, ok := val["id"]; ok {
			var tmp53 int
			if f, ok := v.(float64); ok {
				tmp53 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "int", err)
			}
			target.ID = tmp53
		} else {
			err = goa.MissingAttributeError(`payload`, "id", err)
		}
		if v, ok := val["m2reviews"]; ok {
			var tmp54 string
			if val, ok := v.(string); ok {
				tmp54 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.M2reviews`, v, "string", err)
			}
			target.M2reviews = &tmp54
		}
		if v, ok := val["reviews"]; ok {
			var tmp55 string
			if val, ok := v.(string); ok {
				tmp55 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Reviews`, v, "string", err)
			}
			target.Reviews = &tmp55
		}
		if v, ok := val["title"]; ok {
			var tmp56 string
			if val, ok := v.(string); ok {
				tmp56 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Title`, v, "string", err)
			}
			target.Title = tmp56
		} else {
			err = goa.MissingAttributeError(`payload`, "title", err)
		}
		if v, ok := val["updated_at"]; ok {
			var tmp57 string
			if val, ok := v.(string); ok {
				tmp57 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.UpdatedAt`, v, "string", err)
			}
			target.UpdatedAt = &tmp57
		}
		if v, ok := val["user_id"]; ok {
			var tmp58 int
			if f, ok := v.(float64); ok {
				tmp58 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.UserId`, v, "int", err)
			}
			target.UserId = &tmp58
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp59 bool
			if val, ok := v.(bool); ok {
				tmp59 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = &tmp59
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}

// Validate validates the type instance.
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
	if payload.FirstName != nil {
		if len(*payload.FirstName) < 2 {
			err = goa.InvalidLengthError(`raw.first_name`, *payload.FirstName, len(*payload.FirstName), 2, true, err)
		}
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
	return ctx.Respond(201, nil)
}

// DeleteProposalContext provides the proposal delete action context.
type DeleteProposalContext struct {
	*goa.Context
	ProposalID int
	UserID     int
	Version    string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteProposalContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteProposalContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListProposalContext provides the proposal list action context.
type ListProposalContext struct {
	*goa.Context
	UserID  int
	Version string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListProposalContext) OK(resp ProposalCollection) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.proposal+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowProposalContext provides the proposal show action context.
type ShowProposalContext struct {
	*goa.Context
	ProposalID int
	UserID     int
	Version    string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowProposalContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowProposalContext) OK(resp *Proposal, view ProposalViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.proposal+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateProposalContext provides the proposal update action context.
type UpdateProposalContext struct {
	*goa.Context
	ProposalID int
	UserID     int
	Version    string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewUpdateProposalPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateProposalPayload is the proposal update action payload.
type UpdateProposalPayload struct {
	Abstract  *string
	CreatedAt *string
	DeletedAt *string
	Detail    *string
	FirstName *string
	ID        int
	M2reviews *string
	Reviews   *string
	Title     *string
	UpdatedAt *string
	UserId    *int
	Withdrawn *bool
}

// NewUpdateProposalPayload instantiates a UpdateProposalPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateProposalPayload(raw interface{}) (p *UpdateProposalPayload, err error) {
	p, err = UnmarshalUpdateProposalPayload(raw, err)
	return
}

// UnmarshalUpdateProposalPayload unmarshals and validates a raw interface{} into an instance of UpdateProposalPayload
func UnmarshalUpdateProposalPayload(source interface{}, inErr error) (target *UpdateProposalPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateProposalPayload)
		if v, ok := val["abstract"]; ok {
			var tmp67 string
			if val, ok := v.(string); ok {
				tmp67 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Abstract`, v, "string", err)
			}
			target.Abstract = &tmp67
		}
		if v, ok := val["created_at"]; ok {
			var tmp68 string
			if val, ok := v.(string); ok {
				tmp68 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.CreatedAt`, v, "string", err)
			}
			target.CreatedAt = &tmp68
		}
		if v, ok := val["deleted_at"]; ok {
			var tmp69 string
			if val, ok := v.(string); ok {
				tmp69 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.DeletedAt`, v, "string", err)
			}
			target.DeletedAt = &tmp69
		}
		if v, ok := val["detail"]; ok {
			var tmp70 string
			if val, ok := v.(string); ok {
				tmp70 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Detail`, v, "string", err)
			}
			target.Detail = &tmp70
		}
		if v, ok := val["first_name"]; ok {
			var tmp71 string
			if val, ok := v.(string); ok {
				tmp71 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.FirstName`, v, "string", err)
			}
			target.FirstName = &tmp71
		}
		if v, ok := val["id"]; ok {
			var tmp72 int
			if f, ok := v.(float64); ok {
				tmp72 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "int", err)
			}
			target.ID = tmp72
		} else {
			err = goa.MissingAttributeError(`payload`, "id", err)
		}
		if v, ok := val["m2reviews"]; ok {
			var tmp73 string
			if val, ok := v.(string); ok {
				tmp73 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.M2reviews`, v, "string", err)
			}
			target.M2reviews = &tmp73
		}
		if v, ok := val["reviews"]; ok {
			var tmp74 string
			if val, ok := v.(string); ok {
				tmp74 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Reviews`, v, "string", err)
			}
			target.Reviews = &tmp74
		}
		if v, ok := val["title"]; ok {
			var tmp75 string
			if val, ok := v.(string); ok {
				tmp75 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Title`, v, "string", err)
			}
			target.Title = &tmp75
		}
		if v, ok := val["updated_at"]; ok {
			var tmp76 string
			if val, ok := v.(string); ok {
				tmp76 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.UpdatedAt`, v, "string", err)
			}
			target.UpdatedAt = &tmp76
		}
		if v, ok := val["user_id"]; ok {
			var tmp77 int
			if f, ok := v.(float64); ok {
				tmp77 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.UserId`, v, "int", err)
			}
			target.UserId = &tmp77
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp78 bool
			if val, ok := v.(bool); ok {
				tmp78 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = &tmp78
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}

// Validate validates the type instance.
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
	if payload.FirstName != nil {
		if len(*payload.FirstName) < 2 {
			err = goa.InvalidLengthError(`raw.first_name`, *payload.FirstName, len(*payload.FirstName), 2, true, err)
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
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateProposalContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// CreateReviewContext provides the review create action context.
type CreateReviewContext struct {
	*goa.Context
	ProposalID int
	UserID     int
	Version    string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewCreateReviewPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateReviewPayload is the review create action payload.
type CreateReviewPayload struct {
	User       *string
	Comment    *string
	ID         int
	ProposalId *int
	Rating     int
	Reviewers  *string
}

// NewCreateReviewPayload instantiates a CreateReviewPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateReviewPayload(raw interface{}) (p *CreateReviewPayload, err error) {
	p, err = UnmarshalCreateReviewPayload(raw, err)
	return
}

// UnmarshalCreateReviewPayload unmarshals and validates a raw interface{} into an instance of CreateReviewPayload
func UnmarshalCreateReviewPayload(source interface{}, inErr error) (target *CreateReviewPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateReviewPayload)
		if v, ok := val["User"]; ok {
			var tmp81 string
			if val, ok := v.(string); ok {
				tmp81 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.User`, v, "string", err)
			}
			target.User = &tmp81
		}
		if v, ok := val["comment"]; ok {
			var tmp82 string
			if val, ok := v.(string); ok {
				tmp82 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Comment`, v, "string", err)
			}
			target.Comment = &tmp82
		}
		if v, ok := val["id"]; ok {
			var tmp83 int
			if f, ok := v.(float64); ok {
				tmp83 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "int", err)
			}
			target.ID = tmp83
		} else {
			err = goa.MissingAttributeError(`payload`, "id", err)
		}
		if v, ok := val["proposal_id"]; ok {
			var tmp84 int
			if f, ok := v.(float64); ok {
				tmp84 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ProposalId`, v, "int", err)
			}
			target.ProposalId = &tmp84
		}
		if v, ok := val["rating"]; ok {
			var tmp85 int
			if f, ok := v.(float64); ok {
				tmp85 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Rating`, v, "int", err)
			}
			target.Rating = tmp85
		} else {
			err = goa.MissingAttributeError(`payload`, "rating", err)
		}
		if v, ok := val["reviewers"]; ok {
			var tmp86 string
			if val, ok := v.(string); ok {
				tmp86 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Reviewers`, v, "string", err)
			}
			target.Reviewers = &tmp86
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}

// Validate validates the type instance.
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
	return ctx.Respond(201, nil)
}

// DeleteReviewContext provides the review delete action context.
type DeleteReviewContext struct {
	*goa.Context
	ProposalID int
	ReviewID   int
	UserID     int
	Version    string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteReviewContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteReviewContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListReviewContext provides the review list action context.
type ListReviewContext struct {
	*goa.Context
	ProposalID int
	UserID     int
	Version    string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListReviewContext) OK(resp ReviewCollection) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.review+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowReviewContext provides the review show action context.
type ShowReviewContext struct {
	*goa.Context
	ProposalID int
	ReviewID   int
	UserID     int
	Version    string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowReviewContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowReviewContext) OK(resp *Review, view ReviewViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.review+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateReviewContext provides the review update action context.
type UpdateReviewContext struct {
	*goa.Context
	ProposalID int
	ReviewID   int
	UserID     int
	Version    string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewUpdateReviewPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateReviewPayload is the review update action payload.
type UpdateReviewPayload struct {
	User       *string
	Comment    *string
	ID         int
	ProposalId *int
	Rating     *int
	Reviewers  *string
}

// NewUpdateReviewPayload instantiates a UpdateReviewPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateReviewPayload(raw interface{}) (p *UpdateReviewPayload, err error) {
	p, err = UnmarshalUpdateReviewPayload(raw, err)
	return
}

// UnmarshalUpdateReviewPayload unmarshals and validates a raw interface{} into an instance of UpdateReviewPayload
func UnmarshalUpdateReviewPayload(source interface{}, inErr error) (target *UpdateReviewPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateReviewPayload)
		if v, ok := val["User"]; ok {
			var tmp98 string
			if val, ok := v.(string); ok {
				tmp98 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.User`, v, "string", err)
			}
			target.User = &tmp98
		}
		if v, ok := val["comment"]; ok {
			var tmp99 string
			if val, ok := v.(string); ok {
				tmp99 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Comment`, v, "string", err)
			}
			target.Comment = &tmp99
		}
		if v, ok := val["id"]; ok {
			var tmp100 int
			if f, ok := v.(float64); ok {
				tmp100 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "int", err)
			}
			target.ID = tmp100
		} else {
			err = goa.MissingAttributeError(`payload`, "id", err)
		}
		if v, ok := val["proposal_id"]; ok {
			var tmp101 int
			if f, ok := v.(float64); ok {
				tmp101 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ProposalId`, v, "int", err)
			}
			target.ProposalId = &tmp101
		}
		if v, ok := val["rating"]; ok {
			var tmp102 int
			if f, ok := v.(float64); ok {
				tmp102 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Rating`, v, "int", err)
			}
			target.Rating = &tmp102
		}
		if v, ok := val["reviewers"]; ok {
			var tmp103 string
			if val, ok := v.(string); ok {
				tmp103 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Reviewers`, v, "string", err)
			}
			target.Reviewers = &tmp103
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}

// Validate validates the type instance.
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
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateReviewContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	*goa.Context
	Version string
	Payload *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(c *goa.Context) (*CreateUserContext, error) {
	var err error
	ctx := CreateUserContext{Context: c}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewCreateUserPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	Bio       *string
	City      *string
	Country   *string
	CreatedAt *string
	Email     string
	// First name Description
	FirstName string
	ID        int
	LastName  string
	Role      string
	State     *string
	UpdatedAt *string
}

// NewCreateUserPayload instantiates a CreateUserPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateUserPayload(raw interface{}) (p *CreateUserPayload, err error) {
	p, err = UnmarshalCreateUserPayload(raw, err)
	return
}

// UnmarshalCreateUserPayload unmarshals and validates a raw interface{} into an instance of CreateUserPayload
func UnmarshalCreateUserPayload(source interface{}, inErr error) (target *CreateUserPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateUserPayload)
		if v, ok := val["bio"]; ok {
			var tmp104 string
			if val, ok := v.(string); ok {
				tmp104 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Bio`, v, "string", err)
			}
			target.Bio = &tmp104
		}
		if v, ok := val["city"]; ok {
			var tmp105 string
			if val, ok := v.(string); ok {
				tmp105 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.City`, v, "string", err)
			}
			target.City = &tmp105
		}
		if v, ok := val["country"]; ok {
			var tmp106 string
			if val, ok := v.(string); ok {
				tmp106 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Country`, v, "string", err)
			}
			target.Country = &tmp106
		}
		if v, ok := val["created_at"]; ok {
			var tmp107 string
			if val, ok := v.(string); ok {
				tmp107 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.CreatedAt`, v, "string", err)
			}
			target.CreatedAt = &tmp107
		}
		if v, ok := val["email"]; ok {
			var tmp108 string
			if val, ok := v.(string); ok {
				tmp108 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			target.Email = tmp108
		} else {
			err = goa.MissingAttributeError(`payload`, "email", err)
		}
		if v, ok := val["first_name"]; ok {
			var tmp109 string
			if val, ok := v.(string); ok {
				tmp109 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.FirstName`, v, "string", err)
			}
			target.FirstName = tmp109
		} else {
			err = goa.MissingAttributeError(`payload`, "first_name", err)
		}
		if v, ok := val["id"]; ok {
			var tmp110 int
			if f, ok := v.(float64); ok {
				tmp110 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "int", err)
			}
			target.ID = tmp110
		} else {
			err = goa.MissingAttributeError(`payload`, "id", err)
		}
		if v, ok := val["last_name"]; ok {
			var tmp111 string
			if val, ok := v.(string); ok {
				tmp111 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.LastName`, v, "string", err)
			}
			target.LastName = tmp111
		} else {
			err = goa.MissingAttributeError(`payload`, "last_name", err)
		}
		if v, ok := val["role"]; ok {
			var tmp112 string
			if val, ok := v.(string); ok {
				tmp112 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Role`, v, "string", err)
			}
			target.Role = tmp112
		} else {
			err = goa.MissingAttributeError(`payload`, "role", err)
		}
		if v, ok := val["state"]; ok {
			var tmp113 string
			if val, ok := v.(string); ok {
				tmp113 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.State`, v, "string", err)
			}
			target.State = &tmp113
		}
		if v, ok := val["updated_at"]; ok {
			var tmp114 string
			if val, ok := v.(string); ok {
				tmp114 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.UpdatedAt`, v, "string", err)
			}
			target.UpdatedAt = &tmp114
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}

// Validate validates the type instance.
func (payload *CreateUserPayload) Validate() (err error) {

	if payload.Role == "" {
		err = goa.MissingAttributeError(`raw`, "role", err)
	}

	if payload.FirstName == "" {
		err = goa.MissingAttributeError(`raw`, "first_name", err)
	}

	if payload.LastName == "" {
		err = goa.MissingAttributeError(`raw`, "last_name", err)
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
	return ctx.Respond(201, nil)
}

// DeleteUserContext provides the user delete action context.
type DeleteUserContext struct {
	*goa.Context
	UserID  int
	Version string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteUserContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteUserContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListUserContext provides the user list action context.
type ListUserContext struct {
	*goa.Context
	Version string
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller list action.
func NewListUserContext(c *goa.Context) (*ListUserContext, error) {
	var err error
	ctx := ListUserContext{Context: c}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUserContext) OK(resp UserCollection) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.user+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowUserContext provides the user show action context.
type ShowUserContext struct {
	*goa.Context
	UserID  int
	Version string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(resp *User, view UserViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.user+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	*goa.Context
	UserID  int
	Version string
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
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewUpdateUserPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
	Bio       *string
	City      *string
	Country   *string
	CreatedAt *string
	Email     string
	// First name Description
	FirstName *string
	ID        int
	LastName  *string
	Role      string
	State     *string
	UpdatedAt *string
}

// NewUpdateUserPayload instantiates a UpdateUserPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateUserPayload(raw interface{}) (p *UpdateUserPayload, err error) {
	p, err = UnmarshalUpdateUserPayload(raw, err)
	return
}

// UnmarshalUpdateUserPayload unmarshals and validates a raw interface{} into an instance of UpdateUserPayload
func UnmarshalUpdateUserPayload(source interface{}, inErr error) (target *UpdateUserPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateUserPayload)
		if v, ok := val["bio"]; ok {
			var tmp118 string
			if val, ok := v.(string); ok {
				tmp118 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Bio`, v, "string", err)
			}
			target.Bio = &tmp118
		}
		if v, ok := val["city"]; ok {
			var tmp119 string
			if val, ok := v.(string); ok {
				tmp119 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.City`, v, "string", err)
			}
			target.City = &tmp119
		}
		if v, ok := val["country"]; ok {
			var tmp120 string
			if val, ok := v.(string); ok {
				tmp120 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Country`, v, "string", err)
			}
			target.Country = &tmp120
		}
		if v, ok := val["created_at"]; ok {
			var tmp121 string
			if val, ok := v.(string); ok {
				tmp121 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.CreatedAt`, v, "string", err)
			}
			target.CreatedAt = &tmp121
		}
		if v, ok := val["email"]; ok {
			var tmp122 string
			if val, ok := v.(string); ok {
				tmp122 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			target.Email = tmp122
		} else {
			err = goa.MissingAttributeError(`payload`, "email", err)
		}
		if v, ok := val["first_name"]; ok {
			var tmp123 string
			if val, ok := v.(string); ok {
				tmp123 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.FirstName`, v, "string", err)
			}
			target.FirstName = &tmp123
		}
		if v, ok := val["id"]; ok {
			var tmp124 int
			if f, ok := v.(float64); ok {
				tmp124 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "int", err)
			}
			target.ID = tmp124
		} else {
			err = goa.MissingAttributeError(`payload`, "id", err)
		}
		if v, ok := val["last_name"]; ok {
			var tmp125 string
			if val, ok := v.(string); ok {
				tmp125 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.LastName`, v, "string", err)
			}
			target.LastName = &tmp125
		}
		if v, ok := val["role"]; ok {
			var tmp126 string
			if val, ok := v.(string); ok {
				tmp126 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Role`, v, "string", err)
			}
			target.Role = tmp126
		} else {
			err = goa.MissingAttributeError(`payload`, "role", err)
		}
		if v, ok := val["state"]; ok {
			var tmp127 string
			if val, ok := v.(string); ok {
				tmp127 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.State`, v, "string", err)
			}
			target.State = &tmp127
		}
		if v, ok := val["updated_at"]; ok {
			var tmp128 string
			if val, ok := v.(string); ok {
				tmp128 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.UpdatedAt`, v, "string", err)
			}
			target.UpdatedAt = &tmp128
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}

// Validate validates the type instance.
func (payload *UpdateUserPayload) Validate() (err error) {

	if payload.Role == "" {
		err = goa.MissingAttributeError(`raw`, "role", err)
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

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateUserContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUserContext) NotFound() error {
	return ctx.Respond(404, nil)
}
