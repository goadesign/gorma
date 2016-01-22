package gorma

import (
	"strings"

	"bitbucket.org/pkg/inflect"
)

func (m *ManyToManyDefinition) LeftNamePlural() string {
	return inflect.Pluralize(m.Left.Name)
}
func (m *ManyToManyDefinition) RightNamePlural() string {
	return inflect.Pluralize(m.Right.Name)
}

func (m *ManyToManyDefinition) LeftName() string {
	return m.Left.Name
}
func (m *ManyToManyDefinition) RightName() string {
	return m.Right.Name
}

func (m *ManyToManyDefinition) LowerLeftName() string {
	return strings.ToLower(m.Left.Name)
}
func (m *ManyToManyDefinition) LowerRightName() string {
	return strings.ToLower(m.Right.Name)
}
