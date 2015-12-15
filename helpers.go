package gorma

import (
	"path/filepath"
	"strings"
	"unicode"

	"github.com/qor/inflection"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

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

// JSONSchemaDir is the path to the directory where the schema controller is generated.
func ModelDir() string {
	return filepath.Join(codegen.OutputDir, "models")
}

func DeModel(s string) string {
	return strings.Replace(s, "Model", "", -1)
}

func Lower(s string) string {
	return strings.ToLower(s)
}
func Upper(s string) string {
	return strings.ToUpper(s)
}

func StorageDefinition(res *design.UserTypeDefinition) string {

	var associations string
	if assoc, ok := res.Metadata["github.com/bketelsen/gorma#many2many"]; ok {
		children := strings.Split(assoc, ",")

		for _, child := range children {
			pieces := strings.Split(child, ":")
			associations = associations + "List" + pieces[0] + "(ctx *app.List" + strings.ToLower(pieces[0]) + DeModel(res.TypeName) + "Context) []" + pieces[1] + "\n"
			associations = associations + "Add" + pieces[1] + "(ctx *app.Add" + strings.ToLower(pieces[1]) + DeModel(res.TypeName) + "Context) (" + pieces[1] + " error)\n"
			associations = associations + "Delete" + pieces[1] + "(ctx *app.Delete" + strings.ToLower(pieces[1]) + DeModel(res.TypeName) + "Context) error \n"
		}
	}
	return associations
}
func IncludeForeignKey(res *design.UserTypeDefinition) string {
	if assoc, ok := res.Metadata["github.com/bketelsen/gorma#belongsto"]; ok {
		return assoc + "ID int\n"
	}
	return ""
}
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

func IncludeMany2Many(res *design.UserTypeDefinition) string {
	var associations string
	if assoc, ok := res.Metadata["github.com/bketelsen/gorma#many2many"]; ok {
		children := strings.Split(assoc, ",")

		for _, child := range children {
			pieces := strings.Split(child, ":")
			associations = associations + pieces[0] + "\t" + pieces[1] + "\t" + "`gorm:\"many2many:" + pieces[2] + ";\"`\n"
		}
	}
	return associations
}
func Authboss(res *design.UserTypeDefinition) string {
	if _, ok := res.Metadata["github.com/bketelsen/gorma#authboss"]; ok {
		fields := `	// Auth
	Email    string

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

func Split(s string, sep string) []string {

	return strings.Split(s, sep)
}

func MakeModelDef(s string, res *design.UserTypeDefinition) string {

	start := s[0:strings.Index(s, "{")+1] + "\n  	ID        int `gorm:\"primary_key\"`\nCreatedAt time.Time\nUpdatedAt time.Time\nDeletedAt *time.Time\n" + IncludeMany2Many(res) + IncludeForeignKey(res) + IncludeChildren(res) + Authboss(res) + s[strings.Index(s, "{")+2:]
	newstrings := make([]string, 0)
	chunks := strings.Split(start, "\n")
	// Good lord, shoot me for this hack - remove the ID field in the model if it exists
	for _, chunk := range chunks {
		var isId, isEmail, isAuthboss bool
		if strings.HasPrefix(chunk, "\tID ") {
			isId = true
		}
		if _, ok := res.Metadata["github.com/bketelsen/gorma#authboss"]; ok && strings.HasPrefix(chunk, "\tEmail") {
			isAuthboss = true
			isEmail = true
		}
		if isAuthboss {
			if !isEmail && !isId {
				newstrings = append(newstrings, chunk)
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
