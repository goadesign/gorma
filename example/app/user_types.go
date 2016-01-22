//************************************************************************//
// API "cellar": Application User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// BottlePayload type
type BottlePayload struct {
	Color     *string
	Country   *string
	Name      *string
	Region    *string
	Review    *string
	Sweetness *int
	Varietal  *string
	Vineyard  *string
	Vintage   *int
}

// Validate validates the type instance.
func (ut *BottlePayload) Validate() (err error) {
	if ut.Color != nil {
		if !(*ut.Color == "red" || *ut.Color == "white" || *ut.Color == "rose" || *ut.Color == "yellow" || *ut.Color == "sparkling") {
			err = goa.InvalidEnumValueError(`response.color`, *ut.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
		}
	}
	if ut.Country != nil {
		if len(*ut.Country) < 2 {
			err = goa.InvalidLengthError(`response.country`, *ut.Country, len(*ut.Country), 2, true, err)
		}
	}
	if ut.Name != nil {
		if len(*ut.Name) < 2 {
			err = goa.InvalidLengthError(`response.name`, *ut.Name, len(*ut.Name), 2, true, err)
		}
	}
	if ut.Review != nil {
		if len(*ut.Review) < 3 {
			err = goa.InvalidLengthError(`response.review`, *ut.Review, len(*ut.Review), 3, true, err)
		}
	}
	if ut.Review != nil {
		if len(*ut.Review) > 300 {
			err = goa.InvalidLengthError(`response.review`, *ut.Review, len(*ut.Review), 300, false, err)
		}
	}
	if ut.Sweetness != nil {
		if *ut.Sweetness < 1 {
			err = goa.InvalidRangeError(`response.sweetness`, *ut.Sweetness, 1, true, err)
		}
	}
	if ut.Sweetness != nil {
		if *ut.Sweetness > 5 {
			err = goa.InvalidRangeError(`response.sweetness`, *ut.Sweetness, 5, false, err)
		}
	}
	if ut.Varietal != nil {
		if len(*ut.Varietal) < 4 {
			err = goa.InvalidLengthError(`response.varietal`, *ut.Varietal, len(*ut.Varietal), 4, true, err)
		}
	}
	if ut.Vineyard != nil {
		if len(*ut.Vineyard) < 2 {
			err = goa.InvalidLengthError(`response.vineyard`, *ut.Vineyard, len(*ut.Vineyard), 2, true, err)
		}
	}
	if ut.Vintage != nil {
		if *ut.Vintage < 1900 {
			err = goa.InvalidRangeError(`response.vintage`, *ut.Vintage, 1900, true, err)
		}
	}
	if ut.Vintage != nil {
		if *ut.Vintage > 2020 {
			err = goa.InvalidRangeError(`response.vintage`, *ut.Vintage, 2020, false, err)
		}
	}
	return
}

// MarshalBottlePayload validates and renders an instance of BottlePayload into a interface{}
func MarshalBottlePayload(source *BottlePayload, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp25 := map[string]interface{}{
		"color":     source.Color,
		"country":   source.Country,
		"name":      source.Name,
		"region":    source.Region,
		"review":    source.Review,
		"sweetness": source.Sweetness,
		"varietal":  source.Varietal,
		"vineyard":  source.Vineyard,
		"vintage":   source.Vintage,
	}
	target = tmp25
	return
}
