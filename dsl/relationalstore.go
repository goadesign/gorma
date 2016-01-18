package dsl

import (
	"fmt"

	"github.com/bketelsen/gorma"
)

// StorageGroup implements the top level Gorma DSL
// Examples and more docs here later
func RelationalStore(name string, dsl func()) {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.
	checkInit()
	if s, ok := storageGroupDefinition(true); ok {
		if s.RelationalStores == nil {
			s.RelationalStores = make(map[string]*gorma.RelationalStoreDefinition)
		}
		store, ok := s.RelationalStores[name]
		if !ok {
			store := &RelationalStoreDefinition{
				Name: name,
				DSL:  dsl,
			}
		}
		if !executeDSL(dsl, store) { // @raphael - who is executing this?
			return
		}
		s.RelationalStores[name] = store
	}

}

// Context returns the generic definition name used in error messages.
func (a *RelationalStoreDefinition) Context() string {
	if a.Name != "" {
		return fmt.Sprintf("RelationalStore %#v", a.Name)
	}
	return "unnamed RelationalStore"
}
