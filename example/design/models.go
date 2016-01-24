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
			Description("This is the bottle model")

			Field("Vintage", gorma.Integer, func() {
				SQLTag("index")
			})
			Field("vinyard_county", gorma.String, func() {
				Alias("vinyardcounty")
			})
			Field("CreatedAt", gorma.Timestamp, func() {})
			Field("UpdatedAt", gorma.Timestamp, func() {})
			Field("DeletedAt", gorma.NullableTimestamp, func() {})
		})
	})
})
