package design

import (
	"github.com/goadesign/goa/design"
	gdsl "github.com/goadesign/goa/design/apidsl"
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var sg = StorageGroup("MyStorageGroup", func() {
	Description("This is the global storage group")
	Store("mysql", gorma.MySQL, func() {
		Description("This is the mysql relational store")
		Model("Container", func() {
			BuiltFrom(Bottle, func() {
				Map("id", "id", gorma.PKInteger)
				//Map("account", "account", gorma.ForeignKey)
				gdsl.Attribute("container_type", design.Integer, func() {
					gdsl.Default(1)
				})
			})
			RenderTo(Bottle, func() {})
			BuiltFrom(Box, func() {
				Map("id", "id", gorma.PKInteger)
				//Map("account", "account", gorma.ForeignKey)
				gdsl.Attribute("container_type", design.Integer, func() {
					gdsl.Default(2)
				})
			})
			RenderTo(Box, func() {})
			BelongsTo("Account")
			Description("This is the Container model")
			Field("DeletedAt", gorma.NullableTimestamp, func() {})
		})
		Model("Account", func() {
			BuiltFrom(Account, func() {
				Map("id", "id", gorma.PKInteger)
			})
			RenderTo(Account, func() {})
			Description("This is the Account model")
		})
	})
})
