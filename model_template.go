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
{{ if ne $belongsto "" }}{{$barray := split $belongsto ","}}{{ range $idx, $bt := $barray}}
	m.{{ $bt}}ID=int(ctx.{{ demodel $bt}}ID){{end}}{{end}}
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
{{ $tablename := index .Metadata "github.com/bketelsen/gorma#tablename" }}
{{ if ne $tablename "" }}
func (m {{$typeName}}) TableName() string {
	return "{{ $tablename }}"
}
{{ end }}
{{ $roler := index .Metadata "github.com/bketelsen/gorma#roler" }}
{{ if ne $roler "" }}
func (m {{$typeName}}) GetRole() string {
	return m.Role
}
{{end}}



type {{$typeName}}Storage interface {
	List(ctx context.Context) []{{$typeName}}
	One(ctx context.Context, id int) ({{$typeName}}, error)
	Add(ctx context.Context, o {{$typeName}}) ({{$typeName}}, error)
	Update(ctx context.Context, o {{$typeName}}) (error)
	Delete(ctx context.Context, id int) (error)
	{{ storagedef . }}
}
{{ $cached := index .Metadata "github.com/bketelsen/gorma#cached" }}
type {{$typeName}}DB struct {
	DB gorm.DB
	{{ if ne $cached "" }}cache *cache.Cache{{end}}
}
{{ if ne $belongsto "" }}{{$barray := split $belongsto ","}}{{ range $idx, $bt := $barray}}
// would prefer to just pass a context in here, but they're all different, so can't
func {{$typeName}}FilterBy{{$bt}}(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if parentid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("{{ snake $bt }}_id = ?", parentid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

func (m *{{$typeName}}DB) ListBy{{$bt}}(ctx context.Context, parentid int) []{{$typeName}} {

	var objs []{{$typeName}}
    m.DB.Scopes({{$typeName}}FilterBy{{$bt}}(parentid, &m.DB)).Find(&objs)
	return objs
}


{{end}}{{end}}

func New{{$typeName}}DB(db gorm.DB) *{{$typeName}}DB {
	{{ if ne $cached "" }}
	return &{{$typeName}}DB{
		DB: db,
		cache: cache.New(5*time.Minute, 30*time.Second),
	}
	{{ else  }}
	return &{{$typeName}}DB{DB: db}

	{{ end  }}
}

func (m *{{$typeName}}DB) List(ctx context.Context) []{{$typeName}} {

	var objs []{{$typeName}}
    m.DB.Find(&objs)
	return objs
}

func (m *{{$typeName}}DB) One(ctx context.Context, id int) ({{$typeName}}, error) {
	{{ if ne $cached "" }}//first attempt to retrieve from cache
	o,found := m.cache.Get(strconv.Itoa(id))
	if found {
		return o.({{$typeName}}), nil
	} 
	// fallback to database if not found{{ end }}
	var obj {{$typeName}}

	err := m.DB.Find(&obj, id).Error
	{{ if ne $cached "" }} go m.cache.Set(strconv.Itoa(id), obj, cache.DefaultExpiration) {{ end }}
	return obj, err
}

func (m *{{$typeName}}DB) Add(ctx context.Context, model {{$typeName}}) ({{$typeName}}, error) {
	err := m.DB.Create(&model).Error
	{{ if ne $cached "" }} go m.cache.Set(strconv.Itoa(model.ID), model, cache.DefaultExpiration) {{ end }}
	return model, err
}

func (m *{{$typeName}}DB) Update(ctx context.Context, model {{$typeName}}) error {
	obj, err := m.One(ctx, model.ID)
	if err != nil {
		return  err
	}
	err = m.DB.Model(&obj).Updates(model).Error
	{{ if ne $cached "" }} 
	go func(){
	obj, err := m.One(ctx, model.ID)
	if err == nil {
		m.cache.Set(strconv.Itoa(model.ID), obj, cache.DefaultExpiration)
	}
	}()	
	{{ end }}

	return err
}

func (m *{{$typeName}}DB) Delete(ctx context.Context, id int)  error {
	var obj {{$typeName}}
	err := m.DB.Delete(&obj, id).Error
	if err != nil {
		return  err
	}
	{{ if ne $cached "" }} go m.cache.Delete(strconv.Itoa(id)) {{ end }}
	return  nil
}

{{ if ne $m2m "" }}{{$barray := split $m2m ","}}{{ range $idx, $bt := $barray}}
{{ $pieces := split $bt ":" }} {{ $lowertype := index $pieces 1  }} {{ $lower := lower $lowertype }}  {{ $lowerplural := index $pieces 0  }} {{ $lowerplural := lower $lowerplural}}
func (m *{{$typeName}}DB) Delete{{index $pieces 1}}(ctx context.Context,{{lower $typeName}}ID,  {{$lower}}ID int)  error {
	var obj {{$typeName}}
	obj.ID = {{lower $typeName}}ID
	var assoc {{index $pieces 1}}
	var err error
	assoc.ID = {{$lower}}ID
	if err != nil {
		return err
	}
	err = m.DB.Model(&obj).Association("{{index $pieces 0}}").Delete(assoc).Error
	if err != nil {
		return  err
	}
	return  nil
}
func (m *{{$typeName}}DB) Add{{index $pieces 1}}(ctx context.Context, {{lower $typeName}}ID, {{$lower}}ID int) error {
	var {{lower $typeName}} {{$typeName}}
	{{lower $typeName}}.ID = {{lower $typeName}}ID
	var assoc {{index $pieces 1}}
	assoc.ID = {{$lower}}ID
	err := m.DB.Model(&{{lower $typeName}}).Association("{{index $pieces 0}}").Append(assoc).Error
	if err != nil {
		return  err
	}
	return  nil
}
func (m *{{$typeName}}DB) List{{index $pieces 0}}(ctx context.Context, {{lower $typeName}}ID int)  []{{index $pieces 1}} {
	list := make([]{{index $pieces 1}}, 0)
	var obj {{$typeName}}
	obj.ID = {{lower $typeName}}ID
	m.DB.Model(&obj).Association("{{index $pieces 0}}").Find(&list)
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
