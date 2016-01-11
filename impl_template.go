package gorma

const implTmpl = `// {{if .TypeDef.Description}}{{.TypeDef.Description}}{{else}}app.{{ .TypeName}} model type{{end}}
// Identifier: {{ .TypeName}}
{{ $dynamictable := .DoDynamicTableName }}
{{ $typename  := .TypeName }}
{{ $typedef := .TypeDef  }}
{{ $pks := .PrimaryKeys }}
type {{$typename}}Storage interface {
	DB() interface{}
	List(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}) []{{$typename}}
	One(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, {{ pkattributes $pks  }}) ({{$typename}}, error)
	Add(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, o {{$typename}}) ({{$typename}}, error)
	Update(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, o {{$typename}}) (error)
	Delete(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, {{ pkattributes $pks }}) (error)
{{ range $idx, $bt := .BelongsTo}}
	ListBy{{$bt.Parent}}(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, parentid int) []{{$typename}}
	OneBy{{$bt.Parent}}(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, parentid, id int) ({{$typename}}, error)
{{end}}
	{{ storagedef $typedef }}
}

func New{{.TypeName}}DB(db gorm.DB) *{{.TypeName}}DB {

	return &{{.ModelLower}}.{{.TypeName}}DB{db: db}

}
type {{.TypeName}} struct {
	{{.ModelLower}}.{{.TypeName}}
}
type {{.TypeName}}DB struct {
	{{.ModelLower}}.{{.TypeName}}DB
}
`
