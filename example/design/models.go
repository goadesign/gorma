package design

import (
	"github.com/bketelsen/gorma"
	. "github.com/bketelsen/gorma/dsl"
)

var sg = StorageGroup("MyStorageGroup", func() {
	Description("This is the global storage group")
	RelationalStore("mysql", gorma.MySQL, func() {
		Description("This is the mysql relational store")
		RelationalModel("Bottles", BottlePayload, func() {
			Description("This is the bottles model")
			RelationalField("ID", gorma.PKInteger, func() {
				Description("This is the ID PK field")

			})
			RelationalField("CreatedAt", gorma.Timestamp, func() {})
			RelationalField("UpdatedAt", gorma.Timestamp, func() {})
			RelationalField("DeletedAt", gorma.NullableTimestamp, func() {})
		})
	})
})
