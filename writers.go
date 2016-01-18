package gorma

import (
	"fmt"
	"regexp"
	"strings"
	"text/template"

	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

// WildcardRegex is the regex used to capture path parameters.
var WildcardRegex = regexp.MustCompile("(?:[^/]*/:([^/]+))+")
var mUserImplT *template.Template

func init() {
	var err error
	fm := template.FuncMap{
		"gotypename":   codegen.GoTypeName,
		"gotyperef":    GoTypeRef,
		"gopkgtyperef": GoPackageTypeRef,
		"goify":        codegen.Goify,
		"gonative":     codegen.GoNativeType,
		"add":          func(a, b int) int { return a + b },
	}
	if mUserImplT, err = template.New("user marshaler").Funcs(fm).Parse(mUserImplTmpl); err != nil {
		panic(err)
	}
}

type (
	// ContextsWriter generate codes for a goa application contexts.
	ContextsWriter struct {
		*codegen.GoGenerator
		CtxTmpl        *template.Template
		CtxNewTmpl     *template.Template
		CtxRespTmpl    *template.Template
		PayloadTmpl    *template.Template
		NewPayloadTmpl *template.Template
		ConversionTmpl *template.Template
	}

	// MediaTypesWriter generate code for a goa application media types.
	// Media types are data structures used to render the response bodies.
	MediaTypesWriter struct {
		*codegen.GoGenerator
		MediaTypeTmpl *template.Template
	}

	// UserTypesWriter generate code for a goa application user types.
	// User types are data structures defined in the DSL with "Type".
	UserTypesWriter struct {
		*codegen.GoGenerator
		UserTypeTmpl *template.Template
	}

	// ContextTemplateData contains all the information used by the template to render the context
	// code for an action.
	ContextTemplateData struct {
		Name         string // e.g. "ListBottleContext"
		ResourceName string // e.g. "bottles"
		ActionName   string // e.g. "list"
		Params       *design.AttributeDefinition
		Payload      *design.UserTypeDefinition
		Headers      *design.AttributeDefinition
		Routes       []*design.RouteDefinition
		Responses    map[string]*design.ResponseDefinition
		API          *design.APIDefinition
		Version      *design.APIVersionDefinition
		DefaultPkg   string
	}

	// MediaTypeTemplateData contains all the information used by the template to redner the
	// media types code.
	MediaTypeTemplateData struct {
		MediaType  *design.UserTypeDefinition
		Versioned  bool
		DefaultPkg string
	}

	// UserTypeTemplateData contains all the information used by the template to redner the
	// media types code.
	UserTypeTemplateData struct {
		UserType   *RelationalModel
		BaseType   *design.UserTypeDefinition
		DefaultPkg string
	}

	BelongsTo struct {
		Parent        string
		DatabaseField string
	}
	Many2Many struct {
		Relation            string
		LowerRelation       string
		PluralRelation      string
		LowerPluralRelation string
		TableName           string
	}
	// ControllerTemplateData contains the information required to generate an action handler.
	ControllerTemplateData struct {
		Resource string                       // Lower case plural resource name, e.g. "bottles"
		Actions  []map[string]interface{}     // Array of actions, each action has keys "Name", "Routes" and "Context"
		Version  *design.APIVersionDefinition // Controller API version
	}

	// ResourceData contains the information required to generate the resource GoGenerator
	ResourceData struct {
		Name              string                      // Name of resource
		Identifier        string                      // Identifier of resource media type
		Description       string                      // Description of resource
		Type              *design.MediaTypeDefinition // Type of resource media type
		CanonicalTemplate string                      // CanonicalFormat represents the resource canonical path in the form of a fmt.Sprintf format.
		CanonicalParams   []string                    // CanonicalParams is the list of parameter names that appear in the resource canonical path in order.
	}

	ConversionData struct {
		TypeDef          *design.ResourceDefinition
		TypeName         string
		MediaUpper       string
		MediaLower       string
		BelongsTo        []BelongsTo
		DoMedia          bool
		APIVersion       string
		RequiredPackages map[string]bool
		Version          *design.APIVersionDefinition
		DefaultPkg       string
	}
)

// Versioned returns true if the context was built from an API version.
func (c *ConversionData) Versioned() bool {
	return !c.Version.IsDefault()
}

// Versioned returns true if the context was built from an API version.
func (c *ContextTemplateData) Versioned() bool {
	return !c.Version.IsDefault()
}
func NewConversionData(version *design.APIVersionDefinition, utd *design.ResourceDefinition) ConversionData {
	md := ConversionData{
		TypeDef:          utd,
		RequiredPackages: make(map[string]bool, 0),
	}
	md.Version = version
	md.TypeName = codegen.Goify(utd.Name, true)
	md.MediaUpper = strings.Title(utd.Name)
	md.MediaLower = lower(utd.Name)
	if version.Version != "" {
		md.APIVersion = codegen.VersionPackage(version.Version)
	} else {
		// import the default package instead of nothing
		md.APIVersion = "app"
	}

	var belongs []BelongsTo
	if bt, ok := metaLookup(utd.Metadata, "#belongsto"); ok {
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
	if _, ok := metaLookup(utd.Metadata, "#nomedia"); ok {
		md.DoMedia = !ok
	}
	return md
}

// IsPathParam returns true if the given parameter name corresponds to a path parameter for all
// the context action routes. Such parameter is required but does not need to be validated as
// httprouter takes care of that.
func (c *ContextTemplateData) IsPathParam(param string) bool {
	params := c.Params
	pp := false
	if params.Type.IsObject() {
		for _, r := range c.Routes {
			pp = false
			for _, p := range r.Params(c.Version) {
				if p == param {
					pp = true
					break
				}
			}
			if !pp {
				break
			}
		}
	}
	return pp
}

// MustValidate returns true if code that checks for the presence of the given param must be
// generated.
func (c *ContextTemplateData) MustValidate(name string) bool {
	return c.Params.IsRequired(name) && !c.IsPathParam(name)
}

// NewContextsWriter returns a contexts code writer.
// Contexts provide the glue between the underlying request data and the user controller.
// Used
func NewContextsWriter(filename string) (*ContextsWriter, error) {
	cw := codegen.NewGoGenerator(filename)
	funcMap := cw.FuncMap
	funcMap["gotyperef"] = GoTypeRef
	funcMap["gotypedef"] = codegen.GoTypeDef
	funcMap["goify"] = codegen.Goify
	funcMap["gotypename"] = codegen.GoTypeName
	funcMap["gopkgtypename"] = codegen.GoPackageTypeName
	funcMap["typeUnmarshaler"] = codegen.TypeUnmarshaler
	funcMap["userTypeUnmarshalerImpl"] = codegen.UserTypeUnmarshalerImpl
	funcMap["userTypeMarshalerImpl"] = UserTypeMarshalerImpl
	funcMap["validationChecker"] = codegen.ValidationChecker
	funcMap["tabs"] = codegen.Tabs
	funcMap["add"] = func(a, b int) int { return a + b }
	funcMap["gopkgtyperef"] = codegen.GoPackageTypeRef
	funcMap["hasusertype"] = hasUserType
	funcMap["version"] = versionize
	funcMap["title"] = titleCase
	conversionTmpl, err := template.New("conversion").Funcs(funcMap).Parse(resourceT)
	if err != nil {
		return nil, err
	}

	w := ContextsWriter{
		GoGenerator:    cw,
		ConversionTmpl: conversionTmpl,
	}
	return &w, nil
}

// Execute writes the code for the context types to the writer.
func (w *ContextsWriter) Execute(data *ConversionData) error {

	if err := w.ConversionTmpl.Execute(w, data); err != nil {
		return err
	}

	return nil
}

// NewMediaTypesWriter returns a contexts code writer.
// Media types contain the data used to render response bodies.
func NewMediaTypesWriter(filename string) (*MediaTypesWriter, error) {
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
	funcMap["newDumpData"] = newDumpData
	funcMap["userTypeUnmarshalerImpl"] = codegen.UserTypeUnmarshalerImpl
	funcMap["userTypeMarshalerImpl"] = UserTypeMarshalerImpl
	mediaTypeTmpl, err := template.New("media type").Funcs(funcMap).Parse(mediaTypeT)
	if err != nil {
		return nil, err
	}
	w := MediaTypesWriter{
		GoGenerator:   cw,
		MediaTypeTmpl: mediaTypeTmpl,
	}
	return &w, nil
}

// Execute writes the code for the context types to the writer.
func (w *MediaTypesWriter) Execute(data *MediaTypeTemplateData) error {
	data.MediaType.TypeName = deModel(data.MediaType.TypeName)
	return w.MediaTypeTmpl.Execute(w, data)
}

// NewUserTypesWriter returns a contexts code writer.
// User types contain custom data structured defined in the DSL with "Type".
func NewUserTypesWriter(filename string) (*UserTypesWriter, error) {
	cw := codegen.NewGoGenerator(filename)
	funcMap := cw.FuncMap
	funcMap["gotypedef"] = codegen.GoTypeDef
	funcMap["gotyperef"] = codegen.GoTypeRef
	funcMap["goify"] = codegen.Goify
	funcMap["gotypename"] = codegen.GoTypeName
	funcMap["recursiveValidate"] = codegen.RecursiveChecker
	funcMap["userTypeUnmarshalerImpl"] = codegen.UserTypeUnmarshalerImpl
	funcMap["userTypeMarshalerImpl"] = codegen.UserTypeMarshalerImpl
	funcMap["pkattributes"] = pkAttributes
	funcMap["pkwhere"] = pkWhere
	funcMap["pkwherefields"] = pkWhereFields
	funcMap["pkupdatefields"] = pkUpdateFields
	funcMap["lower"] = lower
	funcMap["demodel"] = deModel
	funcMap["typeUnmarshaler"] = codegen.TypeUnmarshaler
	funcMap["storagedef"] = storageDef
	userTypeTmpl, err := template.New("user type").Funcs(funcMap).Parse(userTypeT)
	if err != nil {
		return nil, err
	}
	w := UserTypesWriter{
		GoGenerator:  cw,
		UserTypeTmpl: userTypeTmpl,
	}
	return &w, nil
}

// Execute writes the code for the context types to the writer.
func (w *UserTypesWriter) Execute(model *UserTypeTemplateData) error {
	return w.UserTypeTmpl.Execute(w, model)
}

// newCoerceData is a helper function that creates a map that can be given to the "Coerce" template.
func newCoerceData(name string, att *design.AttributeDefinition, pointer bool, pkg string, depth int) map[string]interface{} {
	return map[string]interface{}{
		"Name":      name,
		"VarName":   codegen.Goify(name, false),
		"Pointer":   pointer,
		"Attribute": att,
		"Pkg":       pkg,
		"Depth":     depth,
	}
}

// newDumpData is a helper function that creates a map that can be given to the "Dump" template.
func newDumpData(mt *design.MediaTypeDefinition, versioned bool, defaultPkg, context, source, target, view string) map[string]interface{} {
	return map[string]interface{}{
		"MediaType":  mt,
		"Context":    context,
		"Source":     source,
		"Target":     target,
		"View":       view,
		"Versioned":  versioned,
		"DefaultPkg": defaultPkg,
	}
}

// arrayAttribute returns the array element attribute definition.
func arrayAttribute(a *design.AttributeDefinition) *design.AttributeDefinition {
	return a.Type.(*design.Array).ElemType
}

// UserTypeMarshalerImpl returns the Go code for a function that marshals and validates instances
// of the given user type into raw values using the given view to render the attributes.
func UserTypeMarshalerImpl(u *design.UserTypeDefinition, versioned bool, defaultPkg string) string {
	var impl string
	impl = codegen.AttributeMarshaler(u.AttributeDefinition, versioned, defaultPkg, "", "source", "target")
	data := map[string]interface{}{
		"Name":           userTypeMarshalerFuncName(u),
		"Type":           u,
		"Impl":           impl,
		"Versioned":      versioned,
		"DefaultPackage": defaultPkg,
	}
	return codegen.RunTemplate(mUserImplT, data)
}

// userTypeMarshalerFuncName returns the name for the given media type marshaler function.
func userTypeMarshalerFuncName(u *design.UserTypeDefinition) string {
	return fmt.Sprintf("Marshal%s", deModel(codegen.GoTypeName(u, u.AllRequired(), 0)))
}

// GoTypeRef returns the Go code that refers to the Go type which matches the given data type
// (the part that comes after `var foo`)
// required only applies when referring to a user type that is an object defined inline. In this
// case the type (Object) does not carry the required field information defined in the parent
// (anonymous) attribute.
// tabs is used to properly tabulate the object struct fields and only applies to this case.
// This function assumes the type is in the same package as the code accessing it.
// copied in from goa so we can override package versioning
func GoTypeRef(t design.DataType, required []string, versioned bool, defaultPkg string, tabs int) string {
	return GoPackageTypeRef(t, required, versioned, defaultPkg, tabs)
}

// GoPackageTypeRef returns the Go code that refers to the Go type which matches the given data type.
// versioned indicates whether the type is being referenced from a version package (true) or the
// default package defPkg (false).
// required only applies when referring to a user type that is an object defined inline. In this
// case the type (Object) does not carry the required field information defined in the parent
// (anonymous) attribute.
// tabs is used to properly tabulate the object struct fields and only applies to this case.
// copied in from goa so we can override package versioning
func GoPackageTypeRef(t design.DataType, required []string, versioned bool, defPkg string, tabs int) string {
	switch t.(type) {
	case *design.UserTypeDefinition, *design.MediaTypeDefinition:
		var prefix string
		if t.IsObject() {
			prefix = "*"
		}
		return prefix + GoPackageTypeName(t, required, versioned, defPkg, tabs)
	case design.Object:
		return "*" + GoPackageTypeName(t, required, versioned, defPkg, tabs)
	default:
		return GoPackageTypeName(t, required, versioned, defPkg, tabs)
	}
}

// GoPackageTypeName returns the Go type name for a data type.
// versioned indicates whether the type is being referenced from a version package (true) or the
// default package defPkg (false).
// required only applies when referring to a user type that is an object defined inline. In this
// case the type (Object) does not carry the required field information defined in the parent
// (anonymous) attribute.
// tabs is used to properly tabulate the object struct fields and only applies to this case.
// Copied in from Goa so we can override package versioning functionality
func GoPackageTypeName(t design.DataType, required []string, versioned bool, defPkg string, tabs int) string {
	switch actual := t.(type) {
	case design.Primitive:
		return codegen.GoNativeType(t)
	case *design.Array:
		return "[]" + GoPackageTypeRef(actual.ElemType.Type, actual.ElemType.AllRequired(), versioned, defPkg, tabs+1)
	case design.Object:
		att := &design.AttributeDefinition{Type: actual}
		if len(required) > 0 {
			requiredVal := &design.RequiredValidationDefinition{Names: required}
			att.Validations = append(att.Validations, requiredVal)
		}
		return codegen.GoTypeDef(att, versioned, defPkg, tabs, false)
	case *design.Hash:
		return fmt.Sprintf(
			"map[%s]%s",
			GoPackageTypeRef(actual.KeyType.Type, actual.KeyType.AllRequired(), versioned, defPkg, tabs+1),
			GoPackageTypeRef(actual.ElemType.Type, actual.ElemType.AllRequired(), versioned, defPkg, tabs+1),
		)
	case *design.UserTypeDefinition:
		pkgPrefix := PackagePrefix(actual, versioned, defPkg)
		return pkgPrefix + codegen.Goify(actual.TypeName, true)
	case *design.MediaTypeDefinition:
		pkgPrefix := PackagePrefix(actual.UserTypeDefinition, versioned, defPkg)
		return pkgPrefix + codegen.Goify(actual.TypeName, true)
	default:
		panic(fmt.Sprintf("goa bug: unknown type %#v", actual))
	}
}

// PackagePrefix returns the package prefix to use to access ut from ver given it lives in the
// package pkg.
func PackagePrefix(ut *design.UserTypeDefinition, versioned bool, pkg string) string {
	if !versioned {
		// If the version is the default version then the user type is in the same package
		// (otherwise the DSL would not be valid).
		return ""
	}
	if len(ut.APIVersions) == 0 {
		// If the type is not versioned but we are accessing it from the non-default version
		// then we need to qualify it with the default version package.
		return pkg + "."
	}
	// If the type is versioned then we must be accessing it from the current version
	// (unversioned definitions cannot use versioned definitions)
	return pkg
}

const (
	mUserImplTmpl = `// {{.Name}} validates and renders an instance of {{gotypename .Type nil 0}} into a interface{}{{if .View}}
// using view "{{.View}}".{{end}}
func {{.Name}}(source {{gotyperef .Type .Type.AllRequired .Versioned .DefaultPackage 0}}, inErr error) (target {{gonative .Type}}, err error) {
	err = inErr
{{.Impl}}
	return
}`
	// ctxT generates the code for the context data type.
	// template input: *ContextTemplateData
	ctxT = `// {{.Name}} provides the {{.ResourceName}} {{.ActionName}} action context.
type {{.Name}} struct {
	*goa.Context
{{if .Params}}{{$ctx := .}}{{range $name, $att := .Params.Type.ToObject}}{{/*
*/}}	{{goify $name true}} {{if and $att.Type.IsPrimitive ($ctx.Params.IsPrimitivePointer $name)}}*{{end}}{{gotyperef .Type nil 0}}
{{end}}{{end}}{{if .Payload}}	Payload {{gotyperef .Payload nil 0}}
{{end}}}
`
	// coerceT generates the code that coerces the generic deserialized
	// data to the actual type.
	// template input: map[string]interface{} as returned by newCoerceData
	coerceT = `{{if eq .Attribute.Type.Kind 1}}{{/*

*/}}{{/* BooleanType */}}{{/*
*/}}{{$varName := or (and (not .Pointer) .VarName) tempvar}}{{/*
*/}}{{tabs .Depth}}if {{.VarName}}, err2 := strconv.ParseBool(raw{{goify .Name true}}); err2 == nil {
{{if .Pointer}}{{tabs .Depth}}	{{$varName}} := &{{.VarName}}
{{end}}{{tabs .Depth}}	{{.Pkg}} = {{$varName}}
{{tabs .Depth}}} else {
{{tabs .Depth}}	err = goa.InvalidParamTypeError("{{.Name}}", raw{{goify .Name true}}, "boolean", err)
{{tabs .Depth}}}
{{end}}{{if eq .Attribute.Type.Kind 2}}{{/*

*/}}{{/* IntegerType */}}{{/*
*/}}{{$tmp := tempvar}}{{/*
*/}}{{tabs .Depth}}if {{.VarName}}, err2 := strconv.Atoi(raw{{goify .Name true}}); err2 == nil {
{{if .Pointer}}{{$tmp2 := tempvar}}{{tabs .Depth}}	{{$tmp2}} := int({{.VarName}})
{{tabs .Depth}}	{{$tmp}} := &{{$tmp2}}
{{tabs .Depth}}	{{.Pkg}} = {{$tmp}}
{{else}}{{tabs .Depth}}	{{.Pkg}} = int({{.VarName}})
{{end}}{{tabs .Depth}}} else {
{{tabs .Depth}}	err = goa.InvalidParamTypeError("{{.Name}}", raw{{goify .Name true}}, "integer", err)
{{tabs .Depth}}}
{{end}}{{if eq .Attribute.Type.Kind 3}}{{/*

*/}}{{/* NumberType */}}{{/*
*/}}{{$varName := or (and (not .Pointer) .VarName) tempvar}}{{/*
*/}}{{tabs .Depth}}if {{.VarName}}, err2 := strconv.ParseFloat(raw{{goify .Name true}}, 64); err2 == nil {
{{if .Pointer}}{{tabs .Depth}}	{{$varName}} := &{{.VarName}}
{{end}}{{tabs .Depth}}	{{.Pkg}} = {{$varName}}
{{tabs .Depth}}} else {
{{tabs .Depth}}	err = goa.InvalidParamTypeError("{{.Name}}", raw{{goify .Name true}}, "number", err)
{{tabs .Depth}}}
{{end}}{{if eq .Attribute.Type.Kind 4}}{{/*

*/}}{{/* StringType */}}{{/*
*/}}{{tabs .Depth}}{{.Pkg}} = {{if .Pointer}}&{{end}}raw{{goify .Name true}}
{{end}}{{if eq .Attribute.Type.Kind 5}}{{/*

*/}}{{/* AnyType */}}{{/*
*/}}{{tabs .Depth}}{{.Pkg}} = {{if .Pointer}}&{{end}}raw{{goify .Name true}}
{{end}}{{if eq .Attribute.Type.Kind 6}}{{/*

*/}}{{/* ArrayType */}}{{/*
*/}}{{tabs .Depth}}elems{{goify .Name true}} := strings.Split(raw{{goify .Name true}}, ",")
{{if eq (arrayAttribute .Attribute).Type.Kind 4}}{{tabs .Depth}}{{.Pkg}} = elems{{goify .Name true}}
{{else}}{{tabs .Depth}}elems{{goify .Name true}}2 := make({{gotyperef .Attribute.Type nil .Depth}}, len(elems{{goify .Name true}}))
{{tabs .Depth}}for i, rawElem := range elems{{goify .Name true}} {
{{template "Coerce" (newCoerceData "elem" (arrayAttribute .Attribute) false (printf "elems%s2[i]" (goify .Name true)) (add .Depth 1))}}{{tabs .Depth}}}
{{tabs .Depth}}{{.Pkg}} = elems{{goify .Name true}}2
{{end}}{{end}}`

	// ctxNewT generates the code for the context factory method.
	// template input: *ContextTemplateData
	ctxNewT = `{{define "Coerce"}}` + coerceT + `{{end}}` + `
// New{{goify .Name true}} parses the incoming request URL and body, performs validations and creates the
// context used by the {{.ResourceName}} controller {{.ActionName}} action.
func New{{.Name}}(c *goa.Context) (*{{.Name}}, error) {
	var err error
	ctx := {{.Name}}{Context: c}
{{if .Headers}}{{$headers := .Headers}}{{range $name, $_ := $headers.Type.ToObject}}{{if ($headers.IsRequired $name)}}	if c.Request().Header.Get("{{$name}}") == "" {
		err = goa.MissingHeaderError("{{$name}}", err)
	}{{end}}{{end}}
{{end}}{{if.Params}}{{$ctx := .}}{{range $name, $att := .Params.Type.ToObject}}	raw{{goify $name true}} := c.Get("{{$name}}")
{{$mustValidate := $ctx.MustValidate $name}}{{if $mustValidate}}	if raw{{goify $name true}} == "" {
		err = goa.MissingParamError("{{$name}}", err)
	} else {
{{else}}	if raw{{goify $name true}} != "" {
{{end}}{{template "Coerce" (newCoerceData $name $att ($ctx.Params.IsPrimitivePointer $name) (printf "ctx.%s" (goify $name true)) 2)}}{{/*
*/}}{{$validation := validationChecker $att ($ctx.Params.IsNonZero $name) ($ctx.Params.IsRequired $name) (printf "ctx.%s" (goify $name true)) $name 2}}{{/*
*/}}{{if $validation}}{{$validation}}
{{end}}	}
{{end}}{{end}}{{/* if .Params */}}{{if .Payload}}	p, err := New{{gotypename .Payload nil 0}}(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
{{end}}	return &ctx, err
}

`
	// ctxRespT generates response helper methods GoGenerator
	// template input: *ContextTemplateData
	ctxRespT = `{{$ctx := .}}{{range .Responses}}{{$mt := $ctx.API.MediaTypeWithIdentifier .MediaType}}{{/*
*/}}// {{goify .Name true}} sends a HTTP response with status code {{.Status}}.
func (ctx *{{$ctx.Name}}) {{goify .Name true}}({{/*
*/}}{{if $mt}}resp {{gopkgtyperef $mt $mt.AllRequired $ctx.Versioned $ctx.DefaultPkg 0}}{{if gt (len $mt.ComputeViews) 1}}, view {{gopkgtypename $mt $mt.AllRequired $ctx.Versioned $ctx.DefaultPkg 0}}ViewEnum{{end}}{{/*
*/}}{{else if .MediaType}}resp []byte{{end}}) error {
{{if $mt}}	r, err := resp.Dump({{if gt (len $mt.ComputeViews) 1}}view{{end}})
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "{{$mt.Identifier}}; charset=utf-8")
	return ctx.JSON({{.Status}}, r){{else}}	return ctx.Respond({{.Status}}, {{if and (not $mt) .MediaType}}resp{{else}}nil{{end}}){{end}}
}

{{end}}`

	// payloadT generates the payload type definition GoGenerator
	// template input: *ContextTemplateData
	payloadT = `{{$payload := .Payload}}// {{gotypename .Payload nil 0}} is the {{.ResourceName}} {{.ActionName}} action payload.
type {{gotypename .Payload nil 1}} {{gotypedef .Payload .Versioned .DefaultPkg 0 false}}
`
	// newPayloadT generates the code for the payload factory method.
	// template input: *ContextTemplateData
	newPayloadT = `{{$typeName := gotypename .Payload nil 0}}{{if (not .Payload.IsPrimitive)}}

{{userTypeUnmarshalerImpl .Payload .Versioned .DefaultPkg "payload"}}{{end}}
`

	// ctrlT generates the controller interface for a given resource.
	// template input: *ControllerTemplateData
	ctrlT = `// {{.Resource}}Controller is the controller interface for the {{.Resource}} actions.
type {{.Resource}}Controller interface {
	goa.Controller
{{range .Actions}}	{{.Name}}(*{{.Context}}) error
{{end}}}
`

	// mountT generates the code for a resource "Mount" function.
	// template input: *ControllerTemplateData
	mountT = `
// Mount{{.Resource}}Controller "mounts" a {{.Resource}} resource controller on the given service.
func Mount{{.Resource}}Controller(service goa.Service, ctrl {{.Resource}}Controller) {
	var h goa.Handler
	mux := service.ServeMux(){{if not .Version.IsDefault}}.Version("{{.Version.Version}}"){{end}}
{{$res := .Resource}}{{$ver := .Version}}{{range .Actions}}{{$action := .}}	h = func(c *goa.Context) error {
		ctx, err := New{{.Context}}(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.{{.Name}}(ctx)
	}
{{range .Routes}}	mux.Handle("{{.Verb}}", "{{.FullPath $ver}}", ctrl.HandleFunc("{{$action.Name}}", h))
	service.Info("mount", "ctrl", "{{$res}}",{{if not $ver.IsDefault}} "version", "{{$ver.Version}}",{{end}} "action", "{{$action.Name}}", "route", "{{.Verb}} {{.FullPath $ver}}")
{{end}}{{end}}}
`

	// resourceT generates the code for a resource.
	// template input: *ResourceData
	resourceT = `{{ $typename  := .TypeName }}
{{ $belongs := .BelongsTo }}
{{ if .DoMedia }}
{{ $version := .APIVersion }}
{{ $keep := . }}
{{ range $idx, $action := .TypeDef.Actions  }}
{{ if hasusertype $action }}
func {{$typename}}From{{version $version}}{{title $action.Name}}Payload(ctx *{{$version}}.{{title $action.Name}}{{$typename}}Context) (*{{$typename}}, error) {
	payload := ctx.Payload
	var err error
	middle,err  := Marshal{{title $action.Name}}{{$typename}}Payload(payload, err)
	target, err := Unmarshal{{$typename}}(middle)
	return target, err
}
{{ end }}{{end}}
{{end}}
`

	// mediaTypeT generates the code for a media type.
	// template input: MediaTypeTemplateData
	mediaTypeT = `{{define "Dump"}}` + dumpT + `{{end}}` + `{{$typeName := gotypename .MediaType .MediaType.AllRequired 0}}
// Load{{$typeName}} loads raw data into an instance of {{$typeName}}
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func Load{{$typeName}}(raw interface{}) (res {{gotyperef .MediaType .MediaType.AllRequired 1}}, err error) {
	{{typeUnmarshaler .MediaType .Versioned .DefaultPkg "" "raw" "res"}}
	return
}

{{userTypeMarshalerImpl .MediaType .Versioned .DefaultPkg }}
{{userTypeUnmarshalerImpl .MediaType .Versioned .DefaultPkg "load"}}
`

	// dumpT generates the code for dumping a media type or media type collection element.
	dumpT = `{{if .MediaType.IsArray}}	{{.Target}} = make({{gonative .MediaType}}, len({{.Source}}))
{{$tmp := tempvar}}	for i, {{$tmp}} := range {{.Source}} {
{{$tmpel := tempvar}}		var {{$tmpel}} {{gonative .MediaType.ToArray.ElemType.Type}}
		{{template "Dump" (newDumpData .MediaType.ToArray.ElemType.Type .Versioned .DefaultPkg (printf "%s[*]" .Context) $tmp $tmpel .View)}}
		{{.Target}}[i] = {{$tmpel}}
	}{{else}}{{typeMarshaler .MediaType .Versioned .DefaultPkg .Context .Source .Target .View}}{{end}}`

	// userTypeT generates the code for a user type.
	// template input: UserTypeTemplateData
	userTypeT = `// {{if .UserType.Description}}{{.UserType.Description}}{{else}}{{.UserType.Name }} type{{end}}
	{{.UserType.Definition}}
{{ if ne .UserType.TableName "" }}
// TableName overrides the table name settings in gorm
func (m {{.UserType.Name}}) TableName() string {
	return "{{ .UserType.TableName}}"
}{{end}}
// {{.UserType.Name}}DB is the implementation of the storage interface for {{.UserType.Name}}
type {{.UserType.Name}}DB struct {
	Db gorm.DB
	{{ if .UserType.Cached }}cache *cache.Cache{{end}}
}
// New{{.UserType.Name}}DB creates a new storage type
func New{{.UserType.Name}}DB(db gorm.DB) *{{.UserType.Name}}DB {
	{{ if .UserType.Cached }}return &{{.UserType.Name}}DB{
		Db: db,
		cache: cache.New(5*time.Minute, 30*time.Second),
	}
	{{ else  }}return &{{.UserType.Name}}DB{Db: db}{{ end  }}
}
// DB returns  the underlying database
func (m *{{.UserType.Name}}DB) DB() interface{} {
	return &m.Db
}
{{ if .UserType.Roler }}
// GetRole returns the value of the role field and satisfies the Roler interface
func (m {{.UserType.Name}}) GetRole() string {
	return {{$f := .UserType.Fields.role}}{{if $f.Nullable}}*{{end}}m.Role
}
{{end}}
// Storage Interface
type {{.UserType.Name}}Storage interface {
	DB() interface{}
	List(ctx context.Context{{ if .UserType.DynamicTableName}}, tableName string{{ end }}) []{{.UserType.Name}}
	One(ctx context.Context{{ if .UserType.DynamicTableName }}, tableName string{{ end }}, {{.UserType.PKAttributes}}) ({{.UserType.Name}}, error)
	Add(ctx context.Context{{ if .UserType.DynamicTableName }}, tableName string{{ end }}, o {{.UserType.Name}}) ({{.UserType.Name}}, error)
	Update(ctx context.Context{{ if .UserType.DynamicTableName }}, tableName string{{ end }}, o {{.UserType.Name}}) (error)
	Delete(ctx context.Context{{ if .UserType.DynamicTableName }}, tableName string{{ end }}, {{ .UserType.PKAttributes}}) (error) 
	{{$typename:= .UserType.Name}}{{$dtn:=.UserType.DynamicTableName}}{{ range $idx, $bt := .UserType.BelongsTo}}ListBy{{$bt.Name}}(ctx context.Context{{ if $dtn}}, tableName string{{ end }},{{lower $bt.Name}}_id int) []{{$typename}}
	OneBy{{$bt.Name}}(ctx context.Context{{ if $dtn}}, tableName string{{ end }}, {{lower $bt.Name}}_id, id int) ({{$typename}}, error){{end}}
	{{range $i, $m2m := .UserType.ManyToMany}}List{{$m2m.RightNamePlural}}(context.Context, int) []{{lower $m2m.RightName}}.{{$m2m.RightName}}
	Add{{$m2m.RightNamePlural}}(context.Context, int, int) (error)
	Delete{{$m2m.RightNamePlural}}(context.Context, int, int) error{{end}}
}

// stub for marshal/unmarshal functions
func (m *{{$typename}}) Validate() error {
	return nil
}
// CRUD Functions

// One returns a single record by ID
func (m *{{$typename}}DB) One(ctx context.Context{{ if .UserType.DynamicTableName }}, tableName string{{ end }}, {{.UserType.PKAttributes}}) ({{$typename}}, error) {
	{{ if .UserType.Cached }}//first attempt to retrieve from cache
	o,found := m.cache.Get(strconv.Itoa(id))
	if found {
		return o.({{$typename}}), nil
	}
	// fallback to database if not found{{ end }}
	var obj {{$typename}}{{ $l := len $.UserType.PrimaryKeys }}{{ if eq $l 1 }}
	err := m.Db{{ if .UserType.DynamicTableName }}.Table(tableName){{ end }}.Find(&obj, id).Error{{ else  }}err := m.Db{{ if .UserType.DynamicTableName }}.Table(tableName){{ end }}.Find(&obj).Where("{{.UserType.PKWhere}}", {{.UserType.PKWhereFields }} id).Error{{ end }}
	{{ if .UserType.Cached }} go m.cache.Set(strconv.Itoa(id), obj, cache.DefaultExpiration) {{ end }}
	return obj, err
}
// Add creates a new record
func (m *{{$typename}}DB) Add(ctx context.Context{{ if .UserType.DynamicTableName }}, tableName string{{ end }}, model {{$typename}}) ({{$typename}}, error) {
	err := m.Db{{ if .UserType.DynamicTableName }}.Table(tableName){{ end }}.Create(&model).Error
	{{ if .UserType.Cached }} go m.cache.Set(strconv.Itoa(model.ID), model, cache.DefaultExpiration) {{ end }}
	return model, err
}

// Update modifies a single record
func (m *{{$typename}}DB) Update(ctx context.Context{{ if .UserType.DynamicTableName }}, tableName string{{ end }}, model {{$typename}}) error {
	obj, err := m.One(ctx{{ if .UserType.DynamicTableName }}, tableName{{ end }}, {{.UserType.PKUpdateFields}})
	if err != nil {
		return  err
	}
	err = m.Db{{ if .UserType.DynamicTableName }}.Table(tableName){{ end }}.Model(&obj).Updates(model).Error
	{{ if .UserType.Cached }}
	go func(){
	obj, err := m.One(ctx, model.ID)
	if err == nil {
		m.cache.Set(strconv.Itoa(model.ID), obj, cache.DefaultExpiration)
	}
	}()
	{{ end }}
	return err
}
// Delete removes a single record
func (m *{{$typename}}DB) Delete(ctx context.Context{{ if .UserType.DynamicTableName }}, tableName string{{ end }}, {{.UserType.PKAttributes}})  error {
	var obj {{$typename}}
	{{ $l := len .UserType.PrimaryKeys }}
	{{ if eq $l 1 }}
	err := m.Db{{ if .UserType.DynamicTableName }}.Table(tableName){{ end }}.Delete(&obj, id).Error
	{{ else  }}
	err := m.Db{{ if .UserType.DynamicTableName }}.Table(tableName){{ end }}.Delete(&obj).Where("{{.UserType.PKWhere}}", {{.UserType.PKWhereFields}}).Error
	{{ end }}
	if err != nil {
		return  err
	}
	{{ if .UserType.Cached }} go m.cache.Delete(strconv.Itoa(id)) {{ end }}
	return  nil
}

{{$ut := .UserType}}{{$typename := .UserType.Name}}{{ range $idx, $bt := .UserType.BelongsTo}}
// Belongs To Relationships

// {{$typename}}FilterBy{{$bt.Name}} is a gorm filter for a Belongs To relationship
func {{$typename}}FilterBy{{$bt.Name}}(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if parentid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("{{lower $bt.Name}}_id", parentid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

// ListBy{{$bt.Name}} returns an array of associated {{$bt.Name}} models
func (m *{{$typename}}DB) ListBy{{$bt.Name}}(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, parentid int) []{{$typename}} {
	var objs []{{$typename}}
	m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Scopes({{$typename}}FilterBy{{$bt.Name}}(parentid, &m.Db)).Find(&objs)
	return objs
}

// OneBy{{$bt.Name}} returns a single associated {{$bt.Name}} model
func (m *{{$typename}}DB) OneBy{{$bt.Name}}(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, parentid, {{ $ut.PKAttributes}}) ({{$typename}}, error) {
	{{ if $ut.Cached }}//first attempt to retrieve from cache
	o,found := m.cache.Get(strconv.Itoa(id))
	if found {
		return o.({{$typename}}), nil
	}
	// fallback to database if not found{{ end }}
	var obj {{$typename}}
	err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Scopes({{$typename}}FilterBy{{$bt.Name}}(parentid, &m.Db)).Find(&obj, id).Error
	{{ if $ut.Cached }} go m.cache.Set(strconv.Itoa(id), obj, cache.DefaultExpiration) {{ end }}
	return obj, err
}
{{end}}

{{$ut := .UserType }}{{$typeName := .UserType.Name}}{{ range $idx, $bt := .UserType.ManyToMany}}
// Many To Many Relationships

// Delete{{goify $bt.RightName true}} removes a {{$bt.RightName}}/{{$bt.LeftName}} entry from the join table
func (m *{{$typeName}}DB) Delete{{goify $bt.RightName true}}(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{lower $typeName}}ID,  {{lower $bt.RightName}}ID int)  error {
	var obj {{$typeName}}
	obj.ID = {{lower $typeName}}ID
	var assoc {{lower $bt.RightName}}.{{$bt.RightName}}
	var err error
	assoc.ID = {{lower $bt.RightName}}ID
	if err != nil {
		return err
	}
	err = m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Model(&obj).Association("{{$bt.RightNamePlural}}").Delete(assoc).Error
	if err != nil {
		return  err
	}
	return  nil
}  

// Add{{goify $bt.RightName true}} creates a new {{$bt.RightName}}/{{$bt.LeftName}} entry in the join table
func (m *{{$typeName}}DB) Add{{goify $bt.RightName true}}(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{lower $typeName}}ID, {{lower $bt.RightName}}ID int) error {
	var {{lower $typeName}} {{$typeName}}
	{{lower $typeName}}.ID = {{lower $typeName}}ID
	var assoc {{lower $bt.RightName}}.{{$bt.RightName}}
	assoc.ID = {{lower $bt.RightName}}ID
	err := m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Model(&{{lower $typeName}}).Association("{{$bt.RightNamePlural}}").Append(assoc).Error
	if err != nil {
		return  err
	}
	return  nil
}

// List{{goify $bt.RightName true}} returns a list of the {{$bt.RightName}} models related to this {{$bt.LeftName}}
func (m *{{$typeName}}DB) List{{goify $bt.RightName true}}(ctx context.Context{{ if $ut.DynamicTableName }}, tableName string{{ end }}, {{lower $typeName}}ID int)  []{{lower $bt.RightName}}.{{$bt.RightName}} {
	var list []{{lower $bt.RightName}}.{{$bt.RightName}}
	var obj {{$typeName}}
	obj.ID = {{lower $typeName}}ID
	m.Db{{ if $ut.DynamicTableName }}.Table(tableName){{ end }}.Model(&obj).Association("{{$bt.RightNamePlural}}").Find(&list)
	return  list
}
{{end}}
{{ range $idx, $bt := .UserType.BelongsTo}}
// Filter{{$typename}}By{{$bt.Name}} iterates a list and returns only those with the foreign key provided
func Filter{{$typename}}By{{$bt.Name}}(parent *int, list []{{$typename}}) []{{$typename}} {
	var filtered []{{$typename}}
	for _,o := range list {
		if o.{{$bt.Name}}ID == int(*parent) {
			filtered = append(filtered,o)
		}
	}
	return filtered
}
{{end}}
`
)
const resourceTmpl = `
{{ $typename  := .TypeName }}
{{ $belongs := .BelongsTo }}
{{ if .DoMedia }}
{{ $version := .APIVersion }}
{{ range $idx, $action := .TypeDef.Actions  }}
{{ if hasusertype $action }}
func {{$typename}}From{{version $version}}{{title $action.Name}}Payload(ctx *{{$version}}.{{title $action.Name}}{{$typename}}Context) {{$typename}} {
	payload := ctx.Payload
	var err error
	target, _ := Marshal{{title $action.Name}}{{$typename}}Payload(payload, err)
	
	return target
}
{{ end }}{{end}}{{end}}
`

/*


 */
