package dsl

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/dslengine"
	"github.com/goadesign/gorma"
)

func init() {
	gorma.GormaDesign = gorma.NewStorageGroupDefinition()
	dslengine.Register(gorma.GormaDesign)
}

// storageDefinition returns true and current context if it is an StorageGroupDefinition,
// nil and false otherwise.
func storageGroupDefinition(failIfNotSD bool) (*gorma.StorageGroupDefinition, bool) {
	a, ok := dslengine.CurrentDefinition().(*gorma.StorageGroupDefinition)
	if !ok && failIfNotSD {
		dslengine.IncompatibleDSL()
	}
	return a, ok
}

// relationalStoreDefinition returns true and current context if it is an RelationalStoreDefinition,
// nil and false otherwise.
func relationalStoreDefinition(failIfNotSD bool) (*gorma.RelationalStoreDefinition, bool) {
	a, ok := dslengine.CurrentDefinition().(*gorma.RelationalStoreDefinition)
	if !ok && failIfNotSD {
		dslengine.IncompatibleDSL()
	}
	return a, ok
}

// relationalModelDefinition returns true and current context if it is an RelationalModelDefinition,
// nil and false otherwise.
func relationalModelDefinition(failIfNotSD bool) (*gorma.RelationalModelDefinition, bool) {
	a, ok := dslengine.CurrentDefinition().(*gorma.RelationalModelDefinition)
	if !ok && failIfNotSD {
		dslengine.IncompatibleDSL()
	}
	return a, ok
}

// relationalFieldDefinition returns true and current context if it is an RelationalFieldDefinition,
// nil and false otherwise.
func relationalFieldDefinition(failIfNotSD bool) (*gorma.RelationalFieldDefinition, bool) {
	a, ok := dslengine.CurrentDefinition().(*gorma.RelationalFieldDefinition)
	if !ok && failIfNotSD {
		dslengine.IncompatibleDSL()
	}
	return a, ok
}

// buildSourceDefinition returns true and current context if it is an BuildSource
// nil and false otherwise.
func buildSourceDefinition(failIfNotSD bool) (*gorma.BuildSource, bool) {
	a, ok := dslengine.CurrentDefinition().(*gorma.BuildSource)
	if !ok && failIfNotSD {
		dslengine.IncompatibleDSL()
	}
	return a, ok
}

// attributeDefinition returns true and current context if it is an AttributeDefinition
// nil and false otherwise.
func attributeDefinition(failIfNotSD bool) (*design.AttributeDefinition, bool) {
	a, ok := dslengine.CurrentDefinition().(*design.AttributeDefinition)
	if !ok && failIfNotSD {
		dslengine.IncompatibleDSL()
	}
	return a, ok
}
