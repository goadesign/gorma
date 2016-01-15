package gorma

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/bketelsen/gorma/gengorma"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

// NewRelationalModel instantiates and populates a new relational model structure
func NewRelationalModel(name string, t *design.UserTypeDefinition) (*RelationalModel, error) {
	var pks []*RelationalField
	rm := &RelationalModel{
		utd: t,
		//TableName:   codegen.Goify(deModel(name), false), // may be overridden later
		Name:        codegen.Goify(deModel(name), true),
		Fields:      make(map[string]*RelationalField),
		HasMany:     make(map[string]*RelationalModel),
		HasOne:      make(map[string]*RelationalModel),
		ManyToMany:  make(map[string]*ManyToMany),
		BelongsTo:   make(map[string]*RelationalModel),
		PrimaryKeys: pks,
	}
	err := rm.Parse()
	return rm, err
}
func (f *RelationalModel) PKAttributes() string {
	var attr []string
	for _, pk := range f.PrimaryKeys {
		attr = append(attr, fmt.Sprintf("%s %s", strings.ToLower(pk.Name), pk.Datatype))
	}
	return strings.Join(attr, ",")
}

func (f *RelationalModel) PKWhere() string {
	var pkwhere []string
	for _, pk := range f.PrimaryKeys {
		def := fmt.Sprintf("%s = ?", pk.DatabaseFieldName)
		pkwhere = append(pkwhere, def)
	}
	return strings.Join(pkwhere, "and")
}
func (f *RelationalModel) PKWhereFields() string {
	var pkwhere []string
	for _, pk := range f.PrimaryKeys {
		def := fmt.Sprintf("%s", pk.DatabaseFieldName)
		pkwhere = append(pkwhere, def)
	}
	return strings.Join(pkwhere, ",")
}
func (f *RelationalModel) PKUpdateFields() string {

	var pkwhere []string
	for _, pk := range f.PrimaryKeys {
		def := fmt.Sprintf("model.%s", codegen.Goify(pk.Name, true))
		pkwhere = append(pkwhere, def)
	}

	pkw := strings.Join(pkwhere, ",")
	return pkw
}

func (rm *RelationalModel) Definition() string {
	header := fmt.Sprintf("type %s struct {\n", rm.Name)
	var output string
	rm.IterateFields(func(f *RelationalField) error {
		output = output + f.Definition()
		return nil
	})
	footer := "}\n"
	return header + output + footer

}

// Parse populates the RelationalModel based on the defintions in the DSL
func (rm *RelationalModel) Parse() error {
	err := rm.ParseOptions()
	if err != nil {
		return err
	}

	err = rm.ParseFields()
	if err != nil {
		return err
	}
	return nil

}

func (rm *RelationalModel) ParseFields() error {

	var ds design.DataStructure
	ds = rm.utd
	def := ds.Definition()
	t := def.Type
	switch actual := t.(type) {
	case design.Object:
		keys := make([]string, len(actual))
		i := 0
		for n := range actual {
			keys[i] = n
			i++
		}
		sort.Strings(keys)
		for _, name := range keys {
			field, err := NewRelationalField(name, actual[name])
			if err != nil {
				return err
			}
			if rm.utd.IsRequired(name) {
				field.Nullable = false
			}
			rm.Fields[name] = field
			if field.PrimaryKey {
				rm.PrimaryKeys = append(rm.PrimaryKeys, field)
			}
			if field.BelongsTo != "" {
				rm.belongsto = append(rm.belongsto, field.BelongsTo)
			}
			if field.HasOne != "" {
				rm.hasone = append(rm.hasone, field.HasOne)
			}
			if field.HasMany != "" {
				rm.hasmany = append(rm.hasmany, field.HasMany)
			}
			if field.Many2Many != "" {
				rm.many2many = append(rm.many2many, field.Many2Many)
			}
		}

	default:
		return errors.New("Unexpected type")
	}

	return nil
}

func (rm *RelationalModel) ParseOptions() error {

	def := rm.utd.Definition()
	t := def.Type
	switch t.(type) {
	case design.Object:
		if val, ok := metaLookup(rm.utd.Metadata, gengorma.MetaCached); ok {
			rm.Cached = true
			duration, err := strconv.Atoi(val)
			if err != nil {
				return errors.New("Cache duration must be a string that can be parsed as an integer")
			}
			rm.CacheDuration = duration
		}
		if val, ok := metaLookup(rm.utd.Metadata, gengorma.MetaSQLTag); ok {
			rm.SQLTag = val
		}
		if _, ok := metaLookup(rm.utd.Metadata, gengorma.MetaDynamicTableName); ok {
			rm.DynamicTableName = true
		}
		if _, ok := metaLookup(rm.utd.Metadata, gengorma.MetaRoler); ok {
			rm.Roler = true
		}
		if _, ok := metaLookup(rm.utd.Metadata, gengorma.MetaNoMedia); ok {
			rm.NoMedia = true
		}
		if val, ok := metaLookup(rm.utd.Metadata, gengorma.MetaTableName); ok {
			rm.TableName = val
		}
		if val, ok := metaLookup(rm.utd.Metadata, gengorma.MetaGormTag); ok {
			rm.Alias = val
		}

		return nil
	default:
		return errors.New("gorma bug: unexpected data structure type")
	}
}

func (rm *RelationalModel) IterateFields(it FieldIterator) error {

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
