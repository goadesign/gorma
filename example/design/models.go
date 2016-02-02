package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var sg = StorageGroup("MyStorageGroup", func() {
	Description("This is the global storage group")
	Store("mysql", gorma.MySQL, func() {
		Description("This is the mysql relational store")
		Model("Bottle", func() {
			BuiltFrom(BottlePayload)
			RenderTo(Bottle)
			BelongsTo("Account")
			Description("This is the bottle model")
			Field("ID", gorma.PKInteger, func() {
				RenderTo("id")
				SQLTag("index")
			})
			Field("Vintage", gorma.String, func() {
				BuiltFrom("myvintage")
				RenderTo("vintage")
			})
			Field("vinyard_county", gorma.String, func() {
				BuiltFrom("vinyard_county")
				RenderTo("vinyard_county")
				Alias("vinyardcounty")
			})
			Field("CreatedAt", gorma.Timestamp, func() {})
			Field("UpdatedAt", gorma.Timestamp, func() {})
			Field("DeletedAt", gorma.NullableTimestamp, func() {})
		})
		Model("Account", func() {
			RenderTo(Account)
			HasMany("Bottles", "Bottle")
			Description("This is the Account model")
			Field("ID", gorma.PKInteger, func() {
				RenderTo("id")
				SQLTag("index")
			})
			Field("Href", gorma.String, func() {
				RenderTo("href")
			})
			Field("Name", gorma.String, func() {
				RenderTo("name")
			})
			Field("CreatedAt", gorma.Timestamp, func() {})
			Field("CreatedBy", gorma.String, func() {
				RenderTo("created_by")
			})
		})
	})
})
