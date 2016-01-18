package gorma

import (
	"runtime"
	"strings"
)

// topLevelDefinition returns true if the currently evaluated DSL is a root
// DSL (i.e. is not being run in the context of another definition).
func topLevelDefinition(failItNotTopLevel bool) bool {
	top := dsl.Current() == nil
	if failItNotTopLevel && !top {
		incompatibleDSL(caller())
	}
	return top
}

// storageDefinition returns true and current context if it is an StorageGroupDefinition,
// nil and false otherwise.
func storageGroupDefinition(failIfNotSD bool) (*StorageGroupDefinition, bool) {
	a, ok := dsl.Current().(*StorageGroupDefinition)
	if !ok && failIfNotSD {
		incompatibleDSL(caller())
	}
	return a, ok
}

// incompatibleDSL should be called by DSL functions when they are
// invoked in an incorrect context (e.g. "Params" in "Resource").
func incompatibleDSL(dslFunc string) {
	elems := strings.Split(dslFunc, ".")
	ReportError("invalid use of %s", elems[len(elems)-1])
}

// Name of calling function.
func caller() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "<unknown>"
	}
	return runtime.FuncForPC(pc).Name()
}
