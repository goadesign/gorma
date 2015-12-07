package gorma

const rbacTmpl = `
const (
	ADMIN        = "Admin"
	USER         = "User"

		{{  range $idx, $res := .Resources }}{{$resname := $res.Name}}{{  range $actidx, $act := .Actions}}{{$actname := $act.Name}}{{ upper $res.Name}}{{upper $actname}} = "{{ lower $res.Name}}.{{lower $actname}}"
		{{end}}{{end}}
)

var RBAC *gorbac.RBAC

// These are provided as a template.  Edit to suit as required by your applicaton
// Test roles in your controllers with models.RBAC.isGranted(ROLE, SOMEPERMISSION, nil)
func init() {
	RBAC  = gorbac.New()
	RBAC.Add(USER, []string{
		{{  range $idx, $res := .Resources }}{{$resname := $res.Name}}{{  range $actidx, $act := .Actions}}{{$actname := $act.Name}}{{ upper $res.Name}}{{upper $actname}},
		{{end}}{{end}}
		}, nil)
	RBAC.Add(ADMIN, []string{
		{{  range $idx, $res := .Resources }}{{$resname := $res.Name}}{{  range $actidx, $act := .Actions}}{{$actname := $act.Name}}{{ upper $res.Name}}{{upper $actname}},
		{{end}}{{end}}
	}, []string{USER})
}
`
