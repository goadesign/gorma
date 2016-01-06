package gorma

import (
	"strings"
	"text/template"

	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

// ModelWriter generate code for a goa application media types.
// Media types are data structures used to render the response bodies.
type ModelWriter struct {
	*codegen.GoGenerator
	ModelTmpl *template.Template
}
type Field struct {
	Column  string
	Coltype string
}
type BelongsTo struct {
	Parent        string
	DatabaseField string
}
type Many2Many struct {
	Relation            string
	LowerRelation       string
	PluralRelation      string
	LowerPluralRelation string
	TableName           string
}
type ModelData struct {
	TypeDef            *design.UserTypeDefinition
	TypeName           string
	ModelUpper         string
	ModelLower         string
	BelongsTo          []BelongsTo
	M2M                []Many2Many
	CustomTableName    string
	DynamicTableName   bool
	DoMedia            bool
	DoRoler            bool
	DoCustomTableName  bool
	DoDynamicTableName bool
	DoCache            bool
}

func NewModelData(utd *design.UserTypeDefinition) ModelData {
	md := ModelData{
		TypeDef: utd,
	}
	tn := deModel(codegen.GoTypeName(utd, 0))
	md.TypeName = tn
	md.ModelUpper = upper(tn)
	md.ModelLower = lower(tn)

	belongs := make([]BelongsTo, 0)
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

	m2m := make([]Many2Many, 0)
	if m2, ok := metaLookup(utd.Metadata, M2M); ok {
		mlist := strings.Split(m2, ",")
		for _, s := range mlist {
			parms := strings.Split(s, ":")

			minst := Many2Many{
				Relation:            parms[0],
				LowerRelation:       lower(parms[0]),
				PluralRelation:      parms[1],
				LowerPluralRelation: lower(parms[1]),
				TableName:           parms[2],
			}
			m2m = append(m2m, minst)
		}

	}
	md.M2M = m2m

	if _, ok := metaLookup(utd.Metadata, ROLER); ok {
		md.DoRoler = ok
	}

	if ctn, ok := metaLookup(utd.Metadata, TABLENAME); ok {
		md.CustomTableName = ctn
		md.DoCustomTableName = ok
	}

	if _, ok := metaLookup(utd.Metadata, DYNAMICTABLE); ok {
		md.DynamicTableName = ok
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

// NewModelWriter returns a contexts code writer.
// Media types contain the data used to render response bodies.
func NewModelWriter(filename string) (*ModelWriter, error) {
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

	modelTmpl, err := template.New("models").Funcs(funcMap).Parse(modelTmpl)
	if err != nil {
		return nil, err
	}
	w := ModelWriter{
		GoGenerator: cw,
		ModelTmpl:   modelTmpl,
	}
	return &w, nil
}

// Execute writes the code for the context types to the writer.
func (w *ModelWriter) Execute(mt *design.UserTypeDefinition) error {
	md := NewModelData(mt)
	return w.ModelTmpl.Execute(w, md)
}
