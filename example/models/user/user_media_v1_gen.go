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

// User views
type UserViewEnum string

const (
	// User default view
	UserDefaultView UserViewEnum = "default"

	// User link view
	UserLinkView UserViewEnum = "link"
)

// LoadUser loads raw data into an instance of User
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadUser(raw interface{}) (res *User, err error) {
	res, err = UnmarshalUser(raw, err)
	return
}

// Dump produces raw data from an instance of User running all the
// validations. See LoadUser for the definition of raw data.

func (mt *User) Dump(view UserViewEnum) (res map[string]interface{}, err error) {
	if view == UserDefaultView {
		res, err = MarshalUser(mt, err)
	}
	if view == UserLinkView {
		res, err = MarshalUserLink(mt, err)
	}
	return
}

// MarshalUser validates and renders an instance of User into a interface{}
// using view "default".
func MarshalUser(source *User, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp34 := map[string]interface{}{
		"bio":       source.Bio,
		"city":      source.City,
		"country":   source.Country,
		"email":     source.Email,
		"firstname": source.Firstname,
		"href":      source.Href,
		"id":        source.ID,
		"lastname":  source.Lastname,
		"role":      source.Role,
		"state":     source.State,
	}
	target = tmp34
	return
}

// MarshalUserLink validates and renders an instance of User into a interface{}
// using view "link".
func MarshalUserLink(source *User, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp35 := map[string]interface{}{
		"email": source.Email,
		"href":  source.Href,
		"id":    source.ID,
	}
	target = tmp35
	return
}

// UnmarshalUser unmarshals and validates a raw interface{} into an instance of User
func UnmarshalUser(source interface{}, inErr error) (target *User, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(User)
		if v, ok := val["bio"]; ok {
			var tmp36 string
			if val, ok := v.(string); ok {
				tmp36 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Bio`, v, "string", err)
			}
			target.Bio = &tmp36
		}
		if v, ok := val["city"]; ok {
			var tmp37 string
			if val, ok := v.(string); ok {
				tmp37 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.City`, v, "string", err)
			}
			target.City = &tmp37
		}
		if v, ok := val["country"]; ok {
			var tmp38 string
			if val, ok := v.(string); ok {
				tmp38 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			target.Country = &tmp38
		}
		if v, ok := val["email"]; ok {
			var tmp39 string
			if val, ok := v.(string); ok {
				tmp39 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			target.Email = &tmp39
		}
		if v, ok := val["firstname"]; ok {
			var tmp40 string
			if val, ok := v.(string); ok {
				tmp40 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			target.Firstname = &tmp40
		}
		if v, ok := val["href"]; ok {
			var tmp41 string
			if val, ok := v.(string); ok {
				tmp41 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = &tmp41
		}
		if v, ok := val["id"]; ok {
			var tmp42 int
			if f, ok := v.(float64); ok {
				tmp42 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = &tmp42
		}
		if v, ok := val["lastname"]; ok {
			var tmp43 string
			if val, ok := v.(string); ok {
				tmp43 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Lastname`, v, "string", err)
			}
			target.Lastname = &tmp43
		}
		if v, ok := val["role"]; ok {
			var tmp44 string
			if val, ok := v.(string); ok {
				tmp44 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Role`, v, "string", err)
			}
			target.Role = &tmp44
		}
		if v, ok := val["state"]; ok {
			var tmp45 string
			if val, ok := v.(string); ok {
				tmp45 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.State`, v, "string", err)
			}
			target.State = &tmp45
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}
