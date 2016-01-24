package gorma

import (
	"fmt"

	"github.com/goadesign/goa/design"
)

// Validate tests whether the StorageGroup definition is consistent.
func (a *StorageGroupDefinition) Validate() *design.ValidationErrors {
	fmt.Println("Validating Group")
	verr := new(design.ValidationErrors)
	if a.Name == "" {
		verr.Add(a, "storage group name not defined")
	}
	a.IterateStores(func(store *RelationalStoreDefinition) error {
		verr.Merge(store.Validate())
		return nil
	})

	return verr.AsError()
}

// Validate tests whether the RelationalStore definition is consistent.
func (a *RelationalStoreDefinition) Validate() *design.ValidationErrors {
	fmt.Println("Validating Store")
	verr := new(design.ValidationErrors)
	if a.Name == "" {
		verr.Add(a, "store name not defined")
	}
	if a.Parent == nil {
		verr.Add(a, "missing storage group parent")
	}
	a.IterateModels(func(model *RelationalModelDefinition) error {
		verr.Merge(model.Validate())
		return nil
	})

	return verr.AsError()
}

// Validate tests whether the RelationalModel definition is consistent.
func (a *RelationalModelDefinition) Validate() *design.ValidationErrors {
	fmt.Println("Validating Model")
	verr := new(design.ValidationErrors)
	if a.Name == "" {
		verr.Add(a, "model name not defined")
	}
	if a.Parent == nil {
		verr.Add(a, "missing relational store parent")
	}
	a.IterateFields(func(field *RelationalFieldDefinition) error {
		verr.Merge(field.Validate())
		return nil
	})

	return verr.AsError()
}

// Validate tests whether the RelationalField definition is consistent.
func (field *RelationalFieldDefinition) Validate() *design.ValidationErrors {
	fmt.Println("Validing Field")
	verr := new(design.ValidationErrors)

	if field.Parent == nil {
		verr.Add(field, "missing relational model parent")
	}
	if field.Name == "" {
		verr.Add(field, "field name not defined")
	}
	return verr.AsError()
}
