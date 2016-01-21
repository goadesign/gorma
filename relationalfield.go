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

	def := fmt.Sprintf("%s\t%s %s %s\n", f.Name, goDatatype(f), tags(f), "// "+f.Description)
	return def

}

// Tags returns teh sql and gorm struct tags for the Definition
func (f *RelationalFieldDefinition) Tags() string {

	return ""
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
	default:
		if f.BelongsTo != "" {
			return fmt.Sprintf("%s.%s", strings.ToLower(f.BelongsTo), f.BelongsTo)
		}
		if f.HasMany != "" {

			return fmt.Sprintf("[]%s.%s", strings.ToLower(f.HasMany), f.HasMany)
		}
		if f.HasOne != "" {
			return fmt.Sprintf("%s.%s", strings.ToLower(f.HasOne), f.HasOne)
		}
		if f.Many2Many != "" {
			return fmt.Sprintf("[]%s.%s", strings.ToLower(f.Many2Many), f.Many2Many)
		}
	}

	return "UNKNOWN TYPE"

}

func tags(f *RelationalFieldDefinition) string {
	var sqltag, stag, atag, gormtags string
	var sqltags []string
	var tags []string
	if f.SQLTag != "" {
		stag = f.SQLTag
		sqltags = append(sqltags, stag)
	}
	if f.Alias != "" {
		atag = f.Alias
		sqltags = append(sqltags, atag)
	}
	if f.Alias != "" || f.SQLTag != "" {
		sqltag = "sql:\"" + strings.Join(sqltags, ";") + "\""
		tags = append(tags, sqltag)

	}

	if f.PrimaryKey {
		gormtags = "gorm:\"" + "primary_key" + "\""
		tags = append(tags, gormtags)
	}
	output := strings.Join(tags, " ")
	if sqltag != "" || gormtags != "" {
		return fmt.Sprintf("`%s`", output)
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
