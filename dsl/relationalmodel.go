package dsl

import (
	"strconv"
	"strings"

	"bitbucket.org/pkg/inflect"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/dslengine"
	"github.com/goadesign/goa/goagen/codegen"
	"github.com/goadesign/gorma"
)

// Model is the DSL that represents a Relational Model.
// Model name should be Title cased.  Use BuildsFrom() and RendersTo() DSL
// to define the mapping between a Model and a Goa Type.
// Models may contain multiple instances of the `Field` DSL to
// add fields to the model.
//
// To control whether the ID field is auto-generated, use `NoAutomaticIDFields()`
// Similarly, use NoAutomaticTimestamps() and NoAutomaticSoftDelete() to
// prevent CreatedAt, UpdatedAt and DeletedAt fields from being created.
func Model(name string, dsl func()) {
	if s, ok := relationalStoreDefinition(true); ok {
		var model *gorma.RelationalModelDefinition
		var ok bool
		model, ok = s.RelationalModels[name]
		if !ok {
			model = gorma.NewRelationalModelDefinition()
			model.ModelName = name
			model.DefinitionDSL = dsl
			model.Parent = s
			model.RelationalFields = make(map[string]*gorma.RelationalFieldDefinition)
		} else {
			dslengine.ReportError("Model %s already exists", name)
			return
		}
		s.RelationalModels[name] = model
		model.UserTypeDefinition.TypeName = model.ModelName
		// much stutter here
		// TODO(BJK) refactor
		if !s.NoAutoIDFields {
			field := gorma.NewRelationalFieldDefinition()
			field.FieldName = SanitizeFieldName("ID")
			field.Parent = model
			field.Datatype = gorma.Integer
			field.PrimaryKey = true
			field.Nullable = false
			field.DatabaseFieldName = SanitizeDBFieldName("ID")
			model.RelationalFields[field.FieldName] = field
		}
		if !s.NoAutoTimestamps {
			// add createdat
			field := gorma.NewRelationalFieldDefinition()
			field.FieldName = SanitizeFieldName("CreatedAt")
			field.Parent = model
			field.Datatype = gorma.Timestamp
			field.DatabaseFieldName = SanitizeDBFieldName("CreatedAt")
			model.RelationalFields[field.FieldName] = field

			// add updatedat
			field = gorma.NewRelationalFieldDefinition()
			field.FieldName = SanitizeFieldName("UpdatedAt")
			field.Parent = model
			field.Datatype = gorma.Timestamp
			field.DatabaseFieldName = SanitizeDBFieldName("UpdatedAt")
			model.RelationalFields[field.FieldName] = field
		}
		if !s.NoAutoSoftDelete {
			// Add softdelete
			field := gorma.NewRelationalFieldDefinition()
			field.FieldName = SanitizeFieldName("DeletedAt")
			field.Parent = model
			field.Nullable = true
			field.Datatype = gorma.NullableTimestamp
			field.DatabaseFieldName = SanitizeDBFieldName("DeletedAt")
			model.RelationalFields[field.FieldName] = field
		}
	}
}

// RendersTo informs Gorma that this model will need to be
// rendered to a Goa type.  Conversion functions
// will be generated to convert to/from the model.
//
// Usage: RendersTo(MediaType)
func RendersTo(rt interface{}) {
	if m, ok := relationalModelDefinition(false); ok {
		mts, ok := rt.(*design.MediaTypeDefinition)
		if ok {
			m.RenderTo[mts.TypeName] = mts
		}

	}
}

// BuildsFrom informs Gorma that this model will be populated
// from a Goa UserType.  Conversion functions
// will be generated to convert from the payload to the model.
//
// Usage:  BuildsFrom(YourType)
//
// Fields not in `YourType` that you want in your model must be
// added explicitly with the `Field` DSL.
func BuildsFrom(dsl func()) {
	if m, ok := relationalModelDefinition(false); ok {
		/*		mts, ok := bf.(*design.UserTypeDefinition)
				if ok {
					m.BuiltFrom[mts.TypeName] = mts
				} else if mts, ok := bf.(*design.MediaTypeDefinition); ok {
					m.BuiltFrom[mts.TypeName] = mts.UserTypeDefinition
				}
				m.PopulateFromModeledType()
		*/
		bf := gorma.NewBuildSource()
		bf.DefinitionDSL = dsl
		bf.Parent = m
		m.BuildSources = append(m.BuildSources, bf)
	}

}

// Payload specifies the Resource and Action containing
// a User Type (Payload).
// Gorma will generate a conversion function for the Payload to
// the Model.
func Payload(r interface{}, act string) {
	if bs, ok := buildSourceDefinition(true); ok {
		var res *design.ResourceDefinition
		var resName string
		if n, ok := r.(string); ok {
			res = design.Design.Resources[n]
			resName = n
		} else {
			res, _ = r.(*design.ResourceDefinition)
		}
		if res == nil {
			dslengine.ReportError("There is no resource %q", resName)
			return
		}
		a, ok := res.Actions[act]
		if !ok {
			dslengine.ReportError("There is no action")
			return
		}
		payload := a.Payload

		// Set UTD in BuildsFrom parent context

		bs.Parent.BuiltFrom[payload.TypeName] = payload
		bs.Parent.PopulateFromModeledType()
	}
}

// BelongsTo signifies a relationship between this model and a
// Parent.  The Parent has the child, and the Child belongs
// to the Parent.
//
// Usage:  BelongsTo("User")
func BelongsTo(parent string) {
	if r, ok := relationalModelDefinition(false); ok {
		idfield := gorma.NewRelationalFieldDefinition()
		idfield.FieldName = codegen.Goify(inflect.Singularize(parent), true) + "ID"
		idfield.Description = "Belongs To " + codegen.Goify(inflect.Singularize(parent), true)
		idfield.Parent = r
		idfield.Datatype = gorma.BelongsTo
		idfield.DatabaseFieldName = SanitizeDBFieldName(codegen.Goify(inflect.Singularize(parent), true) + "ID")
		r.RelationalFields[idfield.FieldName] = idfield
		bt, ok := r.Parent.RelationalModels[codegen.Goify(inflect.Singularize(parent), true)]
		if ok {
			r.BelongsTo[bt.ModelName] = bt
		} else {
			model := gorma.NewRelationalModelDefinition()
			model.ModelName = codegen.Goify(inflect.Singularize(parent), true)
			model.Parent = r.Parent
			r.BelongsTo[model.ModelName] = model
		}
	}
}

// HasOne signifies a relationship between this model and another model.
// If this model HasOne(OtherModel), then OtherModel is expected
// to have a ThisModelID field as a Foreign Key to this model's
// Primary Key.  ThisModel will have a field named OtherModel of type
// OtherModel.
//
// Usage:  HasOne("Proposal")
func HasOne(child string) {
	if r, ok := relationalModelDefinition(false); ok {
		field := gorma.NewRelationalFieldDefinition()
		field.FieldName = codegen.Goify(inflect.Singularize(child), true)
		field.HasOne = child
		field.Description = "has one " + child
		field.Datatype = gorma.HasOne
		field.Parent = r
		r.RelationalFields[field.FieldName] = field
		bt, ok := r.Parent.RelationalModels[child]

		// Refactor (BJK)
		if ok {
			r.HasOne[child] = bt
			// create the fk field
			f := gorma.NewRelationalFieldDefinition()
			f.FieldName = codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID"
			f.HasOne = child
			f.Description = "has one " + child
			f.Datatype = gorma.HasOneKey
			f.Parent = bt
			f.DatabaseFieldName = SanitizeDBFieldName(codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID")
			bt.RelationalFields[f.FieldName] = f
		} else {
			model := gorma.NewRelationalModelDefinition()
			model.ModelName = child
			model.Parent = r.Parent
			r.HasOne[child] = model

			// create the fk field
			f := gorma.NewRelationalFieldDefinition()
			f.FieldName = codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID"
			f.HasOne = child
			f.Description = "has one " + child
			f.Datatype = gorma.HasOneKey
			f.Parent = bt
			f.DatabaseFieldName = SanitizeDBFieldName(codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID")
			model.RelationalFields[f.FieldName] = f
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
//
// Usage:  HasMany("Orders", "Order")
//
// Generated struct field definition:  Children	[]Child
func HasMany(name, child string) {
	if r, ok := relationalModelDefinition(false); ok {
		field := gorma.NewRelationalFieldDefinition()
		field.FieldName = codegen.Goify(name, true)
		field.HasMany = child
		field.Description = "has many " + inflect.Pluralize(child)
		field.Datatype = gorma.HasMany
		field.Parent = r
		r.RelationalFields[field.FieldName] = field

		var model *gorma.RelationalModelDefinition
		model, ok := r.Parent.RelationalModels[child]
		if ok {
			r.HasMany[child] = model
			// create the fk field
			f := gorma.NewRelationalFieldDefinition()
			f.FieldName = codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID"
			f.HasMany = child
			f.Description = "has many " + child
			f.Datatype = gorma.HasManyKey
			f.Parent = model
			f.DatabaseFieldName = SanitizeDBFieldName(codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID")
			model.RelationalFields[f.FieldName] = f
		} else {
			model = gorma.NewRelationalModelDefinition()
			model.ModelName = child
			model.Parent = r.Parent
		}
		r.HasMany[child] = model
		// create the fk field
		f := gorma.NewRelationalFieldDefinition()
		f.FieldName = codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID"
		f.HasMany = child
		f.Description = "has many " + child
		f.Datatype = gorma.HasManyKey
		f.Parent = model
		f.DatabaseFieldName = SanitizeDBFieldName(codegen.Goify(inflect.Singularize(r.ModelName), true) + "ID")
		model.RelationalFields[f.FieldName] = f
	}
}

// ManyToMany creates a join table to store the intersection relationship
// between this model and another model.  For example, in retail an Order can
// contain many products, and a product can belong to many orders.  To express
// this relationship use the following syntax:
//
//		Model("Order", func(){
//			ManyToMany("Product", "order_lines")
//		})
//
// This specifies that the Order and Product tables have a "junction" table
// called `order_lines` that contains the order and product information.
// The generated model will have a field called `Products` that will
// be an array of type `Product`.
func ManyToMany(other, tablename string) {
	if r, ok := relationalModelDefinition(false); ok {
		field := gorma.NewRelationalFieldDefinition()
		field.FieldName = inflect.Pluralize(other)
		field.TableName = tablename
		field.Many2Many = other
		field.Datatype = gorma.Many2Many
		field.Description = "many to many " + r.ModelName + "/" + strings.Title(other)
		field.Parent = r
		r.RelationalFields[field.FieldName] = field
		var model *gorma.RelationalModelDefinition
		model, ok := r.Parent.RelationalModels[other]
		var m2m *gorma.ManyToManyDefinition

		// refactor (BJK)
		if ok {
			m2m = &gorma.ManyToManyDefinition{
				Left:          r,
				Right:         model,
				DatabaseField: tablename,
			}
			r.ManyToMany[other] = m2m
		} else {
			model := gorma.NewRelationalModelDefinition()
			model.ModelName = other
			model.Parent = r.Parent
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
		f.DatabaseFieldName = d
	}
}

// Cached caches the models for `duration` seconds.
// Not fully implemented yet, and not guaranteed to stay
// in Gorma long-term because of the complex rendering
// that happens in the conversion functions.
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
// Creates a "Role" field in the table if it doesn't already exist
// as a string type
func Roler() {
	if r, ok := relationalModelDefinition(false); ok {
		r.Roler = true
		if _, ok := r.RelationalFields["Role"]; !ok {
			field := gorma.NewRelationalFieldDefinition()
			field.FieldName = "Role"
			field.Datatype = gorma.String
			r.RelationalFields["Role"] = field
		}
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
