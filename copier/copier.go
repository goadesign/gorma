package copier

import (
	"encoding/json"
	"errors"
)

// Copy populates all compatible/matching fields from the
// left struct into the right struct.
// right struct MUST be a pointer value
// Copy uses json marshal/unmarshal (right now)
// therefore will be as compatible as your structs
// are with JSON.  In Goa/Gorma this should be fine
func Copy(left, right interface{}) error {
	b, err := json.Marshal(left)
	if err != nil {
		return errors.New("Error marshaling left:" + err.Error())
	}
	err = json.Unmarshal(b, right)
	if err != nil {
		return errors.New("Error populating right:" + err.Error())
	}

	return nil
}
