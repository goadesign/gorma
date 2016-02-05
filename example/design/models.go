package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var sg = StorageGroup("MyStorageGroup", func() {
	Description("This is the global storage group")
	Store("mysql", gorma.MySQL, func() {
		Description("This is the mysql relational store")
		Model("Bottle", func() {
			Description("This is the Bottle model")
			BuildsFrom(BottlePayload)
			RendersTo(Bottle)
			BelongsTo("Account")
			Field("id", gorma.PKInteger, func() {
				PrimaryKey()
				MapsFrom(BottlePayload, "id")
			})
			apidsl.Attribute("oauth_source", design.String) // manually specify one that doesn't exist
			// everything else is auto populated from BuildsFrom()
		})
		Model("Account", func() {
			Description("This is the Account model")
			HasMany("Bottles", "Bottle")
			Field("id", gorma.PKInteger, func() {
				PrimaryKey()
			})
			Field("oauth_source", gorma.String) // manually specify one that doesn't exist
			// everything else is auto populated from BuildsFrom()
		})
	})
})
