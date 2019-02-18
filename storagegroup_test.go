package gorma_test

import (
	"fmt"
	"testing"

	"github.com/goadesign/gorma"
)

func TestStorageGroupContext(t *testing.T) {
	sg := &gorma.StorageGroupDefinition{}
	sg.Name = "SG"

	c := sg.Context()
	exp := fmt.Sprintf("StorageGroup %#v", sg.Name)
	if c != exp {
		t.Errorf("Expected %s, got %s", exp, c)
	}

	sg.Name = ""

	c = sg.Context()
	exp = "unnamed Storage Group"
	if c != exp {
		t.Errorf("Expected %s, got %s", exp, c)
	}
}

func TestStorageGroupDSL(t *testing.T) {
	sg := &gorma.StorageGroupDefinition{}
	f := func() {
		return
	}
	sg.DefinitionDSL = f
	c := sg.DSL()
	if c == nil {
		t.Errorf("Expected %T, got nil", f)
	}

}
