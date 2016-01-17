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

package review

import "github.com/raphael/goa"

// LoadReview loads raw data into an instance of Review
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadReview(raw interface{}) (res *Review, err error) {
	res, err = UnmarshalReview(raw, err)
	return
}

// MarshalReview validates and renders an instance of Review into a interface{}
func MarshalReview(source *Review, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp14 := map[string]interface{}{
		"User":        source.User,
		"comment":     source.Comment,
		"id":          source.ID,
		"proposal_id": source.ProposalId,
		"rating":      source.Rating,
		"reviewers":   source.Reviewers,
	}
	target = tmp14
	return
}

// UnmarshalReview unmarshals and validates a raw interface{} into an instance of Review
func UnmarshalReview(source interface{}, inErr error) (target *Review, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Review)
		if v, ok := val["User"]; ok {
			var tmp15 string
			if val, ok := v.(string); ok {
				tmp15 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.User`, v, "string", err)
			}
			target.User = &tmp15
		}
		if v, ok := val["comment"]; ok {
			var tmp16 string
			if val, ok := v.(string); ok {
				tmp16 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Comment`, v, "string", err)
			}
			target.Comment = &tmp16
		}
		if v, ok := val["id"]; ok {
			var tmp17 int
			if f, ok := v.(float64); ok {
				tmp17 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp17
		} else {
			err = goa.MissingAttributeError(`load`, "id", err)
		}
		if v, ok := val["proposal_id"]; ok {
			var tmp18 int
			if f, ok := v.(float64); ok {
				tmp18 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ProposalId`, v, "int", err)
			}
			target.ProposalId = &tmp18
		}
		if v, ok := val["rating"]; ok {
			var tmp19 int
			if f, ok := v.(float64); ok {
				tmp19 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			target.Rating = &tmp19
		}
		if v, ok := val["reviewers"]; ok {
			var tmp20 string
			if val, ok := v.(string); ok {
				tmp20 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Reviewers`, v, "string", err)
			}
			target.Reviewers = &tmp20
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}
