package gorma

import (
	"sort"

	"github.com/raphael/goa/design"
)

// NewStorageGroup creates a StorageGroup structure by parsing
// the APIDefinition and creating all the necessary Stores and Models
func NewStorageGroup(a *design.APIDefinition) (*StorageGroup, error) {
	sg := &StorageGroup{}
	sg.api = a
	sg.RelationalStore = make([]*RelationalStoreDefinition, 0)
	NewRelationalStoreDefinition()
	err := sg.Parse()
	return sg, err
}

// Parse is the function that should be called to parse a full Goa API definition.
// It populates the StorageGroup and recursively all children objects
func (sg *StorageGroup) Parse() error {

	err := sg.api.IterateVersions(func(v *design.APIVersionDefinition) error {
		err := v.IterateUserTypes(func(t *design.UserTypeDefinition) error {
			if t.Type.IsObject() {
				name := t.TypeName
				m, err := NewRelationalModel(name, t)
				if err != nil {
					return err
				}
				sg.RelationalStoreDefinition.Models[name] = m
			}
			return nil
		}) // IterateUserTypes
		return err
	}) // IterateVersions
	if err != nil {
		return err
	}
	err = sg.RelationalStoreDefinition.ResolveRelationships()

	return err
}

// IterateStores runs an iterator function once per Relational Store in the StorageGroup's Store list
func (sd *StorageGroupDefinition) IterateStores(it StoreIterator) error {
	names := make([]string, len(sd.Stores))
	i := 0
	for n := range sd.Stores {
		names[i] = n
		i++
	}
	sort.Strings(names)
	for _, n := range names {
		if err := it(sd.Stores[n]); err != nil {
			return err
		}
	}
	return nil
}
