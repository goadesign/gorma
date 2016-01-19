package dsl_test

import (
	"strings"
	"testing"

	"github.com/bketelsen/gorma"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/design/dsl"
)

func TestRelationalModel(t *testing.T) {

	sg := setup()
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
	if user.Parent == nil || user.Parent != msql {
		t.Errorf("expected parent to be set")
	}
	if user.ModeledType == nil {
		t.Errorf("expected representative type to be set")
	}
	if !strings.Contains(user.Description, "Users") {
		t.Errorf("expected description, got ", user.Description)
	}

}
