package dsl

import (
	"fmt"

	"github.com/bketelsen/gorma"
)

// StorageGroup implements the top level Gorma DSL
// Examples and more docs here later
func RelationalField(name string, dsl func()) {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.
	checkInit()
	if s, ok := relationalModelDefinition(true); ok {
		if s.RelationalFields == nil {
			s.RelationalFields = make(map[string]*gorma.RelationalFieldDefinition)
		}
		store, ok := s.RelationalFields[name]
		if !ok {
			store := &RelationalFieldDefinition{
				Name: name,
				DSL:  dsl,
			}
		}
		if !executeDSL(dsl, store) { // @raphael - who is executing this?
			return
		}
		s.RelationalFields[name] = store
	}

}

// Context returns the generic definition name used in error messages.
func (a *RelationalFieldDefinition) Context() string {
	if a.Name != "" {
		return fmt.Sprintf("RelationalField %#v", a.Name)
	}
	return "unnamed RelationalField"
}
