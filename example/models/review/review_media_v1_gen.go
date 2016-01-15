//************************************************************************//
// API "congo" version v1: Application Media Helpers
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

// Review views
type ReviewViewEnum string

const (
	// Review default view
	ReviewDefaultView ReviewViewEnum = "default"

	// Review link view
	ReviewLinkView ReviewViewEnum = "link"
)

// LoadReview loads raw data into an instance of Review
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadReview(raw interface{}) (res *Review, err error) {
	res, err = UnmarshalReview(raw, err)
	return
}

// Dump produces raw data from an instance of Review running all the
// validations. See LoadReview for the definition of raw data.

func (mt *Review) Dump(view ReviewViewEnum) (res map[string]interface{}, err error) {
	if view == ReviewDefaultView {
		res, err = MarshalReview(mt, err)
	}
	if view == ReviewLinkView {
		res, err = MarshalReviewLink(mt, err)
	}
	return
}

// MarshalReview validates and renders an instance of Review into a interface{}
// using view "default".
func MarshalReview(source *Review, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp25 := map[string]interface{}{
		"comment": source.Comment,
		"href":    source.Href,
		"id":      source.ID,
		"rating":  source.Rating,
	}
	target = tmp25
	return
}

// MarshalReviewLink validates and renders an instance of Review into a interface{}
// using view "link".
func MarshalReviewLink(source *Review, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp26 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
	}
	target = tmp26
	return
}

// UnmarshalReview unmarshals and validates a raw interface{} into an instance of Review
func UnmarshalReview(source interface{}, inErr error) (target *Review, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Review)
		if v, ok := val["comment"]; ok {
			var tmp27 string
			if val, ok := v.(string); ok {
				tmp27 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Comment`, v, "string", err)
			}
			target.Comment = &tmp27
		}
		if v, ok := val["href"]; ok {
			var tmp28 string
			if val, ok := v.(string); ok {
				tmp28 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = &tmp28
		}
		if v, ok := val["id"]; ok {
			var tmp29 int
			if f, ok := v.(float64); ok {
				tmp29 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = &tmp29
		}
		if v, ok := val["rating"]; ok {
			var tmp30 int
			if f, ok := v.(float64); ok {
				tmp30 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			target.Rating = &tmp30
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}
