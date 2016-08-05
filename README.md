
# gorma
Gorma is a storage generator for [goa](http://goa.design).

[![GoDoc](https://godoc.org/github.com/goadesign/gorma?status.svg)](http://godoc.org/github.com/goadesign/gorma) [![Build Status](https://travis-ci.org/goadesign/gorma.svg?branch=master)](https://travis-ci.org/goadesign/gorma) [![Go Report Card](https://goreportcard.com/badge/github.com/goadesign/gorma)](https://goreportcard.com/report/github.com/goadesign/gorma)

## Table of Contents

- [Purpose](#purpose)
- [Opinionated](#opinionated)
- [Translations](#translations)
- [Use](#use)


## Purpose
Gorma uses a custom `goa` DSL to generate a working storage system for your API.


## Opinionated
Gorma generates Go code that uses [gorm](https://github.com/jinzhu/gorm) to access your database, therefore it is quite opinionated about how the data access layer is generated.

By default, a primary key field is created as type `int` with name ID.  Also Gorm's magic date stamp fields `created_at`, `updated_at` and `deleted_at` are created.  Override this behavior with the Automatic* DSL functions on the Store.


## Translations
Use the `BuildsFrom` and `RendersTo` DSL to have Gorma generate translation functions to translate your model
to Media Types and from Payloads (User Types).  If you don't have any complex business logic in your controllers, this makes a typical controller function 3-4 lines long.

## Use
Write a storage definition using DSL from the `dsl` package.  Example:

```

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

				Field("ID", gorma.Integer, func() {    // Required for CRUD getters to take a PK argument!
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


```

See the `dsl` GoDoc for all the details and options.

From the root of your application, issue the `goagen` command as follows:

```
	goagen --design=github.com/gopheracademy/congo/design gen --pkg-path=github.com/goadesign/gorma
```
Be sure to replace `github.com/gopheracademy/congo/design` with the design package of your `goa` application.



