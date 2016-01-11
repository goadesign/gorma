package gorma

const implTmpl = `// {{if .TypeDef.Description}}{{.TypeDef.Description}}{{else}}app.{{ .TypeName}} model type{{end}}
// Identifier: {{ .TypeName}}
type {{.TypeName}} struct {
	{{.ModelLower}}.{{.TypeName}}
}
{{end}}
type {{.TypeName}}DB struct {
	{{.ModelLower}}.{{.TypeName}}DB
}
{{end}}
`
