package dsl_test

import (
	"testing"

	"github.com/bketelsen/gorma"
	gdsl "github.com/bketelsen/gorma/dsl"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/design/dsl"
)

var SG *gorma.StorageGroupDefinition

func TestStorageGroup(t *testing.T) {

	sg := setup()
	design := design.Design
	sd, ok := design.Constructs["gorma"][gorma.StorageGroup].(*gorma.StorageGroupDefinition)
	if !ok {
		t.Errorf("expected %#v to be %#v ", sd, sg)
	}

}
func TestBadStorageGroup(t *testing.T) {

	sg := badSetup()
	design := design.Design
	sd, ok := design.Constructs["gorma"][gorma.StorageGroup].(*gorma.StorageGroupDefinition)
	if !ok {
		t.Errorf("expected %#v to be %#v ", sd, sg)
	}
	if err := sg.Validate(); err == nil {
		t.Errorf("Expected errors with bad Storage Group Definition, got none")
	} else {
		if len(err.Errors) != 1 {
			t.Errorf("Expected 1 error, got %d", len(err.Errors))
		}
	}

}

func TestStorageGroupChildren(t *testing.T) {

	sg := setup()
	des := design.Design
	dsl.RunDSL()
	sd, ok := des.Constructs["gorma"][gorma.StorageGroup].(*gorma.StorageGroupDefinition)
	if !ok {
		t.Errorf("expected %#v to be %#v ", sd, sg)
	}
	if len(sd.RelationalStores) != 1 {
		t.Errorf("expected %d relational store, got %d", 1, len(sd.RelationalStores))
	}

}

func setup() *gorma.StorageGroupDefinition {
	sg := gdsl.StorageGroup("MyStorageGroup", func() {
		gdsl.RelationalStore("mysql", func() {
			gdsl.RelationalModel("Users", func() {
				gdsl.RelationalField("FirstName", func() {

				})
			})
		})
	})
	return sg
}

func badSetup() *gorma.StorageGroupDefinition {
	sg := gdsl.StorageGroup("", func() {
		gdsl.RelationalStore("mysql", func() {
			gdsl.RelationalModel("Users", func() {
				gdsl.RelationalField("FirstName", func() {

				})
			})
		})
	})
	return sg
}
