package dsl

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/bketelsen/gorma"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/design/dsl"
)

func checkInit() {

	if design.Design == nil {
		dsl.InitDesign()
	}
	// check to see if this type is registered
	//set, ok := design.Design.ConstructSet["gorma"] // later!
	if gorma.GormaDesign == nil {
		//if !ok {
		// There is no registered gorma construct set
		fmt.Println("NO Design, creating")
		gorma.Init()
	}
}

// topLevelDefinition returns true if the currently evaluated DSL is a root
// DSL (i.e. is not being run in the context of another definition).
func topLevelDefinition(failItNotTopLevel bool) bool {
	top := dsl.CurrentStack() == nil
	if failItNotTopLevel && !top {
		incompatibleDSL(caller())
	}
	return top
}

// storageDefinition returns true and current context if it is an StorageGroupDefinition,
// nil and false otherwise.
func storageGroupDefinition(failIfNotSD bool) (*gorma.StorageGroupDefinition, bool) {
	a, ok := dsl.CurrentStack().(*gorma.StorageGroupDefinition)
	if !ok && failIfNotSD {
		incompatibleDSL(caller())
	}
	return a, ok
}

// relationalStoreDefinition returns true and current context if it is an RelationalStoreDefinition,
// nil and false otherwise.
func relationalStoreDefinition(failIfNotSD bool) (*gorma.RelationalStoreDefinition, bool) {
	a, ok := dsl.CurrentStack().(*gorma.RelationalStoreDefinition)
	if !ok && failIfNotSD {
		incompatibleDSL(caller())
	}
	return a, ok
}

// relationalModelDefinition returns true and current context if it is an RelationalModelDefinition,
// nil and false otherwise.
func relationalModelDefinition(failIfNotSD bool) (*gorma.RelationalModelDefinition, bool) {
	a, ok := dsl.CurrentStack().(*gorma.RelationalModelDefinition)
	if !ok && failIfNotSD {
		incompatibleDSL(caller())
	}
	return a, ok
}

// relationalFieldDefinition returns true and current context if it is an RelationalFieldDefinition,
// nil and false otherwise.
func relationalFieldDefinition(failIfNotSD bool) (*gorma.RelationalFieldDefinition, bool) {
	a, ok := dsl.CurrentStack().(*gorma.RelationalFieldDefinition)
	if !ok && failIfNotSD {
		incompatibleDSL(caller())
	}
	return a, ok
}

// incompatibleDSL should be called by DSL functions when they are
// invoked in an incorrect context (e.g. "Params" in "Resource").
func incompatibleDSL(dslFunc string) {
	elems := strings.Split(dslFunc, ".")
	dsl.ReportError("invalid use of %s", elems[len(elems)-1])
}

// Name of calling function.
func caller() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "<unknown>"
	}
	return runtime.FuncForPC(pc).Name()
}
