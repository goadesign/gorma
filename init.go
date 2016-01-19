package gorma

import "github.com/raphael/goa/design"

// GormaConstructs is a map of Gorma types (currently only a StorageGroupDefinition
var GormaConstructs design.Construct

const (
	// Gorma is the constant string used as the index in the
	// goa DesignConstructs map
	Gorma = "gorma"
	// StorageGroup is the constant string used as the index in the
	// GormaConstructs map
	StorageGroup                       = "storagegroup"
	MySQL        RelationalStorageType = "mysql"
	Postgres     RelationalStorageType = "postgres"
)

// Init creates the necessary data structures for parsing a DSL
func Init() {
	// 	GormaConstructs = design.Design.NewConstructsSet("gorma") // later
	GormaConstructs = design.NewConstruct(Gorma)
	sg := &StorageGroupDefinition{}
	GormaConstructs[StorageGroup] = sg

}
