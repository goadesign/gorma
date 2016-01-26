package gorma

import (
	"fmt"
	"sort"
	"strings"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/dsl"
	"github.com/goadesign/goa/goagen/codegen"
)

// Context returns the generic definition name used in error messages.
func (f *RelationalModelDefinition) Context() string {
	if f.Name != "" {
		return fmt.Sprintf("RelationalModel %#v", f.Name)
	}
	return "unnamed RelationalModel"
}

// DSL returns this object's DSL.
func (f *RelationalModelDefinition) DSL() func() {
	return f.DefinitionDSL
}

// Children returns a slice of this objects children.
func (f RelationalModelDefinition) Children() []design.Definition {
	var stores []design.Definition
	for _, s := range f.RelationalFields {
		stores = append(stores, s)
	}
	return stores
}

// PKAttributes constructs a pair of field + definition strings
// useful for method parameters.
func (f *RelationalModelDefinition) PKAttributes() string {
	var attr []string
	for _, pk := range f.PrimaryKeys {
		attr = append(attr, fmt.Sprintf("%s %s", pk.DatabaseFieldName, goDatatype(pk)))
	}
	return strings.Join(attr, ",")
}

// PKWhere returns an array of strings representing the where clause
// of a retrieval by primary key(s) -- x = ? and y = ?.
func (f *RelationalModelDefinition) PKWhere() string {
	var pkwhere []string
	for _, pk := range f.PrimaryKeys {
		def := fmt.Sprintf("%s = ?", pk.DatabaseFieldName)
		pkwhere = append(pkwhere, def)
	}
	return strings.Join(pkwhere, " and ")
}

// PKWhereFields returns the fields for a where clause for the primary
// keys of a model.
func (f *RelationalModelDefinition) PKWhereFields() string {
	var pkwhere []string
	for _, pk := range f.PrimaryKeys {
		def := fmt.Sprintf("%s", pk.DatabaseFieldName)
		pkwhere = append(pkwhere, def)
	}
	return strings.Join(pkwhere, ",")
}

// PKUpdateFields returns something?  This function doesn't look useful in
// current form.  Perhaps it isn't.
func (f *RelationalModelDefinition) PKUpdateFields(modelname string) string {
	var pkwhere []string
	for _, pk := range f.PrimaryKeys {
		def := fmt.Sprintf("%s.%s", modelname, codegen.Goify(pk.Name, true))
		pkwhere = append(pkwhere, def)
	}

	pkw := strings.Join(pkwhere, ",")
	return pkw
}

// StructDefinition returns the struct definition for the model.
func (f *RelationalModelDefinition) StructDefinition() string {
	header := fmt.Sprintf("type %s struct {\n", f.Name)
	var output string
	f.IterateFields(func(field *RelationalFieldDefinition) error {
		output = output + field.FieldDefinition()
		return nil
	})
	footer := "}\n"
	return header + output + footer

}

// LowerName returns the model name as a lowercase string.
func (f *RelationalModelDefinition) LowerName() string {
	return strings.ToLower(f.Name)
}

// IterateFields returns an iterator function useful for iterating through
// this model's field list.
func (f *RelationalModelDefinition) IterateFields(it FieldIterator) error {
	// Break out each type of fields

	pks := make(map[string]string)
	for n := range f.RelationalFields {
		if f.RelationalFields[n].PrimaryKey {
			//	names[i] = n
			pks[n] = n
		}
	}

	names := make(map[string]string)
	for n := range f.RelationalFields {
		if !f.RelationalFields[n].PrimaryKey && !f.RelationalFields[n].Timestamp {
			names[n] = n
		}
	}

	dates := make(map[string]string)
	for n := range f.RelationalFields {
		if f.RelationalFields[n].Timestamp {
			dates[n] = n
		}
	}

	// Sort only the fields that aren't pk or date
	j := 0
	sortnames := make([]string, len(names))
	for n := range names {
		sortnames[j] = n
		j++
	}
	sort.Strings(sortnames)

	// Put them back together
	j = 0
	i := len(pks) + len(names) + len(dates)
	fields := make([]string, i)
	for _, pk := range pks {
		fields[j] = pk
		j++
	}
	for _, name := range sortnames {
		fields[j] = name
		j++
	}
	for _, date := range dates {
		fields[j] = date
		j++
	}

	// Iterate them
	for _, n := range fields {
		if err := it(f.RelationalFields[n]); err != nil {
			return err
		}
	}
	return nil
}

// PopulateFromModeledType creates fields for the model
// based on the goa UserTypeDefinition it models.
// This happens before fields are processed, so it's
// ok to just assign without testing.
func (f *RelationalModelDefinition) PopulateFromModeledType() {
	if f.BuiltFrom == nil {
		return
	}
	for _, mt := range f.BuiltFrom {
		obj := mt.ToObject()
		obj.IterateAttributes(func(name string, att *design.AttributeDefinition) error {
			rf := &RelationalFieldDefinition{}
			rf.Parent = f
			rf.Name = codegen.Goify(name, true)
			if strings.HasSuffix(rf.Name, "Id") {
				rf.Name = strings.TrimSuffix(rf.Name, "Id")
				rf.Name = rf.Name + "ID"
			}
			switch att.Type.Kind() {
			case design.BooleanKind:
				rf.Datatype = Boolean
			case design.IntegerKind:
				rf.Datatype = Integer
			case design.NumberKind:
				rf.Datatype = Decimal
			case design.StringKind:
				rf.Datatype = String
			case design.DateTimeKind:
				rf.Datatype = Timestamp
			default:
				dsl.ReportError("Unsupported type: %#v ", att.Type.Kind())
			}
			f.RelationalFields[rf.Name] = rf
			return nil
		})
	}
	return
}
