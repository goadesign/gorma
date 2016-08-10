package gorma

import (
	"fmt"
	"sort"
	"strings"
	"text/template"

	"bitbucket.org/pkg/inflect"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/goagen/codegen"
	"github.com/kr/pretty"
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
		UserTypeTmpl   *template.Template
		UserHelperTmpl *template.Template
	}

	// UserHelperWriter generate code for a goa application user types.
	// User types are data structures defined in the DSL with "Type".
	UserHelperWriter struct {
		*codegen.SourceFile
		UserHelperTmpl *template.Template
	}
)

func fieldAssignmentPayloadToModel(model *RelationalModelDefinition, ut *design.UserTypeDefinition, verpkg, v, mtype, utype string) string {
	// Get a sortable slice of field names
	var keys []string
	for k := range model.RelationalFields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var fieldAssignments []string
	for _, fname := range keys {
		field := model.RelationalFields[fname]

		var mpointer, upointer bool
		mpointer = field.Nullable
		obj := ut.Type.ToObject()
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
					// set it explicitly because we're reusing the same bool
					upointer = false
				}

				prefix := ""
				if upointer && !mpointer {
					// ufield = &mfield
					prefix = "*"
				} else if mpointer && !upointer {
					// ufield = *mfield (rare if never?)
					prefix = "&"
				} else if !upointer && !mpointer {
					prefix = ""
				}

				if upointer {
					ifa := fmt.Sprintf("if %s.%s != nil {", v, codegen.Goify(key, true))
					fieldAssignments = append(fieldAssignments, ifa)
				}

				fa := fmt.Sprintf("\t%s.%s = %s%s.%s", utype, fname, prefix, v, codegen.Goify(key, true))
				fieldAssignments = append(fieldAssignments, fa)

				if upointer {
					ifa := fmt.Sprintf("}")
					fieldAssignments = append(fieldAssignments, ifa)
				}
			}
		}
	}
	return strings.Join(fieldAssignments, "\n")
}

func fieldAssignmentModelToType(model *RelationalModelDefinition, ut *design.ViewDefinition, v, mtype, utype string) string {
	tmp := 1
	var fieldAssignments []string

	if !strings.Contains(ut.Name, "link") {
		for ln, lnd := range ut.Parent.Links {
			ln = codegen.Goify(ln, true)
			s := inflect.Singularize(ln)
			var ifb string
			if lnd.MediaType().IsArray() {
				mt := codegen.Goify(lnd.MediaType().ToArray().ElemType.Type.(*design.MediaTypeDefinition).TypeName, true) + "LinkCollection"
				fa := make([]string, 4)
				fa[0] = fmt.Sprintf("tmp%d := make(app.%s, len(%s.%s))", tmp, mt, v, ln)
				fa[1] = fmt.Sprintf("for i, elem := range %s.%s {", v, ln)
				fa[2] = fmt.Sprintf("	tmp%d[i] = elem.%sTo%sLink()", tmp, s, s)
				fa[3] = fmt.Sprintf("}")
				ifb = strings.Join(fa, "\n")
			} else {
				ifb = fmt.Sprintf("tmp%d := %s.%s.%sTo%sLink()", tmp, v, ln, s, s)
			}

			fieldAssignments = append(fieldAssignments, ifb)
			ifd := fmt.Sprintf("%s.Links = &app.%sLinks{%s: tmp%d}", utype, codegen.Goify(utype, true), codegen.Goify(ln, true), tmp)
			fieldAssignments = append(fieldAssignments, ifd)
			tmp++
		}
	}

	// Get a sortable slice of field names
	var keys []string
	for k := range model.RelationalFields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, fname := range keys {
		field := model.RelationalFields[fname]

		var mpointer, upointer bool
		mpointer = field.Nullable
		obj := ut.Type.ToObject()
		definition := ut.Parent.Definition()

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
					// set it explicitly because we're reusing the same bool
					upointer = false
				}

				if field.Datatype == HasOne {
					fa := fmt.Sprintf("%s.%s = %s.%s.%sTo%s()", utype, codegen.Goify(field.FieldName, true), v, codegen.Goify(field.FieldName, true), codegen.Goify(field.FieldName, true), codegen.Goify(field.FieldName, true))
					fieldAssignments = append(fieldAssignments, fa)
					continue
				}

				prefix := ""
				if upointer && !mpointer {
					// ufield = &mfield
					prefix = "&"
				} else if mpointer && !upointer {
					// ufield = *mfield (rare if never?)
					prefix = "*"
				} else if !upointer && !mpointer {
					prefix = ""
				}
				/// test to see if it's a go object here and add the appending stuff

				if gfield.Type.IsObject() || gfield.Type.IsArray() {
					tmp++
					ifa := fmt.Sprintf("for i%d := range %s.%s {", tmp, v, codegen.Goify(fname, true))
					fieldAssignments = append(fieldAssignments, ifa)
					ifd := fmt.Sprintf("tmp%d := &%s.%s[i%d]", tmp, v, codegen.Goify(fname, true), tmp)
					fieldAssignments = append(fieldAssignments, ifd)
					ifb := fmt.Sprintf("%s.%s = append(%s.%s, tmp%d.%sTo%s())", utype, codegen.Goify(key, true), utype, codegen.Goify(key, true), tmp, inflect.Singularize(codegen.Goify(key, true)), inflect.Singularize(codegen.Goify(key, true)))
					fieldAssignments = append(fieldAssignments, ifb)
					ifc := fmt.Sprintf("}")
					fieldAssignments = append(fieldAssignments, ifc)

				} else {
					fa := fmt.Sprintf("\t%s.%s = %s%s.%s", utype, codegen.Goify(key, true), prefix, v, codegen.Goify(fname, true))
					fieldAssignments = append(fieldAssignments, fa)
				}
			} else {
				fn := codegen.Goify(strings.Replace(field.FieldName, "ID", "", -1), false)
				if fn == key {
					gfield, ok := obj[fn]
					if ok {
						fa := fmt.Sprintf("tmp%d := &%s.%s", tmp, mtype, codegen.Goify(fn, true))
						fieldAssignments = append(fieldAssignments, fa)
						var view string
						if gfield.View != "" {
							view = gfield.View
						}
						fa = fmt.Sprintf("%s.%s = tmp%d.%sTo%s%s() // %s", utype, codegen.Goify(fn, true), tmp, codegen.Goify(fn, true), codegen.Goify(fn, true), codegen.Goify(view, true))

						fieldAssignments = append(fieldAssignments, fa)
						tmp++
					}
				}
			}
		}
	}
	return strings.Join(fieldAssignments, "\n")
}

func fieldAssignmentTypeToModel(model *RelationalModelDefinition, ut *design.UserTypeDefinition, utype, mtype string) string {
	// Get a sortable slice of field names
	var keys []string
	for k := range model.RelationalFields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var fieldAssignments []string
	for _, fname := range keys {
		field := model.RelationalFields[fname]

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
					// set it explicitly because we're reusing the same bool
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

func viewSelect(ut *RelationalModelDefinition, v *design.ViewDefinition) string {
	obj := v.Type.(design.Object)
	var fields []string
	for name := range obj {
		if obj[name].Type.IsPrimitive() {
			if strings.TrimSpace(name) != "" && name != "links" {
				bf, ok := ut.RelationalFields[codegen.Goify(name, true)]
				if ok {
					fields = append(fields, bf.DatabaseFieldName)
				}
			}
		}
	}
	sort.Strings(fields)
	return strings.Join(fields, ",")
}

func viewFields(ut *RelationalModelDefinition, v *design.ViewDefinition) []*RelationalFieldDefinition {
	obj := v.Type.(design.Object)
	var fields []*RelationalFieldDefinition
	for name := range obj {
		if obj[name].Type.IsPrimitive() {
			if strings.TrimSpace(name) != "" && name != "links" {
				bf, ok := ut.RelationalFields[codegen.Goify(name, true)]
				if ok {
					fields = append(fields, bf)
				}
			} else if name == "links" {
				for _, ld := range v.Parent.Links {
					pretty.Println(ld.Name, ld.View)
				}
			}
		}
	}

	return fields
}

func viewFieldNames(ut *RelationalModelDefinition, v *design.ViewDefinition) []string {
	obj := v.Type.(design.Object)
	var fields []string
	for name := range obj {
		if obj[name].Type.IsPrimitive() {
			if strings.TrimSpace(name) != "" && name != "links" {
				bf, ok := ut.RelationalFields[codegen.Goify(name, true)]

				if ok {
					fields = append(fields, "&"+codegen.Goify(bf.FieldName, false))
				}
			}
		}
	}

	sort.Strings(fields)
	return fields
}

// NewUserHelperWriter returns a contexts code writer.
// User types contain custom data structured defined in the DSL with "Type".
func NewUserHelperWriter(filename string) (*UserHelperWriter, error) {
	file, err := codegen.SourceFileFor(filename)
	if err != nil {
		return nil, err
	}
	return &UserHelperWriter{SourceFile: file}, nil
}

// Execute writes the code for the context types to the writer.
func (w *UserHelperWriter) Execute(data *UserTypeTemplateData) error {
	fm := make(map[string]interface{})
	fm["famt"] = fieldAssignmentModelToType
	fm["fatm"] = fieldAssignmentTypeToModel
	fm["viewSelect"] = viewSelect
	fm["viewFields"] = viewFields
	fm["viewFieldNames"] = viewFieldNames
	fm["goDatatype"] = goDatatype
	fm["plural"] = inflect.Pluralize
	fm["gtt"] = codegen.GoTypeTransform
	fm["gttn"] = codegen.GoTypeTransformName
	fm["gptn"] = codegen.GoTypeName
	fm["newMediaTemplate"] = newMediaTemplate
	return w.ExecuteTemplate("types", userHelperT, fm, data)
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
	fm["fapm"] = fieldAssignmentPayloadToModel
	fm["viewSelect"] = viewSelect
	fm["viewFields"] = viewFields
	fm["viewFieldNames"] = viewFieldNames
	fm["goDatatype"] = goDatatype
	fm["plural"] = inflect.Pluralize
	fm["gtt"] = codegen.GoTypeTransform
	fm["gttn"] = codegen.GoTypeTransformName
	return w.ExecuteTemplate("types", userTypeT, fm, data)
}

// arrayAttribute returns the array element attribute definition.
func arrayAttribute(a *design.AttributeDefinition) *design.AttributeDefinition {
	return a.Type.(*design.Array).ElemType
}

type mediaTemplate struct {
	Media    *design.MediaTypeDefinition
	ViewName string
	Model    *RelationalModelDefinition
	View     *design.ViewDefinition
}

// {{ template "Media" (newMediaTemplate $rmt $vname $ut $vp $vpn)}}
func newMediaTemplate(mtd *design.MediaTypeDefinition, vn string, view *design.ViewDefinition, model *RelationalModelDefinition) *mediaTemplate {
	return &mediaTemplate{
		Media:    mtd,
		ViewName: vn,
		View:     view,
		Model:    model,
	}
}

const (
	// userTypeT generates the code for a user type.
	// template input: UserTypeTemplateData
	userTypeT = `{{$ut := .UserType}}{{$ap := .AppPkg}}// {{if $ut.Description}}{{$ut.Description}}{{else}}{{$ut.ModelName}} Relational Model{{end}}
{{$ut.StructDefinition}}
// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m {{$ut.ModelName}}) TableName() string {
{{ if ne $ut.Alias "" }}
return "{{ $ut.Alias}}" {{ else }} return "{{ $ut.TableName }}"
{{end}}
}
// {{$ut.ModelName}}DB is the implementation of the storage interface for
// {{$ut.ModelName}}.
type {{$ut.ModelName}}DB struct {
	Db *gorm.DB
	{{ if $ut.Cached }}cache *cache.Cache{{end}}
}
// New{{$ut.ModelName}}DB creates a new storage type.
func New{{$ut.ModelName}}DB(db *gorm.DB) *{{$ut.ModelName}}DB {
	{{ if $ut.Cached }}return &{{$ut.ModelName}}DB{
		Db: db,
		cache: cache.New(5*time.Minute, 30*time.Second),
	}
	{{ else  }}return &{{$ut.ModelName}}DB{Db: db}{{ end  }}
}
// DB returns the underlying database.
func (m *{{$ut.ModelName}}DB) DB() interface{} {
	return m.Db
}

// {{$ut.ModelName}}Storage represents the storage interface.
type {{$ut.ModelName}}Storage interface {
	DB() interface{}
	List(ctx context.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}) ([]*{{$ut.ModelName}}, error)
	Get(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.PKAttributes}}) (*{{$ut.ModelName}}, error)
	Add(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.LowerName}} *{{$ut.ModelName}}) (error)
	Update(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.LowerName}} *{{$ut.ModelName}}) (error)
	Delete(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{ $ut.PKAttributes}}) (error)
{{range $rname, $rmt := $ut.RenderTo}}{{/*

*/}}{{range $vname, $view := $rmt.Views}}{{ $mtd := $ut.Project $rname $vname }}
	List{{goify $rmt.TypeName true}}{{if not (eq $vname "default")}}{{goify $vname true}}{{end}} (ctx context.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}{{/*
*/}}{{range $nm, $bt := $ut.BelongsTo}}, {{goify (printf "%s%s" $bt.ModelName "ID") false}} int{{end}}) []*app.{{goify $rmt.TypeName true}}{{if not (eq $vname "default")}}{{goify $vname true}}{{end}}
	One{{goify $rmt.TypeName true}}{{if not (eq $vname "default")}}{{goify $vname true}}{{end}} (ctx context.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}{{/*
*/}}, {{$ut.PKAttributes}}{{range $nm, $bt := $ut.BelongsTo}},{{goify (printf "%s%s" $bt.ModelName "ID") false}} int{{end}}){{/*
*/}} (*app.{{goify $rmt.TypeName true}}{{if not (eq $vname "default")}}{{goify $vname true}}{{end}}, error)
{{end}}{{/*

*/}}{{end}}
{{range $bfn, $bf := $ut.BuiltFrom}}
	UpdateFrom{{$bfn}}(ctx context.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }},payload *app.{{goify $bfn true}}, {{$ut.PKAttributes}}) error
{{end}}
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *{{$ut.ModelName}}DB) TableName() string {
{{ if ne $ut.Alias "" }}
	return "{{ $ut.Alias}}" {{ else }} return "{{ $ut.TableName }}"
{{end}}
}

{{ range $idx, $bt := $ut.BelongsTo}}
// Belongs To Relationships

// {{$ut.ModelName}}FilterBy{{$bt.ModelName}} is a gorm filter for a Belongs To relationship.
func {{$ut.ModelName}}FilterBy{{$bt.ModelName}}({{goify (printf "%s%s" $bt.ModelName "ID") false}} int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if {{goify (printf "%s%s" $bt.ModelName "ID") false}} > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("{{if $bt.RelationalFields.ID.DatabaseFieldName}}{{ if ne $bt.RelationalFields.ID.DatabaseFieldName "id" }}{{$bt.RelationalFields.ID.DatabaseFieldName}} = ?", {{goify (printf "%s%s" $bt.ModelName "ID") false}}){{else}}{{$bt.Underscore}}_id = ?", {{goify (printf "%s%s" $bt.ModelName "ID") false}}){{end}}
			{{ else }}{{$bt.Underscore}}_id = ?", {{goify (printf "%s%s" $bt.ModelName "ID") false}}){{ end }}
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}
{{end}}

// CRUD Functions

// Get returns a single {{$ut.ModelName}} as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *{{$ut.ModelName}}DB) Get(ctx context.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}, {{$ut.PKAttributes}}) (*{{$ut.ModelName}}, error){
	defer goa.MeasureSince([]string{"goa","db","{{goify $ut.ModelName false}}", "get"}, time.Now())

	var native {{$ut.ModelName}}
	err := m.Db.Table({{ if $ut.DynamicTableName }}tableName{{else}}m.TableName(){{ end }}).Where("{{$ut.PKWhere}}",{{$ut.PKWhereFields}} ).Find(&native).Error
	if err ==  gorm.ErrRecordNotFound {
		return nil, err
	}
	{{ if $ut.Cached }}go m.cache.Set(strconv.Itoa(native.ID), &native, cache.DefaultExpiration)
	{{end}}
	return &native, err
}

// List returns an array of {{$ut.ModelName}}
func (m *{{$ut.ModelName}}DB) List(ctx context.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}) ([]*{{$ut.ModelName}}, error) {
	defer goa.MeasureSince([]string{"goa","db","{{goify $ut.ModelName false}}", "list"}, time.Now())

	var objs []*{{$ut.ModelName}}
	err := m.Db.Table({{ if $ut.DynamicTableName }}tableName{{else}}m.TableName(){{ end }}).Find(&objs).Error
	if err != nil && err !=  gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *{{$ut.ModelName}}DB) Add(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, model *{{$ut.ModelName}}) (error) {
	defer goa.MeasureSince([]string{"goa","db","{{goify $ut.ModelName false}}", "add"}, time.Now())

{{ range $l, $pk := $ut.PrimaryKeys }}
	{{ if eq $pk.Datatype "uuid" }}model.{{$pk.FieldName}} = uuid.NewV4(){{ end }}
{{ end }}
	err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding {{$ut.ModelName}}", "error", err.Error())
		return err
	}
	{{ if $ut.Cached }}
	go m.cache.Set(strconv.Itoa(model.ID), model, cache.DefaultExpiration) {{ end }}
	return nil
}

// Update modifies a single record.
func (m *{{$ut.ModelName}}DB) Update(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, model *{{$ut.ModelName}}) error {
	defer goa.MeasureSince([]string{"goa","db","{{goify $ut.ModelName false}}", "update"}, time.Now())

	obj, err := m.Get(ctx{{ if $ut.DynamicTableName }}, tableName{{ end }}, {{$ut.PKUpdateFields "model"}})
	if err != nil {
		goa.LogError(ctx, "error updating {{$ut.ModelName}}", "error", err.Error())
		return  err
	}
	err = m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Model(obj).Updates(model).Error
	{{ if $ut.Cached }}go func(){
		m.cache.Set(strconv.Itoa(model.ID), obj, cache.DefaultExpiration)
	}()
	{{ end }}
	return err
}

// Delete removes a single record.
func (m *{{$ut.ModelName}}DB) Delete(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.PKAttributes}})  error {
	defer goa.MeasureSince([]string{"goa","db","{{goify $ut.ModelName false}}", "delete"}, time.Now())

	var obj {{$ut.ModelName}}{{ $l := len $ut.PrimaryKeys }}
	{{ if eq $l 1 }}
	err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Delete(&obj, {{$ut.PKWhereFields}}).Error
	{{ else  }}err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Delete(&obj).Where("{{$ut.PKWhere}}", {{$ut.PKWhereFields}}).Error
	{{ end }}
	if err != nil {
		goa.LogError(ctx, "error deleting {{$ut.ModelName}}", "error", err.Error())
		return  err
	}
	{{ if $ut.Cached }} go m.cache.Delete(strconv.Itoa(id)) {{ end }}
	return  nil
}

{{ range $bfn, $bf := $ut.BuiltFrom }}
// {{$ut.ModelName}}From{{$bfn}} Converts source {{goify $bfn true}} to target {{$ut.ModelName}} model
// only copying the non-nil fields from the source.
func {{$ut.ModelName}}From{{$bfn}}(payload *app.{{goify $bfn true}}) *{{$ut.ModelName}} {
	{{$ut.LowerName}} := &{{$ut.ModelName}}{}
 	{{ fapm $ut $bf "app" "payload" "payload" $ut.LowerName}}

 	 return {{$ut.LowerName}}
}

// UpdateFrom{{$bfn}} applies non-nil changes from {{goify $bfn true}} to the model and saves it
func (m *{{$ut.ModelName}}DB)UpdateFrom{{$bfn}}(ctx context.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }},payload *app.{{goify $bfn true}}, {{$ut.PKAttributes}}) error {
	defer goa.MeasureSince([]string{"goa","db","{{goify $ut.ModelName false}}", "updatefrom{{goify $bfn false}}"}, time.Now())

	var obj {{$ut.ModelName}}
	 err := m.Db.Table({{ if $ut.DynamicTableName }}tableName{{else}}m.TableName(){{ end }}).Where("{{$ut.PKWhere}}",{{$ut.PKWhereFields}} ).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving {{$ut.ModelName}}", "error", err.Error())
		return  err
	}
 	{{ fapm $ut $bf "app" "payload" "payload" "obj"}}

	err = m.Db.Save(&obj).Error
 	 return err
}
{{ end  }}


`

	userHelperT = `{{define "Media"}}` + mediaT + `{{end}}` + `{{$ut := .UserType}}{{$ap := .AppPkg}}
{{ if $ut.Roler }}
// GetRole returns the value of the role field and satisfies the Roler interface.
func (m {{$ut.ModelName}}) GetRole() string {
	return {{$f := $ut.Fields.role}}{{if $f.Nullable}}*{{end}}m.Role
}
{{end}}

{{ range $rname, $rmt := $ut.RenderTo }}
{{ range $vname, $view := $rmt.Views}}
{{ $mtd := $ut.Project $rname $vname }}

{{template "Media" (newMediaTemplate $rmt $vname $view $ut)}}
{{end}}{{end}}

`

	mediaT = `// MediaType Retrieval Functions

// List{{goify .Media.TypeName true}}{{if not (eq .ViewName "default")}}{{goify .ViewName true}}{{end}} returns an array of view: {{.ViewName}}.
func (m *{{.Model.ModelName}}DB) List{{goify .Media.TypeName true}}{{if not (eq .ViewName "default")}}{{goify .ViewName true}}{{end}}{{/*
*/}} (ctx context.Context{{ if .Model.DynamicTableName}}, tableName string{{ end }}{{/*
*/}} {{range $nm, $bt := .Model.BelongsTo}},{{goify (printf "%s%s" $bt.ModelName "ID") false}} int{{end}}){{/*
*/}} []*app.{{goify .Media.TypeName true}}{{if not (eq .ViewName "default")}}{{goify .ViewName true}}{{end}}{
	defer goa.MeasureSince([]string{"goa","db","{{goify .Media.TypeName false}}", "list{{goify .Media.TypeName false}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName false}}{{end}}"}, time.Now())

	var native []*{{goify .Model.ModelName true}}
	var objs []*app.{{goify .Media.TypeName true}}{{if not (eq .ViewName "default")}}{{goify .ViewName true}}{{end}}{{$ctx:= .}}
	err := m.Db.Scopes({{range $nm, $bt := .Model.BelongsTo}}{{/*
*/}}{{$ctx.Model.ModelName}}FilterBy{{goify $bt.ModelName true}}({{goify (printf "%s%s" $bt.ModelName "ID") false}}, m.Db), {{end}}){{/*
*/}}.Table({{ if .Model.DynamicTableName }}tableName{{else}}m.TableName(){{ end }}).{{ range $ln, $lv := .Media.Links }}Preload("{{goify $ln true}}").{{end}}Find(&native).Error
{{/* //	err := m.Db.Table({{ if .Model.DynamicTableName }}tableName{{else}}m.TableName(){{ end }}).{{ range $ln, $lv := .Media.Links }}Preload("{{goify $ln true}}").{{end}}Find(&objs).Error */}}
	if err != nil {
		goa.LogError(ctx, "error listing {{.Model.ModelName}}", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.{{.Model.ModelName}}To{{goify .Media.UserTypeDefinition.TypeName true}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}}())
	}

	return objs
}

// {{$.Model.ModelName}}To{{goify .Media.UserTypeDefinition.TypeName true}}{{if not (eq .ViewName "default")}}{{goify .ViewName true}}{{end}}{{/*
*/}} loads a {{.Model.ModelName}} and builds the {{.ViewName}} view of media type {{.Media.TypeName}}.
func (m *{{.Model.ModelName}}) {{$.Model.ModelName}}To{{goify .Media.UserTypeDefinition.TypeName true}}{{if not (eq .ViewName "default")}}{{goify .ViewName true}}{{end}}(){{/*
*/}} *app.{{goify .Media.TypeName true}}{{if not (eq .ViewName "default")}}{{goify .ViewName true}}{{end}} {
	{{.Model.LowerName}} := &app.{{goify .Media.TypeName true}}{{if not (eq .ViewName "default")}}{{goify .ViewName true}}{{end}}{}
 	{{ famt .Model .View "m" "m" .Model.LowerName}}

 	 return {{.Model.LowerName}}
}

// One{{goify .Media.TypeName true}}{{if not (eq .ViewName "default")}}{{goify .ViewName true}}{{end}} loads a {{.Model.ModelName}} and builds the {{.ViewName}} view of media type {{.Media.TypeName}}.
func (m *{{.Model.ModelName}}DB) One{{goify .Media.TypeName true}}{{if not (eq .ViewName "default")}}{{goify .ViewName true}}{{end}}{{/*
*/}} (ctx context.Context{{ if .Model.DynamicTableName}}, tableName string{{ end }},{{.Model.PKAttributes}}{{/*
*/}}{{range $nm, $bt := .Model.BelongsTo}},{{goify (printf "%s%s" $bt.ModelName "ID") false}} int{{end}}){{/*
*/}} (*app.{{goify .Media.TypeName true}}{{if not (eq .ViewName "default")}}{{goify .ViewName true}}{{end}}, error){
	defer goa.MeasureSince([]string{"goa","db","{{goify .Media.TypeName false}}", "one{{goify .Media.TypeName false}}{{if not (eq .ViewName "default")}}{{goify .ViewName false}}{{end}}"}, time.Now())

	var native {{.Model.ModelName}}
	err := m.Db.Scopes({{range $nm, $bt := .Model.BelongsTo}}{{$ctx.Model.ModelName}}FilterBy{{goify $bt.ModelName true}}({{goify (printf "%s%s" $bt.ModelName "ID") false}}, m.Db), {{end}}).Table({{ if .Model.DynamicTableName }}tableName{{else}}m.TableName(){{ end }}){{range $na, $hm:= .Model.HasMany}}.Preload("{{plural $hm.ModelName}}"){{end}}{{range $nm, $bt := .Model.BelongsTo}}.Preload("{{$bt.ModelName}}"){{end}}.Where("{{.Model.PKWhere}}",{{.Model.PKWhereFields}}).Find(&native).Error

	if err != nil && err !=  gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting {{.Model.ModelName}}", "error", err.Error())
		return nil, err
	}
	{{ if .Model.Cached }} go func(){
		m.cache.Set(strconv.Itoa(native.ID), &native, cache.DefaultExpiration)
	}() {{ end }}
	view := *native.{{.Model.ModelName}}To{{goify .Media.UserTypeDefinition.TypeName true}}{{if not (eq .ViewName "default")}}{{goify .ViewName true}}{{end}}()
	return &view, err
}
`
)
