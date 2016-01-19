package dsl_test

import (
	"testing"

	"github.com/bketelsen/gorma"
	gdsl "github.com/bketelsen/gorma/dsl"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/design/dsl"
)

func TestRelationalStore(t *testing.T) {

	var sg = gdsl.StorageGroup("MyStorageGroup", func() {
		gdsl.RelationalStore("mysql", func() {
		})
	})
	des := design.Design
	dsl.RunDSL()
	sd, ok := des.Constructs["gorma"][gorma.StorageGroup].(*gorma.StorageGroupDefinition)
	if !ok {
		t.Errorf("expected %#v to be %#v ", sd, sg)
	}
	msql := sd.RelationalStores["mysql"]
	if msql == nil {
		t.Errorf("expected %s relational store, got nil", "mysql")
	}
}
