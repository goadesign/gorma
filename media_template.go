package gorma

const mediaTmpl = `
{{ $typename  := .TypeName }}
{{ if .DoMedia }}
func (m {{$typename}}) To{{version .APIVersion}}() *{{.APIVersion}}.{{$typename}} {
	target := {{.APIVersion}}.{{$typename}}{}
	copier.Copy(&target, &m)
	return &target
}
{{ end }}
`
