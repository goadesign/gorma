package dsl

import "github.com/bketelsen/gorma"

func checkInit() {

	if design.Design == nil {
		design.InitDesign()
	}
	// check to see if this type is registered
	//set, ok := design.Design.ConstructSet["gorma"] // later!
	if gorma.Design == nil {
		//if !ok {
		// There is no registered gorma construct set
		gorma.Init()
	}
}
