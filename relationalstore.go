package gorma

import (
	"fmt"
	"sort"

	"github.com/goadesign/goa/dslengine"
)

// NewRelationalStoreDefinition returns an initialized
// RelationalStoreDefinition.
func NewRelationalStoreDefinition() *RelationalStoreDefinition {
	m := &RelationalStoreDefinition{
		RelationalModels: make(map[string]*RelationalModelDefinition),
	}
	return m
}

// Context returns the generic definition name used in error messages.
func (sd *RelationalStoreDefinition) Context() string {
	if sd.Name != "" {
		return fmt.Sprintf("RelationalStore %#v", sd.Name)
	}
	return "unnamed RelationalStore"
}

// DSL returns this object's DSL.
func (sd *RelationalStoreDefinition) DSL() func() {
	return sd.DefinitionDSL
}

// Children returns a slice of this objects children.
func (sd RelationalStoreDefinition) Children() []dslengine.Definition {
	var stores []dslengine.Definition
	for _, s := range sd.RelationalModels {
		stores = append(stores, s)
	}
	return stores
}

// IterateModels runs an iterator function once per Model in the Store's model list.
func (sd *RelationalStoreDefinition) IterateModels(it ModelIterator) error {
	names := make([]string, len(sd.RelationalModels))
	i := 0
	for n := range sd.RelationalModels {
		names[i] = n
		i++
	}
	sort.Strings(names)
	for _, n := range names {
		if err := it(sd.RelationalModels[n]); err != nil {
			return err
		}
	}
	return nil
}
