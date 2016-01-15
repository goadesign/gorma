package gorma

import (
	"sort"
	"strings"

	"bitbucket.org/pkg/inflect"
)

// NewRelationalStore returns an initialzed RelationalStore
func NewRelationalStore() *RelationalStore {
	return &RelationalStore{
		Models: make(map[string]*RelationalModel),
	}
}

func (rs *RelationalStore) ResolveRelationships() error {
	for name, model := range rs.Models {
		if len(model.belongsto) > 0 {
			for _, belong := range model.belongsto {
				belong = strings.Title(belong)
				for i, m := range rs.Models {
					if deModel(strings.Title(i)) == belong {
						rs.Models[name].BelongsTo[i] = m
					}
				}
			}
		}
		if len(model.hasone) > 0 {
			for _, one := range model.hasone {
				one = strings.Title(one)
				for i, m := range rs.Models {
					if i == one {
						rs.Models[name].HasOne[i] = m
					}
				}
			}
		}
		if len(model.hasmany) > 0 {
			for _, many := range model.hasmany {
				many = strings.Title(many)
				for i, m := range rs.Models {
					if i == many {
						rs.Models[name].HasMany[i] = m
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
				for i, m := range rs.Models {
					if i == n {
						m2m := &ManyToMany{
							LeftNamePlural:   inflect.Pluralize(model.Name),
							RightNamePlural:  inflect.Pluralize(m.Name),
							LeftName:         model.Name,
							RightName:        m.Name,
							RelationshipName: s[0],
							DatabaseField:    s[1],
						}
						rs.Models[name].ManyToMany[i] = m2m
					}
				}
			}
		}
	}
	return nil
}

func (rs *RelationalStore) IterateModels(it ModelIterator) error {
	names := make([]string, len(rs.Models))
	i := 0
	for n := range rs.Models {
		names[i] = n
		i++
	}
	sort.Strings(names)
	for _, n := range names {
		if err := it(rs.Models[n]); err != nil {
			return err
		}
	}
	return nil
}
