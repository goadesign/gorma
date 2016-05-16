package gorma

// NewRolesDefinition returns an initialized RoleDefinition
func NewRolesDefinition() *RolesDefinition {
	r := &RolesDefinition{}
	r.Roles = make(map[string]*RoleDefinition)
	return r
}

// Context returns the generic definition name used in error messages.
func (f *RolesDefinition) Context() string {
	return "unnamed RolesDefinition"
}

// DSL returns this object's DSL.
func (f *RolesDefinition) DSL() func() {
	return f.DefinitionDSL
}
