package dsl_test

import (
	"github.com/goadesign/gorma"
	gdsl "github.com/goadesign/gorma/dsl"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/dslengine"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RelationalStore", func() {
	var sgname, name string
	var storetype gorma.RelationalStorageType
	var dsl, storedsl func()
	var store *gorma.RelationalStoreDefinition

	BeforeEach(func() {
		Reset()
		sgname = "production"
		storedsl = nil
		dsl = nil
		name = ""
	})

	JustBeforeEach(func() {

		gdsl.StorageGroup(sgname, func() {
			gdsl.Store(name, storetype, dsl)
		})

		Run()

		store = gorma.GormaDesign.RelationalStores[name]
	})

	Context("with no name", func() {
		BeforeEach(func() {
			name = ""
		})

		It("does not produce a valid Relational Store definition", func() {
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
			sg := gorma.GormaDesign
			Ω(sg.RelationalStores).Should(BeEmpty())
		})
	})

	Context("with no DSL and no type", func() {
		BeforeEach(func() {
			name = "mysql"
		})

		It("does not produce a valid Relational Store definition", func() {
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
			sg := gorma.GormaDesign
			Ω(sg.RelationalStores).Should(BeEmpty())
		})
	})

	Context("with an already defined Relational Store with the same name", func() {
		BeforeEach(func() {
			name = "mysql"
		})

		It("produce an error", func() {
			gdsl.StorageGroup(sgname, func() {
				gdsl.Store(name, gorma.MySQL, dsl)
			})
			Ω(Errors).Should(HaveOccurred())
		})
	})
	Context("with an already defined Relational Store with a different name", func() {
		BeforeEach(func() {
			sgname = "mysql"
			storetype = gorma.Postgres
			name = "model"
			dsl = func() {}
			storedsl = func() {
				gdsl.Store("media", gorma.MySQL, dsl)
			}
		})

		It("doesn't return an error", func() {
			gdsl.StorageGroup("news", storedsl)
			Ω(Errors).Should(Not(HaveOccurred()))
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
			It("auto id generation defaults to true", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoIDFields).Should(BeFalse())
			})
			It("auto timestamps defaults to true", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoTimestamps).Should(BeFalse())
			})
			It("auto soft delete defaults to true", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoSoftDelete).Should(BeFalse())
			})
		})

		Context("with NoAutomaticIDFields", func() {
			BeforeEach(func() {
				name = "mysql"
				dsl = func() {
					gdsl.NoAutomaticIDFields()
				}
			})

			It("auto id generation should be off", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoIDFields).Should(BeTrue())
			})
			It("auto timestamps defaults to true", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoTimestamps).Should(BeFalse())
			})
			It("auto soft delete defaults to true", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoSoftDelete).Should(BeFalse())
			})
		})

		Context("with NoAutomaticTimestamps", func() {
			BeforeEach(func() {
				name = "mysql"
				dsl = func() {
					gdsl.NoAutomaticTimestamps()
				}
			})

			It("auto id generation should be on", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoIDFields).Should(BeFalse())
			})
			It("auto timestamps should be off", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoTimestamps).Should(BeTrue())
			})
			It("auto soft delete should be on", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoSoftDelete).Should(BeFalse())
			})
		})

	})

})
