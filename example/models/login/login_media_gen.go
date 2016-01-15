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

package login

import "github.com/raphael/goa"

// LoadLogin loads raw data into an instance of Login
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadLogin(raw interface{}) (res *Login, err error) {
	res, err = UnmarshalLogin(raw, err)
	return
}

// Dump produces raw data from an instance of Login running all the
// validations. See LoadLogin for the definition of raw data.

func (mt *Login) Dump() (res map[string]interface{}, err error) {
	res, err = MarshalLogin(mt, err)
	return
}

// MarshalLogin validates and renders an instance of Login into a interface{}
// using view "default".
func MarshalLogin(source *Login, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp5 := map[string]interface{}{
		"application": source.Application,
		"email":       source.Email,
		"password":    source.Password,
	}
	target = tmp5
	return
}

// UnmarshalLogin unmarshals and validates a raw interface{} into an instance of Login
func UnmarshalLogin(source interface{}, inErr error) (target *Login, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Login)
		if v, ok := val["application"]; ok {
			var tmp6 string
			if val, ok := v.(string); ok {
				tmp6 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Application`, v, "string", err)
			}
			target.Application = &tmp6
		}
		if v, ok := val["email"]; ok {
			var tmp7 string
			if val, ok := v.(string); ok {
				tmp7 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			target.Email = &tmp7
		}
		if v, ok := val["password"]; ok {
			var tmp8 string
			if val, ok := v.(string); ok {
				tmp8 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Password`, v, "string", err)
			}
			target.Password = &tmp8
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
