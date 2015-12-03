package gorma

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"text/template"

	"github.com/jinzhu/inflection"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Generator is the application code generator.
type Generator struct {
	genfiles []string
}

// ModelWriter generate code for a goa application media types.
// Media types are data structures used to render the response bodies.
type ModelWriter struct {
	*codegen.GoGenerator
	ModelTmpl *template.Template
}

type UserTypeMetadata struct {
	Filter string
}

// NewModelWriter returns a contexts code writer.
// Media types contain the data used to render response bodies.
func NewModelWriter(filename string) (*ModelWriter, error) {
	cw := codegen.NewGoGenerator(filename)
	funcMap := cw.FuncMap
	funcMap["gotypedef"] = codegen.GoTypeDef
	funcMap["gotyperef"] = codegen.GoTypeRef
	funcMap["goify"] = codegen.Goify
	funcMap["gotypename"] = codegen.GoTypeName
	funcMap["gonative"] = codegen.GoNativeType
	funcMap["typeUnmarshaler"] = codegen.TypeUnmarshaler
	funcMap["typeMarshaler"] = codegen.MediaTypeMarshaler
	funcMap["recursiveValidate"] = codegen.RecursiveChecker
	funcMap["tempvar"] = codegen.Tempvar
	funcMap["demodel"] = DeModel
	funcMap["modeldef"] = MakeModelDef
	funcMap["snake"] = CamelToSnake

	modelTmpl, err := template.New("models").Funcs(funcMap).Parse(modelTmpl)
	if err != nil {
		return nil, err
	}
	w := ModelWriter{
		GoGenerator: cw,
		ModelTmpl:   modelTmpl,
	}
	return &w, nil
}

// Execute writes the code for the context types to the writer.
func (w *ModelWriter) Execute(mt *design.UserTypeDefinition) error {
	return w.ModelTmpl.Execute(w, mt)
}

// Generate is the generator entry point called by the meta generator.
func Generate(api *design.APIDefinition) ([]string, error) {
	g, err := NewGenerator()
	if err != nil {
		return nil, err
	}
	return g.Generate(api)
}

// NewGenerator returns the application code generator.
func NewGenerator() (*Generator, error) {
	return new(Generator), nil
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

func IncludeForeignKey(res *design.UserTypeDefinition) string {
	if assoc, ok := res.Metadata["github.com/bketelsen/gorma#belongsto"]; ok {
		return assoc + "ID uint\n"
	}
	return ""
}
func IncludeChildren(res *design.UserTypeDefinition) string {
	if assoc, ok := res.Metadata["github.com/bketelsen/gorma#hasmany"]; ok {
		children := strings.Split(assoc, ",")
		var associations string
		for _, child := range children {
			associations = associations + inflection.Plural(child) + " []" + child + "\n"
		}
		return associations
	}
	return ""
}
func MakeModelDef(s string, res *design.UserTypeDefinition) string {
	return s[0:strings.Index(s, "{")+1] + "\n  gorm.Model\n" + IncludeForeignKey(res) + IncludeChildren(res) + s[strings.Index(s, "{")+2:]
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

// Generate produces the skeleton main.
func (g *Generator) Generate(api *design.APIDefinition) ([]string, error) {
	os.RemoveAll(ModelDir())
	os.MkdirAll(ModelDir(), 0755)
	app := kingpin.New("Model generator", "model generator")
	codegen.RegisterFlags(app)
	_, err := app.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
	imp, err := filepath.Rel(filepath.Join(os.Getenv("GOPATH"), "src"), codegen.OutputDir)
	if err != nil {
		return nil, err
	}
	imp = filepath.Join(imp, "app")
	imports := []*codegen.ImportSpec{
		codegen.SimpleImport(imp),
		codegen.SimpleImport("github.com/jinzhu/gorm"),
		codegen.SimpleImport("github.com/jinzhu/copier"),
	}

	title := fmt.Sprintf("%s: Models", api.Name)
	filename := filepath.Join(ModelDir(), "models.go")
	mtw, err := NewModelWriter(filename)
	if err != nil {
		panic(err)
	}
	mtw.WriteHeader(title, "models", imports)
	err = api.IterateUserTypes(func(res *design.UserTypeDefinition) error {
		if res.Type.IsObject() {

			if md, ok := res.Metadata["github.com/bketelsen/gorma"]; ok && md == "Model" {
				fmt.Println("Found Gorma Metadata:", md)
				if err != nil {
					panic(err)
				}
				err = mtw.Execute(res)
				if err != nil {
					g.Cleanup()
					return err
				}
			}
		}

		return nil
	})
	if err := mtw.FormatCode(); err != nil {
		g.Cleanup()
		return nil, err
	}
	if err != nil {
		g.genfiles = append(g.genfiles, filename)
	}
	return g.genfiles, err
}

// Cleanup removes all the files generated by this generator during the last invokation of Generate.
func (g *Generator) Cleanup() {
	for _, f := range g.genfiles {
		os.Remove(f)
	}
	g.genfiles = nil
}

const modelTmpl = `// {{if .Description}}{{.Description}}{{else}}app.{{gotypename . 0}} storage type{{end}}
// Identifier: {{ $typeName :=  gotypename . 0}}{{$typeName := demodel $typeName}}
{{$td := gotypedef . 0 true false}}type {{$typeName}} {{modeldef $td .}}
{{ $belongsto := index .Metadata "github.com/bketelsen/gorma#belongsto" }}
func {{$typeName}}FromCreatePayload(ctx *app.Create{{demodel $typeName}}Context) {{$typeName}} {
	payload := ctx.Payload
	m := {{$typeName}}{}
	copier.Copy(&m, payload)
	{{ if ne $belongsto "" }} m.{{ $belongsto }}ID=uint(ctx.{{ demodel $belongsto }}ID){{end}}
	return m
}

func {{$typeName}}FromUpdatePayload(ctx *app.Update{{demodel $typeName}}Context) {{$typeName}} {
	payload := ctx.Payload
	m := {{$typeName}}{}
	copier.Copy(&m, payload)
	return m
}
func (m {{$typeName}}) ToApp() *app.{{demodel $typeName}} {
	target := app.{{demodel $typeName}}{}
	copier.Copy(&target, &m)
	return &target 
}



type {{$typeName}}Storage interface {
	List(ctx *app.List{{demodel $typeName}}Context) []{{$typeName}}
	Get(ctx *app.Show{{demodel $typeName }}Context) ({{$typeName}}, error)
	Add(ctx *app.Create{{demodel $typeName}}Context) ({{$typeName}}, error)
	Update(ctx *app.Update{{demodel $typeName}}Context) (error)
	Delete(ctx *app.Delete{{demodel $typeName}}Context) (error)
}

type {{$typeName}}DB struct {
	DB gorm.DB
}
{{ if ne $belongsto "" }}
// would prefer to just pass a context in here, but they're all different, so can't
func {{$typeName}}Filter(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if parentid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("{{ snake $belongsto }}_id = ?", parentid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}{{end}}
func New{{$typeName}}DB(db gorm.DB) *{{$typeName}}DB {
	return &{{$typeName}}DB{DB: db}
}

func (m *{{$typeName}}DB) List(ctx *app.List{{demodel $typeName}}Context) []{{$typeName}} {

	var objs []{{$typeName}}
    {{ if ne $belongsto "" }}m.DB.Scopes({{$typeName}}Filter(ctx.{{demodel $belongsto}}ID, &m.DB)).Find(&objs){{ else }} m.DB.Find(&objs) {{end}}
	return objs
}

func (m *{{$typeName}}DB) Get(ctx *app.Show{{demodel $typeName}}Context) ({{$typeName}}, error) {

	var obj {{$typeName}}

	err := m.DB.Find(&obj, ctx.{{demodel $typeName}}ID).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return obj, err
}

func (m *{{$typeName}}DB) Add(ctx *app.Create{{demodel $typeName}}Context) ({{$typeName}}, error) {
	model := {{$typeName}}FromCreatePayload(ctx)
	err := m.DB.Create(&model).Error
	return model, err
}
func (m *{{$typeName}}DB) Update(ctx *app.Update{{demodel $typeName}}Context) error {
	getCtx, err := app.NewShow{{demodel $typeName}}Context(ctx.Context)
	if err != nil {
		return  err
	}
	obj, err := m.Get(getCtx)
	if err != nil {
		return  err
	}
	err = m.DB.Model(&obj).Updates({{$typeName}}FromUpdatePayload(ctx)).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return err
}
func (m *{{$typeName}}DB) Delete(ctx *app.Delete{{demodel $typeName}}Context)  error {
	var obj {{$typeName}}
	err := m.DB.Delete(&obj, ctx.{{demodel $typeName}}ID).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return  err
	}
	return  nil
}



type Mock{{$typeName}}Storage struct {
	{{$typeName}}List  map[uint]{{$typeName}}
	nextID uint
	mut sync.Mutex
}

func NewMock{{$typeName}}Storage() *Mock{{$typeName}}Storage {
	ml := make(map[uint]{{$typeName}}, 0)
	return &Mock{{$typeName}}Storage{ {{$typeName}}List: ml}
}

func (db *Mock{{$typeName}}Storage) List(ctx *app.List{{demodel $typeName}}Context) []{{$typeName}} {
	var list []{{$typeName}} = make([]{{$typeName}}, 0)
	for _, v := range db.{{$typeName}}List {
		list = append(list, v)
	}
	return list
}

func (db *Mock{{$typeName}}Storage) Get(ctx *app.Show{{demodel $typeName}}Context) ({{$typeName}}, error) {

	var obj {{$typeName}}

	obj, ok := db.{{$typeName}}List[uint(ctx.{{demodel $typeName}}ID)]
	if ok {
		return obj, nil
	} else {
		return obj, errors.New("{{$typeName}} does not exist")
	}
}

func (db *Mock{{$typeName}}Storage) Add(ctx *app.Create{{demodel $typeName}}Context)  ({{$typeName}}, error) {
	u := {{$typeName}}FromCreatePayload(ctx)
	db.mut.Lock()
	db.nextID = db.nextID + 1
	u.ID = db.nextID
	db.mut.Unlock()

	db.{{$typeName}}List[u.ID] = u
	return u, nil
}

func (db *Mock{{$typeName}}Storage) Update(ctx *app.Update{{demodel $typeName}}Context) error {
	id := uint(ctx.{{demodel $typeName}}ID)
	_, ok := db.{{$typeName}}List[id]
	if ok {
		db.{{$typeName}}List[id] = {{$typeName}}FromUpdatePayload(ctx)
		return  nil
	} else {
		return errors.New("{{$typeName}} does not exist")
	}
}

func (db *Mock{{$typeName}}Storage) Delete(ctx *app.Delete{{demodel $typeName}}Context)  error {
	_, ok := db.{{$typeName}}List[uint(ctx.{{demodel $typeName}}ID)]
	if ok {
		delete(db.{{$typeName}}List, uint(ctx.{{demodel $typeName}}ID))
		return  nil
	} else {
		return  errors.New("Could not delete this user")
	}
}
`
