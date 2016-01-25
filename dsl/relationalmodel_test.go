package dsl_test

import (
	"github.com/goadesign/gorma"
	gdsl "github.com/goadesign/gorma/dsl"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/dsl"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RelationalModel", func() {
	var sgname, storename, name string
	var dsl func()
	var RandomPayload *UserTypeDefinition
	var ChildPayload *UserTypeDefinition
	var HasOnePayload *UserTypeDefinition
	var HasManyPayload *UserTypeDefinition

	BeforeEach(func() {
		Design = nil
		Errors = nil
		sgname = "production"
		dsl = nil
		storename = "mysql"
		name = ""
		gorma.GormaDesign = nil
		InitDesign()

		RandomPayload = Type("RandomPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

		ChildPayload = Type("ChildPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})
		HasOnePayload = Type("HasOnePayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

		HasManyPayload = Type("HasManyPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

	})

	JustBeforeEach(func() {
		gdsl.StorageGroup(sgname, func() {
			gdsl.Store(storename, gorma.MySQL, func() {
				gdsl.Model(name, dsl)
				gdsl.Model("Child", func() {
					gdsl.BuiltFrom(ChildPayload)
					gdsl.BelongsTo(name)
				})
				gdsl.Model("One", func() {
					gdsl.BuiltFrom(HasOnePayload)
					gdsl.HasOne("Child")
				})
				gdsl.Model("Many", func() {
					gdsl.BuiltFrom(HasManyPayload)
					gdsl.HasMany("Children", "Child")
				})

			})
		})

		RunDSL()

	})

	Context("with no DSL", func() {
		BeforeEach(func() {
			name = "Users"
		})

		It("produces a valid Relational Model definition", func() {
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
			sg := gorma.GormaDesign
			rs := sg.RelationalStores[storename]
			Ω(rs.RelationalModels[name].Name).Should(Equal(name))
		})
	})

	Context("with an already defined Relational Model with the same name", func() {
		BeforeEach(func() {
			name = "duplicate"
		})

		It("produces an error", func() {
			gdsl.StorageGroup(sgname, func() {
				gdsl.Store(storename, gorma.MySQL, func() {
					gdsl.Model(name, dsl)
				})
			})
			Ω(Errors).Should(HaveOccurred())
		})
	})

	Context("with an already defined Relational model with a different name", func() {
		BeforeEach(func() {
			name = "Users"
		})

		It("returns an error", func() {
			gdsl.StorageGroup(sgname, func() {
				gdsl.Store(storename, gorma.MySQL, func() {
					gdsl.Model(name, dsl)
				})
			})
			Ω(Errors).Should(HaveOccurred())
		})
	})

	Context("with valid DSL", func() {
		JustBeforeEach(func() {
			Ω(Errors).ShouldNot(HaveOccurred())
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
		})

		Context("with a description", func() {
			const description = "description"

			BeforeEach(func() {
				name = "Users"
				dsl = func() {
					gdsl.Description(description)
				}
			})

			It("sets the relational store description", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				Ω(rs.RelationalModels[name].Description).Should(Equal(description))
			})
		})
		Context("with a table name", func() {
			const tablename = "user_table"

			BeforeEach(func() {
				name = "Users"
				dsl = func() {
					gdsl.TableName(tablename)
				}
			})

			It("sets the relational store table name", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				Ω(rs.RelationalModels[name].TableName).Should(Equal(tablename))
			})
		})
		Context("with an alias", func() {
			const alias = "user_table"

			BeforeEach(func() {
				name = "Users"
				dsl = func() {
					gdsl.Alias(alias)
				}
			})

			It("sets the relational store alias", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				Ω(rs.RelationalModels[name].Alias).Should(Equal(alias))
			})
		})

		Context("cached", func() {
			const duration = "50"

			BeforeEach(func() {
				name = "Users"
				dsl = func() {
					gdsl.Cached(duration)
				}
			})

			It("sets the relational store cache values", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				Ω(rs.RelationalModels[name].Cached).Should(Equal(true))
				Ω(rs.RelationalModels[name].CacheDuration).Should(Equal(50))

			})
		})
		Context("with nomedia", func() {

			BeforeEach(func() {
				name = "Users"
				dsl = func() {
					gdsl.NoMedia()
				}
			})

			It("sets the relational store alias", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				Ω(rs.RelationalModels[name].NoMedia).Should(Equal(true))
			})
		})

		Context("with roler", func() {

			BeforeEach(func() {
				name = "Users"
				dsl = func() {
					gdsl.Roler()
				}
			})

			It("sets the relational store alias", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				Ω(rs.RelationalModels[name].Roler).Should(Equal(true))
			})
		})

		Context("with dynamic table name", func() {

			BeforeEach(func() {
				name = "Users"
				dsl = func() {
					gdsl.DynamicTableName()
				}
			})

			It("sets the relational store alias", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				Ω(rs.RelationalModels[name].DynamicTableName).Should(Equal(true))
			})
		})
		Context("with an sql tag", func() {
			const tag = "unique"

			BeforeEach(func() {
				name = "Users"
				dsl = func() {
					gdsl.SQLTag(tag)
				}
			})

			It("sets the relational store alias", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				Ω(rs.RelationalModels[name].SQLTag).Should(Equal(tag))
			})
		})

		Context("with a has one relationaship", func() {

			It("sets the creates the foreign key in the child model", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				f, ok := rs.RelationalModels["Child"].RelationalFields["OneID"]

				Ω(ok).Should(Equal(true))
				Ω(f.DatabaseFieldName).Should(Equal("one_id"))
			})
			It("the relationship is added to the HasOne list", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				ch, ok := rs.RelationalModels["One"].HasOne["Child"]

				Ω(ok).Should(Equal(true))
				Ω(ch).Should(Equal(rs.RelationalModels["Child"]))
			})

			It("sets the field definition correctly for the owning model", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				f, ok := rs.RelationalModels["One"].RelationalFields["Child"]

				Ω(ok).Should(Equal(true))
				Ω(f.Name).Should(Equal("Child"))
			})
		})

		Context("with a belongs to relationship", func() {

			BeforeEach(func() {
				name = "User"
			})

			It("sets the creates the foreign key in the child model", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				f, ok := rs.RelationalModels["Child"].RelationalFields["UserID"]

				Ω(ok).Should(Equal(true))
				Ω(f.DatabaseFieldName).Should(Equal("user_id"))
			})
			It("the relationship is added to the BelongsTo list", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				ch, ok := rs.RelationalModels["Child"].BelongsTo["User"]

				Ω(ok).Should(Equal(true))
				Ω(ch).Should(Equal(rs.RelationalModels["User"]))
			})

			It("sets the field definition correctly for the child model", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				f, ok := rs.RelationalModels["Child"].RelationalFields["UserID"]

				Ω(ok).Should(Equal(true))
				Ω(f.Name).Should(Equal("UserID"))
			})
		})

		Context("with a has many relationship", func() {

			It("sets the creates the foreign key in the child model", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				f, ok := rs.RelationalModels["Child"].RelationalFields["ManyID"]

				Ω(ok).Should(Equal(true))
				Ω(f.DatabaseFieldName).Should(Equal("many_id"))
			})
			It("the relationship is added to the Has Many list", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				ch, ok := rs.RelationalModels["Many"].HasMany["Child"]

				Ω(ok).Should(Equal(true))
				Ω(ch).Should(Equal(rs.RelationalModels["Child"]))
			})

			It("sets the field definition correctly for the child model", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				f, ok := rs.RelationalModels["Many"].RelationalFields["Children"]

				Ω(ok).Should(Equal(true))
				Ω(f.Name).Should(Equal("Children"))
			})
		})

	})
})

var _ = Describe("RelationalModel with auto fields enabled and auto fields set in dsl", func() {
	var sgname, storename, name string
	var dsl func()
	var RandomPayload *UserTypeDefinition
	var ChildPayload *UserTypeDefinition
	var HasOnePayload *UserTypeDefinition
	var HasManyPayload *UserTypeDefinition

	BeforeEach(func() {
		Design = nil
		Errors = nil
		sgname = "production"
		dsl = nil
		storename = "mysql"
		name = ""
		gorma.GormaDesign = nil
		InitDesign()

		RandomPayload = Type("RandomPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

		ChildPayload = Type("ChildPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})
		HasOnePayload = Type("HasOnePayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

		HasManyPayload = Type("HasManyPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

	})

	JustBeforeEach(func() {
		gdsl.StorageGroup(sgname, func() {
			gdsl.Store(storename, gorma.MySQL, func() {
				gdsl.Model(name, dsl)
				gdsl.Model("Child", func() {
					gdsl.BuiltFrom(ChildPayload)
					gdsl.BelongsTo(name)
				})
				gdsl.Model("One", func() {
					gdsl.BuiltFrom(HasOnePayload)
					gdsl.HasOne("Child")
				})
				gdsl.Model("Many", func() {
					gdsl.BuiltFrom(HasManyPayload)
					gdsl.HasMany("Children", "Child")
				})

			})
		})

		RunDSL()

	})

	Context("with no DSL", func() {
		BeforeEach(func() {
			name = "Users"
			dsl = func() {
				gdsl.Field("ID", gorma.PKInteger)
				gdsl.Field("CreatedAt", gorma.Timestamp)
				gdsl.Field("UpdatedAt", gorma.Timestamp)
				gdsl.Field("DeletedAt", gorma.NullableTimestamp)
			}
		})

		It("generates auto fields", func() {
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
			sg := gorma.GormaDesign
			rs := sg.RelationalStores[storename]
			Ω(rs.RelationalModels[name].Name).Should(Equal(name))
			_, ok := rs.RelationalModels[name].RelationalFields["ID"]
			Ω(ok).Should(Equal(true))
			_, ok = rs.RelationalModels[name].RelationalFields["UpdatedAt"]
			Ω(ok).Should(Equal(true))
			_, ok = rs.RelationalModels[name].RelationalFields["CreatedAt"]
			Ω(ok).Should(Equal(true))
			_, ok = rs.RelationalModels[name].RelationalFields["DeletedAt"]
			Ω(ok).Should(Equal(true))
			Ω(rs.RelationalModels[name].Name).Should(Equal(name))
			Ω(len(rs.RelationalModels[name].RelationalFields)).Should(Equal(4))

		})
	})
})
var _ = Describe("RelationalModel with auto fields explicitly enabled", func() {
	var sgname, storename, name string
	var dsl func()
	var RandomPayload *UserTypeDefinition
	var ChildPayload *UserTypeDefinition
	var HasOnePayload *UserTypeDefinition
	var HasManyPayload *UserTypeDefinition

	BeforeEach(func() {
		Design = nil
		Errors = nil
		sgname = "production"
		dsl = nil
		storename = "mysql"
		name = ""
		gorma.GormaDesign = nil
		InitDesign()

		RandomPayload = Type("RandomPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

		ChildPayload = Type("ChildPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})
		HasOnePayload = Type("HasOnePayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

		HasManyPayload = Type("HasManyPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

	})

	JustBeforeEach(func() {
		gdsl.StorageGroup(sgname, func() {
			gdsl.Store(storename, gorma.MySQL, func() {
				gdsl.Model(name, dsl)
				gdsl.Model("Child", func() {
					gdsl.BuiltFrom(ChildPayload)
					gdsl.BelongsTo(name)
				})
				gdsl.Model("One", func() {
					gdsl.BuiltFrom(HasOnePayload)
					gdsl.HasOne("Child")
				})
				gdsl.Model("Many", func() {
					gdsl.BuiltFrom(HasManyPayload)
					gdsl.HasMany("Children", "Child")
				})

			})
		})

		RunDSL()

	})

	Context("with no DSL", func() {
		BeforeEach(func() {
			name = "Users"
		})

		It("generates auto fields", func() {
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
			sg := gorma.GormaDesign
			rs := sg.RelationalStores[storename]
			Ω(rs.RelationalModels[name].Name).Should(Equal(name))
			_, ok := rs.RelationalModels[name].RelationalFields["ID"]
			Ω(ok).Should(Equal(true))
			_, ok = rs.RelationalModels[name].RelationalFields["UpdatedAt"]
			Ω(ok).Should(Equal(true))
			_, ok = rs.RelationalModels[name].RelationalFields["CreatedAt"]
			Ω(ok).Should(Equal(true))
			_, ok = rs.RelationalModels[name].RelationalFields["DeletedAt"]
			Ω(ok).Should(Equal(true))
			Ω(rs.RelationalModels[name].Name).Should(Equal(name))
			Ω(len(rs.RelationalModels[name].RelationalFields)).Should(Equal(4))
		})
	})
})

var _ = Describe("RelationalModel with auto fields disabled", func() {
	var sgname, storename, name string
	var dsl func()
	var RandomPayload *UserTypeDefinition
	var ChildPayload *UserTypeDefinition
	var HasOnePayload *UserTypeDefinition
	var HasManyPayload *UserTypeDefinition

	BeforeEach(func() {
		Design = nil
		Errors = nil
		sgname = "production"
		dsl = nil
		storename = "mysql"
		name = ""
		gorma.GormaDesign = nil
		InitDesign()

		RandomPayload = Type("RandomPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

		ChildPayload = Type("ChildPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})
		HasOnePayload = Type("HasOnePayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

		HasManyPayload = Type("HasManyPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

	})

	JustBeforeEach(func() {
		gdsl.StorageGroup(sgname, func() {
			gdsl.Store(storename, gorma.MySQL, func() {
				gdsl.NoAutomaticIDFields()
				gdsl.NoAutomaticTimestamps()
				gdsl.NoAutomaticSoftDelete()
				gdsl.Model(name, dsl)
				gdsl.Model("Child", func() {
					gdsl.BuiltFrom(ChildPayload)
					gdsl.BelongsTo(name)
				})
				gdsl.Model("One", func() {
					gdsl.BuiltFrom(HasOnePayload)
					gdsl.HasOne("Child")
				})
				gdsl.Model("Many", func() {
					gdsl.BuiltFrom(HasManyPayload)
					gdsl.HasMany("Children", "Child")
				})

			})
		})

		RunDSL()

	})

	Context("with no DSL", func() {
		BeforeEach(func() {
			name = "Users"
		})

		It("doesn't generate auto fields", func() {
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
			sg := gorma.GormaDesign
			rs := sg.RelationalStores[storename]
			Ω(rs.RelationalModels[name].Name).Should(Equal(name))
			_, ok := rs.RelationalModels[name].RelationalFields["ID"]
			Ω(ok).Should(Equal(false))
			_, ok = rs.RelationalModels[name].RelationalFields["UpdatedAt"]
			Ω(ok).Should(Equal(false))
			_, ok = rs.RelationalModels[name].RelationalFields["CreatedAt"]
			Ω(ok).Should(Equal(false))
			_, ok = rs.RelationalModels[name].RelationalFields["DeletedAt"]
			Ω(ok).Should(Equal(false))
			Ω(rs.RelationalModels[name].Name).Should(Equal(name))
			Ω(len(rs.RelationalModels[name].RelationalFields)).Should(Equal(0))
		})
	})
})

var _ = Describe("RelationalModel with auto fields unset", func() {
	var sgname, storename, name string
	var dsl func()
	var RandomPayload *UserTypeDefinition
	var ChildPayload *UserTypeDefinition
	var HasOnePayload *UserTypeDefinition
	var HasManyPayload *UserTypeDefinition

	BeforeEach(func() {
		Design = nil
		Errors = nil
		sgname = "production"
		dsl = nil
		storename = "mysql"
		name = ""
		gorma.GormaDesign = nil
		InitDesign()

		RandomPayload = Type("RandomPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

		ChildPayload = Type("ChildPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})
		HasOnePayload = Type("HasOnePayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

		HasManyPayload = Type("HasManyPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

	})

	JustBeforeEach(func() {
		gdsl.StorageGroup(sgname, func() {
			gdsl.Store(storename, gorma.MySQL, func() {
				gdsl.Model(name, dsl)
				gdsl.Model("Child", func() {
					gdsl.BuiltFrom(ChildPayload)
					gdsl.BelongsTo(name)
				})
				gdsl.Model("One", func() {
					gdsl.BuiltFrom(HasOnePayload)
					gdsl.HasOne("Child")
				})
				gdsl.Model("Many", func() {
					gdsl.BuiltFrom(HasManyPayload)
					gdsl.HasMany("Children", "Child")
				})

			})
		})

		RunDSL()

	})

	Context("with no DSL", func() {
		BeforeEach(func() {
			name = "Users"
		})

		It("generates auto fields", func() {
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
			sg := gorma.GormaDesign
			rs := sg.RelationalStores[storename]
			Ω(rs.RelationalModels[name].Name).Should(Equal(name))
			f, ok := rs.RelationalModels[name].RelationalFields["ID"]
			Ω(ok).Should(Equal(true))
			Ω(f.Datatype).Should(Equal(gorma.PKInteger))
			f, ok = rs.RelationalModels[name].RelationalFields["UpdatedAt"]
			Ω(ok).Should(Equal(true))
			Ω(f.Datatype).Should(Equal(gorma.Timestamp))
			f, ok = rs.RelationalModels[name].RelationalFields["CreatedAt"]
			Ω(ok).Should(Equal(true))
			Ω(f.Datatype).Should(Equal(gorma.Timestamp))
			f, ok = rs.RelationalModels[name].RelationalFields["DeletedAt"]
			Ω(ok).Should(Equal(true))
			Ω(f.Datatype).Should(Equal(gorma.NullableTimestamp))
			Ω(rs.RelationalModels[name].Name).Should(Equal(name))
		})
	})
})
