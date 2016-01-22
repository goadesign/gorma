/*
Package dsl uses the Goa DSL engine to generate a data storage layer
for your Goa API.

Using a few DSL definitions you can extend the Goa API to include
database persistence.

An example:

	var sg = StorageGroup("MyStorageGroup", func() {
		Description("This is the global storage group")
		Store("mysql", gorma.MySQL, func() {
			Description("This is the mysql relational store")
			Model("Bottle", func() {
				BuiltFrom(BottlePayload)
				RenderTo(Bottle)
				Description("This is the bottle model")
				Field("ID", gorma.PKInteger, func() {
					Description("This is the ID PK field")
				})
				Field("Vintage", gorma.Integer, func() {
					SQLTag("index")
				})
				Field("CreatedAt", gorma.Timestamp, func() {})
				Field("UpdatedAt", gorma.Timestamp, func() {})
				Field("DeletedAt", gorma.NullableTimestamp, func() {})
			})
		})
	})

Gorma uses Gorm (https://github.com/jinzhu/gorm) for database access.  Gorm was chosen
as the best of the 'light-ORM' libraries available for Go.  It does the mundane work and
allows you to do anything manually if you choose.

The base Gorma definition is a `StorageGroup` which represents all the storage needs for an
application.  A StorageGroup will contain one or more `Store`, which represends a database or
other persistence mechanism.  Gorma supports all the databases that Gorm supports, and
it is possible in the future to support others -- like Key/Value stores.

Every `Store` will have one or more `Model` which maps a Go structure to a table in the database.
Use the `BuiltFrom` and `RenderTo` DSL to tell the model which Goa types will be the payload (input)
and return types.

A `Model` will contain one or more fields.  Gorma will use the `BuiltFrom` definition to populate
a base set of fields.  Custom DSL is provided to add additional fields:

Each table will likely want a primary key.  To add one to your `Model`, create a Field definition
with a type of `gorma.PKInteger` or `gorma.PKBigInteger`.  Gorma will support UUID primary keys
at some point in the future.

	Field("ID", gorma.PKInteger, func() {
		Description("This is the ID PK field")
	})

Gorma generates all the helpers you need to translate to and from the Goa types (media types and payloads).
This makes wiring up your Goa controllers almost too easy to be considered programming.

*/
package dsl
