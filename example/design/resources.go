package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var _ = Resource("ui", func() {
	BasePath("/")
	Action("bootstrap", func() {
		Routing(
			GET("//"),
		)
		Description("Render single page app HTML")
		Response(OK, func() {
			Media("text/html")
		})
	})
})

var _ = Resource("auth", func() {
	DefaultMedia(Authorize)
	BasePath("/auth")
	Action("token", func() {
		Routing(
			POST("/token"),
		)
		Description("Obtain an access token")
		Payload(Login)
		Response(Created, func() {
			Media(Authorize)
		})
	})
	Action("refresh", func() {
		Routing(
			POST("/refresh"),
		)
		Description("Obtain a refreshed access token")
		Payload(Login)
		Response(Created, func() {
			Media(Authorize)
		})
	})
	Action("callback", func() {
		Routing(
			GET("/:provider/callback"),
		)
		Description("OAUTH2 callback endpoint")
		Params(func() {
			Param("provider", String)
		})
		Response(OK, func() {
			Media("text/html")
		})
	})
	Action("oauth", func() {
		Routing(
			GET("/:provider"),
		)
		Description("OAUTH2 login endpoint")
		Params(func() {
			Param("provider", String)
		})
		Response(OK)
	})
})
var _ = Resource("user", func() {
	APIVersion("v1")

	DefaultMedia(User)
	BasePath("/users")
	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all users in account")
		Response(OK, func() {
			Media(CollectionOf(User, func() {
				View("default")
			}))
		})
	})

	Action("show", func() {
		Routing(
			GET("/:userID"),
		)
		Description("Retrieve user with given id")
		Params(func() {
			Param("userID", Integer)
		})
		Metadata("action", "123")
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Record new user")
		Payload(UserModel, func() {
			Required("first_name")
			Required("last_name")
			Required("email")
		})
		Response(Created, "^/accounts/[0-9]+/users/[0-9]+$")
	})

	Action("update", func() {
		Routing(
			PATCH("/:userID"),
		)
		Params(func() {
			Param("userID", Integer)
		})
		Payload(UserModel, func() {
			Required("email")
		})
		Response(NoContent)
		Response(NotFound)
	})
	Action("delete", func() {
		Routing(
			DELETE("/:userID"),
		)
		Params(func() {
			Param("userID", Integer, "User ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
})

var _ = Resource("proposal", func() {
	APIVersion("v1")

	Parent("user")
	DefaultMedia(Proposal)
	BasePath("/proposals")
	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all proposals for a user")
		Response(OK, func() {
			Media(CollectionOf(Proposal, func() {
				View("default")
			}))
		})
	})

	Action("show", func() {
		Routing(
			GET("/:proposalID"),
		)
		Description("Retrieve proposal with given id")
		Params(func() {
			Param("proposalID", Integer)
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create a new proposal")
		Payload(ProposalModel, func() {
			Required("title")
			Required("abstract")
			Required("detail")
		})
		Response(Created, "^/users/[0-9]+/proposals/[0-9]+$")
	})

	Action("update", func() {
		Routing(
			PATCH("/:proposalID"),
		)
		Params(func() {
			Param("proposalID", Integer)
		})
		Payload(ProposalModel)
		Response(NoContent)
		Response(NotFound)
	})
	Action("delete", func() {
		Routing(
			DELETE("/:proposalID"),
		)
		Params(func() {
			Param("proposalID", Integer, "Proposal ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
})

var _ = Resource("review", func() {
	APIVersion("v1")
	Parent("proposal")
	DefaultMedia(Review)
	BasePath("/review")
	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all reviews for a proposal")
		Response(OK, func() {
			Media(CollectionOf(Review, func() {
				View("default")
			}))
		})
	})

	Action("show", func() {
		Routing(
			GET("/:reviewID"),
		)
		Description("Retrieve review with given id")
		Params(func() {
			Param("reviewID", Integer)
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create a new review")
		Payload(ReviewModel, func() {
			Required("rating")
		})
		Response(Created, "^/users/[0-9]+/proposals/[0-9]+/reviews/[0-9]+$")
	})

	Action("update", func() {
		Routing(
			PATCH("/:reviewID"),
		)
		Params(func() {
			Param("reviewID", Integer)
		})
		Payload(ReviewModel)
		Response(NoContent)
		Response(NotFound)
	})
	Action("delete", func() {
		Routing(
			DELETE("/:reviewID"),
		)
		Params(func() {
			Param("reviewID", Integer, "Review ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
})
