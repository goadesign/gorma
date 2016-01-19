package dsl

import (
	"github.com/bketelsen/gorma"
	"github.com/raphael/goa/design/dsl"
)

// StorageGroup implements the top level Gorma DSL
// Examples and more docs here later
func StorageGroup(name string, dsli func()) *gorma.StorageGroupDefinition {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.

	checkInit()

	sg := &gorma.StorageGroupDefinition{
		Name:             name,
		RelationalStores: make(map[string]*gorma.RelationalStoreDefinition),
		DefinitionDSL:    dsli,
	}

	if !topLevelDefinition(true) {
		return nil
	}
	if name == "" {
		dsl.ReportError("Storage Group name cannot be empty")
	}

	gorma.GormaConstructs[gorma.StorageGroup] = sg
	return sg
}

// Description sets the definition description.
// Description can be called inside StorageGroup, RelationalStore, RelationalModel, RelationalField
func Description(d string) {
	if a, ok := storageGroupDefinition(false); ok {
		a.Description = d
	} else if v, ok := relationalStoreDefinition(false); ok {
		v.Description = d
	} else if r, ok := relationalModelDefinition(false); ok {
		r.Description = d
	} else if a, ok := relationalFieldDefinition(false); ok {
		a.Description = d
	}
}
