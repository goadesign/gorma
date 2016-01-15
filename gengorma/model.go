package gengorma

import (
	"fmt"
	"strings"

	"bitbucket.org/pkg/inflect"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/design/dsl"
)

const (
	MetaHasOne           = "github.com/bketelsen/gorma#hasone"
	MetaHasMany          = "github.com/bketelsen/gorma#hasmany"
	MetaBelongsTo        = "github.com/bketelsen/gorma#belongsto"
	MetaCached           = "github.com/bketelsen/gorma#cached"
	MetaPrimaryKey       = "github.com/bketelsen/gorma#primarykey"
	MetaManyToMany       = "github.com/bketelsen/gorma#manytomany"
	MetaDynamicTableName = "github.com/bketelsen/gorma#dyntablename"
	MetaRoler            = "github.com/bketelsen/gorma#roler"
	MetaGormTag          = "github.com/bketelsen/gorma#gormtag"
	MetaNoMedia          = "github.com/bketelsen/gorma#nomedia"
	MetaSQLTag           = "github.com/bketelsen/gorma#sqltag"
	MetaTableName        = "github.com/bketelsen/gorma#tablename"
	MetaTimestampCreated = "github.com/bketelsen/gorma#created"
	MetaTimestampUpdated = "github.com/bketelsen/gorma#updated"
	MetaTimestampDeleted = "github.com/bketelsen/gorma#deleted"
)

// Model is the function that makes models happen.  Witty comment here
// This function returns the model definition so it can be referred to throughout the DSL.
func Model(name string, dsla func()) *design.UserTypeDefinition {
	if design.Design == nil {
		dsl.InitDesign()
	}
	if design.Design.Types == nil {
		design.Design.Types = make(map[string]*design.UserTypeDefinition)
	} else if _, ok := design.Design.Types[name]; ok {
		dsl.ReportError("model %#v defined twice", name)
		return nil
	}
	var t *design.UserTypeDefinition
	// This is a Model, so annotate it with the correct metadata on initialization
	meta := make(map[string]string)
	meta["github.com/bketelsen/gorma"] = "Model"

	t = &design.UserTypeDefinition{
		TypeName:            name,
		AttributeDefinition: &design.AttributeDefinition{Metadata: meta},
		DSL:                 dsla,
	}
	if dsla == nil {
		t.Type = design.String
	}
	design.Design.Types[name] = t

	return t
}

// DynamicTableName annotates the model with the correct
// metadata for a dynamic tablename
func DynamicTableName() {
	dsl.Metadata(MetaDynamicTableName, "true")
}

// NoMedia annotates the model with the correct
// metadata to skip media definition in the models
func NoMedia() {
	dsl.Metadata(MetaNoMedia, "true")
}

// As annotates the field with the correct
// metadata for a custom column name in the database
func As(alias string) {
	dsl.Metadata(MetaGormTag, alias)
}

// TableName annotates the model with the correct
// metadata for a custom table name
func TableName(name string) {
	dsl.Metadata(MetaTableName, name)
}

// SQLTag annotates the model with the correct
// metadata for a custom parameters to the SQL Storage engine
func SQLTag(tag string) {
	dsl.Metadata(MetaSQLTag, tag)
}

// HasOne annotates the model with the correct
// metadata for a HasOne relationship
func HasOne(model *design.UserTypeDefinition) {
	dsl.Attribute(deModel(model.TypeName),
		func() { dsl.Metadata(MetaHasOne, model.TypeName) })
}

// Timestamps creates the created_at and  updated_at
// fields
func Timestamps() {
	dsl.Attribute("created_at", design.Any,
		func() { dsl.Metadata(MetaTimestampCreated, "true") })
	dsl.Attribute("updated_at", design.Any,
		func() { dsl.Metadata(MetaTimestampUpdated, "true") })

}

// SoftDelete creates the deleted_at
// field, and informs gorm that it should use soft-deletes
func SoftDelete() {
	dsl.Attribute("deleted_at", design.Any,
		func() { dsl.Metadata(MetaTimestampDeleted, "true") })
}

// HasMany annotates the model with the correct
// metadata for a HasMany relationship
func HasMany(model string, rel *design.UserTypeDefinition) {
	dsl.Attribute(strings.ToLower(inflect.Pluralize(model)), design.String,
		func() { dsl.Metadata(MetaHasMany, rel.TypeName) })
}

//BelongsTo annotates the model with the correct
// metadata for a BelongsTO relationship
func BelongsTo(model string) {

	modelName := fmt.Sprintf("%s_id", model)
	dsl.Attribute(modelName, design.Integer,
		func() { dsl.Metadata(MetaBelongsTo, strings.Title(model)) })
}

// Cached annotates the model with the correct
// metadata to make individual records cached
func Cached(seconds string) {
	dsl.Metadata(MetaCached, seconds)
}

// Roler annotates the model with the correct
// metadata to create a Role() function and Role field
func Roler() {
	dsl.Attribute("role", design.String)
	dsl.Required("role")
	dsl.Metadata(MetaRoler, "true")
}

// PrimaryKey annotates the model with the correct
// gorm tag to make the model a primary key in the database
func PrimaryKey(field string) {
	dsl.Required(field)
	dsl.Attribute(field, design.Integer,
		func() { dsl.Metadata(MetaPrimaryKey, "primary_key") })
}

func ManyToMany(relation, tablename string, rel *design.UserTypeDefinition) {
	val := fmt.Sprintf("%s:%s:%s", relation, tablename, rel.TypeName)

	dsl.Attribute(inflect.Pluralize(relation), design.String,
		func() { dsl.Metadata(MetaManyToMany, val) })

}

// Temp - refactor later?

// deModel removes the word "Model" from the string.
func deModel(s string) string {
	return strings.Replace(s, "Model", "", -1)
}
