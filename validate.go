package gorma

import "github.com/raphael/goa/design"

// Validate tests whether the StorageGroup definition is consistent
func (a *StorageGroupDefinition) Validate() *design.ValidationErrors {
	verr := new(goa.ValidationErrors)

	a.IterateStores(func(store *RelationalStoreDefinition) error {
		verr.Merge(store.Validate())
	})

	return verr.AsError()
}

// Validate tests whether the RelationalStore definition is consistent
func (a *RelationalStoreDefinition) Validate() *design.ValidationErrors {
	verr := new(goa.ValidationErrors)

	a.IterateModels(func(model *RelationalModelDefinition) error {
		verr.Merge(model.Validate())
	})

	return verr.AsError()
}

// Validate tests whether the RelationalModel definition is consistent
func (a *RelationalModelDefinition) Validate() *design.ValidationErrors {

	verr := new(goa.ValidationErrors)

	a.IterateFields(func(field *RelationalFieldDefinition) error {
		verr.Merge(field.Validate())
	})

	return verr.AsError()
}

// Validate tests whether the RelationalField definition is consistent
func (a *RelationalFieldDefinition) Validate() *design.ValidationErrors {
	verr := new(goa.ValidationErrors)
	if field.Name == "" {
		verr.Add(a, "field name not defined")
	}
	return verr.AsError()
}
