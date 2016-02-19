package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("CongoStorageGroup", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")
		Model("User", func() {
			BuildsFrom(func() {
				Payload("user", "create")
				Payload("user", "update")
			})
			RendersTo(User)
			Description("User Model Description")
			HasMany("Reviews", "Review")
			HasMany("Proposals", "Proposal")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the User Model PK field")
			})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Proposal", func() {
			BuildsFrom(func() {
				Payload("proposal", "create")
				Payload("proposal", "update")
			})
			RendersTo(Proposal)
			Description("Proposal Model")
			BelongsTo("User")
			HasMany("Reviews", "Review")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the Payload Model PK field")
			})
			Field("title", func() {
				Alias("proposal_title")
			})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Review", func() {
			BuildsFrom(func() {
				Payload("review", "create")
				Payload("review", "update")
			})
			RendersTo(Review)
			Description("Review Model")
			BelongsTo("User")
			BelongsTo("Proposal")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the Review Model PK field")
			})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Test", func() {
			Description("TestModel")
			NoAutomaticIDFields()
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})
		Model("TestToo", func() {
			Description("TestTooModel")
			NoAutomaticIDFields()
			Field("idone", gorma.Integer, func() {
				PrimaryKey()
				Description("This is one of the TestToo Model PK fields")
			})
			Field("idtwo", gorma.Integer, func() {
				PrimaryKey()
				Description("This is one of the TestToo Model PK fields")
			})
			BuildsFrom(func() {
				Payload("user", "create")
				Payload("user", "update")
			})
			RendersTo(User)
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})
	})
})
