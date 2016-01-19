package dsl

import "github.com/bketelsen/gorma"

// StorageGroup implements the top level Gorma DSL
// Examples and more docs here later
func RelationalStore(name string, storeType gorma.RelationalStorageType, dsl func()) {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.
	checkInit()
	if s, ok := storageGroupDefinition(true); ok {
		if s.RelationalStores == nil {
			s.RelationalStores = make(map[string]*gorma.RelationalStoreDefinition)
		}
		store, ok := s.RelationalStores[name]
		if !ok {
			store = &gorma.RelationalStoreDefinition{
				Name:             name,
				DefinitionDSL:    dsl,
				Parent:           s,
				Type:             storeType,
				RelationalModels: make(map[string]*gorma.RelationalModelDefinition),
			}
		}
		s.RelationalStores[name] = store
	}

}
