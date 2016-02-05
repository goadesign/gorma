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

	// UserHelpersWriter generate code for a goa application user types.
	// User types are data structures defined in the DSL with "Type".
	UserHelperWriter struct {
		*codegen.SourceFile
		UserHelperTmpl *template.Template
	}
)

func fieldAssignmentModelToType(model *RelationalModelDefinition, ut *design.ViewDefinition, v, mtype, utype string) string {
	//utPackage := "app"
	var fieldAssignments []string
	// type.Field = model.Field
	for fname, field := range model.RelationalFields {

		var mpointer, upointer bool
		mpointer = field.Nullable
		obj := ut.Type.ToObject()
		definition := ut.Parent.Definition()

		fmt.Println(fname, field.Datatype)
		if field.Datatype == "" {
			continue
		}
		// Set the relational field
		// if the view has one of them
		if field.Datatype == BelongsTo {
			fn := strings.Replace(field.FieldName, "ID", "", -1)
			_, ok := obj[codegen.Goify(fn, false)]
			if ok {
				fa := fmt.Sprintf("%s.%s = %s.%s.%sTo%s()", utype, codegen.Goify(fn, true), mtype, codegen.Goify(fn, true), codegen.Goify(fn, true), codegen.Goify(fn, true))
				fieldAssignments = append(fieldAssignments, fa)
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

				fa := fmt.Sprintf("\t%s.%s = %s%s.%s", utype, codegen.Goify(key, true), prefix, v, codegen.Goify(fname, true))
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
					fmt.Println(n)
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
	List(ctx goa.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}) []{{$ut.ModelName}}
	Get(ctx goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.PKAttributes}}) ({{$ut.ModelName}}, error)
	Add(ctx goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.LowerName}} {{$ut.ModelName}}) ({{$ut.ModelName}}, error)
	Update(ctx goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.LowerName}} {{$ut.ModelName}}) (error)
	Delete(ctx goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{ $ut.PKAttributes}}) (error) 	
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *{{$ut.ModelName}}DB) TableName() string {
{{ if ne $ut.Alias "" }}
return "{{ $ut.Alias}}" {{ else }} return "{{ $ut.TableName }}"
{{end}}
}



// CRUD Functions

// Get returns a single {{$ut.ModelName}} as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *{{$ut.ModelName}}DB) Get(ctx goa.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}, id int) {{$ut.ModelName}}{	
	now := time.Now()
	defer ctx.Info("{{$ut.ModelName}}:Get", "duration", time.Since(now))
	var native {{$ut.ModelName}}
	m.Db.Table({{ if $ut.DynamicTableName }}.Table(tableName){{else}}m.TableName(){{ end }}).Where("id = ?", id).Find(&native)
	return native 
}

// List returns an array of {{$ut.ModelName}}
func (m *{{$ut.ModelName}}DB) List{{$ut.TypeName}}(ctx goa.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}) []{{$ut.ModelName}}{
	now := time.Now()
	defer ctx.Info("{{$ut.ModelName}}:List", "duration", time.Since(now))
	var objs []{{$ut.ModelName}}
	err := m.Db.Table({{ if $ut.DynamicTableName }}.Table(tableName){{else}}m.TableName(){{ end }}).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing {{$ut.ModelName}}", "error", err.Error())
		return objs
	}

	return objs
}
// Add creates a new record.
func (m *{{$ut.ModelName}}DB) Add(ctx goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, model {{$ut.ModelName}}) ({{$ut.ModelName}}, error) {
	now := time.Now()
	defer ctx.Info("{{$ut.ModelName}}:Add", "duration", time.Since(now))
	err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Create(&model).Error
	if err != nil {
		ctx.Error("error updating {{$ut.ModelName}}", "error", err.Error())
		return model, err
	}
	{{ if $ut.Cached }}
	go m.cache.Set(strconv.Itoa(model.ID), model, cache.DefaultExpiration) {{ end }}
	return model, err
}
// Update modifies a single record.
func (m *{{$ut.ModelName}}DB) Update(ctx goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, model {{$ut.ModelName}}) error {
	now := time.Now()
	defer ctx.Info("{{$ut.ModelName}}:Update", "duration", time.Since(now))
	obj := m.Get(ctx{{ if $ut.DynamicTableName }}, tableName{{ end }}, {{$ut.PKUpdateFields "model"}})
	err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Model(&obj).Updates(model).Error
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
func (m *{{$ut.ModelName}}DB) Delete(ctx goa.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{$ut.PKAttributes}})  error {
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
`

	userHelperT = `{{$ut := .UserType}}{{$ap := .AppPkg}}
{{ if $ut.Roler }}
// GetRole returns the value of the role field and satisfies the Roler interface.
func (m {{$ut.ModelName}}) GetRole() string {
	return {{$f := $ut.Fields.role}}{{if $f.Nullable}}*{{end}}m.Role
}
{{end}}


{{ range $rname, $rmt := $ut.RenderTo }}
{{ range $vname, $view := $rmt.Views}}
{{ $mtd := $ut.Project $rname $vname }}


// MediaType Retrieval Functions
// List{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}} returns an array of view: {{$vname}}
func (m *{{$ut.ModelName}}DB) List{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}} (ctx goa.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}) []app.{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}}{
	now := time.Now()
	defer ctx.Info("List{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}}", "duration", time.Since(now))
	var objs []app.{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}}
	err := m.Db.Table({{ if $ut.DynamicTableName }}.Table(tableName){{else}}m.TableName(){{ end }}).{{ range $ln, $lv := $rmt.Links }}Preload("{{goify $ln true}}").{{end}}Find(&objs).Error
	if err != nil {
		ctx.Error("error listing {{$ut.ModelName}}", "error", err.Error())
		return objs
	}

	return objs
}

func (m *{{$ut.ModelName}}) {{$ut.ModelName}}To{{$rmt.UserTypeDefinition.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}}() *app.{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}} {
	{{$ut.LowerName}} := &app.{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}}{}		
 	{{ famt $ut $view "m" "m" $ut.LowerName}}		

 	 return {{$ut.LowerName}}
}

// One{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}} returns an array of view: {{$vname}}
func (m *{{$ut.ModelName}}DB) One{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}} (ctx goa.Context{{ if $ut.DynamicTableName}}, tableName string{{ end }}, id int) *app.{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}}{	
	now := time.Now()
	defer ctx.Info("One{{$rmt.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}}", "duration", time.Since(now))

	var native {{$ut.ModelName}}

	m.Db.Table({{ if $ut.DynamicTableName }}.Table(tableName){{else}}m.TableName(){{ end }}){{range $na, $hm:= $ut.HasMany}}.Preload("{{$hm.ModelName}}"){{end}}{{range $nm, $bt := $ut.BelongsTo}}.Preload("{{$bt.ModelName}}"){{end}}.Where("id = ?", id).Find(&native)
	view := native.{{$ut.ModelName}}To{{$rmt.UserTypeDefinition.TypeName}}{{if eq $vname "default"}}{{else}}{{goify $vname true}}{{end}}()
	return view 
	
}
{{end}}{{end}}
`
)

/*
{{$functionName := gttn $rmt.UserTypeDefinition $mtd.UserTypeDefinition "App"}}
{{ gtt $rmt.UserTypeDefinition $mtd.UserTypeDefinition "app" $functionName }}

*/
