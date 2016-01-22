package dsl

import (
	"github.com/goadesign/gorma"
	"github.com/goadesign/goa/design/dsl"
)

// StorageGroup implements the top level Gorma DSL
// There should be one StorageGroup per Goa application.
func StorageGroup(name string, dsli func()) *gorma.StorageGroupDefinition {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.
	checkInit()

	if !topLevelDefinition(true) {
		return nil
	}
	if name == "" {
		dsl.ReportError("Storage Group name cannot be empty")
	}
	if gorma.GormaDesign != nil {
		dsl.ReportError("Only one StorageGroup is allowed")
	}
	gorma.GormaDesign.Name = name
	gorma.GormaDesign.RelationalStores = make(map[string]*gorma.RelationalStoreDefinition)
	gorma.GormaDesign.DefinitionDSL = dsli
	return gorma.GormaDesign
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
	} else if f, ok := relationalFieldDefinition(false); ok {
		f.Description = d
	}
}
