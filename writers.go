package gorma

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/goagen/codegen"
)

type (
	// UserTypeTemplateData contains all the information used by the template to redner the
	// media types code.
	UserTypeTemplateData struct {
		APIDefinition *design.APIDefinition
		UserType      *RelationalModelDefinition
		DefaultPkg    string
		AppPkg        string
	}
	// UserTypesWriter generate code for a goa application user types.
	// User types are data structures defined in the DSL with "Type".
	UserTypesWriter struct {
		*codegen.SourceFile
		UserTypeTmpl *template.Template
	}
)

func fieldAssignmentModelToType(model *RelationalModelDefinition, ut *design.MediaTypeDefinition, mtype, utype string) string {
	//utPackage := "app"
	var fieldAssignments []string
	// type.Field = model.Field
	for fname, field := range model.RelationalFields {
		if field.Datatype == "" {
			continue
		}
		var mpointer, upointer bool
		mpointer = field.Nullable
		obj := ut.ToObject()
		definition := ut.Definition()
		for key := range obj {
			gfield := obj[key]
			if field.Underscore() == key || field.DatabaseFieldName == key {
				// this is our field
				if gfield.Type.IsObject() || definition.IsPrimitivePointer(key) {
					upointer = true
				} else {
					// set it explicity because we're reusing the same bool
					upointer = false
				}

				var prefix string
				if upointer && !mpointer {
					// ufield = &mfield
					prefix = "&"
				} else if mpointer && !upointer {
					// ufield = *mfield (rare if never?)
					prefix = "*"
				}

				fa := fmt.Sprintf("\t%s.%s = %s%s.%s", utype, codegen.Goify(key, true), prefix, mtype, fname)
				fieldAssignments = append(fieldAssignments, fa)
			}
		}
	}
	return strings.Join(fieldAssignments, "\n")
}

func fieldAssignmentTypeToModel(model *RelationalModelDefinition, ut *design.UserTypeDefinition, utype, mtype string) string {
	//utPackage := "app"
	var fieldAssignments []string
	// type.Field = model.Field
	for fname, field := range model.RelationalFields {
		var mpointer, upointer bool
		mpointer = field.Nullable
		obj := ut.ToObject()
		definition := ut.Definition()
		if field.Datatype == "" {
			continue
		}
		for key := range obj {
			gfield := obj[key]
			if field.Underscore() == key || field.DatabaseFieldName == key {
				// this is our field
				if gfield.Type.IsObject() || definition.IsPrimitivePointer(key) {
					upointer = true
				} else {
					// set it explicity because we're reusing the same bool
					upointer = false
				}

				var prefix string
				if upointer != mpointer {
					prefix = "*"
				}

				fa := fmt.Sprintf("\t%s.%s = %s%s.%s", mtype, fname, prefix, utype, codegen.Goify(key, true))
				fieldAssignments = append(fieldAssignments, fa)
			}
		}

	}
	return strings.Join(fieldAssignments, "\n")
}

// NewUserTypesWriter returns a contexts code writer.
// User types contain custom data structured defined in the DSL with "Type".
func NewUserTypesWriter(filename string) (*UserTypesWriter, error) {
	file, err := codegen.SourceFileFor(filename)
	if err != nil {
		return nil, err
	}
	return &UserTypesWriter{SourceFile: file}, nil
}

// Execute writes the code for the context types to the writer.
func (w *UserTypesWriter) Execute(data *UserTypeTemplateData) error {
	fm := make(map[string]interface{})
	fm["famt"] = fieldAssignmentModelToType
	fm["fatm"] = fieldAssignmentTypeToModel
	return w.ExecuteTemplate("types", userTypeT, fm, data)
}

// arrayAttribute returns the array element attribute definition.
func arrayAttribute(a *design.AttributeDefinition) *design.AttributeDefinition {
	return a.Type.(*design.Array).ElemType
}

const (
	// userTypeT generates the code for a user type.
	// template input: UserTypeTemplateData
	userTypeT = `{{$ut := .UserType}}{{$ap := .AppPkg}}// {{if $ut.Description}}{{$ut.Description}} {{end}}
{{$ut.StructDefinition}}
{{ if ne $ut.Alias "" }}
// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m {{$ut.Name}}) TableName() string {
	return "{{ $ut.Alias}}"
}{{end}}
// {{$ut.Name}}DB is the implementation of the storage interface for
// {{$ut.Name}}.
type {{$ut.Name}}DB struct {
	Db gorm.DB
	{{ if $ut.Cached }}cache *cache.Cache{{end}}
}
// New{{$ut.Name}}DB creates a new storage type.
func New{{$ut.Name}}DB(db gorm.DB) *{{$ut.Name}}DB {
	{{ if $ut.Cached }}return &{{$ut.Name}}DB{
		Db: db,
		cache: cache.New(5*time.Minute, 30*time.Second),
	}
	{{ else  }}return &{{$ut.Name}}DB{Db: db}{{ end  }}
}
// DB returns the underlying database.
func (m *{{$ut.Name}}DB) DB() interface{} {
	return &m.Db
}
{{ if $ut.Roler }}
// GetRole returns the value of the role field and satisfies the Roler interface.
func (m {{$ut.Name}}) GetRole() string {
	return {{$f := $ut.Fields.role}}{{if $f.Nullable}}*{{end}}m.Role
}
{{end}}

// {{$ut.Name}}Storage represents the storage interface.
type {{$ut.Name}}Storage interface {
	DB() interface{}
	List(ctx context.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}) []{{$ut.Name}}
	One(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.PKAttributes}}) ({{$ut.Name}}, error)
	Add(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.LowerName}} {{$ut.Name}}) ({{$ut.Name}}, error)
	Update(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.LowerName}} {{$ut.Name}}) (error)
	Delete(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{ $ut.PKAttributes}}) (error) 	{{$dtn:=$ut.DynamicTableName}}{{ range $idx, $bt := $ut.BelongsTo}}
	ListBy{{$bt.Name}}(ctx context.Context{{ if $dtn}}, tableName string{{ end }},{{$bt.LowerName}}_id int) []{{$ut.Name}}
	OneBy{{$bt.Name}}(ctx context.Context{{ if $dtn}}, tableName string{{ end }}, {{$bt.LowerName}}_id, id int) ({{$ut.Name}}, error){{end}}
	{{range $i, $m2m := $ut.ManyToMany}}
	List{{$m2m.RightNamePlural}}(context.Context, int) []{{$m2m.RightName}}
	Add{{$m2m.RightNamePlural}}(context.Context, int, int) (error)
	Delete{{$m2m.RightNamePlural}}(context.Context, int, int) error
	{{end}}
}

// CRUD Functions

// List returns an array of records.
func (m *{{$ut.Name}}DB) List(ctx context.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}) []{{$ut.Name}} {
	var objs []{{$ut.Name}}
	m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Find(&objs)
	return objs
}

// One returns a single record by ID.
func (m *{{$ut.Name}}DB) One(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.PKAttributes}}) ({{$ut.Name}}, error) {
	{{ if $ut.Cached }}//first attempt to retrieve from cache
	o,found := m.cache.Get(strconv.Itoa(id))
	if found {
		return o.({{$ut.Name}}), nil
	}
	// fallback to database if not found{{ end }}
	var obj {{$ut.Name}}{{ $l := len $ut.PrimaryKeys }}
	{{ if eq $l 1 }}
	err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Find(&obj, id).Error{{ else  }}err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Find(&obj).Where("{{$ut.PKWhere}}", {{$ut.PKWhereFields }}).Error{{ end }}
	{{ if $ut.Cached }} go m.cache.Set(strconv.Itoa(id), obj, cache.DefaultExpiration) {{ end }}
	return obj, err
}
// Add creates a new record.
func (m *{{$ut.Name}}DB) Add(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, model {{$ut.Name}}) ({{$ut.Name}}, error) {
	err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Create(&model).Error{{ if $ut.Cached }}
	go m.cache.Set(strconv.Itoa(model.ID), model, cache.DefaultExpiration) {{ end }}
	return model, err
}
// Update modifies a single record.
func (m *{{$ut.Name}}DB) Update(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, model {{$ut.Name}}) error {
	obj, err := m.One(ctx{{ if $ut.DynamicTableName }}, tableName{{ end }}, {{$ut.PKUpdateFields "model"}})
	if err != nil {
		return  err
	}
	err = m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Model(&obj).Updates(model).Error
	{{ if $ut.Cached }}go func(){
	obj, err := m.One(ctx, model.ID)
	if err == nil {
		m.cache.Set(strconv.Itoa(model.ID), obj, cache.DefaultExpiration)
	}
	}()
	{{ end }}
	return err
}
// Delete removes a single record.
func (m *{{$ut.Name}}DB) Delete(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.PKAttributes}})  error {
	var obj {{$ut.Name}}{{ $l := len $ut.PrimaryKeys }}
	{{ if eq $l 1 }}
	err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Delete(&obj, id).Error
	{{ else  }}err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Delete(&obj).Where("{{$ut.PKWhere}}", {{$ut.PKWhereFields}}).Error
	{{ end }}
	if err != nil {
		return  err
	}
	{{ if $ut.Cached }} go m.cache.Delete(strconv.Itoa(id)) {{ end }}
	return  nil
}
{{ range $idx, $bt := $ut.BelongsTo}}
// Belongs To Relationships

// {{$ut.Name}}FilterBy{{$bt.Name}} is a gorm filter for a Belongs To relationship.
func {{$ut.Name}}FilterBy{{$bt.Name}}(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if parentid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("{{$bt.LowerName}}_id = ?", parentid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}
// ListBy{{$bt.Name}} returns an array of associated {{$bt.Name}} models.
func (m *{{$ut.Name}}DB) ListBy{{$bt.Name}}(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, parentid int) []{{$ut.Name}} {
	var objs []{{$ut.Name}}
	m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Scopes({{$ut.Name}}FilterBy{{$bt.Name}}(parentid, &m.Db)).Find(&objs)
	return objs
}
// OneBy{{$bt.Name}} returns a single associated {{$bt.Name}} model.
func (m *{{$ut.Name}}DB) OneBy{{$bt.Name}}(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, parentid, {{ $ut.PKAttributes}}) ({{$ut.Name}}, error) {
	{{ if $ut.Cached }}//first attempt to retrieve from cache
	o,found := m.cache.Get(strconv.Itoa(id))
	if found {
		return o.({{$ut.Name}}), nil
	}
	// fallback to database if not found{{ end }}
	var obj {{$ut.Name}}
	err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Scopes({{$ut.Name}}FilterBy{{$bt.Name}}(parentid, &m.Db)).Find(&obj, id).Error
	{{ if $ut.Cached }} go m.cache.Set(strconv.Itoa(id), obj, cache.DefaultExpiration) {{ end }}
	return obj, err
}
{{end}}

{{ range $idx, $bt := $ut.ManyToMany}}
// Many To Many Relationships

// Delete{{goify $bt.RightName true}} removes a {{$bt.RightName}}/{{$bt.LeftName}} entry from the join table.
func (m *{{$ut.Name}}DB) Delete{{goify $bt.RightName true}}(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.LowerName}}ID,  {{$bt.LowerRightName}}ID int)  error {
	var obj {{$ut.Name}}
	obj.ID = {{$ut.LowerName}}ID
	var assoc {{$bt.RightName}}
	var err error
	assoc.ID = {{$bt.LowerRightName}}ID
	if err != nil {
		return err
	}
	err = m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Model(&obj).Association("{{$bt.RightNamePlural}}").Delete(assoc).Error
	if err != nil {
		return  err
	}
	return  nil
}
// Add{{goify $bt.RightName true}} creates a new {{$bt.RightName}}/{{$bt.LeftName}} entry in the join table.
func (m *{{$ut.Name}}DB) Add{{goify $bt.RightName true}}(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.LowerName}}ID, {{$bt.LowerRightName}}ID int) error {
	var {{$ut.LowerName}} {{$ut.Name}}
	{{$ut.LowerName}}.ID = {{$ut.LowerName}}ID
	var assoc {{$bt.RightName}}
	assoc.ID = {{$bt.LowerRightName}}ID
	err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Model(&{{$ut.LowerName}}).Association("{{$bt.RightNamePlural}}").Append(assoc).Error
	if err != nil {
		return  err
	}
	return  nil
}
// List{{goify $bt.RightName true}} returns a list of the {{$bt.RightName}} models related to this {{$bt.LeftName}}.
func (m *{{$ut.Name}}DB) List{{goify $bt.RightName true}}(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.LowerName}}ID int)  []{{$bt.RightName}} {
	var list []{{$bt.RightName}}
	var obj {{$ut.Name}}
	obj.ID = {{$ut.LowerName}}ID
	m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Model(&obj).Association("{{$bt.RightNamePlural}}").Find(&list)
	return  list
}
{{end}}
{{ range $idx, $bt := $ut.BelongsTo}}
// Filter{{$ut.Name}}By{{$bt.Name}} iterates a list and returns only those with the foreign key provided.
func Filter{{$ut.Name}}By{{$bt.Name}}(parent *int, list []{{$ut.Name}}) []{{$ut.Name}} {
	var filtered []{{$ut.Name}}
	for _,o := range list {
		if o.{{$bt.Name}}ID == int(*parent) {
			filtered = append(filtered,o)
		}
	}
	return filtered
}
{{end}}

{{ if gt (len $ut.RenderTo) 0 }}// Useful conversion functions
{{range  $tcd := $ut.RenderTo }}{{if $tcd.SupportsNoVersion}}{{ $tcdTypeName := goify $tcd.TypeName true }}
// To{{$tcdTypeName}} converts a model {{$ut.Name}} to an app {{$tcdTypeName}}.
func (m *{{$ut.Name}}) To{{$tcdTypeName}}() {{$ap}}.{{$tcdTypeName}} {
	payload := {{$ap}}.{{$tcdTypeName}}{}
	{{ famt $ut $tcd "m" "payload"}}
	return payload
}
{{end}}{{end}}{{end}}

{{ if gt (len $ut.RenderTo) 0 }}{{range  $tcd := $ut.RenderTo }}{{ range $version := $tcd.APIVersions }}{{ $tcdTypeName := goify $tcd.TypeName true }}
// To{{if eq $version ""}}{{title $ap}}{{else}}{{title $version}}{{end}}{{$tcdTypeName}} converts to goa types.
func (m *{{$ut.Name}}) To{{if eq $version ""}}{{title $ap}}{{else}}{{title $version}}{{end}}{{$tcdTypeName}}() {{if eq $version ""}}{{$ap}}{{else}}{{$version}}{{end}}.{{$tcdTypeName}} {
	payload := {{if eq $version ""}}{{$ap}}{{else}}{{$version}}{{end}}.{{$tcdTypeName}}{}
	{{ famt $ut $tcd "m" "payload"}}
	return payload
}
{{end}}{{end}}{{end}}

{{ range $tcd := $ut.BuiltFrom}}{{ range $version := $tcd.APIVersions }} // v{{$version}}
// Convert from	{{if eq $version ""}}default version{{else}}Version {{$version}}{{end}} {{$tcd.TypeName}} to {{$ut.Name}}.
func {{$ut.Name}}From{{$version}}{{$tcd.Name}}(t {{if ne $version ""}}{{$version}}.{{else}}{{$ap}}.{{end}}{{$tcd.Name}}) {{$ut.Name}} {
	{{$ut.LowerName}} := {{$ut.Name}}{}
	{{ fatm $ut $tcd.Type "m" $ut.LowerName}}
	return {{$ut.LowerName}}
}
{{end}}{{end}}
{{ range $tcd := $ut.BuiltFrom}}
// Convert from	default version {{$tcd.TypeName}} to {{$ut.Name}}.
func {{$ut.Name}}From{{$tcd.TypeName}}(t {{$ap}}.{{$tcd.TypeName}}) {{$ut.Name}} {
	{{$ut.LowerName}} := {{$ut.Name}}{}
	{{ fatm $ut $tcd "t" $ut.LowerName}}
	return {{$ut.LowerName}}
}

{{end}}
`
)
