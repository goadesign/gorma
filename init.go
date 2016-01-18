package gorma

import "github.com/raphael/goa/design"

// GormaConstructs is a map of Gorma types (currently only a StorageGroupDefinition
var GormaConstructs map[string]design.DSLDefinition

// Design is the top level StorageGroup Definition
var Design *StorageGroupDefinition

const (
	// StorageGroup is the constant string used as the index in the
	// GormaConstructs map
	StorageGroup = "storagegroup"
)

// Init creates the necessary data structures for parsing a DSL
func Init() {
	// 	GormaConstructs = design.Design.NewConstructsSet("gorma") // later
	GormaConstructs = make(map[string]design.DSLDefinition)
	sg := &StorageGroupDefinition{}
	GormaConstructs[StorageGroup] = sg

}
