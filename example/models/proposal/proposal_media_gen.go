//************************************************************************//
// API "congo": Application Media Helpers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/bketelsen/gorma/example
// --design=github.com/bketelsen/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package proposal

import "github.com/raphael/goa"

// LoadProposal loads raw data into an instance of Proposal
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadProposal(raw interface{}) (res *Proposal, err error) {
	res, err = UnmarshalProposal(raw, err)
	return
}

// MarshalProposal validates and renders an instance of Proposal into a interface{}
func MarshalProposal(source *Proposal, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp1 := map[string]interface{}{
		"abstract":   source.Abstract,
		"created_at": source.CreatedAt,
		"deleted_at": source.DeletedAt,
		"detail":     source.Detail,
		"first_name": source.FirstName,
		"id":         source.ID,
		"m2reviews":  source.M2reviews,
		"reviews":    source.Reviews,
		"title":      source.Title,
		"updated_at": source.UpdatedAt,
		"user_id":    source.UserId,
		"withdrawn":  source.Withdrawn,
	}
	target = tmp1
	return
}

// UnmarshalProposal unmarshals and validates a raw interface{} into an instance of Proposal
func UnmarshalProposal(source interface{}, inErr error) (target *Proposal, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Proposal)
		if v, ok := val["abstract"]; ok {
			var tmp2 string
			if val, ok := v.(string); ok {
				tmp2 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Abstract`, v, "string", err)
			}
			target.Abstract = &tmp2
		}
		if v, ok := val["created_at"]; ok {
			var tmp3 string
			if val, ok := v.(string); ok {
				tmp3 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.CreatedAt`, v, "string", err)
			}
			target.CreatedAt = &tmp3
		}
		if v, ok := val["deleted_at"]; ok {
			var tmp4 string
			if val, ok := v.(string); ok {
				tmp4 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.DeletedAt`, v, "string", err)
			}
			target.DeletedAt = &tmp4
		}
		if v, ok := val["detail"]; ok {
			var tmp5 string
			if val, ok := v.(string); ok {
				tmp5 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Detail`, v, "string", err)
			}
			target.Detail = &tmp5
		}
		if v, ok := val["first_name"]; ok {
			var tmp6 string
			if val, ok := v.(string); ok {
				tmp6 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.FirstName`, v, "string", err)
			}
			target.FirstName = &tmp6
		}
		if v, ok := val["id"]; ok {
			var tmp7 int
			if f, ok := v.(float64); ok {
				tmp7 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp7
		} else {
			err = goa.MissingAttributeError(`load`, "id", err)
		}
		if v, ok := val["m2reviews"]; ok {
			var tmp8 string
			if val, ok := v.(string); ok {
				tmp8 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.M2reviews`, v, "string", err)
			}
			target.M2reviews = &tmp8
		}
		if v, ok := val["reviews"]; ok {
			var tmp9 string
			if val, ok := v.(string); ok {
				tmp9 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Reviews`, v, "string", err)
			}
			target.Reviews = &tmp9
		}
		if v, ok := val["title"]; ok {
			var tmp10 string
			if val, ok := v.(string); ok {
				tmp10 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Title`, v, "string", err)
			}
			target.Title = &tmp10
		}
		if v, ok := val["updated_at"]; ok {
			var tmp11 string
			if val, ok := v.(string); ok {
				tmp11 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.UpdatedAt`, v, "string", err)
			}
			target.UpdatedAt = &tmp11
		}
		if v, ok := val["user_id"]; ok {
			var tmp12 int
			if f, ok := v.(float64); ok {
				tmp12 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.UserId`, v, "int", err)
			}
			target.UserId = &tmp12
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp13 bool
			if val, ok := v.(bool); ok {
				tmp13 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = &tmp13
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}
