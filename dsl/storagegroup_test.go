package dsl_test

import (
	"testing"

	"github.com/bketelsen/gorma"
	gdsl "github.com/bketelsen/gorma/dsl"
	"github.com/kr/pretty"
	"github.com/raphael/goa/design"
)

func TestStorageGroup(t *testing.T) {

	var sg = gdsl.StorageGroup("MyStorageGroup", func() {})
	design := design.Design
	pretty.Println(design)
	pretty.Println(design.Constructs["gorma"])
	pretty.Println(design.Constructs["gorma"][gorma.StorageGroup])
	sd, ok := design.Constructs["gorma"][gorma.StorageGroup].(*gorma.StorageGroupDefinition)
	if !ok {
		t.Errorf("expected %#v to be %#v ", sd, sg)
	}

}
