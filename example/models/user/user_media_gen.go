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

package user

import "github.com/raphael/goa"

// LoadUser loads raw data into an instance of User
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadUser(raw interface{}) (res *User, err error) {
	res, err = UnmarshalUser(raw, err)
	return
}

// MarshalUser validates and renders an instance of User into a interface{}
func MarshalUser(source *User, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp21 := map[string]interface{}{
		"bio":        source.Bio,
		"city":       source.City,
		"country":    source.Country,
		"created_at": source.CreatedAt,
		"email":      source.Email,
		"first_name": source.FirstName,
		"id":         source.ID,
		"last_name":  source.LastName,
		"role":       source.Role,
		"state":      source.State,
		"updated_at": source.UpdatedAt,
	}
	target = tmp21
	return
}

// UnmarshalUser unmarshals and validates a raw interface{} into an instance of User
func UnmarshalUser(source interface{}, inErr error) (target *User, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(User)
		if v, ok := val["bio"]; ok {
			var tmp22 string
			if val, ok := v.(string); ok {
				tmp22 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Bio`, v, "string", err)
			}
			target.Bio = &tmp22
		}
		if v, ok := val["city"]; ok {
			var tmp23 string
			if val, ok := v.(string); ok {
				tmp23 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.City`, v, "string", err)
			}
			target.City = &tmp23
		}
		if v, ok := val["country"]; ok {
			var tmp24 string
			if val, ok := v.(string); ok {
				tmp24 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			target.Country = &tmp24
		}
		if v, ok := val["created_at"]; ok {
			var tmp25 string
			if val, ok := v.(string); ok {
				tmp25 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.CreatedAt`, v, "string", err)
			}
			target.CreatedAt = &tmp25
		}
		if v, ok := val["email"]; ok {
			var tmp26 string
			if val, ok := v.(string); ok {
				tmp26 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			target.Email = &tmp26
		}
		if v, ok := val["first_name"]; ok {
			var tmp27 string
			if val, ok := v.(string); ok {
				tmp27 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.FirstName`, v, "string", err)
			}
			target.FirstName = &tmp27
		}
		if v, ok := val["id"]; ok {
			var tmp28 int
			if f, ok := v.(float64); ok {
				tmp28 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp28
		} else {
			err = goa.MissingAttributeError(`load`, "id", err)
		}
		if v, ok := val["last_name"]; ok {
			var tmp29 string
			if val, ok := v.(string); ok {
				tmp29 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.LastName`, v, "string", err)
			}
			target.LastName = &tmp29
		}
		if v, ok := val["role"]; ok {
			var tmp30 string
			if val, ok := v.(string); ok {
				tmp30 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Role`, v, "string", err)
			}
			target.Role = tmp30
		} else {
			err = goa.MissingAttributeError(`load`, "role", err)
		}
		if v, ok := val["state"]; ok {
			var tmp31 string
			if val, ok := v.(string); ok {
				tmp31 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.State`, v, "string", err)
			}
			target.State = &tmp31
		}
		if v, ok := val["updated_at"]; ok {
			var tmp32 string
			if val, ok := v.(string); ok {
				tmp32 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.UpdatedAt`, v, "string", err)
			}
			target.UpdatedAt = &tmp32
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}
