package dsl

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
	"github.com/goadesign/goa/dslengine"
	"github.com/goadesign/gorma"
)

func checkInit() {

	if design.Design == nil {
		apidsl.InitDesign()
	}
	// check to see if this type is registered
	//set, ok := design.Design.ConstructSet["gorma"] // later!
	if gorma.GormaDesign == nil {
		//if !ok {
		// There is no registered gorma construct set
		gorma.Init()
	}
}

// topLevelDefinition returns true if the currently evaluated DSL is a root
// DSL (i.e. is not being run in the context of another definition).
func topLevelDefinition(failItNotTopLevel bool) bool {
	top := dslengine.CurrentDefinition() == nil
	if failItNotTopLevel && !top {
		incompatibleDSL(caller())
	}
	return top
}

// storageDefinition returns true and current context if it is an StorageGroupDefinition,
// nil and false otherwise.
func storageGroupDefinition(failIfNotSD bool) (*gorma.StorageGroupDefinition, bool) {
	a, ok := dslengine.CurrentDefinition().(*gorma.StorageGroupDefinition)
	if !ok && failIfNotSD {
		incompatibleDSL(caller())
	}
	return a, ok
}

// relationalStoreDefinition returns true and current context if it is an RelationalStoreDefinition,
// nil and false otherwise.
func relationalStoreDefinition(failIfNotSD bool) (*gorma.RelationalStoreDefinition, bool) {
	a, ok := dslengine.CurrentDefinition().(*gorma.RelationalStoreDefinition)
	if !ok && failIfNotSD {
		incompatibleDSL(caller())
	}
	return a, ok
}

// relationalModelDefinition returns true and current context if it is an RelationalModelDefinition,
// nil and false otherwise.
func relationalModelDefinition(failIfNotSD bool) (*gorma.RelationalModelDefinition, bool) {
	a, ok := dslengine.CurrentDefinition().(*gorma.RelationalModelDefinition)
	if !ok && failIfNotSD {
		incompatibleDSL(caller())
	}
	return a, ok
}

// relationalFieldDefinition returns true and current context if it is an RelationalFieldDefinition,
// nil and false otherwise.
func relationalFieldDefinition(failIfNotSD bool) (*gorma.RelationalFieldDefinition, bool) {
	a, ok := dslengine.CurrentDefinition().(*gorma.RelationalFieldDefinition)
	if !ok && failIfNotSD {
		incompatibleDSL(caller())
	}
	return a, ok
}

// buildSourceDefinition returns true and current context if it is an BuildSource
// nil and false otherwise.
func buildSourceDefinition(failIfNotSD bool) (*gorma.BuildSource, bool) {
	a, ok := dslengine.CurrentDefinition().(*gorma.BuildSource)
	if !ok && failIfNotSD {
		incompatibleDSL(caller())
	}
	return a, ok
}

// attributeDefinition returns true and current context if it is an AttributeDefinition
// nil and false otherwise.
func attributeDefinition(failIfNotSD bool) (*design.AttributeDefinition, bool) {
	a, ok := dslengine.CurrentDefinition().(*design.AttributeDefinition)
	if !ok && failIfNotSD {
		incompatibleDSL(caller())
	}
	return a, ok
}

// invalidArgError records an invalid argument error.
// It is used by DSL functions that take dynamic arguments.
func invalidArgError(expected string, actual interface{}) {
	dslengine.ReportError("cannot use %#v (type %s) as type %s",
		actual, reflect.TypeOf(actual), expected)
}

// incompatibleDSL should be called by DSL functions when they are
// invoked in an incorrect context (e.g. "Params" in "Resource").
func incompatibleDSL(dslFunc string) {
	elems := strings.Split(dslFunc, ".")
	dslengine.ReportError("invalid use of %s", elems[len(elems)-1])
}

// Name of calling function.
func caller() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "<unknown>"
	}
	return runtime.FuncForPC(pc).Name()
}
