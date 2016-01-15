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

package user

import "github.com/raphael/goa"

// LoadUserCollection loads raw data into an instance of UserCollection
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadUserCollection(raw interface{}) (res UserCollection, err error) {
	res, err = UnmarshalUserCollection(raw, err)
	return
}

// Dump produces raw data from an instance of UserCollection running all the
// validations. See LoadUserCollection for the definition of raw data.

func (mt UserCollection) Dump() (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, len(mt))
	for i, tmp46 := range mt {
		var tmp47 map[string]interface{}
		tmp47, err = MarshalUser(tmp46, err)
		res[i] = tmp47
	}
	return
}

// MarshalUserCollection validates and renders an instance of UserCollection into a interface{}
// using view "default".
func MarshalUserCollection(source UserCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalUser(res, err)
	}
	return
}

// UnmarshalUserCollection unmarshals and validates a raw interface{} into an instance of UserCollection
func UnmarshalUserCollection(source interface{}, inErr error) (target UserCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*User, len(val))
		for tmp48, v := range val {
			target[tmp48], err = UnmarshalUser(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}
