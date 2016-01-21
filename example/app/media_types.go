//************************************************************************//
// API "cellar": Application Media Types
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

import "github.com/goadesign/goa"

// A tenant account
// Identifier: application/vnd.account+json
type Account struct {
	// Date of creation
	CreatedAt *string
	// Email of account owner
	CreatedBy *string
	// API href of account
	Href string
	// ID of account
	ID int
	// Name of account
	Name string
}

// Account views
type AccountViewEnum string

const (
	// Account default view
	AccountDefaultView AccountViewEnum = "default"
	// Account link view
	AccountLinkView AccountViewEnum = "link"
	// Account tiny view
	AccountTinyView AccountViewEnum = "tiny"
)

// Dump produces raw data from an instance of Account running all the
// validations. See LoadAccount for the definition of raw data.
func (mt *Account) Dump(view AccountViewEnum) (res map[string]interface{}, err error) {
	if view == AccountDefaultView {
		res, err = MarshalAccount(mt, err)
	}
	if view == AccountLinkView {
		res, err = MarshalAccountLink(mt, err)
	}
	if view == AccountTinyView {
		res, err = MarshalAccountTiny(mt, err)
	}
	return
}

// Validate validates the media type instance.
func (mt *Account) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}

	if mt.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.created_at`, *mt.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if mt.CreatedBy != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.CreatedBy); err2 != nil {
			err = goa.InvalidFormatError(`response.created_by`, *mt.CreatedBy, goa.FormatEmail, err2, err)
		}
	}
	return
}

// MarshalAccount validates and renders an instance of Account into a interface{}
// using view "default".
func MarshalAccount(source *Account, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp15 := map[string]interface{}{
		"created_at": source.CreatedAt,
		"created_by": source.CreatedBy,
		"href":       source.Href,
		"id":         source.ID,
		"name":       source.Name,
	}
	target = tmp15
	return
}

// MarshalAccountLink validates and renders an instance of Account into a interface{}
// using view "link".
func MarshalAccountLink(source *Account, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp16 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
	}
	target = tmp16
	return
}

// MarshalAccountTiny validates and renders an instance of Account into a interface{}
// using view "tiny".
func MarshalAccountTiny(source *Account, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp17 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
		"name": source.Name,
	}
	target = tmp17
	return
}

// A bottle of wine
// Identifier: application/vnd.bottle+json
type Bottle struct {
	// Account that owns bottle
	Account *Account
	Color   string
	Country *string
	// Date of creation
	CreatedAt *string
	// API href of bottle
	Href string
	// ID of bottle
	ID   int
	Name string
	// Rating of bottle between 1 and 5
	Rating    *int
	Region    *string
	Review    *string
	Sweetness *int
	// Date of last update
	UpdatedAt *string
	Varietal  string
	Vineyard  string
	Vintage   int
}

// Bottle views
type BottleViewEnum string

const (
	// Bottle default view
	BottleDefaultView BottleViewEnum = "default"
	// Bottle full view
	BottleFullView BottleViewEnum = "full"
	// Bottle tiny view
	BottleTinyView BottleViewEnum = "tiny"
)

// Dump produces raw data from an instance of Bottle running all the
// validations. See LoadBottle for the definition of raw data.
func (mt *Bottle) Dump(view BottleViewEnum) (res map[string]interface{}, err error) {
	if view == BottleDefaultView {
		res, err = MarshalBottle(mt, err)
	}
	if view == BottleFullView {
		res, err = MarshalBottleFull(mt, err)
	}
	if view == BottleTinyView {
		res, err = MarshalBottleTiny(mt, err)
	}
	return
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

	if mt.Color == "" {
		err = goa.MissingAttributeError(`response`, "color", err)
	}

	if mt.Account.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.Account.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.account.created_at`, *mt.Account.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if mt.Account.CreatedBy != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Account.CreatedBy); err2 != nil {
			err = goa.InvalidFormatError(`response.account.created_by`, *mt.Account.CreatedBy, goa.FormatEmail, err2, err)
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
	if mt.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.created_at`, *mt.CreatedAt, goa.FormatDateTime, err2, err)
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
	if mt.UpdatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.UpdatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.updated_at`, *mt.UpdatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if len(mt.Varietal) < 4 {
		err = goa.InvalidLengthError(`response.varietal`, mt.Varietal, len(mt.Varietal), 4, true, err)
	}
	if len(mt.Vineyard) < 2 {
		err = goa.InvalidLengthError(`response.vineyard`, mt.Vineyard, len(mt.Vineyard), 2, true, err)
	}
	if mt.Vintage < 1900 {
		err = goa.InvalidRangeError(`response.vintage`, mt.Vintage, 1900, true, err)
	}
	if mt.Vintage > 2020 {
		err = goa.InvalidRangeError(`response.vintage`, mt.Vintage, 2020, false, err)
	}
	return
}

// MarshalBottle validates and renders an instance of Bottle into a interface{}
// using view "default".
func MarshalBottle(source *Bottle, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp18 := map[string]interface{}{
		"href":     source.Href,
		"id":       source.ID,
		"name":     source.Name,
		"rating":   source.Rating,
		"varietal": source.Varietal,
		"vineyard": source.Vineyard,
		"vintage":  source.Vintage,
	}
	if source.Account != nil {
		tmp18["account"], err = MarshalAccountTiny(source.Account, err)
	}
	target = tmp18
	if err == nil {
		links := make(map[string]interface{})
		links["account"], err = MarshalAccountLink(source.Account, err)
		target["links"] = links
	}
	return
}

// MarshalBottleFull validates and renders an instance of Bottle into a interface{}
// using view "full".
func MarshalBottleFull(source *Bottle, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp19 := map[string]interface{}{
		"color":      source.Color,
		"country":    source.Country,
		"created_at": source.CreatedAt,
		"href":       source.Href,
		"id":         source.ID,
		"name":       source.Name,
		"rating":     source.Rating,
		"region":     source.Region,
		"review":     source.Review,
		"sweetness":  source.Sweetness,
		"updated_at": source.UpdatedAt,
		"varietal":   source.Varietal,
		"vineyard":   source.Vineyard,
		"vintage":    source.Vintage,
	}
	if source.Account != nil {
		tmp19["account"], err = MarshalAccount(source.Account, err)
	}
	target = tmp19
	if err == nil {
		links := make(map[string]interface{})
		links["account"], err = MarshalAccountLink(source.Account, err)
		target["links"] = links
	}
	return
}

// MarshalBottleTiny validates and renders an instance of Bottle into a interface{}
// using view "tiny".
func MarshalBottleTiny(source *Bottle, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp20 := map[string]interface{}{
		"href":   source.Href,
		"id":     source.ID,
		"name":   source.Name,
		"rating": source.Rating,
	}
	target = tmp20
	if err == nil {
		links := make(map[string]interface{})
		links["account"], err = MarshalAccountLink(source.Account, err)
		target["links"] = links
	}
	return
}

// BottleCollection media type
// Identifier: application/vnd.bottle+json; type=collection
type BottleCollection []*Bottle

// BottleCollection views
type BottleCollectionViewEnum string

const (
	// BottleCollection default view
	BottleCollectionDefaultView BottleCollectionViewEnum = "default"
	// BottleCollection tiny view
	BottleCollectionTinyView BottleCollectionViewEnum = "tiny"
)

// Dump produces raw data from an instance of BottleCollection running all the
// validations. See LoadBottleCollection for the definition of raw data.
func (mt BottleCollection) Dump(view BottleCollectionViewEnum) (res []map[string]interface{}, err error) {
	if view == BottleCollectionDefaultView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp21 := range mt {
			var tmp22 map[string]interface{}
			tmp22, err = MarshalBottle(tmp21, err)
			res[i] = tmp22
		}
	}
	if view == BottleCollectionTinyView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp23 := range mt {
			var tmp24 map[string]interface{}
			tmp24, err = MarshalBottleTiny(tmp23, err)
			res[i] = tmp24
		}
	}
	return
}

// Validate validates the media type instance.
func (mt BottleCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Account.CreatedAt != nil {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, *e.Account.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response[*].account.created_at`, *e.Account.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if e.Account.CreatedBy != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Account.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`response[*].account.created_by`, *e.Account.CreatedBy, goa.FormatEmail, err2, err)
			}
		}
		if !(e.Color == "red" || e.Color == "white" || e.Color == "rose" || e.Color == "yellow" || e.Color == "sparkling") {
			err = goa.InvalidEnumValueError(`response[*].color`, e.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
		}
		if e.Country != nil {
			if len(*e.Country) < 2 {
				err = goa.InvalidLengthError(`response[*].country`, *e.Country, len(*e.Country), 2, true, err)
			}
		}
		if e.CreatedAt != nil {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, *e.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response[*].created_at`, *e.CreatedAt, goa.FormatDateTime, err2, err)
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
		if e.Review != nil {
			if len(*e.Review) < 3 {
				err = goa.InvalidLengthError(`response[*].review`, *e.Review, len(*e.Review), 3, true, err)
			}
		}
		if e.Review != nil {
			if len(*e.Review) > 300 {
				err = goa.InvalidLengthError(`response[*].review`, *e.Review, len(*e.Review), 300, false, err)
			}
		}
		if e.Sweetness != nil {
			if *e.Sweetness < 1 {
				err = goa.InvalidRangeError(`response[*].sweetness`, *e.Sweetness, 1, true, err)
			}
		}
		if e.Sweetness != nil {
			if *e.Sweetness > 5 {
				err = goa.InvalidRangeError(`response[*].sweetness`, *e.Sweetness, 5, false, err)
			}
		}
		if e.UpdatedAt != nil {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, *e.UpdatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response[*].updated_at`, *e.UpdatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if len(e.Varietal) < 4 {
			err = goa.InvalidLengthError(`response[*].varietal`, e.Varietal, len(e.Varietal), 4, true, err)
		}
		if len(e.Vineyard) < 2 {
			err = goa.InvalidLengthError(`response[*].vineyard`, e.Vineyard, len(e.Vineyard), 2, true, err)
		}
		if e.Vintage < 1900 {
			err = goa.InvalidRangeError(`response[*].vintage`, e.Vintage, 1900, true, err)
		}
		if e.Vintage > 2020 {
			err = goa.InvalidRangeError(`response[*].vintage`, e.Vintage, 2020, false, err)
		}
	}
	return
}

// MarshalBottleCollection validates and renders an instance of BottleCollection into a interface{}
// using view "default".
func MarshalBottleCollection(source BottleCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalBottle(res, err)
	}
	return
}

// MarshalBottleCollectionTiny validates and renders an instance of BottleCollection into a interface{}
// using view "tiny".
func MarshalBottleCollectionTiny(source BottleCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalBottleTiny(res, err)
	}
	return
}
