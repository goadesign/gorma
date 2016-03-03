package dsl

import (
	"github.com/goadesign/goa/dslengine"
	"github.com/goadesign/gorma"
)

// Store represents a database.  Gorma lets you specify
// a database type, but it's currently not used for any generation
// logic.
func Store(name string, storeType gorma.RelationalStorageType, dsl func()) {
	if name == "" || len(name) == 0 {
		dslengine.ReportError("Relational Store requires a name.")
		return
	}
	if len(storeType) == 0 {
		dslengine.ReportError("Relational Store requires a RelationalStoreType.")
		return
	}
	if dsl == nil {
		dslengine.ReportError("Relational Store requires a dsl.")
		return
	}
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
		} else {
			dslengine.ReportError("Relational Store %s can only be declared once.", name)
		}
		s.RelationalStores[name] = store
	}

}

// NoAutomaticIDFields applies to a `Store` or `Model` type.  It allows you
// to turn off the default behavior that will automatically create
// an ID/int Primary Key for each model.
func NoAutomaticIDFields() {
	if s, ok := relationalStoreDefinition(false); ok {
		s.NoAutoIDFields = true
	} else if m, ok := relationalModelDefinition(false); ok {
		delete(m.RelationalFields, "ID")
	}
}

// NoAutomaticTimestamps applies to a `Store` or `Model` type.  It allows you
// to turn off the default behavior that will automatically create
// an `CreatedAt` and `UpdatedAt` fields for each model.
func NoAutomaticTimestamps() {
	if s, ok := relationalStoreDefinition(false); ok {
		s.NoAutoTimestamps = true
	} else if m, ok := relationalModelDefinition(false); ok {
		delete(m.RelationalFields, "CreatedAt")
		delete(m.RelationalFields, "UpdatedAt")
	}
}

// NoAutomaticSoftDelete applies to a `Store` or `Model` type.  It allows
// you to turn off the default behavior that will automatically
// create a `DeletedAt` field (*time.Time) that acts as a
// soft-delete filter for your models.
func NoAutomaticSoftDelete() {
	if s, ok := relationalStoreDefinition(false); ok {
		s.NoAutoSoftDelete = true
	} else if m, ok := relationalModelDefinition(false); ok {
		delete(m.RelationalFields, "DeletedAt")
	}
}
