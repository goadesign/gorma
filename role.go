package gorma

import "fmt"

// NewRoleDefinition returns an initialized RoleDefinition
func NewRoleDefinition() *RoleDefinition {
	r := &RoleDefinition{}
	return r
}

// Context returns the generic definition name used in error messages.
func (f *RoleDefinition) Context() string {
	if f.Name != "" {
		return fmt.Sprintf("RoleDefinition %#v", f.Name)
	}
	return "unnamed RoleDefinition"
}

// DSL returns this object's DSL.
func (f *RoleDefinition) DSL() func() {
	return f.DefinitionDSL
}
