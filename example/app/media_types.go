//************************************************************************//
// API "congo": Application Media Types
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

import "github.com/raphael/goa"

// Token authorization response
// Identifier: application/vnd.authorize+json
type Authorize struct {
	// access token
	AccessToken *string
	// Time to expiration in seconds
	ExpiresIn *int
	// type of token
	TokenType *string
}

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
	tmp7 := map[string]interface{}{
		"access_token": source.AccessToken,
		"expires_in":   source.ExpiresIn,
		"token_type":   source.TokenType,
	}
	target = tmp7
	return
}

// UnmarshalAuthorize unmarshals and validates a raw interface{} into an instance of Authorize
func UnmarshalAuthorize(source interface{}, inErr error) (target *Authorize, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Authorize)
		if v, ok := val["access_token"]; ok {
			var tmp8 string
			if val, ok := v.(string); ok {
				tmp8 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.AccessToken`, v, "string", err)
			}
			target.AccessToken = &tmp8
		}
		if v, ok := val["expires_in"]; ok {
			var tmp9 int
			if f, ok := v.(float64); ok {
				tmp9 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ExpiresIn`, v, "int", err)
			}
			target.ExpiresIn = &tmp9
		}
		if v, ok := val["token_type"]; ok {
			var tmp10 string
			if val, ok := v.(string); ok {
				tmp10 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.TokenType`, v, "string", err)
			}
			target.TokenType = &tmp10
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// Login media type
// Identifier: application/vnd.login+json
type Login struct {
	// UUID of requesting application
	Application *string
	// email
	Email *string
	// password
	Password *string
}

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
	tmp11 := map[string]interface{}{
		"application": source.Application,
		"email":       source.Email,
		"password":    source.Password,
	}
	target = tmp11
	return
}

// UnmarshalLogin unmarshals and validates a raw interface{} into an instance of Login
func UnmarshalLogin(source interface{}, inErr error) (target *Login, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Login)
		if v, ok := val["application"]; ok {
			var tmp12 string
			if val, ok := v.(string); ok {
				tmp12 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Application`, v, "string", err)
			}
			target.Application = &tmp12
		}
		if v, ok := val["email"]; ok {
			var tmp13 string
			if val, ok := v.(string); ok {
				tmp13 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			target.Email = &tmp13
		}
		if v, ok := val["password"]; ok {
			var tmp14 string
			if val, ok := v.(string); ok {
				tmp14 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Password`, v, "string", err)
			}
			target.Password = &tmp14
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
