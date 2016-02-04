package gorma

import (
	"strings"

	"bitbucket.org/pkg/inflect"
)

func (m *ManyToManyDefinition) LeftNamePlural() string {
	return inflect.Pluralize(m.Left.ModelName)
}
func (m *ManyToManyDefinition) RightNamePlural() string {
	return inflect.Pluralize(m.Right.ModelName)
}

func (m *ManyToManyDefinition) LeftName() string {
	return m.Left.ModelName
}
func (m *ManyToManyDefinition) RightName() string {
	return m.Right.ModelName
}

func (m *ManyToManyDefinition) LowerLeftName() string {
	return strings.ToLower(m.Left.ModelName)
}
func (m *ManyToManyDefinition) LowerRightName() string {
	return strings.ToLower(m.Right.ModelName)
}
