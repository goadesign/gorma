package gorma

import (
	"fmt"
	"sort"
	"strings"

	"github.com/raphael/goa/design"

	"bitbucket.org/pkg/inflect"
)

// Context returns the generic definition name used in error messages.
func (a *RelationalStoreDefinition) Context() string {
	if a.Name != "" {
		return fmt.Sprintf("RelationalStore %#v", a.Name)
	}
	return "unnamed RelationalStore"
}

// DSL returns this object's DSL
func (sd *RelationalStoreDefinition) DSL() func() {
	return sd.DefinitionDSL
}

// Children returnsa slice of this objects children
func (sd RelationalStoreDefinition) Children() []design.ExternalDSLDefinition {
	var stores []design.ExternalDSLDefinition
	for _, s := range sd.RelationalModels {
		stores = append(stores, s)
	}
	return stores
}

// ResolveRelationships should be run after parsing the full
// Goa DSL, it correctly identifies foreign keys and other relationships that
// are stubbed out during initial parsing.
func (rs *RelationalStoreDefinition) ResolveRelationships() error {
	for name, model := range rs.RelationalModels {
		if len(model.belongsto) > 0 {
			for _, belong := range model.belongsto {
				belong = strings.Title(belong)
				for i, m := range rs.RelationalModels {
					if strings.Title(i) == belong {
						rs.RelationalModels[name].BelongsTo[i] = m
					}
				}
			}
		}
		if len(model.hasone) > 0 {
			for _, one := range model.hasone {
				one = strings.Title(one)
				for i, m := range rs.RelationalModels {
					if i == one {
						rs.RelationalModels[name].HasOne[i] = m
					}
				}
			}
		}
		if len(model.hasmany) > 0 {
			for _, many := range model.hasmany {
				many = strings.Title(many)
				for i, m := range rs.RelationalModels {
					if i == many {
						rs.RelationalModels[name].HasMany[i] = m
					}
				}
			}
		}
		if len(model.many2many) > 0 {
			for _, many := range model.many2many {
				s := strings.Split(many, ":")
				// 0 = field name
				// 1 = table name
				// 2 = Relation Model
				n := strings.Title(s[2])
				for i, m := range rs.RelationalModels {
					if i == n {
						m2m := &ManyToManyDefinition{
							LeftNamePlural:   inflect.Pluralize(model.Name),
							RightNamePlural:  inflect.Pluralize(m.Name),
							LeftName:         model.Name,
							RightName:        m.Name,
							Left:             model,
							Right:            m,
							RelationshipName: s[0],
							DatabaseField:    s[1],
						}
						rs.RelationalModels[name].ManyToMany[i] = m2m
					}
				}
			}
		}
	}
	return nil
}

// IterateModels runs an iterator function once per Model in the Store's model list.
func (rs *RelationalStoreDefinition) IterateModels(it ModelIterator) error {
	names := make([]string, len(rs.RelationalModels))
	i := 0
	for n := range rs.RelationalModels {
		names[i] = n
		i++
	}
	sort.Strings(names)
	for _, n := range names {
		if err := it(rs.RelationalModels[n]); err != nil {
			return err
		}
	}
	return nil
}
