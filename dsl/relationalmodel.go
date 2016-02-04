package dsl

import (
	"fmt"
	"strconv"
	"strings"

	"bitbucket.org/pkg/inflect"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/dslengine"
	"github.com/goadesign/goa/goagen/codegen"
	"github.com/goadesign/gorma"
	"github.com/kr/pretty"
)

// Model is the DSL that represents a Relational Model.
// Model name should be Title cased.  Use BuiltFrom() and RenderTo() DSL
// to define the mapping between a Model and a Goa Type.
func Model(name string, dsl func()) {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.
	checkInit()
	if s, ok := relationalStoreDefinition(true); ok {
		model, ok := s.RelationalModels[name]
		if !ok {
			model = &gorma.RelationalModelDefinition{
				ModelName:        name,
				DefinitionDSL:    dsl,
				Parent:           s,
				RelationalFields: make(map[string]*gorma.RelationalFieldDefinition),
				SourceMaps:       make(map[string]*gorma.SourceMapping),
				TargetMaps:       make(map[string]*gorma.TargetMapping),
				BuiltFrom:        make(map[string]*design.UserTypeDefinition),
				RenderTo:         make(map[string]*design.MediaTypeDefinition),
				BelongsTo:        make(map[string]*gorma.RelationalModelDefinition),
				HasMany:          make(map[string]*gorma.RelationalModelDefinition),
				HasOne:           make(map[string]*gorma.RelationalModelDefinition),
				ManyToMany:       make(map[string]*gorma.ManyToManyDefinition),
			}
		} else {
			model.DefinitionDSL = dsl
		}
		s.RelationalModels[name] = model
		if !s.NoAutoIDFields {
			field := &gorma.RelationalFieldDefinition{
				Name:              SanitizeFieldName("ID"),
				Parent:            model,
				Datatype:          gorma.PKInteger,
				PrimaryKey:        true,
				DatabaseFieldName: SanitizeDBFieldName("ID"),
			}
			model.RelationalFields[field.Name] = field
		}
		if !s.NoAutoTimestamps {
			// add createdat
			field := &gorma.RelationalFieldDefinition{
				Name:              SanitizeFieldName("CreatedAt"),
				Parent:            model,
				Datatype:          gorma.Timestamp,
				DatabaseFieldName: SanitizeDBFieldName("CreatedAt"),
			}
			model.RelationalFields[field.Name] = field
			// add updatedat
			field = &gorma.RelationalFieldDefinition{
				Name:              SanitizeFieldName("UpdatedAt"),
				Parent:            model,
				Datatype:          gorma.Timestamp,
				DatabaseFieldName: SanitizeDBFieldName("UpdatedAt"),
			}
			model.RelationalFields[field.Name] = field
		}
		if !s.NoAutoSoftDelete {
			// Add softdelete
			field := &gorma.RelationalFieldDefinition{
				Name:              SanitizeFieldName("DeletedAt"),
				Parent:            model,
				Nullable:          true,
				Datatype:          gorma.NullableTimestamp,
				DatabaseFieldName: SanitizeDBFieldName("DeletedAt"),
			}
			model.RelationalFields[field.Name] = field
		}
	}
}

// RenderTo informs Gorma that this model will need to be
// rendered to a Goa type.  Conversion functions
// will be generated to convert to/from the model.
// Usage:   RenderTo(MediaType)
func RenderTo(rt interface{}, dsl func()) {
	checkInit()
	if m, ok := relationalModelDefinition(true); ok {
		mts, ok := rt.(*design.MediaTypeDefinition)
		if ok {
			m.RenderTo[mts.TypeName] = mts
			m.TargetMaps[mts.TypeName] = &gorma.TargetMapping{
				MappingName:   m.ModelName + ":" + mts.TypeName,
				Description:   "Maps " + m.ModelName + " to " + mts.TypeName,
				Remote:        mts.UserTypeDefinition,
				Parent:        m,
				DefinitionDSL: dsl,
				Mappings:      make(map[string]*gorma.MapDefinition),
			}
		}

	}
}

// BuiltFrom informs Gorma that this model will be populated
// from a Goa UserType.  Conversion functions
// will be generated to convert from the payload to the model.
// Usage:  BuiltFrom(YourType)
func BuiltFrom(bf interface{}, dsl func()) {
	checkInit()
	if m, ok := relationalModelDefinition(true); ok {
		mts, ok := bf.(*design.UserTypeDefinition)
		if ok {
			m.BuiltFrom[mts.TypeName] = mts
			m.SourceMaps[mts.TypeName] = &gorma.SourceMapping{
				AttributeDefinition: &design.AttributeDefinition{DSLFunc: dsl},
				MappingName:         mts.TypeName + ":" + m.ModelName,
				Description:         "Maps " + mts.TypeName + " to " + m.ModelName,
				Remote:              mts,
				Parent:              m,
				DefinitionDSL:       dsl,
				Mappings:            make(map[string]*gorma.MapDefinition),
			}
		} else if mts, ok := bf.(*design.MediaTypeDefinition); ok {
			m.BuiltFrom[mts.TypeName] = mts.UserTypeDefinition
			m.SourceMaps[mts.TypeName] = &gorma.SourceMapping{

				AttributeDefinition: &design.AttributeDefinition{DSLFunc: dsl},
				MappingName:         mts.TypeName + ":" + m.ModelName,
				Description:         "Maps " + mts.TypeName + " to " + m.ModelName,
				Remote:              mts.UserTypeDefinition,
				Parent:              m,
				DefinitionDSL:       dsl,
				Mappings:            make(map[string]*gorma.MapDefinition),
			}
		}
		m.PopulateFromModeledType()

	}
}

// Map defines the mapping between Gorma and Goa types
func Map(left, right string, fieldType gorma.FieldType) {
	checkInit()

	if m, ok := sourceMappingDefinition(false); ok {
		fmt.Println("Source Map:", left, right, fieldType, m.Name)
		fmt.Println(m.Parent, m.Parent.Name)
		if md, ok := m.Mappings[left+":"+right]; ok {
			// we have a mapping already
			pretty.Println("Source Mapping Exists", md)
		} else {
			// no mapping exists
			pretty.Println("Need New Source Mapping")
			md := &gorma.MapDefinition{
				RemoteField: left,
				ParentField: right,
				GormaType:   fieldType,
			}
			m.Mappings[left+":"+right] = md
		}
		field, ok := m.Parent.RelationalFields[codegen.Goify(right, true)]
		if !ok {
			field = &gorma.RelationalFieldDefinition{
				Name:     codegen.Goify(right, true),
				Parent:   m.Parent,
				Datatype: fieldType,
			}
			m.Parent.RelationalFields[codegen.Goify(right, true)] = field
		}

	} else if m2, ok := targetMappingDefinition(true); ok {
		fmt.Println("Target Map:", left, right, fieldType, m2.MappingName)
		fmt.Println(m2.Parent, m2.Parent.Name)
		if md, ok := m2.Mappings[left+":"+right]; ok {
			// we have a mapping already
			pretty.Println("Target Mapping Exists", md)
		} else {
			// no mapping exists
			pretty.Println("Need New Target Mapping")
			md := &gorma.MapDefinition{
				RemoteField: right,
				ParentField: left,
				GormaType:   fieldType,
			}
			m.Mappings[right+":"+left] = md
		}
		field, ok := m2.Parent.RelationalFields[codegen.Goify(right, true)]
		if !ok {
			field = &gorma.RelationalFieldDefinition{
				Name:     codegen.Goify(right, true),
				Parent:   m2.Parent,
				Datatype: fieldType,
			}
			m2.Parent.RelationalFields[codegen.Goify(right, true)] = field
		}
	} else {
		panic("WTF")
	}

}

// BelongsTo signifies a relationship between this model and a
// Parent.  The Parent has the child, and the Child belongs
// to the Parent.
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
			r.BelongsTo[bt.ModelName] = bt

		} else {
			models := &gorma.RelationalModelDefinition{
				ModelName:        codegen.Goify(inflect.Singularize(parent), true),
				Parent:           r.Parent,
				RelationalFields: make(map[string]*gorma.RelationalFieldDefinition),
				SourceMaps:       make(map[string]*gorma.SourceMapping),
				TargetMaps:       make(map[string]*gorma.TargetMapping),
				BuiltFrom:        make(map[string]*design.UserTypeDefinition),
				RenderTo:         make(map[string]*design.MediaTypeDefinition),
				BelongsTo:        make(map[string]*gorma.RelationalModelDefinition),
				HasMany:          make(map[string]*gorma.RelationalModelDefinition),
				HasOne:           make(map[string]*gorma.RelationalModelDefinition),
				ManyToMany:       make(map[string]*gorma.ManyToManyDefinition),
			}
			r.BelongsTo[models.ModelName] = models
		}
	}
}

// HasOne signifies a relationship between this model and another model.
// If this model HasOne(OtherModel), then OtherModel is expected
// to have a ThisModelID field as a Foreign Key to this model's
// Primary Key.  ThisModel will have a field named OtherModel of type
// OtherModel.
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
				Name:              codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID",
				HasOne:            child,
				Description:       "has one " + child,
				Datatype:          gorma.HasOneKey,
				Parent:            bt,
				DatabaseFieldName: SanitizeDBFieldName(codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID"),
			}
			bt.RelationalFields[f.Name] = f
		} else {
			models := &gorma.RelationalModelDefinition{
				ModelName:        child,
				Parent:           r.Parent,
				RelationalFields: make(map[string]*gorma.RelationalFieldDefinition),
				SourceMaps:       make(map[string]*gorma.SourceMapping),
				TargetMaps:       make(map[string]*gorma.TargetMapping),
				BuiltFrom:        make(map[string]*design.UserTypeDefinition),
				RenderTo:         make(map[string]*design.MediaTypeDefinition),
				BelongsTo:        make(map[string]*gorma.RelationalModelDefinition),
				HasMany:          make(map[string]*gorma.RelationalModelDefinition),
				HasOne:           make(map[string]*gorma.RelationalModelDefinition),
				ManyToMany:       make(map[string]*gorma.ManyToManyDefinition),
			}
			r.HasOne[child] = models
			// create the fk field
			f := &gorma.RelationalFieldDefinition{
				Name:              codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID",
				HasOne:            child,
				Description:       "has one " + child,
				Datatype:          gorma.HasOneKey,
				Parent:            bt,
				DatabaseFieldName: SanitizeDBFieldName(codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID"),
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
// the first parameter as the field name in the struct definition.
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
				Name:              codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID",
				HasMany:           child,
				Description:       "has many " + child,
				Datatype:          gorma.HasManyKey,
				Parent:            model,
				DatabaseFieldName: SanitizeDBFieldName(codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID"),
			}
			model.RelationalFields[f.Name] = f
		} else {
			model = &gorma.RelationalModelDefinition{
				ModelName:        child,
				Parent:           r.Parent,
				RelationalFields: make(map[string]*gorma.RelationalFieldDefinition),
				SourceMaps:       make(map[string]*gorma.SourceMapping),
				TargetMaps:       make(map[string]*gorma.TargetMapping),
				BuiltFrom:        make(map[string]*design.UserTypeDefinition),
				RenderTo:         make(map[string]*design.MediaTypeDefinition),
				BelongsTo:        make(map[string]*gorma.RelationalModelDefinition),
				HasMany:          make(map[string]*gorma.RelationalModelDefinition),
				HasOne:           make(map[string]*gorma.RelationalModelDefinition),
				ManyToMany:       make(map[string]*gorma.ManyToManyDefinition),
			}
			r.HasMany[child] = model
			// create the fk field
			f := &gorma.RelationalFieldDefinition{
				Name:              codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID",
				HasMany:           child,
				Description:       "has many " + child,
				Datatype:          gorma.HasManyKey,
				Parent:            model,
				DatabaseFieldName: SanitizeDBFieldName(codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID"),
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
// called `order_lines` that contains the order and product information.
// The generated model will have a field called `Products` that will
// be an array of type `product.Product`.
func ManyToMany(other, tablename string) {
	if r, ok := relationalModelDefinition(false); ok {
		field := &gorma.RelationalFieldDefinition{
			Name:        inflect.Pluralize(other),
			Many2Many:   other,
			Description: "many to many " + r.ModelName + "/" + strings.Title(other),
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
				ModelName:        other,
				Parent:           r.Parent,
				RelationalFields: make(map[string]*gorma.RelationalFieldDefinition),
				SourceMaps:       make(map[string]*gorma.SourceMapping),
				TargetMaps:       make(map[string]*gorma.TargetMapping),
				BuiltFrom:        make(map[string]*design.UserTypeDefinition),
				RenderTo:         make(map[string]*design.MediaTypeDefinition),
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

// Alias overrides the name of the SQL store's table or field.
func Alias(d string) {
	if r, ok := relationalModelDefinition(false); ok {
		r.Alias = d
	} else if f, ok := relationalFieldDefinition(false); ok {
		f.Alias = d
	}
}

// Cached caches the models for `duration` seconds.
func Cached(d string) {
	if r, ok := relationalModelDefinition(false); ok {
		r.Cached = true
		dur, err := strconv.Atoi(d)
		if err != nil {
			dslengine.ReportError("Duration %s couldn't be parsed as integer", d)
		}
		r.CacheDuration = dur
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

// DynamicTableName sets a boolean flag that causes the generator to
// generate function definitions in the database models that specify
// the name of the database table.  Useful when using multiple tables
// with different names but same schema e.g. Users, AdminUsers.
func DynamicTableName() {
	if r, ok := relationalModelDefinition(false); ok {
		r.DynamicTableName = true
	}
}

// SQLTag sets the model's struct tag `sql` value
// for indexing and other purposes.
func SQLTag(d string) {
	if r, ok := relationalModelDefinition(false); ok {
		r.SQLTag = d
	} else if f, ok := relationalFieldDefinition(false); ok {
		f.SQLTag = d
	}
}
