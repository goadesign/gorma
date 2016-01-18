package gorma

import (
	"fmt"
	"sort"
	"strings"

	"github.com/raphael/goa/goagen/codegen"
)

// PKAttributes constructs a pair of field + definition strings
// useful for method parameters
func (f *RelationalModelDefinition) PKAttributes() string {
	var attr []string
	for _, pk := range f.PrimaryKeys {
		attr = append(attr, fmt.Sprintf("%s %s", strings.ToLower(pk.Name), pk.Datatype))
	}
	return strings.Join(attr, ",")
}

// PKWhere returns an array of strings representing the where clause
// of a retrieval by primary key(s) -- x = ? and y = ?
func (f *RelationalModelDefinition) PKWhere() string {
	var pkwhere []string
	for _, pk := range f.PrimaryKeys {
		def := fmt.Sprintf("%s = ?", pk.DatabaseFieldName)
		pkwhere = append(pkwhere, def)
	}
	return strings.Join(pkwhere, "and")
}
func (f *RelationalModelDefinition) PKWhereFields() string {
	var pkwhere []string
	for _, pk := range f.PrimaryKeys {
		def := fmt.Sprintf("%s", pk.DatabaseFieldName)
		pkwhere = append(pkwhere, def)
	}
	return strings.Join(pkwhere, ",")
}

// PKUpdateFields returns something?  This function doesn't look useful in
// current form.  Perhaps it isnt.
func (f *RelationalModelDefinition) PKUpdateFields() string {

	var pkwhere []string
	for _, pk := range f.PrimaryKeys {
		def := fmt.Sprintf("model.%s", codegen.Goify(pk.Name, true))
		pkwhere = append(pkwhere, def)
	}

	pkw := strings.Join(pkwhere, ",")
	return pkw
}

func (rm *RelationalModelDefinition) Definition() string {
	header := fmt.Sprintf("type %s struct {\n", rm.Name)
	var output string
	rm.IterateFields(func(f *RelationalFieldDefinition) error {
		output = output + f.Definition()
		return nil
	})
	footer := "}\n"
	return header + output + footer

}

// IterateFields returns an iterator function useful for iterating through
// this model's field list
func (rm *RelationalModelDefinition) IterateFields(it FieldIterator) error {

	names := make(map[string]string)
	pks := make(map[string]string)
	dates := make(map[string]string)

	// Break out each type of fields
	for n := range rm.Fields {
		if rm.Fields[n].PrimaryKey {
			//	names[i] = n
			pks[n] = n
		}
	}
	for n := range rm.Fields {
		if !rm.Fields[n].PrimaryKey && !rm.Fields[n].Timestamp {
			names[n] = n
		}
	}
	for n := range rm.Fields {
		if rm.Fields[n].Timestamp {
			//	names[i] = n
			dates[n] = n
		}
	}

	// Sort only the fields that aren't pk or date
	j := 0
	sortfields := make([]string, len(names))
	for n := range names {
		sortfields[j] = n
	}
	sort.Strings(sortfields)

	// Put them back together
	j = 0
	i := len(pks) + len(names) + len(dates)
	fields := make([]string, i)
	for _, pk := range pks {
		fields[j] = pk
		j++
	}
	for _, name := range names {
		fields[j] = name
		j++
	}
	for _, date := range dates {
		fields[j] = date
		j++
	}

	// Iterate them
	for _, n := range fields {
		if err := it(rm.Fields[n]); err != nil {
			return err
		}
	}
	return nil
}
