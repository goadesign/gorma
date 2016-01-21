package dsl

import (
	"fmt"
	"strings"

	"github.com/bketelsen/gorma"
	"github.com/raphael/goa/design/dsl"
)

// RelationalField is a DSL definition for a field in a Relational Model
// Examples and more docs here later
//func Field(name string, fieldType gorma.FieldType, dsl func()) {
func Field(name string, args ...interface{}) {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.

	checkInit()
	// standardize field name definitions
	name = strings.Title(name)
	fieldType, dsl := parsemModelArgs(args...)
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
			fmt.Println("AutoAdded field", name)
			field.DefinitionDSL = dsl

		}

		if fieldType == gorma.PKUUID || fieldType == gorma.PKInteger || fieldType == gorma.PKBigInteger {
			field.PrimaryKey = true
			field.Description = "primary key"
			s.PrimaryKeys = append(s.PrimaryKeys, field)
		}

		if fieldType == gorma.Timestamp {
			field.Timestamp = true
			field.Description = "timestamp"
		}
		if fieldType == gorma.NullableTimestamp {
			field.Timestamp = true
			field.Nullable = true
			field.Description = "nullable timestamp (soft delete)"
		}

		s.RelationalFields[name] = field
	}

}

func parsemModelArgs(args ...interface{}) (gorma.FieldType, func()) {

	var (
		fieldType gorma.FieldType
		dslp      func()
		ok        bool
	)

	parseFieldType := func(expected string, index int) {
		if fieldType, ok = args[index].(gorma.FieldType); !ok {
			invalidArgError(expected, args[index])
		}
	}
	parseDSL := func(index int, success, failure func()) {
		if dslp, ok = args[index].(func()); ok {
			success()
		} else {
			failure()
		}
	}

	success := func() {}

	switch len(args) {
	case 0:
		fieldType = gorma.NotFound
	case 1:
		parseDSL(0, success, func() { parseFieldType("DataType or func()", 0) })
	case 2:
		parseFieldType("FieldType", 0)
		parseDSL(1, success, func() { invalidArgError("DSL", args[1]) })

	default:
		dsl.ReportError("too many arguments in call to Attribute")
	}

	return fieldType, dslp

}

// DatabaseFieldName creates sql tag necessary to name the
// database column
func DatabaseFieldName(d string) {
	if r, ok := relationalFieldDefinition(false); ok {
		r.DatabaseFieldName = d
	}
}
