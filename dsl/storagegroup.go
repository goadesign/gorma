package dsl

import (
	"fmt"

	"github.com/bketelsen/gorma"
	"github.com/raphael/goa/design"
)

// StorageGroup implements the top level Gorma DSL
// Examples and more docs here later
func StorageGroup(name string, dsl func()) *gorma.StorageGroupDefinition {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.

	checkInit()

	sg := &StorageGroupDefinition{
		Name:             name,
		RelationalStores: make(map[string]*RelationalStoreDefinition),
		DSL:              dsl,
	}

	if !topLevelDefinition(true) {
		return nil
	}
	if name == "" {
		design.ReportError("Storage Group name cannot be empty")
	}

	gorma.Design = sg
	return gorma.Design
}

// Context returns the generic definition name used in error messages.
func (a *StorageGroupDefinition) Context() string {
	if a.Name != "" {
		return fmt.Sprintf("StorageGroup %#v", a.Name)
	}
	return "unnamed Storage Group"
}
