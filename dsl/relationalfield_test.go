package dsl_test

import (
	"strings"
	"testing"

	"github.com/bketelsen/gorma"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/design/dsl"
)

func TestRelationalField(t *testing.T) {

	sg := setup()
	des := design.Design
	dsl.RunDSL()
	sd, ok := des.Constructs["gorma"][gorma.StorageGroup].(*gorma.StorageGroupDefinition)
	if !ok {
		t.Errorf("expected %#v to be %#v ", sd, sg)
	}

	msql := sd.RelationalStores["mysql"]
	user := msql.RelationalModels["Users"]
	fname, ok := user.RelationalFields["FirstName"]
	if !ok {
		t.Errorf("expected user to have fname field")
	}
	if !strings.Contains(fname.Description, "FirstName") {
		t.Errorf("expected description, got ", fname.Description)
	}

}
