package dsl

import "github.com/bketelsen/gorma"

// RelationalField is a DSL definition for a field in a Relational Model
// Examples and more docs here later
func RelationalField(name string, dsl func()) {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.
	checkInit()
	if s, ok := relationalModelDefinition(true); ok {
		if s.RelationalFields == nil {
			s.RelationalFields = make(map[string]*gorma.RelationalFieldDefinition)
		}
		fields, ok := s.RelationalFields[name]
		if !ok {
			fields = &gorma.RelationalFieldDefinition{
				Name:          name,
				DefinitionDSL: dsl,
			}
		}

		s.RelationalFields[name] = fields
	}

}
