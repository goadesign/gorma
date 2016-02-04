package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Account is the account resource media type.
var Account = MediaType("application/vnd.account+json", func() {
	Description("A tenant account")
	Attributes(func() {
		Attribute("id", Integer, "ID of account")
		Attribute("href", String, "API href of account")
		Attribute("name", String, "Name of account")
		Attribute("created_at", DateTime, "Date of creation")
		Attribute("created_by", String, "Email of account owner", func() {
			Format("email")
		})

		Required("id", "href", "name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("created_at")
		Attribute("created_by")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})
	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})

// Box is the bottle resource media type.
var Box = MediaType("application/vnd.box+json", func() {
	Description("A box of wine")
	Reference(BottlePayload)
	Attributes(func() {
		Attribute("id", Integer, "ID of bottle")
		Attribute("href", String, "API href of bottle")
		Attribute("rating", Integer, "Rating of bottle between 1 and 5", func() {
			Minimum(1)
			Maximum(5)
		})
		Attribute("account", Account, "Account that owns bottle")
		Attribute("created_at", DateTime, "Date of creation")
		Attribute("updated_at", DateTime, "Date of last update")
		// Attributes below inherit from the base type
		Attribute("name")
		Attribute("vineyard")
		Attribute("varietal")
		Attribute("vintage")
		Attribute("color")
		Attribute("sweetness")
		Attribute("country")
		Attribute("region")
		Attribute("review")
		Attribute("vinyard_county")

		Required("id", "href", "name", "vineyard", "varietal", "vintage", "color")
	})

	Links(func() {
		Link("account")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("rating")
		Attribute("vineyard")
		Attribute("varietal")
		Attribute("vintage")
		Attribute("account", func() {
			View("tiny")
		})
		Attribute("links")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("rating")
		Attribute("links")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("account")
		Attribute("rating")
		Attribute("vineyard")
		Attribute("vinyard_county")
		Attribute("varietal")
		Attribute("vintage")
		Attribute("color")
		Attribute("sweetness")
		Attribute("country")
		Attribute("region")
		Attribute("review")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("links")
	})
})

// Bottle is the bottle resource media type.
var Bottle = MediaType("application/vnd.bottle+json", func() {
	Description("A bottle of wine")
	Reference(BottlePayload)
	Attributes(func() {
		Attribute("id", Integer, "ID of bottle")
		Attribute("href", String, "API href of bottle")
		Attribute("rating", Integer, "Rating of bottle between 1 and 5", func() {
			Minimum(1)
			Maximum(5)
		})
		Attribute("account", Account, "Account that owns bottle")
		Attribute("created_at", DateTime, "Date of creation")
		Attribute("updated_at", DateTime, "Date of last update")
		// Attributes below inherit from the base type
		Attribute("name")
		Attribute("vineyard")
		Attribute("varietal")
		Attribute("vintage")
		Attribute("color")
		Attribute("sweetness")
		Attribute("country")
		Attribute("region")
		Attribute("review")
		Attribute("vinyard_county")

		Required("id", "href", "name", "vineyard", "varietal", "vintage", "color")
	})

	Links(func() {
		Link("account")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("rating")
		Attribute("vineyard")
		Attribute("varietal")
		Attribute("vintage")
		Attribute("account", func() {
			View("tiny")
		})
		Attribute("links")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("rating")
		Attribute("links")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("account")
		Attribute("rating")
		Attribute("vineyard")
		Attribute("vinyard_county")
		Attribute("varietal")
		Attribute("vintage")
		Attribute("color")
		Attribute("sweetness")
		Attribute("country")
		Attribute("region")
		Attribute("review")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("links")
	})
})
