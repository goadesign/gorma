package gorma

import (
	"fmt"
	"sort"

	"github.com/raphael/goa/design"
)

// Context returns the generic definition name used in error messages.
func (sd *RelationalStoreDefinition) Context() string {
	if sd.Name != "" {
		return fmt.Sprintf("RelationalStore %#v", sd.Name)
	}
	return "unnamed RelationalStore"
}

// DSL returns this object's DSL
func (sd *RelationalStoreDefinition) DSL() func() {
	fmt.Println("Retrieving Store's DSL")
	return sd.DefinitionDSL
}

// Children returnsa slice of this objects children
func (sd RelationalStoreDefinition) Children() []design.Definition {
	var stores []design.Definition
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
