package gorma

import (
	"fmt"
	"strings"

	"github.com/bketelsen/gorma/gengorma"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

// NewRelationalField creates and parses a field from the design attributes
func NewRelationalField(name string, a *design.AttributeDefinition) (*RelationalField, error) {
	f := &RelationalField{}
	f.a = a
	t := a.Definition().Type
	f.Datatype = codegen.GoNativeType(t)
	f.Name = codegen.Goify(name, true)
	if strings.HasSuffix(f.Name, "Id") {
		f.Name = f.Name[:len(f.Name)-2] + "ID"
	}
	f.DatabaseFieldName = f.Name
	err := f.Parse()
	return f, err

}

// Parse populates all the attributes of the Field
func (f *RelationalField) Parse() error {
	err := f.ParsePrimaryKey()
	if err != nil {
		return err
	}
	err = f.ParseSQLTag()
	if err != nil {
		return err
	}
	err = f.ParseTimestamps()
	if err != nil {
		return err
	}
	err = f.ParseAlias()
	if err != nil {
		return err
	}
	err = f.ParseBelongsTo()
	if err != nil {
		return err
	}
	err = f.ParseHasOne()
	if err != nil {
		return err
	}
	err = f.ParseHasMany()
	if err != nil {
		return err
	}
	return err
}

func (f *RelationalField) Definition() string {

	var fieldType, fieldName, pointer string
	fieldType = codegen.GoTypeName(f.a.Definition().Type, []string{}, 1)
	if f.HasOne != "" {
		fieldType = f.HasOne
	}

	fieldName = f.Name
	if f.Nullable {
		pointer = "*"
	}
	return fmt.Sprintf("%s \t %s%s %s", fieldName, pointer, fieldType, f.Tags())

}

func (f *RelationalField) Tags() string {
	var sqltags, gormtags, jsontags string
	if f.SQLTag != "" {
		sqltags = "sql:\"f.SQLTag\""
	}

	if f.PrimaryKey {
		gormtags = "primary_key"
	}
	return fmt.Sprintf("`json: %s sql: %s gorm: %s`", jsontags, sqltags, gormtags)
}

//ParseTimestamps populates the timestamps field
func (f *RelationalField) ParseTimestamps() error {
	if _, ok := metaLookup(f.a.Metadata, gengorma.MetaTimestampCreated); ok {
		f.Timestamp = true
		f.Datatype = "time.Time"
	}
	if _, ok := metaLookup(f.a.Metadata, gengorma.MetaTimestampUpdated); ok {
		f.Timestamp = true
		f.Datatype = "time.Time"
	}
	if _, ok := metaLookup(f.a.Metadata, gengorma.MetaTimestampDeleted); ok {
		f.Timestamp = true
		f.Nullable = true
		f.Datatype = "*time.Time"
	}

	return nil

}

//ParseSQLTag populates the SQLTag field
func (f *RelationalField) ParseSQLTag() error {
	// is it a primary key?
	if gt, ok := metaLookup(f.a.Metadata, gengorma.MetaSQLTag); ok {
		f.SQLTag = gt
	}
	return nil

}

//ParseBelongsTo populates the SQLTag field
func (f *RelationalField) ParseBelongsTo() error {
	if gt, ok := metaLookup(f.a.Metadata, gengorma.MetaBelongsTo); ok {
		f.BelongsTo = gt
	}
	return nil

}

//ParseHasOne populates the SQLTag field
func (f *RelationalField) ParseHasOne() error {
	if gt, ok := metaLookup(f.a.Metadata, gengorma.MetaHasOne); ok {
		f.HasOne = gt
	}
	return nil

}

//ParseHasMany populates the SQLTag field
func (f *RelationalField) ParseHasMany() error {
	if gt, ok := metaLookup(f.a.Metadata, gengorma.MetaHasMany); ok {
		f.HasMany = gt
	}
	return nil

}

//ParseAlias populates the DatabaseFieldName field
func (f *RelationalField) ParseAlias() error {
	if gt, ok := metaLookup(f.a.Metadata, gengorma.MetaGormTag); ok {
		f.DatabaseFieldName = gt
	}
	return nil

}

//ParsePrimaryKey populates the primary key tag
func (f *RelationalField) ParsePrimaryKey() error {
	// is it a primary key?
	if gt, ok := metaLookup(f.a.Metadata, gengorma.MetaPrimaryKey); ok {
		if strings.Contains(gt, "primary_key") {
			f.PrimaryKey = true
		}
	}
	if f.Name == "ID" || f.Name == "Id" || f.Name == "id" {
		f.PrimaryKey = true
	}
	return nil
}
