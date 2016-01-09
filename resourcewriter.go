package gorma

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

// MediaWriter generate code for a goa application media types.
// Media types are data structures used to render the response bodies.
type ResourceWriter struct {
	*codegen.GoGenerator
	ResourceTmpl *template.Template
}
type ResourceData struct {
	TypeDef    *design.ResourceDefinition
	TypeName   string
	MediaUpper string
	MediaLower string
	BelongsTo  []BelongsTo
	DoMedia    bool
	APIVersion string
}

func NewResourceData(version string, utd *design.ResourceDefinition) ResourceData {
	md := ResourceData{
		TypeDef: utd,
	}
	md.TypeName = codegen.Goify(utd.Name, true)
	md.MediaUpper = upper(utd.Name)
	md.MediaLower = lower(utd.Name)
	if version != "" {
		md.APIVersion = codegen.VersionPackage(version)
	} else {
		// import the default package instead of nothing
		md.APIVersion = "app"
	}

	var belongs []BelongsTo
	if bt, ok := metaLookup(utd.Metadata, BELONGSTO); ok {
		btlist := strings.Split(bt, ",")
		for _, s := range btlist {
			binst := BelongsTo{
				Parent:        s,
				DatabaseField: camelToSnake(s),
			}
			belongs = append(belongs, binst)
		}
	}
	md.BelongsTo = belongs
	md.DoMedia = true
	if _, ok := metaLookup(utd.Metadata, MEDIA); ok {
		md.DoMedia = !ok
	}
	return md
}

// NewMediaWriter returns a contexts code writer.
// Media types contain the data used to render response bodies.
func NewResourceWriter(filename string) (*ResourceWriter, error) {
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
	funcMap["demodel"] = deModel
	funcMap["modeldef"] = ModelDef
	funcMap["snake"] = camelToSnake
	funcMap["split"] = split
	funcMap["storagedef"] = StorageDef
	funcMap["lower"] = lower
	funcMap["upper"] = upper
	funcMap["title"] = titleCase
	funcMap["plural"] = plural
	funcMap["metaLookup"] = metaLookupTmpl
	funcMap["columns"] = GetAttributeColumns
	funcMap["version"] = versionize
	funcMap["hasusertype"] = hasUserType

	modelTmpl, err := template.New("media").Funcs(funcMap).Parse(resourceTmpl)
	if err != nil {
		return nil, err
	}
	w := ResourceWriter{
		GoGenerator:  cw,
		ResourceTmpl: modelTmpl,
	}
	return &w, nil
}

// Execute writes the code for the context types to the writer.
func (w *ResourceWriter) Execute(version string, mt *design.ResourceDefinition) error {
	md := NewResourceData(version, mt)
	err := w.ResourceTmpl.Execute(w, md)
	if err != nil {
		fmt.Println("Error executing template", err.Error())
	}
	return err
}
