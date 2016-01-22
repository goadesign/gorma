package dsl

import (
	"strconv"
	"strings"

	"bitbucket.org/pkg/inflect"

	"github.com/bketelsen/gorma"
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/dsl"
	"github.com/goadesign/goa/goagen/codegen"
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
		//models.PopulateFromModeledType() -- need to do this later
		s.RelationalModels[name] = models
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
			Name:        codegen.Goify(inflect.Singularize(parent), true) + "ID",
			Description: "belongs to " + parent,
			Parent:      r,
			Datatype:    gorma.Integer,
		}

		r.RelationalFields[idfield.Name] = idfield
		bt, ok := r.Parent.RelationalModels[parent]
		if ok {
			r.BelongsTo[parent] = bt
		} else {
			models := &gorma.RelationalModelDefinition{
				Name:             parent,
				Parent:           r.Parent,
				RelationalFields: make(map[string]*gorma.RelationalFieldDefinition),
				BelongsTo:        make(map[string]*gorma.RelationalModelDefinition),
				HasMany:          make(map[string]*gorma.RelationalModelDefinition),
				HasOne:           make(map[string]*gorma.RelationalModelDefinition),
				ManyToMany:       make(map[string]*gorma.ManyToManyDefinition),
			}
			r.BelongsTo[parent] = models
		}

	}
}

// HasOne signifies a relationship between this model and a
// Child.  The Parent has the child, and the Child belongs
// to the Parent.
// Usage:  HasOne("Proposal")
func HasOne(child string) {
	if r, ok := relationalModelDefinition(false); ok {
		field := &gorma.RelationalFieldDefinition{
			Name:        codegen.Goify(inflect.Singularize(child), true),
			HasOne:      child,
			Description: "has one " + child,
			Parent:      r,
		}
		r.RelationalFields[field.Name] = field
		bt, ok := r.Parent.RelationalModels[child]
		if ok {
			// wow!
			r.HasOne[child] = bt
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
		}
	}
}

// HasMany signifies a relationship between this model and a
// set of Children.  The Parent has the children, and the Children belong
// to the Parent.  The first parameter becomes the name of the
// field in the model struct, the second parameter is the name
// of the child model.
// Usage:  HasMany("Orders", "Order")
func HasMany(name, child string) {
	if r, ok := relationalModelDefinition(false); ok {
		field := &gorma.RelationalFieldDefinition{
			Name:        name,
			HasMany:     child,
			Description: "has many " + inflect.Pluralize(child),
			Parent:      r,
		}
		r.RelationalFields[field.Name] = field
		var model *gorma.RelationalModelDefinition
		model, ok := r.Parent.RelationalModels[child]
		if ok {
			r.HasMany[child] = model
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
