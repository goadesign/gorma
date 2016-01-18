package dsl

import (
	"fmt"

	"github.com/bketelsen/gorma"
)

// StorageGroup implements the top level Gorma DSL
// Examples and more docs here later
func RelationalModel(name string, dsl func()) {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.
	checkInit()
	if s, ok := relationalStoreDefinition(true); ok {
		if s.RelationalModels == nil {
			s.RelationalModels = make(map[string]*gorma.RelationalModelDefinition)
		}
		store, ok := s.RelationalModels[name]
		if !ok {
			store := &RelationalModelDefinition{
				Name: name,
				DSL:  dsl,
			}
		}
		if !executeDSL(dsl, store) { // @raphael - who is executing this?
			return
		}
		s.RelationalModels[name] = store
	}

}

// Context returns the generic definition name used in error messages.
func (a *RelationalModelDefinition) Context() string {
	if a.Name != "" {
		return fmt.Sprintf("RelationalModel %#v", a.Name)
	}
	return "unnamed RelationalModel"
}
