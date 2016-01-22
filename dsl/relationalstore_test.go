package dsl_test

import (
	"github.com/goadesign/gorma"
	gdsl "github.com/goadesign/gorma/dsl"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/dsl"
)

var _ = Describe("RelationalStore", func() {
	var sgname, name string
	var dsl func()

	BeforeEach(func() {
		Design = nil
		Errors = nil
		sgname = "production"
		dsl = nil
		name = ""
		gorma.GormaDesign = nil
		InitDesign()

	})

	JustBeforeEach(func() {

		gdsl.StorageGroup(sgname, func() {
			gdsl.Store(name, gorma.MySQL, dsl)
		})

		RunDSL()

	})

	Context("with no DSL", func() {
		BeforeEach(func() {
			name = "mysql"
		})

		It("produces a valid Relational Store definition", func() {
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
			sg := gorma.GormaDesign
			Ω(sg.RelationalStores[name].Name).Should(Equal(name))
		})
	})

	Context("with an already defined Relational Store with the same name", func() {
		BeforeEach(func() {
			name = "mysql"
		})

		It("produces an error", func() {
			gdsl.StorageGroup(sgname, func() {
				gdsl.Store(name, gorma.MySQL, dsl)
			})
			Ω(Errors).Should(HaveOccurred())
		})
	})

	Context("with an already defined Relational Store with a different name", func() {
		BeforeEach(func() {
			sgname = "mysql"
		})

		It("returns an error", func() {
			gdsl.StorageGroup("news", dsl)
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
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].Description).Should(Equal(description))
			})
		})

	})
})
