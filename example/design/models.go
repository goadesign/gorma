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
			gdsl.RelationalField("Name", func() {
				gdsl.Description("This is the Name field")

			})
		})
	})
})
