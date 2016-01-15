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

	"github.com/bketelsen/gorma/example/app"
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
	UserID    *int
	Abstract  string
	CreatedAt *interface{}
	DeletedAt *interface{}
	Detail    string
	Firstname *string
	ID        string
	M2reviews []*app.ReviewModel
	Reviews   []*app.ReviewModel
	Title     string
	UpdatedAt *interface{}
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
		if v, ok := val["UserID"]; ok {
			var tmp58 int
			if f, ok := v.(float64); ok {
				tmp58 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.UserID`, v, "int", err)
			}
			target.UserID = &tmp58
		}
		if v, ok := val["abstract"]; ok {
			var tmp59 string
			if val, ok := v.(string); ok {
				tmp59 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Abstract`, v, "string", err)
			}
			target.Abstract = tmp59
		} else {
			err = goa.MissingAttributeError(`payload`, "abstract", err)
		}
		if v, ok := val["created_at"]; ok {
			var tmp60 interface{}
			tmp60 = v
			target.CreatedAt = &tmp60
		}
		if v, ok := val["deleted_at"]; ok {
			var tmp61 interface{}
			tmp61 = v
			target.DeletedAt = &tmp61
		}
		if v, ok := val["detail"]; ok {
			var tmp62 string
			if val, ok := v.(string); ok {
				tmp62 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Detail`, v, "string", err)
			}
			target.Detail = tmp62
		} else {
			err = goa.MissingAttributeError(`payload`, "detail", err)
		}
		if v, ok := val["firstname"]; ok {
			var tmp63 string
			if val, ok := v.(string); ok {
				tmp63 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Firstname`, v, "string", err)
			}
			target.Firstname = &tmp63
		}
		if v, ok := val["id"]; ok {
			var tmp64 string
			if val, ok := v.(string); ok {
				tmp64 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "string", err)
			}
			target.ID = tmp64
		} else {
			err = goa.MissingAttributeError(`payload`, "id", err)
		}
		if v, ok := val["m2reviews"]; ok {
			var tmp65 []*app.ReviewModel
			if val, ok := v.([]interface{}); ok {
				tmp65 = make([]*app.ReviewModel, len(val))
				for tmp66, v := range val {
					tmp65[tmp66], err = app.UnmarshalReviewModel(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.M2reviews`, v, "array", err)
			}
			target.M2reviews = tmp65
		}
		if v, ok := val["reviews"]; ok {
			var tmp67 []*app.ReviewModel
			if val, ok := v.([]interface{}); ok {
				tmp67 = make([]*app.ReviewModel, len(val))
				for tmp68, v := range val {
					tmp67[tmp68], err = app.UnmarshalReviewModel(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Reviews`, v, "array", err)
			}
			target.Reviews = tmp67
		}
		if v, ok := val["title"]; ok {
			var tmp69 string
			if val, ok := v.(string); ok {
				tmp69 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Title`, v, "string", err)
			}
			target.Title = tmp69
		} else {
			err = goa.MissingAttributeError(`payload`, "title", err)
		}
		if v, ok := val["updated_at"]; ok {
			var tmp70 interface{}
			tmp70 = v
			target.UpdatedAt = &tmp70
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp71 bool
			if val, ok := v.(bool); ok {
				tmp71 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = &tmp71
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
	if payload.ID == "" {
		err = goa.MissingAttributeError(`raw`, "id", err)
	}

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
	if payload.Firstname != nil {
		if len(*payload.Firstname) < 2 {
			err = goa.InvalidLengthError(`raw.firstname`, *payload.Firstname, len(*payload.Firstname), 2, true, err)
		}
	}
	for _, e := range payload.M2reviews {
		if e.Comment != nil {
			if len(*e.Comment) < 10 {
				err = goa.InvalidLengthError(`raw.m2reviews[*].comment`, *e.Comment, len(*e.Comment), 10, true, err)
			}
		}
		if e.Comment != nil {
			if len(*e.Comment) > 200 {
				err = goa.InvalidLengthError(`raw.m2reviews[*].comment`, *e.Comment, len(*e.Comment), 200, false, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`raw.m2reviews[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`raw.m2reviews[*].rating`, *e.Rating, 5, false, err)
			}
		}
		for _, e := range e.Reviewers {
			if e.Bio != nil {
				if len(*e.Bio) > 500 {
					err = goa.InvalidLengthError(`raw.m2reviews[*].reviewers[*].bio`, *e.Bio, len(*e.Bio), 500, false, err)
				}
			}
			if e.Email != nil {
				if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
					err = goa.InvalidFormatError(`raw.m2reviews[*].reviewers[*].email`, *e.Email, goa.FormatEmail, err2, err)
				}
			}
		}
	}
	for _, e := range payload.Reviews {
		if e.Comment != nil {
			if len(*e.Comment) < 10 {
				err = goa.InvalidLengthError(`raw.reviews[*].comment`, *e.Comment, len(*e.Comment), 10, true, err)
			}
		}
		if e.Comment != nil {
			if len(*e.Comment) > 200 {
				err = goa.InvalidLengthError(`raw.reviews[*].comment`, *e.Comment, len(*e.Comment), 200, false, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`raw.reviews[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`raw.reviews[*].rating`, *e.Rating, 5, false, err)
			}
		}
		for _, e := range e.Reviewers {
			if e.Bio != nil {
				if len(*e.Bio) > 500 {
					err = goa.InvalidLengthError(`raw.reviews[*].reviewers[*].bio`, *e.Bio, len(*e.Bio), 500, false, err)
				}
			}
			if e.Email != nil {
				if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
					err = goa.InvalidFormatError(`raw.reviews[*].reviewers[*].email`, *e.Email, goa.FormatEmail, err2, err)
				}
			}
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
	UserID    *int
	Abstract  *string
	CreatedAt *interface{}
	DeletedAt *interface{}
	Detail    *string
	Firstname *string
	ID        string
	M2reviews []*app.ReviewModel
	Reviews   []*app.ReviewModel
	Title     *string
	UpdatedAt *interface{}
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
		if v, ok := val["UserID"]; ok {
			var tmp79 int
			if f, ok := v.(float64); ok {
				tmp79 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.UserID`, v, "int", err)
			}
			target.UserID = &tmp79
		}
		if v, ok := val["abstract"]; ok {
			var tmp80 string
			if val, ok := v.(string); ok {
				tmp80 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Abstract`, v, "string", err)
			}
			target.Abstract = &tmp80
		}
		if v, ok := val["created_at"]; ok {
			var tmp81 interface{}
			tmp81 = v
			target.CreatedAt = &tmp81
		}
		if v, ok := val["deleted_at"]; ok {
			var tmp82 interface{}
			tmp82 = v
			target.DeletedAt = &tmp82
		}
		if v, ok := val["detail"]; ok {
			var tmp83 string
			if val, ok := v.(string); ok {
				tmp83 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Detail`, v, "string", err)
			}
			target.Detail = &tmp83
		}
		if v, ok := val["firstname"]; ok {
			var tmp84 string
			if val, ok := v.(string); ok {
				tmp84 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Firstname`, v, "string", err)
			}
			target.Firstname = &tmp84
		}
		if v, ok := val["id"]; ok {
			var tmp85 string
			if val, ok := v.(string); ok {
				tmp85 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "string", err)
			}
			target.ID = tmp85
		} else {
			err = goa.MissingAttributeError(`payload`, "id", err)
		}
		if v, ok := val["m2reviews"]; ok {
			var tmp86 []*app.ReviewModel
			if val, ok := v.([]interface{}); ok {
				tmp86 = make([]*app.ReviewModel, len(val))
				for tmp87, v := range val {
					tmp86[tmp87], err = app.UnmarshalReviewModel(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.M2reviews`, v, "array", err)
			}
			target.M2reviews = tmp86
		}
		if v, ok := val["reviews"]; ok {
			var tmp88 []*app.ReviewModel
			if val, ok := v.([]interface{}); ok {
				tmp88 = make([]*app.ReviewModel, len(val))
				for tmp89, v := range val {
					tmp88[tmp89], err = app.UnmarshalReviewModel(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Reviews`, v, "array", err)
			}
			target.Reviews = tmp88
		}
		if v, ok := val["title"]; ok {
			var tmp90 string
			if val, ok := v.(string); ok {
				tmp90 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Title`, v, "string", err)
			}
			target.Title = &tmp90
		}
		if v, ok := val["updated_at"]; ok {
			var tmp91 interface{}
			tmp91 = v
			target.UpdatedAt = &tmp91
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp92 bool
			if val, ok := v.(bool); ok {
				tmp92 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = &tmp92
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
	if payload.ID == "" {
		err = goa.MissingAttributeError(`raw`, "id", err)
	}

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
	if payload.Firstname != nil {
		if len(*payload.Firstname) < 2 {
			err = goa.InvalidLengthError(`raw.firstname`, *payload.Firstname, len(*payload.Firstname), 2, true, err)
		}
	}
	for _, e := range payload.M2reviews {
		if e.Comment != nil {
			if len(*e.Comment) < 10 {
				err = goa.InvalidLengthError(`raw.m2reviews[*].comment`, *e.Comment, len(*e.Comment), 10, true, err)
			}
		}
		if e.Comment != nil {
			if len(*e.Comment) > 200 {
				err = goa.InvalidLengthError(`raw.m2reviews[*].comment`, *e.Comment, len(*e.Comment), 200, false, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`raw.m2reviews[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`raw.m2reviews[*].rating`, *e.Rating, 5, false, err)
			}
		}
		for _, e := range e.Reviewers {
			if e.Bio != nil {
				if len(*e.Bio) > 500 {
					err = goa.InvalidLengthError(`raw.m2reviews[*].reviewers[*].bio`, *e.Bio, len(*e.Bio), 500, false, err)
				}
			}
			if e.Email != nil {
				if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
					err = goa.InvalidFormatError(`raw.m2reviews[*].reviewers[*].email`, *e.Email, goa.FormatEmail, err2, err)
				}
			}
		}
	}
	for _, e := range payload.Reviews {
		if e.Comment != nil {
			if len(*e.Comment) < 10 {
				err = goa.InvalidLengthError(`raw.reviews[*].comment`, *e.Comment, len(*e.Comment), 10, true, err)
			}
		}
		if e.Comment != nil {
			if len(*e.Comment) > 200 {
				err = goa.InvalidLengthError(`raw.reviews[*].comment`, *e.Comment, len(*e.Comment), 200, false, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`raw.reviews[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`raw.reviews[*].rating`, *e.Rating, 5, false, err)
			}
		}
		for _, e := range e.Reviewers {
			if e.Bio != nil {
				if len(*e.Bio) > 500 {
					err = goa.InvalidLengthError(`raw.reviews[*].reviewers[*].bio`, *e.Bio, len(*e.Bio), 500, false, err)
				}
			}
			if e.Email != nil {
				if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
					err = goa.InvalidFormatError(`raw.reviews[*].reviewers[*].email`, *e.Email, goa.FormatEmail, err2, err)
				}
			}
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
	ProposalID *int
	UserID     *int
	Comment    *string
	ID         string
	Rating     int
	Reviewers  []*app.UserModel
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
		if v, ok := val["ProposalID"]; ok {
			var tmp95 int
			if f, ok := v.(float64); ok {
				tmp95 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ProposalID`, v, "int", err)
			}
			target.ProposalID = &tmp95
		}
		if v, ok := val["UserID"]; ok {
			var tmp96 int
			if f, ok := v.(float64); ok {
				tmp96 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.UserID`, v, "int", err)
			}
			target.UserID = &tmp96
		}
		if v, ok := val["comment"]; ok {
			var tmp97 string
			if val, ok := v.(string); ok {
				tmp97 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Comment`, v, "string", err)
			}
			target.Comment = &tmp97
		}
		if v, ok := val["id"]; ok {
			var tmp98 string
			if val, ok := v.(string); ok {
				tmp98 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "string", err)
			}
			target.ID = tmp98
		} else {
			err = goa.MissingAttributeError(`payload`, "id", err)
		}
		if v, ok := val["rating"]; ok {
			var tmp99 int
			if f, ok := v.(float64); ok {
				tmp99 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Rating`, v, "int", err)
			}
			target.Rating = tmp99
		} else {
			err = goa.MissingAttributeError(`payload`, "rating", err)
		}
		if v, ok := val["reviewers"]; ok {
			var tmp100 []*app.UserModel
			if val, ok := v.([]interface{}); ok {
				tmp100 = make([]*app.UserModel, len(val))
				for tmp101, v := range val {
					tmp100[tmp101], err = app.UnmarshalUserModel(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Reviewers`, v, "array", err)
			}
			target.Reviewers = tmp100
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
	if payload.ID == "" {
		err = goa.MissingAttributeError(`raw`, "id", err)
	}

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
	for _, e := range payload.Reviewers {
		if e.Bio != nil {
			if len(*e.Bio) > 500 {
				err = goa.InvalidLengthError(`raw.reviewers[*].bio`, *e.Bio, len(*e.Bio), 500, false, err)
			}
		}
		if e.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
				err = goa.InvalidFormatError(`raw.reviewers[*].email`, *e.Email, goa.FormatEmail, err2, err)
			}
		}
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
	ProposalID *int
	UserID     *int
	Comment    *string
	ID         string
	Rating     *int
	Reviewers  []*app.UserModel
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
		if v, ok := val["ProposalID"]; ok {
			var tmp113 int
			if f, ok := v.(float64); ok {
				tmp113 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ProposalID`, v, "int", err)
			}
			target.ProposalID = &tmp113
		}
		if v, ok := val["UserID"]; ok {
			var tmp114 int
			if f, ok := v.(float64); ok {
				tmp114 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.UserID`, v, "int", err)
			}
			target.UserID = &tmp114
		}
		if v, ok := val["comment"]; ok {
			var tmp115 string
			if val, ok := v.(string); ok {
				tmp115 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Comment`, v, "string", err)
			}
			target.Comment = &tmp115
		}
		if v, ok := val["id"]; ok {
			var tmp116 string
			if val, ok := v.(string); ok {
				tmp116 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "string", err)
			}
			target.ID = tmp116
		} else {
			err = goa.MissingAttributeError(`payload`, "id", err)
		}
		if v, ok := val["rating"]; ok {
			var tmp117 int
			if f, ok := v.(float64); ok {
				tmp117 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Rating`, v, "int", err)
			}
			target.Rating = &tmp117
		}
		if v, ok := val["reviewers"]; ok {
			var tmp118 []*app.UserModel
			if val, ok := v.([]interface{}); ok {
				tmp118 = make([]*app.UserModel, len(val))
				for tmp119, v := range val {
					tmp118[tmp119], err = app.UnmarshalUserModel(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Reviewers`, v, "array", err)
			}
			target.Reviewers = tmp118
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
	if payload.ID == "" {
		err = goa.MissingAttributeError(`raw`, "id", err)
	}

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
	for _, e := range payload.Reviewers {
		if e.Bio != nil {
			if len(*e.Bio) > 500 {
				err = goa.InvalidLengthError(`raw.reviewers[*].bio`, *e.Bio, len(*e.Bio), 500, false, err)
			}
		}
		if e.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
				err = goa.InvalidFormatError(`raw.reviewers[*].email`, *e.Email, goa.FormatEmail, err2, err)
			}
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
	Email     string
	Firstname string
	ID        string
	Lastname  string
	Role      *string
	State     *string
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
			var tmp120 string
			if val, ok := v.(string); ok {
				tmp120 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Bio`, v, "string", err)
			}
			target.Bio = &tmp120
		}
		if v, ok := val["city"]; ok {
			var tmp121 string
			if val, ok := v.(string); ok {
				tmp121 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.City`, v, "string", err)
			}
			target.City = &tmp121
		}
		if v, ok := val["country"]; ok {
			var tmp122 string
			if val, ok := v.(string); ok {
				tmp122 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Country`, v, "string", err)
			}
			target.Country = &tmp122
		}
		if v, ok := val["email"]; ok {
			var tmp123 string
			if val, ok := v.(string); ok {
				tmp123 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			target.Email = tmp123
		} else {
			err = goa.MissingAttributeError(`payload`, "email", err)
		}
		if v, ok := val["firstname"]; ok {
			var tmp124 string
			if val, ok := v.(string); ok {
				tmp124 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Firstname`, v, "string", err)
			}
			target.Firstname = tmp124
		} else {
			err = goa.MissingAttributeError(`payload`, "firstname", err)
		}
		if v, ok := val["id"]; ok {
			var tmp125 string
			if val, ok := v.(string); ok {
				tmp125 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "string", err)
			}
			target.ID = tmp125
		} else {
			err = goa.MissingAttributeError(`payload`, "id", err)
		}
		if v, ok := val["lastname"]; ok {
			var tmp126 string
			if val, ok := v.(string); ok {
				tmp126 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Lastname`, v, "string", err)
			}
			target.Lastname = tmp126
		} else {
			err = goa.MissingAttributeError(`payload`, "lastname", err)
		}
		if v, ok := val["role"]; ok {
			var tmp127 string
			if val, ok := v.(string); ok {
				tmp127 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Role`, v, "string", err)
			}
			target.Role = &tmp127
		}
		if v, ok := val["state"]; ok {
			var tmp128 string
			if val, ok := v.(string); ok {
				tmp128 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.State`, v, "string", err)
			}
			target.State = &tmp128
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
	if payload.ID == "" {
		err = goa.MissingAttributeError(`raw`, "id", err)
	}

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
	Email     string
	Firstname *string
	ID        string
	Lastname  *string
	Role      *string
	State     *string
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
			var tmp132 string
			if val, ok := v.(string); ok {
				tmp132 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Bio`, v, "string", err)
			}
			target.Bio = &tmp132
		}
		if v, ok := val["city"]; ok {
			var tmp133 string
			if val, ok := v.(string); ok {
				tmp133 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.City`, v, "string", err)
			}
			target.City = &tmp133
		}
		if v, ok := val["country"]; ok {
			var tmp134 string
			if val, ok := v.(string); ok {
				tmp134 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Country`, v, "string", err)
			}
			target.Country = &tmp134
		}
		if v, ok := val["email"]; ok {
			var tmp135 string
			if val, ok := v.(string); ok {
				tmp135 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			target.Email = tmp135
		} else {
			err = goa.MissingAttributeError(`payload`, "email", err)
		}
		if v, ok := val["firstname"]; ok {
			var tmp136 string
			if val, ok := v.(string); ok {
				tmp136 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Firstname`, v, "string", err)
			}
			target.Firstname = &tmp136
		}
		if v, ok := val["id"]; ok {
			var tmp137 string
			if val, ok := v.(string); ok {
				tmp137 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.ID`, v, "string", err)
			}
			target.ID = tmp137
		} else {
			err = goa.MissingAttributeError(`payload`, "id", err)
		}
		if v, ok := val["lastname"]; ok {
			var tmp138 string
			if val, ok := v.(string); ok {
				tmp138 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Lastname`, v, "string", err)
			}
			target.Lastname = &tmp138
		}
		if v, ok := val["role"]; ok {
			var tmp139 string
			if val, ok := v.(string); ok {
				tmp139 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Role`, v, "string", err)
			}
			target.Role = &tmp139
		}
		if v, ok := val["state"]; ok {
			var tmp140 string
			if val, ok := v.(string); ok {
				tmp140 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.State`, v, "string", err)
			}
			target.State = &tmp140
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
	if payload.ID == "" {
		err = goa.MissingAttributeError(`raw`, "id", err)
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
