package gorma

import (
	"text/template"

	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

// RbacWriter generate code for a goa application media types.
// Media types are data structures used to render the response bodies.
type RbacWriter struct {
	*codegen.GoGenerator
	RbacTmpl *template.Template
}

// NewRbacWriter returns a contexts code writer.
// Media types contain the data used to render response bodies.
func NewRbacWriter(filename string) (*RbacWriter, error) {
	cw := codegen.NewGoGenerator(filename)
	funcMap := cw.FuncMap
	funcMap["gotypedef"] = codegen.GoTypeDef
	funcMap["gotyperef"] = codegen.GoTypeRef
	funcMap["goify"] = codegen.Goify
	funcMap["gotypename"] = codegen.GoTypeName
	funcMap["gonative"] = codegen.GoNativeType
	funcMap["typeUnmarshaler"] = codegen.TypeUnmarshaler
	funcMap["typeMarshaler"] = codegen.MediaTypeMarshaler
	funcMap["recursiveValidate"] = codegen.RecursiveChecker
	funcMap["tempvar"] = codegen.Tempvar
	funcMap["demodel"] = DeModel
	funcMap["modeldef"] = MakeModelDef
	funcMap["snake"] = CamelToSnake
	funcMap["lower"] = Lower
	funcMap["upper"] = Upper

	rbacTmpl, err := template.New("rbac").Funcs(funcMap).Parse(rbacTmpl)
	if err != nil {
		return nil, err
	}
	w := RbacWriter{
		GoGenerator: cw,
		RbacTmpl:    rbacTmpl,
	}
	return &w, nil
}

// Execute writes the code for the context types to the writer.
func (w *RbacWriter) Execute(mt *design.APIDefinition) error {
	return w.RbacTmpl.Execute(w, mt)
}
