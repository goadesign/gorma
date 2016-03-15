package dsl

import (
	"github.com/goadesign/goa/dslengine"
	"github.com/goadesign/gorma"
)

// StorageGroup implements the top level Gorma DSL
// There should be one StorageGroup per Goa application.
func StorageGroup(name string, dsli func()) *gorma.StorageGroupDefinition {
	if !dslengine.IsTopLevelDefinition() {
		return nil
	}
	if name == "" {
		dslengine.ReportError("Storage Group name cannot be empty")
	}

	if gorma.GormaDesign != nil {
		if gorma.GormaDesign.Name == name {
			dslengine.ReportError("Only one StorageGroup is allowed")
		}
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
