package gorma

import (
	"fmt"
	"strings"

	"bitbucket.org/pkg/inflect"

	"github.com/bketelsen/gorma/gengorma"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

// deModel removes the word "Model" from the string.
func deModel(s string) string {
	return strings.Replace(s, "Model", "", -1)
}
func modifyModelDefinition(res *design.UserTypeDefinition) error {
	return nil
}
func modelOptions(res *design.UserTypeDefinition) ModelOptions {
	var options ModelOptions
	def := res.Definition()
	t := def.Type
	switch t.(type) {
	case design.Object:
		if _, ok := metaLookup(res.Metadata, gengorma.MetaCached); ok {
			options.Cached = true
		}
		if val, ok := metaLookup(res.Metadata, gengorma.MetaSQLTag); ok {
			options.SQLTag = val
		}
		if _, ok := metaLookup(res.Metadata, gengorma.MetaDynamicTableName); ok {
			options.DynamicTableName = true
		}
		if _, ok := metaLookup(res.Metadata, gengorma.MetaRoler); ok {
			options.Roler = true
		}
		if _, ok := metaLookup(res.Metadata, gengorma.MetaNoMedia); ok {
			options.NoMedia = true
		}
		if val, ok := metaLookup(res.Metadata, gengorma.MetaTableName); ok {
			options.TableName = val
		}

		return options
	default:
		panic("gorma bug: unexpected data structure type")
	}
}

// metaLookup is a helper function to lookup gorma-namespaced metadata keys in a
// case-insensitive way.
func metaLookup(md design.MetadataDefinition, hashtag string) (result string, ok bool) {
	needle := strings.ToLower(hashtag)
	for k, v := range md {
		k = strings.ToLower(k)
		if k == needle {
			return v, true
		}
	}

	return
}

type PrimaryKey struct {
	Field string
	Type  string
}

func primaryKeys(res *design.UserTypeDefinition) map[string]PrimaryKey {
	pks := make(map[string]PrimaryKey, 0)
	def := res.Definition()
	t := def.Type
	fmt.Println(res.TypeName)
	switch actual := t.(type) {
	case design.Object:
		for n := range actual {
			typ := "int"
			if gt, ok := metaLookup(actual[n].Metadata, gengorma.MetaGormTag); ok {
				if strings.Contains(gt, "primary_key") {
					pk := PrimaryKey{
						Field: n,
						Type:  typ, // TODO(BJK) support others
					}
					pks[n] = pk
				}
			}
			if n == "ID" || n == "Id" || n == "id" {
				pk := PrimaryKey{
					Field: n,
					Type:  typ, //TODO (BJK) support others
				}
				pks[n] = pk
			}
		}

	default:
		panic("gorma bug: expected data structure type")
	}
	if len(pks) == 0 {
		pks["id"] = PrimaryKey{Field: "id", Type: "int"}
	}
	return pks
}
func pkAttributes(pks map[string]PrimaryKey) string {
	var pkdefs []string
	for _, pk := range pks {
		def := fmt.Sprintf("%s %s", pk.Field, pk.Type)
		pkdefs = append(pkdefs, def)
	}

	return strings.Join(pkdefs, ",")
}
func pkWhere(pks map[string]PrimaryKey) string {

	var pkwhere []string
	for _, pk := range pks {
		def := fmt.Sprintf("%s = ?", pk.Field)
		pkwhere = append(pkwhere, def)
	}

	pkw := strings.Join(pkwhere, " and ")
	return pkw
}
func pkWhereFields(pks map[string]PrimaryKey) string {

	var pkwhere []string
	for _, pk := range pks {
		def := fmt.Sprintf("%s", pk.Field)
		pkwhere = append(pkwhere, def)
	}

	pkw := strings.Join(pkwhere, ",")
	return pkw
}
func pkUpdateFields(pks map[string]PrimaryKey) string {

	var pkwhere []string
	for _, pk := range pks {
		def := fmt.Sprintf("model.%s", codegen.Goify(pk.Field, true))
		pkwhere = append(pkwhere, def)
	}

	pkw := strings.Join(pkwhere, ",")
	return pkw
}

// lower returns the string in lowercase.
func lower(s string) string {
	return strings.ToLower(s)
}
func belongsTo(utd *design.UserTypeDefinition) []BelongsTo {
	var belongs []BelongsTo
	def := utd.Definition()
	t := def.Type
	switch actual := t.(type) {
	case design.Object:
		for n := range actual {
			if bt, ok := metaLookup(actual[n].Metadata, gengorma.MetaBelongsTo); ok {
				bel := BelongsTo{
					Parent:        n,
					DatabaseField: bt,
				}
				belongs = append(belongs, bel)
			}
		}

	default:
		panic("gorma bug: expected data structure type")
	}
	return belongs
}

/*
	Many2Many struct {
		Relation            string
		LowerRelation       string
		PluralRelation      string
		LowerPluralRelation string
		TableName           string
	}
*/
func many2Many(utd *design.UserTypeDefinition) []Many2Many {
	var m2m []Many2Many
	def := utd.Definition()
	t := def.Type
	switch actual := t.(type) {
	case design.Object:
		for n := range actual {
			if bt, ok := metaLookup(actual[n].Metadata, gengorma.MetaManyToMany); ok {
				vals := strings.Split(bt, ":")
				if len(vals) < 2 {
					panic("Invalid ManyToMany Definition")
				}
				m := Many2Many{
					Relation:            vals[0],
					LowerRelation:       strings.ToLower(vals[0]),
					PluralRelation:      inflect.Pluralize(vals[0]),
					LowerPluralRelation: strings.ToLower(inflect.Pluralize(vals[0])),
					TableName:           vals[1],
				}
				m2m = append(m2m, m)
			}
		}
	default:
		panic("gorma bug: expected data structure type")
	}
	//md.RequiredPackages[lower(deModel(parms[1]))] = true
	return m2m
}
