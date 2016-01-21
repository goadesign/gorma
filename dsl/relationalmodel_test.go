package dsl_test

import (
	"github.com/bketelsen/gorma"
	gdsl "github.com/bketelsen/gorma/dsl"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
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
				gdsl.Model(name, RandomPayload, dsl)
				gdsl.Model("Child", ChildPayload, func() {
					gdsl.BelongsTo(name)
				})
				gdsl.Model("HasOne", HasOnePayload, func() {
					gdsl.HasOne("Child")
				})
				gdsl.Model("HasMany", HasOnePayload, func() {
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
			name = "mysql"
		})

		It("produces an error", func() {
			gdsl.StorageGroup(sgname, func() {
				gdsl.Store(storename, gorma.MySQL, func() {
					gdsl.Model(name, RandomPayload, dsl)
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
					gdsl.Model(name, RandomPayload, dsl)
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

	})
})
