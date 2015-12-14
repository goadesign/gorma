package gorma

const intTmpl = `
 {{ $typeName :=  gotypename . 0}}{{$typeName := demodel $typeName}}
type {{$typeName}}Storage interface {
	List(ctx *app.List{{demodel $typeName}}Context) []{{$typeName}}
	Get(ctx *app.Show{{demodel $typeName }}Context) ({{$typeName}}, error)
	Add(ctx *app.Create{{demodel $typeName}}Context) ({{$typeName}}, error)
	Update(ctx *app.Update{{demodel $typeName}}Context) (error)
	Delete(ctx *app.Delete{{demodel $typeName}}Context) (error)
}
`
