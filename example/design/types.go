package design

import (
	. "github.com/bketelsen/gorma/gengorma"
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var UserModelAdapter = ModelAdapter("UserModel", "v1.CreateUserPayload", func() {

})

// UserModel defines the data structure used in the create user request body.
// It is also the base type for the user media type used to render users.
var UserModel = Model("UserModel", func() {
	Metadata("github.com/bketelsen/gorma#authboss", "All")
	Metadata("github.com/bketelsen/gorma#roler", "true")
	//	HasMany("Proposal")
	//	HasMany("Review")
	PrimaryKey("id")
	Roler()
	Cached("60")
	Timestamps()
	Attribute("first_name", func() {
		SQLTag("index")
		Description("First name Description")
	})
	Attribute("last_name", func() {
	})
	Attribute("city", func() {
	})
	Attribute("state", func() {
	})
	Attribute("country", func() {
	})
	Attribute("email", func() {
		Format("email")
	})
	Attribute("bio", func() {
		MaxLength(500)
	})

})

// ProposalModel defines the data structure used in the create proposal request body.
// It is also the base type for the proposal media type used to render users.
var ProposalModel = Model("ProposalModel", func() {
	PrimaryKey("id")
	TableName("proposals")
	BelongsTo("User")
	HasMany("reviews", ReviewModel)
	ManyToMany("m2reviews", "proposal_review", ReviewModel)
	Cached("60") // manage in-memory cache with 60 second TTL
	Timestamps() // created_at and updated_at
	SoftDelete() // deleted_at as pointer for soft deletes
	Attribute("first_name", func() {
		As("person_name") // sql column name
		MinLength(2)
	})
	Attribute("title", func() {
		MinLength(10)
		MaxLength(200)
	})
	Attribute("abstract", func() {
		MinLength(50)
		MaxLength(500)
	})
	Attribute("detail", func() {
		MinLength(100)
		MaxLength(2000)
	})
	Attribute("withdrawn", Boolean)
})

// ReviewModel defines the data structure used to create a review request body
// It is also the base type for the review media type used to render reviews
var ReviewModel = Model("ReviewModel", func() {
	PrimaryKey("id")
	BelongsTo("Proposal")
	HasOne(UserModel)
	HasMany("reviewers", UserModel)
	Attribute("comment", func() {
		MinLength(10)
		MaxLength(200)
	})
	Attribute("rating", Integer, func() {
		Minimum(1)
		Maximum(5)
	})
})
