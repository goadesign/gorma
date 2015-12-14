package gorma

import (
	"text/template"

	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

type InterfaceWriter struct {
	*codegen.GoGenerator
	InterfaceTmpl *template.Template
}

func NewInterfaceWriter(filename string) (*InterfaceWriter, error) {
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

	intTmpl, err := template.New("interfaces").Funcs(funcMap).Parse(intTmpl)
	if err != nil {
		return nil, err
	}
	w := InterfaceWriter{
		GoGenerator:   cw,
		InterfaceTmpl: intTmpl,
	}
	return &w, nil
}

// Execute writes the code for the context types to the writer.
func (w *InterfaceWriter) Execute(mt *design.APIDefinition) error {
	return w.InterfaceTmpl.Execute(w, mt)
}
