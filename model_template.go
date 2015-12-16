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
	Get(ctx goa.Context) ({{$typeName}}, error)
	Add(ctx goa.Context) ({{$typeName}}, error)
	Update(ctx goa.Context) (error)
	Delete(ctx goa.Context) (error)
	{{ storagedef . }}
}

type {{$typeName}}DB struct {
	DB gorm.DB
}
{{ if ne $belongsto "" }}{{$barray := split $belongsto ","}}{{ range $idx, $bt := $barray}}
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
func New{{$typeName}}DB(db gorm.DB) *{{$typeName}}DB {
	return &{{$typeName}}DB{DB: db}
}

func (m *{{$typeName}}DB) List(ctx goa.Context) []{{$typeName}} {

	var objs []{{$typeName}}
    {{ if ne $belongsto "" }}m.DB.Scopes({{$typeName}}Filter(ctx.{{demodel $belongsto}}ID, &m.DB)).Find(&objs){{ else }} m.DB.Find(&objs) {{end}}
	return objs
}

func (m *{{$typeName}}DB) Get(ctx goa.Context) ({{$typeName}}, error) {

	var obj {{$typeName}}

	err := m.DB.Find(&obj, ctx.{{demodel $typeName}}ID).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return obj, err
}

func (m *{{$typeName}}DB) Add(ctx goa.Context) ({{$typeName}}, error) {
	model := {{$typeName}}FromCreatePayload(ctx)
	err := m.DB.Create(&model).Error
	return model, err
}
func (m *{{$typeName}}DB) Update(ctx goa.Context) error {
	obj, err := m.Get(ctx)
	if err != nil {
		return  err
	}
	err = m.DB.Model(&obj).Updates({{$typeName}}FromUpdatePayload(ctx)).Error
	if err != nil {
		ctx.Error(err.Error())
	}
	return err
}
func (m *{{$typeName}}DB) Delete(ctx goa.Context)  error {
	var obj {{$typeName}}
	err := m.DB.Delete(&obj, ctx.{{demodel $typeName}}ID).Error
	if err != nil {
		ctx.Logger.Error(err.Error())
		return  err
	}
	return  nil
}

{{ if ne $m2m "" }}{{$barray := split $m2m ","}}{{ range $idx, $bt := $barray}}
{{ $pieces := split $bt ":" }} {{ $lowertype := index $pieces 1  }} {{ $lower := lower $lowertype }}  {{ $lowerplural := index $pieces 0  }} {{ $lowerplural := lower $lowerplural}}
func (m *{{$typeName}}DB) Delete{{index $pieces 1}}(ctx goa.Context)  error {
	var obj {{$typeName}}

	assoc_id := ctx.{{index $pieces 1}}ID
	var assoc {{index $pieces 1}}
	var err error
	assoc.ID = assoc_id
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
func (m *{{$typeName}}DB) Add{{index $pieces 1}}(ctx goa.Context) error {
	var obj {{$typeName}}
	assoc_id, err  := strconv.Atoi(ctx.Payload.{{index $pieces 1}}Id)
	if err != nil {
		return err
	}
	var assoc {{index $pieces 1}}
	assoc.ID = assoc_id
	err = m.DB.Model(&obj).Association("{{index $pieces 0}}").Append(assoc).Error
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

type Mock{{$typeName}}Storage struct {
	{{$typeName}}List  map[int]{{$typeName}}
	nextID int
	mut sync.Mutex
}
{{if ne $belongsto ""}}{{$barray := split $belongsto ","}}{{ range $idx, $bt := $barray}}
func filter{{$typeName}}By{{$bt}}(parent int, list []{{$typeName}}) []{{$typeName}} {
	filtered := make([]{{$typeName}},0)
	for _,o := range list {
		if o.{{$bt}}ID == int(parent) {
			filtered = append(filtered,o)
		}
	}
	return filtered
}
{{end}}{{end}}


func NewMock{{$typeName}}Storage() *Mock{{$typeName}}Storage {
	ml := make(map[int]{{$typeName}}, 0)
	return &Mock{{$typeName}}Storage{ {{$typeName}}List: ml}
}

func (db *Mock{{$typeName}}Storage) List(ctx goa.Context) []{{$typeName}} {
	var list []{{$typeName}} = make([]{{$typeName}}, 0)
	for _, v := range db.{{$typeName}}List {
		list = append(list, v)
	}
{{if ne $belongsto ""}}
return filter{{$typeName}}By{{$belongsto}}(ctx.{{$belongsto}}ID, list) {{else}}return list{{end}}
}

func (db *Mock{{$typeName}}Storage) Get(ctx goa.Context) ({{$typeName}}, error) {

	var obj {{$typeName}}

	obj, ok := db.{{$typeName}}List[int(ctx.{{demodel $typeName}}ID)]
	if ok {
		return obj, nil
	} else {
		return obj, errors.New("{{$typeName}} does not exist")
	}
}

func (db *Mock{{$typeName}}Storage) Add(ctx goa.Context)  ({{$typeName}}, error) {
	u := {{$typeName}}FromCreatePayload(ctx)
	db.mut.Lock()
	db.nextID = db.nextID + 1
	u.ID = db.nextID
	db.mut.Unlock()

	db.{{$typeName}}List[u.ID] = u
	return u, nil
}

func (db *Mock{{$typeName}}Storage) Update(ctx goa.Context) error {
	id := int(ctx.{{demodel $typeName}}ID)
	_, ok := db.{{$typeName}}List[id]
	if ok {
		db.{{$typeName}}List[id] = {{$typeName}}FromUpdatePayload(ctx)
		return  nil
	} else {
		return errors.New("{{$typeName}} does not exist")
	}
}

func (db *Mock{{$typeName}}Storage) Delete(ctx goa.Context)  error {
	_, ok := db.{{$typeName}}List[int(ctx.{{demodel $typeName}}ID)]
	if ok {
		delete(db.{{$typeName}}List, int(ctx.{{demodel $typeName}}ID))
		return  nil
	} else {
		return  errors.New("Could not delete this user")
	}
}
`
