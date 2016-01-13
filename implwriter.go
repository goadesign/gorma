package gorma

import (
	"strings"
	"text/template"

	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

// ImplWriter generate code for a goa application media types.
// Media types are data structures used to render the response bodies.
type ImplWriter struct {
	*codegen.GoGenerator
	ImplTmpl *template.Template
}
type ImplData struct {
	TypeDef            *design.UserTypeDefinition
	TypeName           string
	ModelUpper         string
	ModelLower         string
	BelongsTo          []BelongsTo
	M2M                []Many2Many
	PrimaryKeys        map[string]PrimaryKey
	CustomTableName    string
	DoMedia            bool
	DoRoler            bool
	DoCustomTableName  bool
	DoDynamicTableName bool
	DoCache            bool
	APIVersion         string
	RequiredPackages   map[string]bool
}

func NewImplData(version string, utd *design.UserTypeDefinition) ImplData {
	md := ImplData{
		TypeDef:          utd,
		RequiredPackages: make(map[string]bool, 0),
	}
	tn := deModel(codegen.GoTypeName(utd, 0))
	md.TypeName = tn
	md.ModelUpper = upper(tn)
	md.ModelLower = lower(tn)
	if version != "" {
		md.APIVersion = codegen.VersionPackage(version)
	} else {
		md.APIVersion = "app"
	}
	md.PrimaryKeys = getPrimaryKeys(utd)

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

	var m2m []Many2Many
	if m2, ok := metaLookup(utd.Metadata, M2M); ok {
		mlist := strings.Split(m2, ",")
		for _, s := range mlist {
			parms := strings.Split(s, ":")
			if len(parms) == 3 {

				minst := Many2Many{
					Relation:            parms[1],
					LowerRelation:       lower(parms[1]),
					PluralRelation:      parms[0],
					LowerPluralRelation: lower(parms[0]),
					TableName:           parms[2],
				}
				m2m = append(m2m, minst)

				md.RequiredPackages[lower(deModel(parms[1]))] = true
			}
		}

	}
	md.M2M = m2m

	if many, ok := metaLookup(utd.Metadata, "#hasmany"); ok {
		list := strings.Split(many, ",")
		for _, s := range list {
			md.RequiredPackages[lower(s)] = true
		}
	}

	if children, ok := metaLookup(utd.Metadata, "#hasone"); ok {
		list := strings.Split(children, ",")
		for _, s := range list {
			md.RequiredPackages[lower(s)] = true
		}
	}

	if _, ok := metaLookup(utd.Metadata, ROLER); ok {
		md.DoRoler = ok
	}

	if ctn, ok := metaLookup(utd.Metadata, TABLENAME); ok {
		md.CustomTableName = ctn
		md.DoCustomTableName = ok
	}

	if _, ok := metaLookup(utd.Metadata, DYNAMICTABLE); ok {
		md.DoDynamicTableName = ok
	}
	md.DoMedia = true
	if _, ok := metaLookup(utd.Metadata, MEDIA); ok {
		md.DoMedia = !ok
	}
	if _, ok := metaLookup(utd.Metadata, CACHE); ok {
		md.DoCache = ok
	}
	return md
}

// NewImplWriter returns a contexts code writer.
// Media types contain the data used to render response bodies.
func NewImplWriter(filename string) (*ImplWriter, error) {
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
	funcMap["title"] = titleCase
	funcMap["plural"] = plural
	funcMap["metaLookup"] = metaLookupTmpl
	funcMap["columns"] = GetAttributeColumns
	funcMap["pkattributes"] = pkAttributes
	funcMap["pkwhere"] = pkWhere
	funcMap["pkwherefields"] = pkWhereFields
	funcMap["pkupdatefields"] = pkUpdateFields
	implTmpl, err := template.New("implementations").Funcs(funcMap).Parse(implTmpl)
	if err != nil {
		return nil, err
	}
	w := ImplWriter{
		GoGenerator: cw,
		ImplTmpl:    implTmpl,
	}
	return &w, nil
}

// Execute writes the code for the context types to the writer.
func (w *ImplWriter) Execute(md *ImplData) error {
	return w.ImplTmpl.Execute(w, md)
}
