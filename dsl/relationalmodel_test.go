package dsl_test

import (
	"testing"

	"github.com/bketelsen/gorma"
	gdsl "github.com/bketelsen/gorma/dsl"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/design/dsl"
)

func TestRelationalModel(t *testing.T) {

	var sg = gdsl.StorageGroup("MyStorageGroup", func() {
		gdsl.RelationalStore("mysql", func() {
			gdsl.RelationalModel("Users", func() {
			})
		})
	})
	des := design.Design
	dsl.RunDSL()
	sd, ok := des.Constructs["gorma"][gorma.StorageGroup].(*gorma.StorageGroupDefinition)
	if !ok {
		t.Errorf("expected %#v to be %#v ", sd, sg)
	}
	msql := sd.RelationalStores["mysql"]
	user := msql.RelationalModels["Users"]
	if user == nil {
		t.Errorf("expected %s model, got nil", "Users")
	}

}
