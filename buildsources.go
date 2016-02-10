package gorma

import "fmt"

// NewBuildSource returns an initialized BuildSource
func NewBuildSource() *BuildSource {
	bs := &BuildSource{}
	return bs
}

// Context returns the generic definition name used in error messages.
func (f *BuildSource) Context() string {
	if f.BuildSourceName != "" {
		return fmt.Sprintf("BuildSource %#v", f.BuildSourceName)
	}
	return "unnamed BuildSource"
}

// DSL returns this object's DSL.
func (f *BuildSource) DSL() func() {
	return f.DefinitionDSL
}
