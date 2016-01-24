package dsl

import "github.com/goadesign/gorma"

// Store represents a database.  Gorma lets you specify
// a database type, but it's currently not used for any generation
// logic.
func Store(name string, storeType gorma.RelationalStorageType, dsl func()) {
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
				AutoTimestamps:   true,
				AutoIDGenerate:   true,
				AutoSoftDelete:   true,
			}
		}
		s.RelationalStores[name] = store
	}

}

// AutomaticIDFields applies to a `Store` type.  It allows you
// to turn off the default behavior that will automatically create
// an ID/int Primary Key for each model.  If set to false,
// no ID field will be generated.
func AutomaticIDFields(auto bool) {
	if r, ok := relationalStoreDefinition(false); ok {
		r.AutoIDGenerate = auto
	}
}

// AutomaticTimestamps applies to a `Store` type.  It allows you
// to turn off the default behavior that will automatically create
// an `CreatedAt` and `UpdatedAt` fields for each model.  If set to false,
// these fields won't be created.
func AutomaticTimestamps(auto bool) {
	if r, ok := relationalStoreDefinition(false); ok {
		r.AutoTimestamps = auto
	}
}

// AutomaticSoftDelete applies to a `Store` type.  It allows
// you to turn off the default behavior that will automatically
// create a `DeletedAt` field (*time.Time) that acts as a
// soft-delete filter for your models.  If set to false,
// this field won't be created.
func AutomaticSoftDelete(auto bool) {
	if r, ok := relationalStoreDefinition(false); ok {
		r.AutoSoftDelete = auto
	}
}
