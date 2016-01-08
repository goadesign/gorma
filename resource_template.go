package gorma

const resourceTmpl = `
{{ $typename  := .TypeName }}
{{ $belongs := .BelongsTo }}
{{ if .DoMedia }}
{{ $version := .APIVersion }}
{{ range $idx, $action := .TypeDef.Actions  }}
{{ if hasusertype $action }}
func {{$typename}}From{{version $version}}{{title $action.Name}}Payload(ctx *{{$version}}.{{title $action.Name}}{{$typename}}Context) {{$typename}} {
	payload := ctx.Payload
	m := {{$typename}}{}
	copier.Copy(&m, payload)
{{ range $idx, $bt := $belongs }}
	m.{{ $bt.Parent}}ID=int(ctx.{{ $bt.Parent}}ID){{end}}
	return m
}
{{ end }}{{end}}{{end}}
`
