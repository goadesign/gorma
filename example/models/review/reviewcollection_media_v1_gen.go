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

// LoadReviewCollection loads raw data into an instance of ReviewCollection
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadReviewCollection(raw interface{}) (res ReviewCollection, err error) {
	res, err = UnmarshalReviewCollection(raw, err)
	return
}

// Dump produces raw data from an instance of ReviewCollection running all the
// validations. See LoadReviewCollection for the definition of raw data.

func (mt ReviewCollection) Dump() (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, len(mt))
	for i, tmp31 := range mt {
		var tmp32 map[string]interface{}
		tmp32, err = MarshalReview(tmp31, err)
		res[i] = tmp32
	}
	return
}

// MarshalReviewCollection validates and renders an instance of ReviewCollection into a interface{}
// using view "default".
func MarshalReviewCollection(source ReviewCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalReview(res, err)
	}
	return
}

// UnmarshalReviewCollection unmarshals and validates a raw interface{} into an instance of ReviewCollection
func UnmarshalReviewCollection(source interface{}, inErr error) (target ReviewCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*Review, len(val))
		for tmp33, v := range val {
			target[tmp33], err = UnmarshalReview(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}
