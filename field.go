package gorma

import "fmt"

// Context returns the generic definition name used in error messages.
func (f *RelationalFieldDefinition) Context() string {
	if f.Name != "" {
		return fmt.Sprintf("RelationalField %#v", f.Name)
	}
	return "unnamed RelationalField"
}

// Definition returns the field's struct definition
func (f *RelationalFieldDefinition) Definition() string {

	return ""

}

// Tags returns teh sql and gorm struct tags for the Definition
func (f *RelationalFieldDefinition) Tags() string {

	return ""
}
