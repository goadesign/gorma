# gorma
Gorma is a storage generator for [goa](http://goa.design).

[![GoDoc](https://godoc.org/github.com/goadesign/gorma?status.svg)](http://godoc.org/github.com/goadesign/gorma)

## Table of Contents

- [Purpose](#purpose)
- [Opinionated](#opinionated)
- [Use](#use)


## Purpose
Gorma uses a custom `goa` DSL to generate a working storage system for your API.


## Opinionated
Gorma generates Go code that uses [gorm](https://github.com/jinzhu/gorm) to access your database, therefore it is quite opinionated about how the data access layer is generated.


## Use
Write a storage definition using DSL from the `dsl` package.  Example:

```
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
```

See the `dsl` GoDoc for all the details and options.

From the root of your application, issue the `goagen` command as follows:

```
	goagen --design=github.com/gopheracademy/congo/design gen --pkg-path=github.com/goadesign/gorma
```
Be sure to replace `github.com/gopheracademy/congo/design` with the design package of your `goa` application.



