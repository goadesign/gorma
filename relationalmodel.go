package gorma

import (
	"errors"
	"sort"
	"strconv"

	"github.com/bketelsen/gorma/gengorma"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

// NewRelationalModel instantiates and populates a new relational model structure
func NewRelationalModel(name string, t *design.UserTypeDefinition) (*RelationalModel, error) {
	rm := &RelationalModel{
		utd:         t,
		TableName:   deModel(name), // may be overridden later
		Name:        codegen.Goify(name, true),
		Fields:      make(map[string]*RelationalField),
		HasMany:     make(map[string]*RelationalModel),
		HasOne:      make(map[string]*RelationalModel),
		ManyToMany:  make(map[string]*ManyToMany),
		BelongsTo:   make(map[string]*RelationalModel),
		PrimaryKeys: make([]*RelationalField, 1),
	}
	err := rm.Parse()
	return rm, err
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
