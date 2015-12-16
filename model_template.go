package gorma

const modelTmpl = `// {{if .Description}}{{.Description}}{{else}}app.{{gotypename . 0}} storage type{{end}}
// Identifier: {{ $typeName :=  gotypename . 0}}{{$typeName := demodel $typeName}}
{{$td := gotypedef . 0 true false}}type {{$typeName}} {{modeldef $td .}}
{{ $belongsto := index .Metadata "github.com/bketelsen/gorma#belongsto" }}
{{ $m2m := index .Metadata "github.com/bketelsen/gorma#many2many" }}
func {{$typeName}}FromCreatePayload(ctx *app.Create{{demodel $typeName}}Context) {{$typeName}} {
	payload := ctx.Payload
	m := {{$typeName}}{}
	copier.Copy(&m, payload)
	{{ if ne $belongsto "" }} m.{{ $belongsto }}ID=int(ctx.{{ demodel $belongsto }}ID){{end}}
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
{{ $roler := index .Metadata "github.com/bketelsen/gorma#roler" }}
{{ if ne $roler "" }}
func (m {{$typeName}}) GetRole() string {
	return m.Role
}
{{end}}

type {{$typeName}}Storage interface {
	List(ctx goa.Context) []{{$typeName}}
	Get(ctx goa.Context, id int) ({{$typeName}}, error)
	Add(ctx goa.Context, o {{$typeName}}) ({{$typeName}}, error)
	Update(ctx goa.Context, o {{$typeName}}) (error)
	Delete(ctx goa.Context, id int) (error)
	{{ storagedef . }}
}

type {{$typeName}}DB struct {
	DB gorm.DB
}
/*{{ if ne $belongsto "" }}{{$barray := split $belongsto ","}}{{ range $idx, $bt := $barray}}
// would prefer to just pass a context in here, but they're all different, so can't
func {{$typeName}}Filter(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if parentid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("{{ snake $bt }}_id = ?", parentid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}{{end}}{{end}}
*/
func New{{$typeName}}DB(db gorm.DB) *{{$typeName}}DB {
	return &{{$typeName}}DB{DB: db}
}

func (m *{{$typeName}}DB) List(ctx goa.Context) []{{$typeName}} {

	var objs []{{$typeName}}
    m.DB.Find(&objs)
	return objs
}

func (m *{{$typeName}}DB) Get(ctx goa.Context, id int) ({{$typeName}}, error) {

	var obj {{$typeName}}

	err := m.DB.Find(&obj, id).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return obj, err
}

func (m *{{$typeName}}DB) Add(ctx goa.Context, model {{$typeName}}) ({{$typeName}}, error) {
	err := m.DB.Create(&model).Error
	return model, err
}
func (m *{{$typeName}}DB) Update(ctx goa.Context, model {{$typeName}}) error {
	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return  err
	}
	err = m.DB.Model(&obj).Updates(model).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return err
}
func (m *{{$typeName}}DB) Delete(ctx goa.Context, id int)  error {
	var obj {{$typeName}}
	err := m.DB.Delete(&obj, id).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return  err
	}
	return  nil
}

{{ if ne $m2m "" }}{{$barray := split $m2m ","}}{{ range $idx, $bt := $barray}}
{{ $pieces := split $bt ":" }} {{ $lowertype := index $pieces 1  }} {{ $lower := lower $lowertype }}  {{ $lowerplural := index $pieces 0  }} {{ $lowerplural := lower $lowerplural}}
func (m *{{$typeName}}DB) Delete{{index $pieces 1}}(ctx goa.Context, {{$lower}}ID int)  error {
	var obj {{$typeName}}

	var assoc {{index $pieces 1}}
	var err error
	assoc.ID = {{$lower}}ID
	if err != nil {
		return err
	}
	err = m.DB.Model(&obj).Association("{{index $pieces 0}}").Delete(assoc).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return  err
	}
	return  nil
}
func (m *{{$typeName}}DB) Add{{index $pieces 1}}(ctx goa.Context, {{$lower}}ID int) error {
	var assoc {{index $pieces 1}}
	assoc.ID = {{$lower}}ID
	err = m.DB.Model(&{{$lower}}).Association("{{index $pieces 0}}").Append(assoc).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return  err
	}
	return  nil
}
func (m *{{$typeName}}DB) List{{index $pieces 0}}(ctx goa.Context)  []{{index $pieces 1}} {
	list := make([]{{index $pieces 1}}, 0)
	var obj {{$typeName}}
	obj.ID = ctx.{{$typeName}}ID
	err := m.DB.Model(&obj).Association("{{index $pieces 0}}").Find(&list).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return  list
	}
	return  nil
}
{{end}}{{end}}
{{if ne $belongsto ""}}{{$barray := split $belongsto ","}}{{ range $idx, $bt := $barray}}
func Filter{{$typeName}}By{{$bt}}(parent int, list []{{$typeName}}) []{{$typeName}} {
	filtered := make([]{{$typeName}},0)
	for _,o := range list {
		if o.{{$bt}}ID == int(parent) {
			filtered = append(filtered,o)
		}
	}
	return filtered
}
{{end}}{{end}}
`
