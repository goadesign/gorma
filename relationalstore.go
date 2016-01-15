package gorma

import "strings"

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
