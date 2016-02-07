package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This is the api definition used by goa to generate the api
var _ = API("congo", func() {
	Title("Congo - Conference Management Made Easy")
	Description("Multi-tenant conference management application")
	// Generate roles and permissions for all Resources/Actions
	//		Metadata("github.com/bketelsen/gorma#rbac", "All")
	Contact(func() {
		Name("Gopher Academy")
		Email("social@gopheracademy.com")
		URL("https://gopheracademy.com")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/gopheracademy/congo/blob/master/LICENSE")
	})
	Docs(func() {
		Description("congo guide")
		URL("https://gopheracademy.com/congo/getting-started.html")
	})
	Host("api.gopheracademy.com")
	Scheme("http")
	BasePath("/api")
	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})

var _ = Version("v1", func() {
	Description("Multi-tenant conference management application API v1")
	BasePath("/:api_version/api")
})
