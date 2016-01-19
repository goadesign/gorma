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

	var UserPayload = Type("UserPayload", dsl)
	BeforeEach(func() {
		Design = nil
		Errors = nil
		sgname = "production"
		dsl = nil
		storename = "mysql"
		name = ""
		gorma.GormaConstructs = nil

	})

	JustBeforeEach(func() {
		gdsl.StorageGroup(sgname, func() {
			gdsl.RelationalStore(storename, gorma.MySQL, func() {
				gdsl.RelationalModel(name, UserPayload, dsl)
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
			sg := gorma.GormaConstructs[gorma.StorageGroup].(*gorma.StorageGroupDefinition)
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
				gdsl.RelationalStore(storename, gorma.MySQL, func() {
					gdsl.RelationalModel(name, UserPayload, dsl)
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
				gdsl.RelationalStore(storename, gorma.MySQL, func() {
					gdsl.RelationalModel(name, UserPayload, dsl)
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
				name = "mysql"
				dsl = func() {
					gdsl.Description(description)
				}
			})

			It("sets the relational store description", func() {
				sg := gorma.GormaConstructs[gorma.StorageGroup].(*gorma.StorageGroupDefinition)
				rs := sg.RelationalStores[storename]
				Ω(rs.RelationalModels[name].Description).Should(Equal(description))
			})
		})

	})
})
