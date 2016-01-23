package gorma

import (
	"fmt"
	"strings"

	"github.com/goadesign/goa/design"
)

// Context returns the generic definition name used in error messages.
func (f *RelationalFieldDefinition) Context() string {
	if f.Name != "" {
		return fmt.Sprintf("RelationalField %#v", f.Name)
	}
	return "unnamed RelationalField"
}

// DSL returns this object's DSL
func (f *RelationalFieldDefinition) DSL() func() {
	return f.DefinitionDSL
}

// Children returnsa slice of this objects children
func (f RelationalFieldDefinition) Children() []design.Definition {
	// no children yet
	return []design.Definition{}
}

// Definition returns the field's struct definition
func (f *RelationalFieldDefinition) FieldDefinition() string {
	var comment string
	if f.Description != "" {
		comment = "// " + f.Description
	}
	def := fmt.Sprintf("%s\t%s %s %s\n", f.Name, goDatatype(f), tags(f), comment)
	return def
}

// Tags returns the sql and gorm struct tags for the Definition
func (f *RelationalFieldDefinition) Tags() string {
	return tags(f)
}

func (f *RelationalFieldDefinition) LowerName() string {
	return strings.ToLower(f.Name)
}

func goDatatype(f *RelationalFieldDefinition) string {
	var ptr string
	if f.Nullable {
		ptr = "*"
	}
	switch f.Datatype {
	case Boolean:
		return ptr + "bool"
	case Integer, BigInteger:
		return ptr + "int"
	case AutoInteger, AutoBigInteger:
		return ptr + "int " // sql/gorm tags later
	case Decimal, BigDecimal:
		return ptr + "float"
	case String:
		return ptr + "string"
	case Text:
		return ptr + "string"
	case UUID:
		return ptr + "string" // what to do about UUIDS?
	case PKInteger:
		return ptr + "int"
	case PKBigInteger:
		return ptr + "int"
	case PKUUID:
		return ptr + "string " // TBD
	case Timestamp, NullableTimestamp:
		return ptr + "time.Time"
	case BelongsTo:
		return ptr + "int"
	case HasMany:
		return fmt.Sprintf("[]%s", f.HasMany)
	case HasManyKey:
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
	if f.Alias != "" {
		gormtags = append(gormtags, "column:"+f.Alias)
	}
	if f.PrimaryKey {
		gormtags = append(gormtags, "primary_key")
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

/*
	Boolean           FieldType             = "bool"
	Integer           FieldType             = "integer"
	BigInteger        FieldType             = "biginteger"
	AutoInteger       FieldType             = "auto_integer"
	AutoBigInteger    FieldType             = "auto_biginteger"
	Decimal           FieldType             = "decimal"
	BigDecimal        FieldType             = "bigdecimal"
	String            FieldType             = "string"
	Text              FieldType             = "text"
	UUID              FieldType             = "uuid"
	PKInteger         FieldType             = "pkinteger"
	PKBigInteger      FieldType             = "pkbiginteger"
	PKUUID            FieldType             = "pkuuid"
	Timestamp         FieldType             = "timestamp"
	NullableTimestamp FieldType             = "nulltimestamp"

*/
