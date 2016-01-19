package design

import (
	"github.com/bketelsen/gorma"
	gdsl "github.com/bketelsen/gorma/dsl"
)

var sg = gdsl.StorageGroup("MyStorageGroup", func() {
	gdsl.Description("This is the global storage group")
	gdsl.RelationalStore("mysql", gorma.MySQL, func() {
		gdsl.Description("This is the mysql relational store")
		gdsl.RelationalModel("Bottles", BottlePayload, func() {
			gdsl.Description("This is the bottles model")
			gdsl.RelationalField("ID", gorma.PKInteger, func() {
				gdsl.Description("This is the ID PK field")

			})
			gdsl.RelationalField("CreatedAt", gorma.Timestamp, func() {})
			gdsl.RelationalField("UpdatedAt", gorma.Timestamp, func() {})
			gdsl.RelationalField("DeletedAt", gorma.NullableTimestamp, func() {})
		})
	})
})
