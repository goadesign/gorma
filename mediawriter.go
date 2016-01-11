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
type MediaWriter struct {
	*codegen.GoGenerator
	MediaTmpl *template.Template
}
type MediaData struct {
	TypeDef          *design.MediaTypeDefinition
	TypeName         string
	MediaUpper       string
	MediaLower       string
	BelongsTo        []BelongsTo
	DoMedia          bool
	APIVersion       string
	RequiredPackages map[string]bool
}

func NewMediaData(version string, utd *design.MediaTypeDefinition) MediaData {
	md := MediaData{
		TypeDef:          utd,
		RequiredPackages: make(map[string]bool, 0),
	}
	md.TypeName = codegen.Goify(utd.TypeName, true)
	md.MediaUpper = upper(utd.Name())
	md.MediaLower = lower(utd.Name())
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

			md.RequiredPackages[lower(s)] = true
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
func NewMediaWriter(filename string) (*MediaWriter, error) {
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

	modelTmpl, err := template.New("media").Funcs(funcMap).Parse(mediaTmpl)
	if err != nil {
		return nil, err
	}
	w := MediaWriter{
		GoGenerator: cw,
		MediaTmpl:   modelTmpl,
	}
	return &w, nil
}

// Execute writes the code for the context types to the writer.
func (w *MediaWriter) Execute(md *MediaData) error {
	err := w.MediaTmpl.Execute(w, md)
	if err != nil {
		fmt.Println("Error executing template", err.Error())
	}
	return err
}
