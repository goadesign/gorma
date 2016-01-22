package gorma

import (
	"fmt"
	"sort"

	"github.com/goadesign/goa/design"
)

// IterateStores runs an iterator function once per Relational Store in the StorageGroup's Store list
func (sd *StorageGroupDefinition) IterateStores(it StoreIterator) error {
	if sd.RelationalStores != nil {
		names := make([]string, len(sd.RelationalStores))
		i := 0
		for n := range sd.RelationalStores {
			names[i] = n
			i++
		}
		sort.Strings(names)
		for _, n := range names {
			if err := it(sd.RelationalStores[n]); err != nil {
				return err
			}
		}
	}
	return nil
}

// Context returns the generic definition name used in error messages.
func (sd StorageGroupDefinition) Context() string {
	if sd.Name != "" {
		return fmt.Sprintf("StorageGroup %#v", sd.Name)
	}
	return "unnamed Storage Group"
}

// DSL returns this object's DSL
func (sd StorageGroupDefinition) DSL() func() {
	return sd.DefinitionDSL
}

// Children returnsa slice of this objects children
func (sd StorageGroupDefinition) Children() []design.Definition {
	var stores []design.Definition
	for _, s := range sd.RelationalStores {
		stores = append(stores, s)
	}
	return stores
}

// IterateSets goes over all the definition sets of the StorageGroup: The StorageGroup definition itself, each
// store definition, models and fields.
func (a *StorageGroupDefinition) IterateSets(iterator design.SetIterator) {
	// First run the top level StorageGroup

	iterator([]design.Definition{a})
	a.IterateStores(func(store *RelationalStoreDefinition) error {
		iterator([]design.Definition{store})
		store.IterateModels(func(model *RelationalModelDefinition) error {
			iterator([]design.Definition{model})
			model.IterateFields(func(field *RelationalFieldDefinition) error {
				iterator([]design.Definition{field})
				return nil
			})
			return nil
		})
		return nil
	})
}
