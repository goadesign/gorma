package dsl

import (
	"strings"

	"bitbucket.org/pkg/inflect"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/dslengine"
	"github.com/goadesign/goa/goagen/codegen"
	"github.com/goadesign/gorma"
)

// DatabaseFieldName allows for customization of the column name
// by seting the struct tags. This is necessary to create correlate
// non standard column naming conventions in gorm.
func DatabaseFieldName(name string) {
	if f, ok := relationalFieldDefinition(true); ok {
		f.DatabaseFieldName = name
	}
}

// Field is a DSL definition for a field in a Relational Model.
// Parameter Options:
//
//	// A field called "Title" with default type `String`.
//	Field("Title")
//
//	// Explicitly specify the type.
//	Field("Title", gorma.String)
//
//	// "Title" field, as `String` with other DSL included.
//	Field("Title", func(){... other field level dsl ...})
//
//	// All options specified: name, type and dsl.
//	Field("Title", gorma.String, func(){... other field level dsl ...})
func Field(name string, args ...interface{}) {
	name = codegen.Goify(name, true)
	name = SanitizeFieldName(name)
	fieldType, dsl := parseFieldArgs(args...)
	if s, ok := relationalModelDefinition(true); ok {
		if s.RelationalFields == nil {
			s.RelationalFields = make(map[string]*gorma.RelationalFieldDefinition)
		}
		field, ok := s.RelationalFields[name]
		if !ok {
			field = gorma.NewRelationalFieldDefinition()
			field.FieldName = name
			field.DefinitionDSL = dsl
			field.Parent = s
			field.Datatype = fieldType
		} else {
			// the field was auto-added by the model parser
			// so we need to update whatever we can from this new definition
			field.DefinitionDSL = dsl
			field.Datatype = fieldType

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
		field.DatabaseFieldName = SanitizeDBFieldName(name)

		s.RelationalFields[name] = field
	}
}

// MapsFrom establishes a mapping relationship between a source
// Type field and this model.  The source type must be a UserTypeDefinition "Type"
// in goa.  These are typically Payloads.
func MapsFrom(utd *design.UserTypeDefinition, field string) {
	if f, ok := relationalFieldDefinition(true); ok {
		md := gorma.NewMapDefinition()
		md.RemoteField = field
		md.RemoteType = utd
		f.Mappings[utd.TypeName] = md

	}
}

// MapsTo establishes a mapping relationship between a field in model and
// a MediaType in goa.
func MapsTo(mtd *design.MediaTypeDefinition, field string) {
	if f, ok := relationalFieldDefinition(true); ok {
		md := gorma.NewMapDefinition()
		md.RemoteField = field
		md.RemoteType = mtd.UserTypeDefinition
		f.Mappings[mtd.UserTypeDefinition.TypeName] = md
	}
}

func fixID(s string) string {
	if s == "i_d" {
		return "id"
	}
	return s

}

// Nullable sets the fields nullability
// A Nullable field will be stored as a pointer.  A field that is
// not Nullable won't be stored as a pointer.
func Nullable() {
	if f, ok := relationalFieldDefinition(false); ok {
		f.Nullable = true
	}
}

// PrimaryKey establishes a field as a Primary Key by
// seting the struct tags necessary to create the PK in gorm.
// Valid only for `Integer` datatypes currently
func PrimaryKey() {
	if f, ok := relationalFieldDefinition(true); ok {
		if f.Datatype != gorma.Integer && f.Datatype != gorma.UUID {
			dslengine.ReportError("Integer and UUID are the only supported Primary Key field types.")
		}

		f.PrimaryKey = true
		f.Nullable = false
		f.Description = "primary key"
		f.Parent.PrimaryKeys = append(f.Parent.PrimaryKeys, f)
	}
}

// SanitizeFieldName is exported for testing purposes
func SanitizeFieldName(name string) string {
	name = codegen.Goify(name, true)
	if strings.HasSuffix(name, "Id") {
		name = strings.TrimSuffix(name, "Id")
		name = name + "ID"
	}

	return name
}

// SanitizeDBFieldName is exported for testing purposes
func SanitizeDBFieldName(name string) string {
	name = inflect.Underscore(name)
	if strings.HasSuffix(name, "_i_d") {
		name = strings.TrimSuffix(name, "_i_d")
		name = name + "_id"
	}
	if name == "i_d" {
		name = "id"

	}
	return name
}
func parseFieldArgs(args ...interface{}) (gorma.FieldType, func()) {
	var (
		fieldType gorma.FieldType
		dslp      func()
		ok        bool
	)

	parseFieldType := func(expected string, index int) {
		if fieldType, ok = args[index].(gorma.FieldType); !ok {
			dslengine.InvalidArgError(expected, args[index])
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
		parseDSL(1, success, func() { dslengine.InvalidArgError("DSL", args[1]) })

	default:
		dslengine.ReportError("too many arguments in call to Attribute")
	}

	return fieldType, dslp
}
