package gorma_test

import (
	"fmt"
	"testing"

	"github.com/goadesign/gorma"
)

func TestStoreContext(t *testing.T) {
	sg := &gorma.RelationalStoreDefinition{}
	sg.Name = "SG"

	c := sg.Context()
	exp := fmt.Sprintf("RelationalStore %#v", sg.Name)
	if c != exp {
		t.Errorf("Expected %s, got %s", exp, c)
	}

	sg.Name = ""

	c = sg.Context()
	exp = "unnamed RelationalStore"
	if c != exp {
		t.Errorf("Expected %s, got %s", exp, c)
	}
}

func TestStoreDSL(t *testing.T) {
	sg := &gorma.RelationalStoreDefinition{}
	f := func() {
		return
	}
	sg.DefinitionDSL = f
	c := sg.DSL()
	if c == nil {
		t.Errorf("Expected %s, got nil", f)
	}
}
