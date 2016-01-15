package gorma

import (
	"sort"
	"strings"
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
