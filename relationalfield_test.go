package gorma_test

import (
	"fmt"
	"testing"

	"github.com/goadesign/gorma"
	"github.com/goadesign/gorma/dsl"
)

func TestFieldContext(t *testing.T) {
	sg := &gorma.RelationalFieldDefinition{}
	sg.FieldName = "SG"

	c := sg.Context()
	exp := fmt.Sprintf("RelationalField %#v", sg.FieldName)
	if c != exp {
		t.Errorf("Expected %s, got %s", exp, c)
	}

	sg.FieldName = ""

	c = sg.Context()
	exp = "unnamed RelationalField"
	if c != exp {
		t.Errorf("Expected %s, got %s", exp, c)
	}
}

func TestFieldDSL(t *testing.T) {
	sg := &gorma.RelationalFieldDefinition{}
	f := func() {
		return
	}
	sg.DefinitionDSL = f
	c := sg.DSL()
	if c == nil {
		t.Errorf("Expected %s, got nil", f)
	}

}

func TestFieldDefinitions(t *testing.T) {

	var fieldtests = []struct {
		name        string
		datatype    gorma.FieldType
		description string
		nullable    bool
		belongsto   string
		hasmany     string
		hasone      string
		many2many   string
		expected    string
	}{
		{"id", gorma.Integer, "description", false, "", "", "", "", "ID\tint  // description\n"},
		{"id", gorma.UUID, "description", false, "", "", "", "", "ID\tuuid.UUID  // description\n"},
		{"id", gorma.BigInteger, "description", false, "", "", "", "", "ID\tint64  // description\n"},
		{"name", gorma.String, "name", true, "", "", "", "", "Name\t*string  // name\n"},
		{"user", gorma.HasOne, "has one", false, "", "", "User", "", "User\tUser  // has one\n"},
		{"user_id", gorma.BelongsTo, "belongs to", false, "", "", "", "", "UserID\tint  // belongs to\n"},
	}
	for _, tt := range fieldtests {
		f := &gorma.RelationalFieldDefinition{}
		f.FieldName = dsl.SanitizeFieldName(tt.name)
		f.Datatype = tt.datatype
		f.Description = tt.description
		f.Nullable = tt.nullable
		f.BelongsTo = tt.belongsto
		f.HasMany = tt.hasmany
		f.HasOne = tt.hasone
		f.Many2Many = tt.many2many
		def := f.FieldDefinition()

		if def != tt.expected {
			t.Errorf("expected %s,got %s", tt.expected, def)
		}
	}

}
