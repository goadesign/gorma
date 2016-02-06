//************************************************************************//
// API "cellar": Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"time"

	"github.com/goadesign/goa"
)

// A tenant account, tiny view
// Identifier: application/vnd.account+json
type AccountTiny struct {
	// API href of account
	Href string `json:"href" xml:"href"`
	// ID of account
	ID int `json:"id" xml:"id"`
	// Name of account
	Name string `json:"name" xml:"name"`
}

// Validate validates the media type instance.
func (mt *AccountTiny) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}

	return
}

// A tenant account, link view
// Identifier: application/vnd.account+json
type AccountLink struct {
	// API href of account
	Href string `json:"href" xml:"href"`
	// ID of account
	ID int `json:"id" xml:"id"`
}

// Validate validates the media type instance.
func (mt *AccountLink) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}

	return
}

// A tenant account
// Identifier: application/vnd.account+json
type Account struct {
	// Date of creation
	CreatedAt *time.Time `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Email of account owner
	CreatedBy *string `json:"created_by,omitempty" xml:"created_by,omitempty"`
	// API href of account
	Href string `json:"href" xml:"href"`
	// ID of account
	ID int `json:"id" xml:"id"`
	// Name of account
	Name string `json:"name" xml:"name"`
}

// Validate validates the media type instance.
func (mt *Account) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}

	if mt.CreatedBy != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.CreatedBy); err2 != nil {
			err = goa.InvalidFormatError(`response.created_by`, *mt.CreatedBy, goa.FormatEmail, err2, err)
		}
	}
	return
}

// A bottle of wine
// Identifier: application/vnd.bottle+json
type Bottle struct {
	// Account that owns bottle
	Account *Account `json:"account,omitempty" xml:"account,omitempty"`
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID int `json:"id" xml:"id"`
	// Links to related resources
	Links *BottleLinks `json:"links,omitempty" xml:"links,omitempty"`
	Name  string       `json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating   *int   `json:"rating,omitempty" xml:"rating,omitempty"`
	Varietal string `json:"varietal" xml:"varietal"`
	Vineyard string `json:"vineyard" xml:"vineyard"`
	Vintage  string `json:"vintage" xml:"vintage"`
}

// Validate validates the media type instance.
func (mt *Bottle) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}
	if mt.Vineyard == "" {
		err = goa.MissingAttributeError(`response`, "vineyard", err)
	}
	if mt.Varietal == "" {
		err = goa.MissingAttributeError(`response`, "varietal", err)
	}
	if mt.Vintage == "" {
		err = goa.MissingAttributeError(`response`, "vintage", err)
	}

	if mt.Account != nil {
		if mt.Account.CreatedBy != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Account.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`response.account.created_by`, *mt.Account.CreatedBy, goa.FormatEmail, err2, err)
			}
		}
	}
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 2, true, err)
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false, err)
		}
	}
	if len(mt.Varietal) < 4 {
		err = goa.InvalidLengthError(`response.varietal`, mt.Varietal, len(mt.Varietal), 4, true, err)
	}
	if len(mt.Vineyard) < 2 {
		err = goa.InvalidLengthError(`response.vineyard`, mt.Vineyard, len(mt.Vineyard), 2, true, err)
	}
	return
}

// A bottle of wine, tiny view
// Identifier: application/vnd.bottle+json
type BottleTiny struct {
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID int `json:"id" xml:"id"`
	// Links to related resources
	Links *BottleLinks `json:"links,omitempty" xml:"links,omitempty"`
	Name  string       `json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating *int `json:"rating,omitempty" xml:"rating,omitempty"`
}

// Validate validates the media type instance.
func (mt *BottleTiny) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}

	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 2, true, err)
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false, err)
		}
	}
	return
}

// A bottle of wine, full view
// Identifier: application/vnd.bottle+json
type BottleFull struct {
	// Account that owns bottle
	Account *Account `json:"account,omitempty" xml:"account,omitempty"`
	Color   string   `json:"color" xml:"color"`
	Country *string  `json:"country,omitempty" xml:"country,omitempty"`
	// Date of creation
	CreatedAt *time.Time `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID int `json:"id" xml:"id"`
	// Links to related resources
	Links *BottleLinks `json:"links,omitempty" xml:"links,omitempty"`
	Name  string       `json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating *int    `json:"rating,omitempty" xml:"rating,omitempty"`
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	Review *string `json:"review,omitempty" xml:"review,omitempty"`
	// hello sally
	Sweetness *int `json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	// Date of last update
	UpdatedAt     *time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Varietal      string     `json:"varietal" xml:"varietal"`
	Vineyard      string     `json:"vineyard" xml:"vineyard"`
	Vintage       string     `json:"vintage" xml:"vintage"`
	VinyardCounty *string    `json:"vinyard_county,omitempty" xml:"vinyard_county,omitempty"`
}

// Validate validates the media type instance.
func (mt *BottleFull) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}
	if mt.Vineyard == "" {
		err = goa.MissingAttributeError(`response`, "vineyard", err)
	}
	if mt.Varietal == "" {
		err = goa.MissingAttributeError(`response`, "varietal", err)
	}
	if mt.Vintage == "" {
		err = goa.MissingAttributeError(`response`, "vintage", err)
	}
	if mt.Color == "" {
		err = goa.MissingAttributeError(`response`, "color", err)
	}

	if mt.Account != nil {
		if mt.Account.CreatedBy != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Account.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`response.account.created_by`, *mt.Account.CreatedBy, goa.FormatEmail, err2, err)
			}
		}
	}
	if !(mt.Color == "red" || mt.Color == "white" || mt.Color == "rose" || mt.Color == "yellow" || mt.Color == "sparkling") {
		err = goa.InvalidEnumValueError(`response.color`, mt.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
	}
	if mt.Country != nil {
		if len(*mt.Country) < 2 {
			err = goa.InvalidLengthError(`response.country`, *mt.Country, len(*mt.Country), 2, true, err)
		}
	}
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 2, true, err)
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false, err)
		}
	}
	if mt.Review != nil {
		if len(*mt.Review) < 3 {
			err = goa.InvalidLengthError(`response.review`, *mt.Review, len(*mt.Review), 3, true, err)
		}
	}
	if mt.Review != nil {
		if len(*mt.Review) > 300 {
			err = goa.InvalidLengthError(`response.review`, *mt.Review, len(*mt.Review), 300, false, err)
		}
	}
	if mt.Sweetness != nil {
		if *mt.Sweetness < 1 {
			err = goa.InvalidRangeError(`response.sweetness`, *mt.Sweetness, 1, true, err)
		}
	}
	if mt.Sweetness != nil {
		if *mt.Sweetness > 5 {
			err = goa.InvalidRangeError(`response.sweetness`, *mt.Sweetness, 5, false, err)
		}
	}
	if len(mt.Varietal) < 4 {
		err = goa.InvalidLengthError(`response.varietal`, mt.Varietal, len(mt.Varietal), 4, true, err)
	}
	if len(mt.Vineyard) < 2 {
		err = goa.InvalidLengthError(`response.vineyard`, mt.Vineyard, len(mt.Vineyard), 2, true, err)
	}
	return
}

// BottleLinks contains links to related resources of Bottle.
type BottleLinks struct {
	Account *AccountLink `json:"account,omitempty" xml:"account,omitempty"`
}

// , default view
// Identifier: application/vnd.bottle+json; type=collection
type BottleCollection []*Bottle

// Validate validates the media type instance.
func (mt BottleCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Account != nil {
			if e.Account.CreatedBy != nil {
				if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Account.CreatedBy); err2 != nil {
					err = goa.InvalidFormatError(`response[*].account.created_by`, *e.Account.CreatedBy, goa.FormatEmail, err2, err)
				}
			}
		}
		if len(e.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].name`, e.Name, len(e.Name), 2, true, err)
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 5, false, err)
			}
		}
		if len(e.Varietal) < 4 {
			err = goa.InvalidLengthError(`response[*].varietal`, e.Varietal, len(e.Varietal), 4, true, err)
		}
		if len(e.Vineyard) < 2 {
			err = goa.InvalidLengthError(`response[*].vineyard`, e.Vineyard, len(e.Vineyard), 2, true, err)
		}
	}
	return
}

// , tiny view
// Identifier: application/vnd.bottle+json; type=collection
type BottleTinyCollection []*BottleTiny

// Validate validates the media type instance.
func (mt BottleTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if len(e.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].name`, e.Name, len(e.Name), 2, true, err)
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 5, false, err)
			}
		}
	}
	return
}

// BottleLinksArray contains links to related resources of BottleCollection.
type BottleLinksArray []*BottleLinks

// A box of wine
// Identifier: application/vnd.box+json
type Box struct {
	// Account that owns bottle
	Account *Account `json:"account,omitempty" xml:"account,omitempty"`
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID int `json:"id" xml:"id"`
	// Links to related resources
	Links *BoxLinks `json:"links,omitempty" xml:"links,omitempty"`
	Name  string    `json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating   *int   `json:"rating,omitempty" xml:"rating,omitempty"`
	Varietal string `json:"varietal" xml:"varietal"`
	Vineyard string `json:"vineyard" xml:"vineyard"`
	Vintage  string `json:"vintage" xml:"vintage"`
}

// Validate validates the media type instance.
func (mt *Box) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}
	if mt.Vineyard == "" {
		err = goa.MissingAttributeError(`response`, "vineyard", err)
	}
	if mt.Varietal == "" {
		err = goa.MissingAttributeError(`response`, "varietal", err)
	}
	if mt.Vintage == "" {
		err = goa.MissingAttributeError(`response`, "vintage", err)
	}

	if mt.Account != nil {
		if mt.Account.CreatedBy != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Account.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`response.account.created_by`, *mt.Account.CreatedBy, goa.FormatEmail, err2, err)
			}
		}
	}
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 2, true, err)
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false, err)
		}
	}
	if len(mt.Varietal) < 4 {
		err = goa.InvalidLengthError(`response.varietal`, mt.Varietal, len(mt.Varietal), 4, true, err)
	}
	if len(mt.Vineyard) < 2 {
		err = goa.InvalidLengthError(`response.vineyard`, mt.Vineyard, len(mt.Vineyard), 2, true, err)
	}
	return
}

// A box of wine, tiny view
// Identifier: application/vnd.box+json
type BoxTiny struct {
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID int `json:"id" xml:"id"`
	// Links to related resources
	Links *BoxLinks `json:"links,omitempty" xml:"links,omitempty"`
	Name  string    `json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating *int `json:"rating,omitempty" xml:"rating,omitempty"`
}

// Validate validates the media type instance.
func (mt *BoxTiny) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}

	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 2, true, err)
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false, err)
		}
	}
	return
}

// A box of wine, full view
// Identifier: application/vnd.box+json
type BoxFull struct {
	// Account that owns bottle
	Account *Account `json:"account,omitempty" xml:"account,omitempty"`
	Color   string   `json:"color" xml:"color"`
	Country *string  `json:"country,omitempty" xml:"country,omitempty"`
	// Date of creation
	CreatedAt *time.Time `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID int `json:"id" xml:"id"`
	// Links to related resources
	Links *BoxLinks `json:"links,omitempty" xml:"links,omitempty"`
	Name  string    `json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating    *int    `json:"rating,omitempty" xml:"rating,omitempty"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	// Date of last update
	UpdatedAt     *time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Varietal      string     `json:"varietal" xml:"varietal"`
	Vineyard      string     `json:"vineyard" xml:"vineyard"`
	Vintage       string     `json:"vintage" xml:"vintage"`
	VinyardCounty *string    `json:"vinyard_county,omitempty" xml:"vinyard_county,omitempty"`
}

// Validate validates the media type instance.
func (mt *BoxFull) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}
	if mt.Vineyard == "" {
		err = goa.MissingAttributeError(`response`, "vineyard", err)
	}
	if mt.Varietal == "" {
		err = goa.MissingAttributeError(`response`, "varietal", err)
	}
	if mt.Vintage == "" {
		err = goa.MissingAttributeError(`response`, "vintage", err)
	}
	if mt.Color == "" {
		err = goa.MissingAttributeError(`response`, "color", err)
	}

	if mt.Account != nil {
		if mt.Account.CreatedBy != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Account.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`response.account.created_by`, *mt.Account.CreatedBy, goa.FormatEmail, err2, err)
			}
		}
	}
	if !(mt.Color == "red" || mt.Color == "white" || mt.Color == "rose" || mt.Color == "yellow" || mt.Color == "sparkling") {
		err = goa.InvalidEnumValueError(`response.color`, mt.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
	}
	if mt.Country != nil {
		if len(*mt.Country) < 2 {
			err = goa.InvalidLengthError(`response.country`, *mt.Country, len(*mt.Country), 2, true, err)
		}
	}
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 2, true, err)
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false, err)
		}
	}
	if mt.Review != nil {
		if len(*mt.Review) < 3 {
			err = goa.InvalidLengthError(`response.review`, *mt.Review, len(*mt.Review), 3, true, err)
		}
	}
	if mt.Review != nil {
		if len(*mt.Review) > 300 {
			err = goa.InvalidLengthError(`response.review`, *mt.Review, len(*mt.Review), 300, false, err)
		}
	}
	if mt.Sweetness != nil {
		if *mt.Sweetness < 1 {
			err = goa.InvalidRangeError(`response.sweetness`, *mt.Sweetness, 1, true, err)
		}
	}
	if mt.Sweetness != nil {
		if *mt.Sweetness > 5 {
			err = goa.InvalidRangeError(`response.sweetness`, *mt.Sweetness, 5, false, err)
		}
	}
	if len(mt.Varietal) < 4 {
		err = goa.InvalidLengthError(`response.varietal`, mt.Varietal, len(mt.Varietal), 4, true, err)
	}
	if len(mt.Vineyard) < 2 {
		err = goa.InvalidLengthError(`response.vineyard`, mt.Vineyard, len(mt.Vineyard), 2, true, err)
	}
	return
}

// BoxLinks contains links to related resources of Box.
type BoxLinks struct {
	Account *AccountLink `json:"account,omitempty" xml:"account,omitempty"`
}
