package gorma

import (
	"bytes"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/qor/inflection"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

// TitleCase converts a string to Title case
func TitleCase(s string) string {
	return strings.Title(s)
}

// CamelToSnake converts a given string to snake case
func CamelToSnake(s string) string {
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

// ModelDir is the path to the directory where the schema controller is generated.
func ModelDir() string {
	return filepath.Join(codegen.OutputDir, "models")
}

// DeModel remove the word "Model" from the string
func DeModel(s string) string {
	return strings.Replace(s, "Model", "", -1)
}

// Lower returns the string in lowercase
func Lower(s string) string {
	return strings.ToLower(s)
}

// Upper returns the string in upper case
func Upper(s string) string {
	return strings.ToUpper(s)
}

// StorageDefinition creates the storage interface that will be used
// in place of a concrete type for testability
func StorageDefinition(res *design.UserTypeDefinition) string {

	var associations string
	if assoc, ok := res.Metadata["github.com/bketelsen/gorma#many2many"]; ok {
		children := strings.Split(assoc, ",")

		for _, child := range children {
			pieces := strings.Split(child, ":")
			associations = associations + "List" + pieces[0] + "(context.Context, int) []" + pieces[1] + "\n"
			associations = associations + "Add" + pieces[1] + "(context.Context, int, int) (error)\n"
			associations = associations + "Delete" + pieces[1] + "(context.Context, int, int) error \n"
		}
	}
	return associations
}

// IncludeForeignKey adds foreign key relations to the struct being
// generated
func IncludeForeignKey(res *design.UserTypeDefinition) string {
	var associations string
	if assoc, ok := res.Metadata["github.com/bketelsen/gorma#belongsto"]; ok {
		children := strings.Split(assoc, ",")

		for _, child := range children {
			associations = associations + child + "ID int\n"

		}
	}
	return associations
}

// PKTag adds the pk tag to the model definition
func PKTag(res *design.UserTypeDefinition) string {
	var tag string
	if metatag, ok := res.Metadata["github.com/bketelsen/gorma#gormpktag"]; ok {
		tag = fmt.Sprintf(" `gorm:\"%s\"`\n", metatag)

	} else {
		tag = "`gorm:\"primary_key\"`\n"
	}
	return tag
}

// Plural returns the plural version of a word
func Plural(s string) string {
	return inflection.Plural(s)
}

// IncludeChildren adds the fields to a struct represented
// in a has-many relationship
func IncludeChildren(res *design.UserTypeDefinition) string {
	var associations string
	if assoc, ok := res.Metadata["github.com/bketelsen/gorma#hasmany"]; ok {
		children := strings.Split(assoc, ",")

		for _, child := range children {
			associations = associations + inflection.Plural(child) + " []" + child + "\n"
		}
	}
	if assoc, ok := res.Metadata["github.com/bketelsen/gorma#hasone"]; ok {
		children := strings.Split(assoc, ",")
		for _, child := range children {
			associations = associations + child + " " + child + "\n"
			associations = associations + child + "ID " + "*sql.NullInt64\n"
		}
	}
	return associations
}

// IncludeMany2Many returns the appropriate struct tags
// for a m2m relationship in gorm
func IncludeMany2Many(res *design.UserTypeDefinition) string {
	var associations string
	if assoc, ok := res.Metadata["github.com/bketelsen/gorma#many2many"]; ok {
		children := strings.Split(assoc, ",")

		for _, child := range children {
			pieces := strings.Split(child, ":")
			associations = associations + pieces[0] + "\t []" + pieces[1] + "\t" + "`gorm:\"many2many:" + pieces[2] + ";\"`\n"
		}
	}
	return associations
}

// Authboss returns the tags required to implement authboss storage
// currently experimental and quite unfinished
func Authboss(res *design.UserTypeDefinition) string {
	if _, ok := res.Metadata["github.com/bketelsen/gorma#authboss"]; ok {
		fields := `	// Auth
	Password string

	// OAuth2
	Oauth2Uid      string
	Oauth2Provider string
	Oauth2Token    string
	Oauth2Refresh  string
	Oauth2Expiry   time.Time

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Lock
	AttemptNumber int64
	AttemptTime   time.Time
	Locked        time.Time

	// Recover
	RecoverToken       string
	RecoverTokenExpiry time.Time 
	`
		return fields
	}
	return ""
}

// Split splits a string by separater `sep`
func Split(s string, sep string) []string {

	return strings.Split(s, sep)
}

// GoTypeDef returns the Go code that defines a Go type which matches the data structure
// definition (the part that comes after `type foo`).
// tabs indicates the number of tab character(s) used to tabulate the definition however the first
// line is never indented.
// jsonTags controls whether to produce json tags.
// inner indicates whether to prefix the struct of an attribute of type object with *.
// copied from GOA
func GoTypeDef(ds design.DataStructure, tabs int, jsonTags, inner bool) string {
	return godef(ds, tabs, jsonTags, inner, false)
}

// godef is the common implementation for both GoTypeDef and GoResDef.
// The only difference between the two is how the type names for fields that refer to a media type
// is generated: GoTypeDef uses the type name but GoResDef uses the underlying resource name if the
// type is a media type that corresponds to the canonical representation of a resource.
// copied from GOA and modified
func godef(ds design.DataStructure, tabs int, jsonTags, inner, res bool) string {
	var buffer bytes.Buffer
	def := ds.Definition()
	t := def.Type
	switch actual := t.(type) {
	case design.Primitive:
		return codegen.GoTypeName(t, tabs)
	case *design.Array:
		return "[]" + godef(actual.ElemType, tabs, jsonTags, true, res)
	case *design.Hash:
		keyDef := godef(actual.KeyType, tabs, jsonTags, true, res)
		elemDef := godef(actual.ElemType, tabs, jsonTags, true, res)
		return fmt.Sprintf("map[%s]%s", keyDef, elemDef)
	case design.Object:
		if inner {
			buffer.WriteByte('*')
		}
		buffer.WriteString("struct {\n")
		keys := make([]string, len(actual))
		i := 0
		for n := range actual {
			keys[i] = n
			i++
		}
		sort.Strings(keys)
		for _, name := range keys {
			codegen.WriteTabs(&buffer, tabs+1)
			typedef := godef(actual[name], tabs+1, jsonTags, true, res)
			fname := codegen.Goify(name, true)
			var tags string
			if jsonTags {
				var omit string
				var gorm, sql string
				if !def.IsRequired(name) {
					omit = ",omitempty"
				}
				if val, ok := actual[name].Metadata["github.com/bketelsen/gorma#gormtag"]; ok {
					gorm = fmt.Sprintf(" gorm:\"%s\"", val)
				}
				if val, ok := actual[name].Metadata["github.com/bketelsen/gorma#sqltag"]; ok {
					sql = fmt.Sprintf(" sql:\"%s\"", val)
				}
				tags = fmt.Sprintf(" `json:\"%s%s\"%s%s`", name, omit, gorm, sql)
			}
			desc := actual[name].Description
			if desc != "" {
				desc = fmt.Sprintf("// %s\n", desc)
			}
			buffer.WriteString(fmt.Sprintf("%s%s %s%s\n", desc, fname, typedef, tags))
		}
		codegen.WriteTabs(&buffer, tabs)
		buffer.WriteString("}")
		return buffer.String()
	case *design.UserTypeDefinition:
		name := codegen.GoTypeName(actual, tabs)
		if actual.Type.IsObject() {
			return "*" + name
		}
		return name
	case *design.MediaTypeDefinition:
		if res && actual.Resource != nil {
			return "*" + codegen.Goify(actual.Resource.Name, true)
		}
		name := codegen.GoTypeName(actual, tabs)
		if actual.Type.IsObject() {
			return "*" + name
		}
		return name
	default:
		panic("goa bug: unknown data structure type")
	}
}

// Timestamps returns the timestamp fields if "skipts" isn't set
func TimeStamps(res *design.UserTypeDefinition) string {
	var ts string
	if _, ok := res.Metadata["github.com/bketelsen/gorma#skipts"]; ok {
		ts = ""
	} else {
		ts = "CreatedAt time.Time\nUpdatedAt time.Time\nDeletedAt *time.Time\n"
	}
	return ts
}

// MakeModelDef is the main function to create a struct definition
func MakeModelDef(s string, res *design.UserTypeDefinition) string {

	start := s[0:strings.Index(s, "{")+1] + "\n  	ID        int " + PKTag(res) + TimeStamps(res) + IncludeMany2Many(res) + IncludeForeignKey(res) + IncludeChildren(res) + Authboss(res) + s[strings.Index(s, "{")+2:]
	newstrings := make([]string, 0)
	chunks := strings.Split(start, "\n")
	// Good lord, shoot me for this hack - remove the ID field in the model if it exists
	for _, chunk := range chunks {
		var didEmail, isId, isAuthboss bool
		if strings.HasPrefix(chunk, "\tID ") {
			isId = true
		}
		if _, ok := res.Metadata["github.com/bketelsen/gorma#authboss"]; ok && strings.HasPrefix(chunk, "\tEmail") {
			isAuthboss = true
		}
		if isAuthboss {
			if !didEmail && !isId {
				newstrings = append(newstrings, chunk)
				didEmail = true
			}
		} else {
			if !isId {
				newstrings = append(newstrings, chunk)
			}
		}

	}
	return strings.Join(newstrings, "\n")
}

// Is c an ASCII lower-case letter?
func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func unexport(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
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
