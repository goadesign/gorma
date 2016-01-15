package gorma

import (
	"fmt"
	"strings"
	"unicode"

	"bitbucket.org/pkg/inflect"

	"github.com/bketelsen/gorma/gengorma"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

const META_NAMESPACE = "github.com/bketelsen/gorma"

// deModel removes the word "Model" from the string.
func deModel(s string) string {
	return strings.Replace(s, "Model", "", -1)
}
func modifyModelDefinition(res *design.UserTypeDefinition) error {
	return nil
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
func modelMetadata(ad *design.AttributeDefinition) bool {
	needle := strings.ToLower(META_NAMESPACE)
	for k, v := range ad.Metadata {
		k = strings.ToLower(k)
		v = strings.ToLower(v)
		if k == needle {
			if v == "model" {
				return true
			}
		}
	}
	return false
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
func snakeToCamel(s string) string {
	words := strings.Split(s, "_")
	for _, w := range words {
		w = strings.Title(w)
	}
	return strings.Join(words, "")
}

// camelToSnake converts a given string to snake case.
func camelToSnake(s string) string {
	var result string
	var words []string
	var lastPos int
	rs := []rune(s)

	for i := 0; i < len(rs); i++ {
		if i > 0 && unicode.IsUpper(rs[i]) {
			if initialism := startsWithInitialism(s[lastPos:]); initialism != "" {
				words = append(words, initialism)

				i += len(initialism) - 1
				lastPos = i
				continue
			}

			words = append(words, s[lastPos:i])
			lastPos = i
		}
	}

	// append the last word
	if s[lastPos:] != "" {
		words = append(words, s[lastPos:])
	}

	for k, word := range words {
		if k > 0 {
			result += "_"
		}

		result += strings.ToLower(word)
	}

	return result
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
					Parent:        strings.Replace(n, "ID", "", -1),
					DatabaseField: camelToSnake(bt),
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

func storageDef(res *design.UserTypeDefinition) string {
	var associations string
	def := res.Definition()
	t := def.Type
	switch actual := t.(type) {
	case design.Object:
		for n := range actual {
			if assoc, ok := metaLookup(actual[n].Metadata, gengorma.MetaManyToMany); ok {
				vals := strings.Split(assoc, ":")
				if len(vals) < 2 {
					panic("Invalid ManyToMany Definition")
				}
				associations = associations + "List" + n + "(context.Context, int) []" + lower(vals[1]) + "." + vals[1] + "\n"
				associations = associations + "Add" + vals[0] + "(context.Context, int, int) (error)\n"
				associations = associations + "Delete" + vals[0] + "(context.Context, int, int) error \n"
			}
		}
	}

	return associations
}
func hasUserType(action *design.ActionDefinition) bool {
	if action.Payload != nil {
		a := action.Payload.Definition()
		return modelMetadata(a)
	}
	return false
}

func versionize(s string) string {
	if s == "app" {
		return "Default"
	}
	return strings.Title(s)

}

// titleCase converts a string to Title case.
func titleCase(s string) string {
	return strings.Title(s)
}

// startsWithInitialism returns the initialism if the given string begins with it
func startsWithInitialism(s string) string {
	var initialism string
	// the longest initialism is 5 char, the shortest 2
	for i := 1; i <= 5; i++ {
		if len(s) > i-1 && commonInitialisms[s[:i]] {
			initialism = s[:i]
		}
	}
	return initialism
}

// commonInitialisms, taken from
// https://github.com/golang/lint/blob/3d26dc39376c307203d3a221bada26816b3073cf/lint.go#L482
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
}
