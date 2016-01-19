package dsl_test

import (
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

}
