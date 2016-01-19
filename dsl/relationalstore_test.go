package dsl_test

import (
	"testing"

	"github.com/bketelsen/gorma"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/design/dsl"
)

func TestRelationalStore(t *testing.T) {

	sg := setup()
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
	if msql.Type != gorma.MySQL {
		t.Errorf("expected type to be %s, got %s", gorma.MySQL, msql.Type)
	}
	if msql.Parent == nil || msql.Parent != sd {
		t.Errorf("expected parent to be set")
	}
}
