package gorma

import (
	"fmt"

	"github.com/raphael/goa/design"
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

	return ""

}

// Tags returns teh sql and gorm struct tags for the Definition
func (f *RelationalFieldDefinition) Tags() string {

	return ""
}
