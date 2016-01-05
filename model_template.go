package gorma

const modelTmpl = `// {{if .Description}}{{.Description}}{{else}}app.{{gotypename . 0}} storage type{{end}}
// Identifier: {{ $typeName :=  gotypename . 0}}{{$typeName := demodel $typeName}}
type {{$typeName}} {{ modeldef . }}
{{ $belongsto := metaLookup .Metadata "#belongsto" }}
{{ $m2m := metaLookup .Metadata "#many2many" }}
{{ $nomedia := metaLookup .Metadata "#nomedia" }}
{{ if eq $nomedia "" }}
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
{{ end }}
{{ $tablename := metaLookup .Metadata "#tablename" }}
{{ if ne $tablename "" }}
func (m {{$typeName}}) TableName() string {
	return "{{ $tablename }}"
}
{{ end }}
{{ $roler := metaLookup .Metadata "#roler" }}
{{ if ne $roler "" }}
func (m {{$typeName}}) GetRole() string {
	return m.Role
}
{{end}}

{{ $dyntablename := metaLookup .Metadata "#dyntablename" }}

type {{$typeName}}Storage interface {
	DB() interface{}
	List(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}) []{{$typeName}}
	One(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, id int) ({{$typeName}}, error)
	Add(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, o {{$typeName}}) ({{$typeName}}, error)
	Update(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, o {{$typeName}}) (error)
	Delete(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, id int) (error)
{{ if ne $belongsto "" }}{{$barray := split $belongsto ","}}{{ range $idx, $bt := $barray}}
	ListBy{{$bt}}(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, parentid int) []{{$typeName}}
	OneBy{{$bt}}(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, parentid, id int) ({{$typeName}}, error)
{{end}}{{end}}
	{{ storagedef . }}
}
{{ $cached := metaLookup .Metadata "#cached" }}
type {{$typeName}}DB struct {
	db gorm.DB
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

func (m *{{$typeName}}DB) ListBy{{$bt}}(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, parentid int) []{{$typeName}} {

	var objs []{{$typeName}}
	m.db{{ if ne $dyntablename "" }}.Table(tableName){{ end }}.Scopes({{$typeName}}FilterBy{{$bt}}(parentid, &m.db)).Find(&objs)
	return objs
}

func (m *{{$typeName}}DB) OneBy{{$bt}}(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, parentid, id int) ({{$typeName}}, error) {
	{{ if ne $cached "" }}//first attempt to retrieve from cache
	o,found := m.cache.Get(strconv.Itoa(id))
	if found {
		return o.({{$typeName}}), nil
	}
	// fallback to database if not found{{ end }}
	var obj {{$typeName}}

	err := m.db{{ if ne $dyntablename "" }}.Table(tableName){{ end }}.Scopes({{$typeName}}FilterBy{{$bt}}(parentid, &m.db)).Find(&obj, id).Error
	{{ if ne $cached "" }} go m.cache.Set(strconv.Itoa(id), obj, cache.DefaultExpiration) {{ end }}
	return obj, err
}
{{end}}{{end}}

func New{{$typeName}}DB(db gorm.DB) *{{$typeName}}DB {
	{{ if ne $cached "" }}
	return &{{$typeName}}DB{
		db: db,
		cache: cache.New(5*time.Minute, 30*time.Second),
	}
	{{ else  }}
	return &{{$typeName}}DB{db: db}

	{{ end  }}
}

func (m *{{$typeName}}DB) DB() interface{} {
	return &m.db
}

func (m *{{$typeName}}DB) List(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}) []{{$typeName}} {

	var objs []{{$typeName}}
	m.db{{ if ne $dyntablename "" }}.Table(tableName){{ end }}.Find(&objs)
	return objs
}


{{ range $idx, $col := columns .AttributeDefinition }} 
func (m *{{$typeName}}DB) ListBy{{title $col.Column}}Equal(ctx context.Context, {{lower $col.Column}} {{$col.Coltype}}, {{ if ne $dyntablename "" }}, tableName string{{ end }}) []{{$typeName}} {

	var objs []{{$typeName}}
	m.DB.Where("{{lower $col.Column}} = ?",  {{lower $col.Column}}){{ if ne $dyntablename "" }}.Table(tableName){{ end }}.Find(&objs)
	return objs
}
func (m *{{$typeName}}DB) ListBy{{title $col.Column}}Like(ctx context.Context, {{lower $col.Column}} {{$col.Coltype}}, {{ if ne $dyntablename "" }}, tableName string{{ end }}) []{{$typeName}} {

	var objs []{{$typeName}}
	m.DB.Where("{{lower $col.Column}} like ?",  {{lower $col.Column}}){{ if ne $dyntablename "" }}.Table(tableName){{ end }}.Find(&objs)
	return objs
}
{{ end  }}


func (m *{{$typeName}}DB) One(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, id int) ({{$typeName}}, error) {
	{{ if ne $cached "" }}//first attempt to retrieve from cache
	o,found := m.cache.Get(strconv.Itoa(id))
	if found {
		return o.({{$typeName}}), nil
	}
	// fallback to database if not found{{ end }}
	var obj {{$typeName}}

	err := m.db{{ if ne $dyntablename "" }}.Table(tableName){{ end }}.Find(&obj, id).Error
	{{ if ne $cached "" }} go m.cache.Set(strconv.Itoa(id), obj, cache.DefaultExpiration) {{ end }}
	return obj, err
}

func (m *{{$typeName}}DB) Add(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, model {{$typeName}}) ({{$typeName}}, error) {
	err := m.db{{ if ne $dyntablename "" }}.Table(tableName){{ end }}.Create(&model).Error
	{{ if ne $cached "" }} go m.cache.Set(strconv.Itoa(model.ID), model, cache.DefaultExpiration) {{ end }}
	return model, err
}

func (m *{{$typeName}}DB) Update(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, model {{$typeName}}) error {
	obj, err := m.One(ctx{{ if ne $dyntablename "" }}, tableName{{ end }}, model.ID)
	if err != nil {
		return  err
	}
	err = m.db{{ if ne $dyntablename "" }}.Table(tableName){{ end }}.Model(&obj).Updates(model).Error
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


func (m *{{$typeName}}DB) Delete(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, id int)  error {
	var obj {{$typeName}}
	err := m.db{{ if ne $dyntablename "" }}.Table(tableName){{ end }}.Delete(&obj, id).Error
	if err != nil {
		return  err
	}
	{{ if ne $cached "" }} go m.cache.Delete(strconv.Itoa(id)) {{ end }}
	return  nil
}

{{ if ne $m2m "" }}{{$barray := split $m2m ","}}{{ range $idx, $bt := $barray}}
{{ $pieces := split $bt ":" }} {{ $lowertype := index $pieces 1  }} {{ $lower := lower $lowertype }}  {{ $lowerplural := index $pieces 0  }} {{ $lowerplural := lower $lowerplural}}
func (m *{{$typeName}}DB) Delete{{index $pieces 1}}(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }},{{lower $typeName}}ID,  {{$lower}}ID int)  error {
	var obj {{$typeName}}
	obj.ID = {{lower $typeName}}ID
	var assoc {{index $pieces 1}}
	var err error
	assoc.ID = {{$lower}}ID
	if err != nil {
		return err
	}
	err = m.db{{ if ne $dyntablename "" }}.Table(tableName){{ end }}.Model(&obj).Association("{{index $pieces 0}}").Delete(assoc).Error
	if err != nil {
		return  err
	}
	return  nil
}
func (m *{{$typeName}}DB) Add{{index $pieces 1}}(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, {{lower $typeName}}ID, {{$lower}}ID int) error {
	var {{lower $typeName}} {{$typeName}}
	{{lower $typeName}}.ID = {{lower $typeName}}ID
	var assoc {{index $pieces 1}}
	assoc.ID = {{$lower}}ID
	err := m.db{{ if ne $dyntablename "" }}.Table(tableName){{ end }}.Model(&{{lower $typeName}}).Association("{{index $pieces 0}}").Append(assoc).Error
	if err != nil {
		return  err
	}
	return  nil
}
func (m *{{$typeName}}DB) List{{index $pieces 0}}(ctx context.Context{{ if ne $dyntablename "" }}, tableName string{{ end }}, {{lower $typeName}}ID int)  []{{index $pieces 1}} {
	list := make([]{{index $pieces 1}}, 0)
	var obj {{$typeName}}
	obj.ID = {{lower $typeName}}ID
	m.db{{ if ne $dyntablename "" }}.Table(tableName){{ end }}.Model(&obj).Association("{{index $pieces 0}}").Find(&list)
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
