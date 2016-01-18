package gorma

import "github.com/raphael/goa/design"

var GormaConstructs map[string]design.DSLDefinition

func init() {

	GormaConstructs = design.Design.NewConstructsSet("gorma")

}
