package dsl

import (
	"strconv"
	"strings"

	"bitbucket.org/pkg/inflect"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/dsl"
	"github.com/goadesign/goa/goagen/codegen"
	"github.com/goadesign/gorma"
)

// Model is the DSL that represents a Relational Model
// Model name should be Title cased. Use RenderTo() and BuiltFrom()
// to have Gorma generate conversion helpers for your model.  Gorma
// will create appropriate fields for all of your database relationships
// too, using the BelongsTo(), HasMany(), HasOne(), and ManyToMany() DSL
func Model(name string, dsl func()) {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.
	checkInit()
	if s, ok := relationalStoreDefinition(true); ok {
		if s.RelationalModels == nil {
			s.RelationalModels = make(map[string]*gorma.RelationalModelDefinition)
		}
		models, ok := s.RelationalModels[name]
		if !ok {
			models = &gorma.RelationalModelDefinition{
				Name:             name,
				DefinitionDSL:    dsl,
				Parent:           s,
				RelationalFields: make(map[string]*gorma.RelationalFieldDefinition),
				BelongsTo:        make(map[string]*gorma.RelationalModelDefinition),
				HasMany:          make(map[string]*gorma.RelationalModelDefinition),
				HasOne:           make(map[string]*gorma.RelationalModelDefinition),
				ManyToMany:       make(map[string]*gorma.ManyToManyDefinition),
			}
		} else {
			models.DefinitionDSL = dsl
		}
		///models.PopulateFromModeledType() -- need to do this later
		s.RelationalModels[name] = models
		if s.AutoIDGenerate {
			field := &gorma.RelationalFieldDefinition{
				Name:              SanitizeFieldName("ID"),
				Parent:            models,
				Datatype:          gorma.PKInteger,
				PrimaryKey:        true,
				DatabaseFieldName: SanitizeDBFieldName("ID"),
			}
			models.RelationalFields[field.Name] = field
		}
		if s.AutoTimestamps {
			// add createdat
			field := &gorma.RelationalFieldDefinition{
				Name:              SanitizeFieldName("CreatedAt"),
				Parent:            models,
				Datatype:          gorma.Timestamp,
				DatabaseFieldName: SanitizeDBFieldName("CreatedAt"),
			}
			models.RelationalFields[field.Name] = field
			// add updatedat
			field = &gorma.RelationalFieldDefinition{
				Name:              SanitizeFieldName("UpdatedAt"),
				Parent:            models,
				Datatype:          gorma.Timestamp,
				DatabaseFieldName: SanitizeDBFieldName("UpdatedAt"),
			}
			models.RelationalFields[field.Name] = field
		}
		if s.AutoSoftDelete {
			// Add softdelete
			field := &gorma.RelationalFieldDefinition{
				Name:              SanitizeFieldName("DeletedAt"),
				Parent:            models,
				Nullable:          true,
				Datatype:          gorma.NullableTimestamp,
				DatabaseFieldName: SanitizeDBFieldName("DeletedAt"),
			}
			models.RelationalFields[field.Name] = field
		}
	}
}

// RenderTo informs Gorma that this model will need to be
// rendered to a Goa type.  Conversion functions
// will be generated to convert to/from the model.
// Usage:   RenderTo(SomeGoaMediaType)
func RenderTo(mts ...*design.MediaTypeDefinition) {
	checkInit()
	if s, ok := relationalModelDefinition(true); ok {
		if s.RenderTo == nil {
			s.RenderTo = []*design.MediaTypeDefinition{}
		}

		for _, mt := range mts {
			s.RenderTo = append(s.RenderTo, mt)
		}
	}
}

// BuiltFrom informs Gorma that this model will be populated
// from a Goa payload (User Type).  Conversion functions
// will be generated to convert from the payload to the model.
// Usage:  BuiltFrom(SomeGoaPayload)
func BuiltFrom(mts ...*design.UserTypeDefinition) {
	checkInit()
	if s, ok := relationalModelDefinition(true); ok {
		if s.BuiltFrom == nil {
			s.BuiltFrom = []*design.UserTypeDefinition{}
		}
		for _, mt := range mts {
			s.BuiltFrom = append(s.BuiltFrom, mt)
		}
		s.PopulateFromModeledType()
	}
}

// BelongsTo signifies a relationship between this model and a
// Parent.  The Parent has the child, and the Child belongs
// to the Parent
// Usage:  BelongsTo("User")
func BelongsTo(parent string) {
	if r, ok := relationalModelDefinition(false); ok {
		idfield := &gorma.RelationalFieldDefinition{
			Name:              codegen.Goify(inflect.Singularize(parent), true) + "ID",
			Description:       "Belongs To " + codegen.Goify(inflect.Singularize(parent), true),
			Parent:            r,
			Datatype:          gorma.BelongsTo,
			DatabaseFieldName: SanitizeDBFieldName(codegen.Goify(inflect.Singularize(parent), true) + "ID"),
		}

		r.RelationalFields[idfield.Name] = idfield
		bt, ok := r.Parent.RelationalModels[codegen.Goify(inflect.Singularize(parent), true)]
		if ok {
			r.BelongsTo[bt.Name] = bt
		} else {
			models := &gorma.RelationalModelDefinition{
				Name:             codegen.Goify(inflect.Singularize(parent), true),
				Parent:           r.Parent,
				RelationalFields: make(map[string]*gorma.RelationalFieldDefinition),
				BelongsTo:        make(map[string]*gorma.RelationalModelDefinition),
				HasMany:          make(map[string]*gorma.RelationalModelDefinition),
				HasOne:           make(map[string]*gorma.RelationalModelDefinition),
				ManyToMany:       make(map[string]*gorma.ManyToManyDefinition),
			}
			r.BelongsTo[models.Name] = models
		}
	}
}

// HasOne signifies a relationship between this model and another model.
// If this model HasOne(OtherModel), then OtherModel is expected
// to have a ThisModelID field as a Foreign Key to this model's
// Primary Key.  ThisModel will have a field named OtherModel of type OtherModel
// Usage:  HasOne("Proposal")
func HasOne(child string) {
	if r, ok := relationalModelDefinition(false); ok {
		field := &gorma.RelationalFieldDefinition{
			Name:        codegen.Goify(inflect.Singularize(child), true),
			HasOne:      child,
			Description: "has one " + child,
			Datatype:    gorma.HasOne,
			Parent:      r,
		}
		r.RelationalFields[field.Name] = field
		bt, ok := r.Parent.RelationalModels[child]
		if ok {
			r.HasOne[child] = bt
			// create the fk field
			f := &gorma.RelationalFieldDefinition{
				Name:              codegen.Goify(inflect.Singularize(r.Name), true) + "ID",
				HasOne:            child,
				Description:       "has one " + child,
				Datatype:          gorma.HasOneKey,
				Parent:            bt,
				DatabaseFieldName: SanitizeDBFieldName(codegen.Goify(inflect.Singularize(r.Name), true) + "ID"),
			}
			bt.RelationalFields[f.Name] = f
		} else {
			models := &gorma.RelationalModelDefinition{
				Name:             child,
				Parent:           r.Parent,
				RelationalFields: make(map[string]*gorma.RelationalFieldDefinition),
				BelongsTo:        make(map[string]*gorma.RelationalModelDefinition),
				HasMany:          make(map[string]*gorma.RelationalModelDefinition),
				HasOne:           make(map[string]*gorma.RelationalModelDefinition),
				ManyToMany:       make(map[string]*gorma.ManyToManyDefinition),
			}
			r.HasOne[child] = models
			// create the fk field
			f := &gorma.RelationalFieldDefinition{
				Name:              codegen.Goify(inflect.Singularize(r.Name), true) + "ID",
				HasOne:            child,
				Description:       "has one " + child,
				Datatype:          gorma.HasOneKey,
				Parent:            bt,
				DatabaseFieldName: SanitizeDBFieldName(codegen.Goify(inflect.Singularize(r.Name), true) + "ID"),
			}
			models.RelationalFields[f.Name] = f
		}
	}
}

// HasMany signifies a relationship between this model and a
// set of Children.  The Parent has the children, and the Children belong
// to the Parent.  The first parameter becomes the name of the
// field in the model struct, the second parameter is the name
// of the child model.  The Child model will have a ParentID field
// appended to the field list.  The Parent model definition will use
// the first parameter as the field name in the struct definition
// Usage:  HasMany("Orders", "Order")
// Struct field definition:  Children	[]Child
func HasMany(name, child string) {
	if r, ok := relationalModelDefinition(false); ok {
		field := &gorma.RelationalFieldDefinition{
			Name:        codegen.Goify(name, true),
			HasMany:     child,
			Description: "has many " + inflect.Pluralize(child),
			Datatype:    gorma.HasMany,
			Parent:      r,
		}
		r.RelationalFields[field.Name] = field
		var model *gorma.RelationalModelDefinition
		model, ok := r.Parent.RelationalModels[child]
		if ok {
			r.HasMany[child] = model
			// create the fk field
			f := &gorma.RelationalFieldDefinition{
				Name:              codegen.Goify(inflect.Singularize(r.Name), true) + "ID",
				HasMany:           child,
				Description:       "has many " + child,
				Datatype:          gorma.HasManyKey,
				Parent:            model,
				DatabaseFieldName: SanitizeDBFieldName(codegen.Goify(inflect.Singularize(r.Name), true) + "ID"),
			}
			model.RelationalFields[f.Name] = f
		} else {
			model = &gorma.RelationalModelDefinition{
				Name:             child,
				Parent:           r.Parent,
				RelationalFields: make(map[string]*gorma.RelationalFieldDefinition),
				BelongsTo:        make(map[string]*gorma.RelationalModelDefinition),
				HasMany:          make(map[string]*gorma.RelationalModelDefinition),
				HasOne:           make(map[string]*gorma.RelationalModelDefinition),
				ManyToMany:       make(map[string]*gorma.ManyToManyDefinition),
			}
			r.HasMany[child] = model
			// create the fk field
			f := &gorma.RelationalFieldDefinition{
				Name:              codegen.Goify(inflect.Singularize(r.Name), true) + "ID",
				HasMany:           child,
				Description:       "has many " + child,
				Datatype:          gorma.HasManyKey,
				Parent:            model,
				DatabaseFieldName: SanitizeDBFieldName(codegen.Goify(inflect.Singularize(r.Name), true) + "ID"),
			}
			model.RelationalFields[f.Name] = f
		}
	}
}

// ManyToMany creates a join table to store the intersection relationship
// between this model and another model.  For example, in retail an Order can
// contain many products, and a product can belong to many orders.  To express
// this relationship use the following syntax:
// Model("Order", func(){
//    ManyToMany("Product", "order_lines")
// })
// This specifies that the Order and Product tables have a "junction" table
// called `order_lines` that contains the order and product information
// The generated model will have a field called `Products` that will
// be an array of type `product.Product`
func ManyToMany(other, tablename string) {
	if r, ok := relationalModelDefinition(false); ok {
		field := &gorma.RelationalFieldDefinition{
			Name:        inflect.Pluralize(other),
			Many2Many:   other,
			Description: "many to many " + r.Name + "/" + strings.Title(other),
			Parent:      r,
		}
		r.RelationalFields[field.Name] = field
		var model *gorma.RelationalModelDefinition
		model, ok := r.Parent.RelationalModels[other]
		var m2m *gorma.ManyToManyDefinition

		if ok {
			m2m = &gorma.ManyToManyDefinition{
				Left:          r,
				Right:         model,
				DatabaseField: tablename,
			}
			r.ManyToMany[other] = m2m
		} else {
			model = &gorma.RelationalModelDefinition{
				Name:             other,
				Parent:           r.Parent,
				RelationalFields: make(map[string]*gorma.RelationalFieldDefinition),
				BelongsTo:        make(map[string]*gorma.RelationalModelDefinition),
				HasMany:          make(map[string]*gorma.RelationalModelDefinition),
				HasOne:           make(map[string]*gorma.RelationalModelDefinition),
				ManyToMany:       make(map[string]*gorma.ManyToManyDefinition),
			}
			m2m = &gorma.ManyToManyDefinition{
				Left:          r,
				Right:         model,
				DatabaseField: tablename,
			}
			r.ManyToMany[other] = m2m
		}
	}
}

// TableName creates a TableName() function that returns
// the name of the table to query. Useful for pre-existing
// schemas
func TableName(d string) {
	if r, ok := relationalModelDefinition(false); ok {
		r.TableName = d
	}
}

// Alias overrides the name of the sql store's table or field
func Alias(d string) {
	if r, ok := relationalModelDefinition(false); ok {
		r.Alias = d
	} else if f, ok := relationalFieldDefinition(false); ok {
		f.Alias = d
	}
}

// Cached caches the models for `duration` seconds
func Cached(d string) {
	if r, ok := relationalModelDefinition(false); ok {
		r.Cached = true
		dur, err := strconv.Atoi(d)
		if err != nil {
			dsl.ReportError("Duration %s couldn't be parsed as integer", d)
		}
		r.CacheDuration = dur
	}
}

// NoMedia sets a boolean flag that prevents the generation
// of media helpers
func NoMedia() {
	if r, ok := relationalModelDefinition(false); ok {
		r.NoMedia = true
	}
}

// Roler sets a boolean flag that cause the generation of a
// Role() function that returns the model's Role value
// Requires a field in the model named Role, type String
func Roler() {
	if r, ok := relationalModelDefinition(false); ok {
		r.Roler = true
	}
}

// DynamicTableName sets a boolean flag that cause the generator
// generate function definitions in the database models that specify
// the name of the database table.  Useful when using multiple tables
// with different names but same schema e.g. Users, AdminUsers
func DynamicTableName() {
	if r, ok := relationalModelDefinition(false); ok {
		r.DynamicTableName = true
	}
}

// SQLTag sets the model's struct tag `sql` value
// for indexing and other purposes
func SQLTag(d string) {
	if r, ok := relationalModelDefinition(false); ok {
		r.SQLTag = d
	} else if f, ok := relationalFieldDefinition(false); ok {
		f.SQLTag = d
	}
}
