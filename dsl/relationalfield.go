package dsl

import "github.com/bketelsen/gorma"

// RelationalField is a DSL definition for a field in a Relational Model
// Examples and more docs here later
func RelationalField(name string, fieldType gorma.FieldType, dsl func()) {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.
	checkInit()
	if s, ok := relationalModelDefinition(true); ok {
		if s.RelationalFields == nil {
			s.RelationalFields = make(map[string]*gorma.RelationalFieldDefinition)
		}
		field, ok := s.RelationalFields[name]
		if !ok {
			field = &gorma.RelationalFieldDefinition{
				Name:          name,
				DefinitionDSL: dsl,
				Parent:        s,
				Datatype:      fieldType,
			}
		} else {
			// the field was auto-added by the model parser
			// so we need to update whatever we can from this new definition
			field.DefinitionDSL = dsl
			field.Datatype = fieldType
		}

		if fieldType == gorma.PKUUID || fieldType == gorma.PKInteger || fieldType == gorma.PKBigInteger {
			field.PrimaryKey = true
		}

		if fieldType == gorma.Timestamp {
			field.Timestamp = true
		}
		if fieldType == gorma.NullableTimestamp {
			field.Timestamp = true
			field.Nullable = true
		}

		s.RelationalFields[name] = field
	}

}

// DatabaseFieldName creates sql tag necessary to name the
// database column
func DatabaseFieldName(d string) {
	if r, ok := relationalFieldDefinition(false); ok {
		r.DatabaseFieldName = d
	}
}
