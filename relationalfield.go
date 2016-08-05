package gorma

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/dslengine"
)

// NewRelationalFieldDefinition returns an initialized
// RelationalFieldDefinition.
func NewRelationalFieldDefinition() *RelationalFieldDefinition {
	m := &RelationalFieldDefinition{
		Mappings: make(map[string]*MapDefinition),
	}
	return m
}

// Context returns the generic definition name used in error messages.
func (f *RelationalFieldDefinition) Context() string {
	if f.FieldName != "" {
		return fmt.Sprintf("RelationalField %#v", f.FieldName)
	}
	return "unnamed RelationalField"
}

// DSL returns this object's DSL.
func (f *RelationalFieldDefinition) DSL() func() {
	return f.DefinitionDSL
}

// Children returns a slice of this objects children.
func (f RelationalFieldDefinition) Children() []dslengine.Definition {
	// no children yet
	return []dslengine.Definition{}
}

// Attribute implements the Container interface of the goa Attribute
// model.
func (f *RelationalFieldDefinition) Attribute() *design.AttributeDefinition {
	return f.a
}

// FieldDefinition returns the field's struct definition.
func (f *RelationalFieldDefinition) FieldDefinition() string {
	var comment string
	if f.Description != "" {
		comment = "// " + f.Description
	}
	def := fmt.Sprintf("%s\t%s %s %s\n", f.FieldName, goDatatype(f, true), tags(f), comment)
	return def
}

// Tags returns the sql and gorm struct tags for the Definition.
func (f *RelationalFieldDefinition) Tags() string {
	return tags(f)
}

// LowerName returns the field name as a lowercase string.
func (f *RelationalFieldDefinition) LowerName() string {
	return strings.ToLower(f.FieldName)
}

// Underscore returns the field name as a lowercase string in snake case.
func (f *RelationalFieldDefinition) Underscore() string {
	runes := []rune(f.FieldName)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

func goDatatype(f *RelationalFieldDefinition, includePtr bool) string {
	var ptr string
	if f.Nullable && includePtr {
		ptr = "*"
	}
	switch f.Datatype {
	case Boolean:
		return ptr + "bool"
	case Integer:
		return ptr + "int"
	case BigInteger:
		return ptr + "int64"
	case AutoInteger, AutoBigInteger:
		return ptr + "int " // sql/gorm tags later
	case Decimal:
		return ptr + "float32"
	case BigDecimal:
		return ptr + "float64"
	case String:
		return ptr + "string"
	case Text:
		return ptr + "string"
	case UUID:
		return ptr + "uuid.UUID"
	case Timestamp, NullableTimestamp:
		return ptr + "time.Time"
	case BelongsTo:
		return ptr + "int"
	case HasMany:
		return fmt.Sprintf("[]%s", f.HasMany)
	case HasManyKey, HasOneKey:
		return ptr + "int"
	case HasOne:
		return fmt.Sprintf("%s", f.HasOne)
	default:

		if f.Many2Many != "" {
			return fmt.Sprintf("[]%s", f.Many2Many)
		}
	}

	return "UNKNOWN TYPE"
}

func tags(f *RelationalFieldDefinition) string {
	var sqltags []string
	if f.SQLTag != "" {
		sqltags = append(sqltags, f.SQLTag)
	}

	var gormtags []string
	if f.DatabaseFieldName != "" && f.DatabaseFieldName != f.Underscore() {
		gormtags = append(gormtags, "column:"+f.DatabaseFieldName)
	}
	if f.PrimaryKey {
		gormtags = append(gormtags, "primary_key")
	}
	if f.Many2Many != "" {
		gormtags = append(gormtags, "many2many:"+f.TableName)
	}

	var tags []string
	if len(sqltags) > 0 {
		sqltag := "sql:\"" + strings.Join(sqltags, ";") + "\""
		tags = append(tags, sqltag)
	}
	if len(gormtags) > 0 {
		gormtag := "gorm:\"" + strings.Join(gormtags, ";") + "\""
		tags = append(tags, gormtag)
	}

	if len(tags) > 0 {
		return "`" + strings.Join(tags, " ") + "`"
	}
	return ""
}
