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
	//utPackage := "app"
	var fieldAssignments []string
	// type.Field = model.Field
	for fname, field := range model.RelationalFields {

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
					// set it explicity because we're reusing the same bool
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

func fieldAssignmentModelToType(model *RelationalModelDefinition, ut *design.ViewDefinition, verpkg, v, mtype, utype string) string {
	//utPackage := "app"
	var tmp int
	tmp = 1
	var fieldAssignments []string
	// type.Field = model.Field

	if !strings.Contains(ut.Name, "link") {
		for ln, _ := range ut.Parent.Links {
			iff := fmt.Sprintf("var tmp%dCollection %s.%sLinkCollection", tmp, codegen.Goify(verpkg, false), inflect.Singularize(codegen.Goify(ln, true)))
			fieldAssignments = append(fieldAssignments, iff)
			ifa := fmt.Sprintf("for _,k := range %s.%s {", v, codegen.Goify(ln, true))
			fieldAssignments = append(fieldAssignments, ifa)
			ifb := fmt.Sprintf("tmp%dCollection = append(tmp%dCollection,  k.%sTo%s%sLink())", tmp, tmp, inflect.Singularize(codegen.Goify(ln, true)), verpkg, inflect.Singularize(codegen.Goify(ln, true)))

			fieldAssignments = append(fieldAssignments, ifb)
			ifc := fmt.Sprintf("}")
			fieldAssignments = append(fieldAssignments, ifc)
			ifd := fmt.Sprintf("%s.Links = &%s.%sLinks{%s: tmp%dCollection}", utype, codegen.Goify(verpkg, false), codegen.Goify(utype, true), codegen.Goify(ln, true), tmp)
			fieldAssignments = append(fieldAssignments, ifd)
			tmp++

		}

	}
	for fname, field := range model.RelationalFields {

		var mpointer, upointer bool
		mpointer = field.Nullable
		obj := ut.Type.ToObject()
		definition := ut.Parent.Definition()

		if field.Datatype == "" {
			continue
		}
		// Set the relational field
		// if the view has one of them
		if field.Datatype == BelongsTo {
			fn := strings.Replace(field.FieldName, "ID", "", -1)
			_, ok := obj[codegen.Goify(fn, false)]
			if ok {
				fa := fmt.Sprintf("tmp%d := &%s.%s", tmp, mtype, codegen.Goify(fn, true))
				fieldAssignments = append(fieldAssignments, fa)
				fa = fmt.Sprintf("%s.%s = tmp%d.%sTo%s%s()", utype, codegen.Goify(fn, true), tmp, codegen.Goify(fn, true), verpkg, codegen.Goify(fn, true))
				fieldAssignments = append(fieldAssignments, fa)
				tmp++
			}
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

				if gfield.Type.IsObject() {
					ifa := fmt.Sprintf("for _,k := range %s.%s {", v, codegen.Goify(fname, true))
					fieldAssignments = append(fieldAssignments, ifa)
					ifb := fmt.Sprintf("%s.%s = append(%s.%s, k.%sTo%s%s())", utype, codegen.Goify(key, true), utype, codegen.Goify(key, true), inflect.Singularize(codegen.Goify(key, true)), verpkg, inflect.Singularize(codegen.Goify(key, true)))
					fieldAssignments = append(fieldAssignments, ifb)
					ifc := fmt.Sprintf("}")
					fieldAssignments = append(fieldAssignments, ifc)

				} else {
					fa := fmt.Sprintf("\t%s.%s = %s%s.%s", utype, codegen.Goify(key, true), prefix, v, codegen.Goify(fname, true))
					fieldAssignments = append(fieldAssignments, fa)
				}
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

func viewSelect(ut *RelationalModelDefinition, v *design.ViewDefinition) string {
	obj := v.Type.(design.Object)
	var fields []string
	for name := range obj {
		if obj[name].Type.IsPrimitive() {
			if strings.TrimSpace(name) != "" && name != "links" {
				bf, ok := ut.RelationalFields[codegen.Goify(name, true)]
				if ok {
					if bf.Alias != "" {
						fields = append(fields, bf.Alias)
					} else {
						fields = append(fields, bf.DatabaseFieldName)
					}
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
				for n, ld := range v.Parent.Links {
					fmt.Println("LINK:", n)
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
	fm["gptn"] = codegen.GoPackageTypeName
	fm["vp"] = codegen.VersionPackage
	fm["newMediaTemplateVersion"] = newMediaTemplateVersion
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
	fm["vp"] = codegen.VersionPackage
	return w.ExecuteTemplate("types", userTypeT, fm, data)
}

// arrayAttribute returns the array element attribute definition.
func arrayAttribute(a *design.AttributeDefinition) *design.AttributeDefinition {
	return a.Type.(*design.Array).ElemType
}

type mediaTemplateVersion struct {
	Media              *design.MediaTypeDefinition
	ViewName           string
	Model              *RelationalModelDefinition
	View               *design.ViewDefinition
	VersionPackage     string //v1
	VersionPackageName string // goified version package V1
}

// {{ template "MediaVersion" (newMediaTemplateVersion $rmt $vname $ut $vp $vpn)}}
func newMediaTemplateVersion(mtd *design.MediaTypeDefinition, vn string, view *design.ViewDefinition, model *RelationalModelDefinition, vp, vpn string) *mediaTemplateVersion {
	return &mediaTemplateVersion{
		Media:              mtd,
		ViewName:           vn,
		View:               view,
		Model:              model,
		VersionPackage:     vp,
		VersionPackageName: vpn,
	}
}

const (
	// userTypeT generates the code for a user type.
	// template input: UserTypeTemplateData
	userTypeT = `{{$ut := .UserType}}{{$ap := .AppPkg}}// {{if $ut.Description}}{{$ut.Description}} {{end}}
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
	Db gorm.DB
	log.Logger
	{{ if $ut.Cached }}cache *cache.Cache{{end}}
}
// New{{$ut.ModelName}}DB creates a new storage type.
func New{{$ut.ModelName}}DB(db gorm.DB, logger log.Logger) *{{$ut.ModelName}}DB {
	glog := logger.New("db", "{{$ut.ModelName}}")
	{{ if $ut.Cached }}return &{{$ut.ModelName}}DB{
		Db: db,
		Logger: glog,
		cache: cache.New(5*time.Minute, 30*time.Second),
	}
	{{ else  }}return &{{$ut.ModelName}}DB{Db: db, Logger: glog}{{ end  }}
}
// DB returns the underlying database.
func (m *{{$ut.ModelName}}DB) DB() interface{} {
	return &m.Db
}


// {{$ut.ModelName}}Storage represents the storage interface.
type {{$ut.ModelName}}Storage interface {
	DB() interface{}
	List(ctx *goa.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}) []{{$ut.ModelName}}
	Get(ctx *goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.PKAttributes}}) ({{$ut.ModelName}}, error)
	Add(ctx *goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.LowerName}} *{{$ut.ModelName}}) (*{{$ut.ModelName}}, error)
	Update(ctx *goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.LowerName}} *{{$ut.ModelName}}) (error)
	Delete(ctx *goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{ $ut.PKAttributes}}) (error) 	
{{ range $rname, $rmt := $ut.RenderTo }}{{ range $vname, $view := $rmt.Views}}{{ $mtd := $ut.Project $rname $vname }}{{$vp := vp "app"}}{{$vpn := goify $vp true}}
{{ if $rmt.UserTypeDefinition.SupportsNoVersion }}
List{{$vpn}}{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}} (ctx *goa.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }} {{range $nm, $bt := $ut.BelongsTo}},{{goify $bt.ModelName false}}id int{{end}}) []*{{goify $vpn false}}.{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}}
One{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}} (ctx *goa.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}, id int{{range $nm, $bt := $ut.BelongsTo}},{{goify $bt.ModelName false}}id int{{end}}) (*{{goify $vpn false}}.{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}}, error)
{{ end }}


{{ range $version :=  $rmt.Versions }} {{$vp := vp $version}}{{$vpn := goify $vp true}}
// {{$version}} I don't remember why I put this here.  Don't delete until I remember.  What versioned things might we add to the Interface? 
{{end}}{{end}}{{end}}
{{range $bfn, $bf := $ut.BuiltFrom}}
UpdateFrom{{$bfn}}(ctx *goa.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }},payload *app.{{goify $bfn true}}, {{$ut.PKAttributes}}) error
{{end }}
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
func {{$ut.ModelName}}FilterBy{{$bt.ModelName}}({{goify $bt.ModelName false}}id int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if {{goify $bt.ModelName false}}id > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("{{$bt.LowerName}}_id = ?", {{goify $bt.ModelName false}}id)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}
{{end}}

// CRUD Functions

// Get returns a single {{$ut.ModelName}} as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *{{$ut.ModelName}}DB) Get(ctx *goa.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}, {{$ut.PKAttributes}}) ({{$ut.ModelName}}, error){	
	now := time.Now()
	defer ctx.Info("{{$ut.ModelName}}:Get", "duration", time.Since(now))
	var native {{$ut.ModelName}}
	err := m.Db.Table({{ if $ut.DynamicTableName }}tableName{{else}}m.TableName(){{ end }}).Where("{{$ut.PKWhere}}",{{$ut.PKWhereFields}} ).Find(&native).Error
	if err == gorm.RecordNotFound {
		return {{$ut.ModelName}}{}, nil
	}
	{{ if $ut.Cached }}go m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration) 
	{{end}}
	return native, err
}

// List returns an array of {{$ut.ModelName}}
func (m *{{$ut.ModelName}}DB) List{{$ut.TypeName}}(ctx *goa.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}) []{{$ut.ModelName}}{
	now := time.Now()
	defer ctx.Info("{{$ut.ModelName}}:List", "duration", time.Since(now))
	var objs []{{$ut.ModelName}}
	err := m.Db.Table({{ if $ut.DynamicTableName }}tableName{{else}}m.TableName(){{ end }}).Find(&objs).Error
	if err != nil && err != gorm.RecordNotFound {
		ctx.Error("error listing {{$ut.ModelName}}", "error", err.Error())
		return objs
	}

	return objs
}
// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *{{$ut.ModelName}}DB) Add(ctx *goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, model *{{$ut.ModelName}}) (*{{$ut.ModelName}}, error) {
	now := time.Now()
	defer ctx.Info("{{$ut.ModelName}}:Add", "duration", time.Since(now))
	err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Create(model).Error
	if err != nil {
		ctx.Error("error updating {{$ut.ModelName}}", "error", err.Error())
		return model, err
	}
	{{ if $ut.Cached }}
	go m.cache.Set(strconv.Itoa(model.ID), model, cache.DefaultExpiration) {{ end }}
	return model, err
}
// Update modifies a single record.
func (m *{{$ut.ModelName}}DB) Update(ctx *goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, model *{{$ut.ModelName}}) error {
	now := time.Now()
	defer ctx.Info("{{$ut.ModelName}}:Update", "duration", time.Since(now))
	obj, err := m.Get(ctx{{ if $ut.DynamicTableName }}, tableName{{ end }}, {{$ut.PKUpdateFields "model"}})
	if err != nil {
		return  err
	}
	err = m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Model(&obj).Updates(model).Error
	{{ if $ut.Cached }}go func(){
		m.cache.Set(strconv.Itoa(model.ID), obj, cache.DefaultExpiration)
	}()
	{{ end }}
	return err
}
// Delete removes a single record.
func (m *{{$ut.ModelName}}DB) Delete(ctx *goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.PKAttributes}})  error {
	now := time.Now()
	defer ctx.Info("{{$ut.ModelName}}:Delete", "duration", time.Since(now))
	var obj {{$ut.ModelName}}{{ $l := len $ut.PrimaryKeys }}
	{{ if eq $l 1 }}
	err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Delete(&obj, id).Error
	{{ else  }}err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Delete(&obj).Where("{{$ut.PKWhere}}", {{$ut.PKWhereFields}}).Error
	{{ end }}
	if err != nil {
		ctx.Error("error retrieving {{$ut.ModelName}}", "error", err.Error())
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

	// UpdateFrom{{$bfn}} applies non-nil changes from {{goify $bfn true}} to the model
	// and saves it
	func (m *{{$ut.ModelName}}DB)UpdateFrom{{$bfn}}(ctx *goa.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }},payload *app.{{goify $bfn true}}, {{$ut.PKAttributes}}) error {
	now := time.Now()
	defer ctx.Info("{{$ut.ModelName}}:Update", "duration", time.Since(now))
	var obj {{$ut.ModelName}}
	 err := m.Db.Table({{ if $ut.DynamicTableName }}tableName{{else}}m.TableName(){{ end }}).Where("{{$ut.PKWhere}}",{{$ut.PKWhereFields}} ).Find(&obj).Error
	if err != nil {
		ctx.Error("error retrieving {{$ut.ModelName}}", "error", err.Error())
		return  err
	}
 	{{ fapm $ut $bf "app" "payload" "payload" "obj"}}		
	
	err = m.Db.Save(&obj).Error
 	 return err
}
{{ end  }}


`

	userHelperT = `{{define "MediaVersion"}}` + mediaVersionT + `{{end}}` + `{{$ut := .UserType}}{{$ap := .AppPkg}}
{{ if $ut.Roler }}
// GetRole returns the value of the role field and satisfies the Roler interface.
func (m {{$ut.ModelName}}) GetRole() string {
	return {{$f := $ut.Fields.role}}{{if $f.Nullable}}*{{end}}m.Role
}
{{end}}


{{ range $rname, $rmt := $ut.RenderTo }}
{{ range $vname, $view := $rmt.Views}}
{{ $mtd := $ut.Project $rname $vname }}

{{$vp := vp "app"}}{{$vpn := goify $vp true}}
{{ if $rmt.UserTypeDefinition.SupportsNoVersion }}
{{ template "MediaVersion" (newMediaTemplateVersion $rmt $vname $view $ut $vp $vpn)}}
{{ end }}
{{ range $version :=  $rmt.Versions }} {{$vp := vp $version}}{{$vpn := goify $vp true}}
// {{$version}}  
{{ template "MediaVersion" (newMediaTemplateVersion $rmt $vname $view $ut $vp $vpn)}}
{{end}}{{end}}{{end}}

`

	mediaVersionT = `// MediaType Retrieval Functions
// List{{.Media.TypeName}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}} returns an array of view: {{.ViewName}}
func (m *{{.Model.ModelName}}DB) List{{.VersionPackageName}}{{.Media.TypeName}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}} (ctx *goa.Context{{ if .Model.DynamicTableName}}, tableName string{{ end }} {{range $nm, $bt := .Model.BelongsTo}},{{goify $bt.ModelName false}}id int{{end}}) []*{{.VersionPackage}}.{{.Media.TypeName}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}}{
	now := time.Now()
	defer ctx.Info("List{{.Media.TypeName}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}}", "duration", time.Since(now))
	var objs []*{{.VersionPackage}}.{{.Media.TypeName}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}} {{$ctx:= .}}
	err := m.Db.Scopes({{range $nm, $bt := .Model.BelongsTo}}{{$ctx.Model.ModelName}}FilterBy{{goify $bt.ModelName true}}({{goify $bt.ModelName false}}id, &m.Db), {{end}}).Table({{ if .Model.DynamicTableName }}tableName{{else}}m.TableName(){{ end }}).{{ range $ln, $lv := .Media.Links }}Preload("{{goify $ln true}}").{{end}}Find(&objs).Error


//	err := m.Db.Table({{ if .Model.DynamicTableName }}tableName{{else}}m.TableName(){{ end }}).{{ range $ln, $lv := .Media.Links }}Preload("{{goify $ln true}}").{{end}}Find(&objs).Error
	if err != nil {
		ctx.Error("error listing {{.Model.ModelName}}", "error", err.Error())
		return objs
	}

	return objs
}

func (m *{{.Model.ModelName}}) {{$.Model.ModelName}}To{{.VersionPackageName}}{{.Media.UserTypeDefinition.TypeName}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}}() *{{.VersionPackage}}.{{.Media.TypeName}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}} {
	{{.Model.LowerName}} := &{{.VersionPackage}}.{{.Media.TypeName}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}}{}		
 	{{ famt .Model .View .VersionPackageName "m" "m" .Model.LowerName}}		

 	 return {{.Model.LowerName}}
}

// One{{.VersionPackageName}}{{.Media.TypeName}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}} returns an array of view: {{.ViewName}}
func (m *{{.Model.ModelName}}DB) One{{.Media.TypeName}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}} (ctx *goa.Context{{ if .Model.DynamicTableName}}, tableName string{{ end }}, id int{{range $nm, $bt := .Model.BelongsTo}},{{goify $bt.ModelName false}}id int{{end}}) (*{{.VersionPackage}}.{{.Media.TypeName}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}}, error){	
	now := time.Now()
	var native {{.Model.ModelName}}
	defer ctx.Info("One{{.Media.TypeName}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}}", "duration", time.Since(now))
	err := m.Db.Scopes({{range $nm, $bt := .Model.BelongsTo}}{{$ctx.Model.ModelName}}FilterBy{{goify $bt.ModelName true}}({{goify $bt.ModelName false}}id, &m.Db), {{end}}).Table({{ if .Model.DynamicTableName }}tableName{{else}}m.TableName(){{ end }}){{range $na, $hm:= .Model.HasMany}}.Preload("{{plural $hm.ModelName}}"){{end}}{{range $nm, $bt := .Model.BelongsTo}}.Preload("{{$bt.ModelName}}"){{end}}.Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		ctx.Error("error getting {{.Model.ModelName}}", "error", err.Error())
		return nil , err 
	}
	{{ if .Model.Cached }} go func(){
		m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)
	}() {{ end }}
	view := *native.{{.Model.ModelName}}To{{.VersionPackageName}}{{.Media.UserTypeDefinition.TypeName}}{{if eq .ViewName "default"}}{{else}}{{goify .ViewName true}}{{end}}()
	return &view, err 
	
}
`
)

/*
{{$functionName := gttn $rmt.UserTypeDefinition $mtd.UserTypeDefinition "App"}}
{{ gtt $rmt.UserTypeDefinition $mtd.UserTypeDefinition "app" $functionName }}

*/
