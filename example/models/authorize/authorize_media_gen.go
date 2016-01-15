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

package authorize

import "github.com/raphael/goa"

// LoadAuthorize loads raw data into an instance of Authorize
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadAuthorize(raw interface{}) (res *Authorize, err error) {
	res, err = UnmarshalAuthorize(raw, err)
	return
}

// Dump produces raw data from an instance of Authorize running all the
// validations. See LoadAuthorize for the definition of raw data.

func (mt *Authorize) Dump() (res map[string]interface{}, err error) {
	res, err = MarshalAuthorize(mt, err)
	return
}

// MarshalAuthorize validates and renders an instance of Authorize into a interface{}
// using view "default".
func MarshalAuthorize(source *Authorize, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp1 := map[string]interface{}{
		"access_token": source.AccessToken,
		"expires_in":   source.ExpiresIn,
		"token_type":   source.TokenType,
	}
	target = tmp1
	return
}

// UnmarshalAuthorize unmarshals and validates a raw interface{} into an instance of Authorize
func UnmarshalAuthorize(source interface{}, inErr error) (target *Authorize, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Authorize)
		if v, ok := val["access_token"]; ok {
			var tmp2 string
			if val, ok := v.(string); ok {
				tmp2 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.AccessToken`, v, "string", err)
			}
			target.AccessToken = &tmp2
		}
		if v, ok := val["expires_in"]; ok {
			var tmp3 int
			if f, ok := v.(float64); ok {
				tmp3 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ExpiresIn`, v, "int", err)
			}
			target.ExpiresIn = &tmp3
		}
		if v, ok := val["token_type"]; ok {
			var tmp4 string
			if val, ok := v.(string); ok {
				tmp4 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.TokenType`, v, "string", err)
			}
			target.TokenType = &tmp4
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
