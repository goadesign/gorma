package dsl

import (
	"github.com/bketelsen/gorma"
	"github.com/raphael/goa/design"
)

// StorageGroup implements the top level Gorma DSL
// Examples and more docs here later
func StorageGroup(name string, dsl func()) *gorma.StorageGroupDefinition {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.
	if design.Design == nil {
		design.InitDesign()
	}

	// check to see if this type is registered
	set, ok := design.Design.ConstructSet["gorma"]
	if !ok {
		// There is no registered gorma construct set
	} else {
		// There is a registered construct set
	}
	if !topLevelDefinition(true) {
		return nil
	}
	if name == "" {
		design.ReportError("Storage Group name cannot be empty")
	}
	design.Design.Name = name
	design.Design.DSL = dsl
	return design.Design
}
