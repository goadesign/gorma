package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var sg = StorageGroup("MyStorageGroup", func() {
	Description("This is the global storage group")
	Store("mysql", gorma.MySQL, func() {
		Description("This is the mysql relational store")
		Model("Container", func() {
			BuildsFrom(Bottle, func() {
				Map("id", "id", gorma.PKInteger)
				//Map("account", "account", gorma.ForeignKey)
			})
			RendersTo(Bottle, func() {})
			BuildsFrom(Box, func() {
				Map("id", "id", gorma.PKInteger)
				//Map("account", "account", gorma.ForeignKey)
			})
			RendersTo(Box, func() {})
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

Model("AccountModel", func() {
    BuildsFrom(Account, func() {
        Attribute("id", ... // usual Attribute DSL
        PrimaryKey("id")
    }
}
