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
				BuildsFrom(func() {
					Payload("myresource","actionname")  // e.g. "bottle", "create" resource definition
				})
				RendersTo(Bottle)						// a Media Type definition
				Description("This is the bottle model")
				Field("ID", gorma.Integer, func() {    //  redundant
					PrimaryKey()
					Description("This is the ID PK field")
				})
				Field("Vintage", gorma.Integer, func() {
					SQLTag("index")						// Add an index
				})
				Field("CreatedAt", gorma.Timestamp)
				Field("UpdatedAt", gorma.Timestamp)			 // Shown for demonstration
				Field("DeletedAt", gorma.NullableTimestamp)  // These are added by default
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

Every `Store` will have one or more `Model`s which maps a Go structure to a table in the database.
Use the `BuildsFrom` and `RendersTo` DSL to build generated functions to convert the model to Media
Type definitions and from User Type definitions.

A `Model` will contain one or more fields.  Gorma will use the `BuildsFrom` definition to populate
a base set of fields.  Custom DSL is provided to add additional fields:

Each table will likely want a primary key. Gorma automatically adds one to your table called "ID" if
there isn't one already.  Gorma supports Integer primary keys currently, but support for UUID and string
primary keys is in the plan for the future. [github](https://github.com/goadesign/gorma/issues/57)

In the event that the `BuildsFrom` types don't contain all the fields that you want to include in your
model, you can add extra fields using the `Field` DSL:

	Field("foo", gorma.String, func() {
		Description("This is the foo field")
	})

You can also specify modifications to fields that you know will be inherited from the `BuildsFrom` DSL.  For
example if the type specified in `BuildsFrom` contains a field called `Author` and you want to ensure that
it gets indexed, you can specify the field explicitly and add a `SQLTag` declaration:

	Field("author", gorma.String, func() {
		SQLTag("index")
	})

In general the only time you need to declare fields explicitly is if you want to modify the type or attributes
of the fields that are inherited from the `BuildsFrom` type, or if you want to add extra fields not included
in the `BuildsFrom` types.

You may specify more than one `BuildsFrom` type.

You can control naming between the `BuildsFrom` and `RendersTo` models by using the `MapsTo` and `MapsFrom` DSL:

	Field("Title", func(){
		MapsFrom(UserPayload, "position")
	})

This creates a field in the Gorma model called "Title", which is populated from the "position" field in the UserPayload.

The `MapsTo` DSL can be used similarly to change output field mapping to Media Types.


Gorma generates all the helpers you need to translate to and from the Goa types (media types and payloads).
This makes wiring up your Goa controllers almost too easy to be considered programming.

*/
package dsl
