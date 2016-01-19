package dsl_test

import (
	"github.com/bketelsen/gorma"
	gdsl "github.com/bketelsen/gorma/dsl"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var _ = Describe("RelationalField", func() {
	var sgname, storename, modelname, name string
	var ft gorma.FieldType
	var dsl func()
	var RandomPayload *UserTypeDefinition
	BeforeEach(func() {
		Design = nil
		Errors = nil
		sgname = "production"
		dsl = nil
		storename = "mysql"
		modelname = "Users"
		name = ""
		ft = gorma.String
		gorma.GormaConstructs = nil
		RandomPayload = Type("UserPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

	})

	JustBeforeEach(func() {
		gdsl.StorageGroup(sgname, func() {
			gdsl.RelationalStore(storename, gorma.MySQL, func() {
				gdsl.RelationalModel(modelname, RandomPayload, func() {
					gdsl.RelationalField(name, ft, dsl)
				})
			})
		})

		RunDSL()

	})

	Context("with no DSL", func() {
		BeforeEach(func() {
			name = "FirstName"
		})

		It("produces a valid Relational Field definition", func() {
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
			sg := gorma.GormaConstructs[gorma.StorageGroup].(*gorma.StorageGroupDefinition)
			rs := sg.RelationalStores[storename]
			rm := rs.RelationalModels[modelname]
			Ω(rm.RelationalFields[name].Name).Should(Equal(name))
		})
	})

	Context("with an already defined Relational Field  with the same name", func() {
		BeforeEach(func() {
			name = "FirstName"
		})

		It("produces an error", func() {
			gdsl.StorageGroup(sgname, func() {
				gdsl.RelationalStore(storename, gorma.MySQL, func() {
					gdsl.RelationalModel(modelname, RandomPayload, func() {
						gdsl.RelationalField(name, ft, dsl)
					})
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
				name = "FirstName"
				dsl = func() {
					gdsl.Description(description)
				}
			})

			It("sets the relational field description", func() {
				sg := gorma.GormaConstructs[gorma.StorageGroup].(*gorma.StorageGroupDefinition)
				rs := sg.RelationalStores[storename]
				rm := rs.RelationalModels[modelname]
				Ω(rm.RelationalFields[name].Description).Should(Equal(description))
			})
		})

	})
})
