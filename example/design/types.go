package design

import (
	. "github.com/bketelsen/gorma/dsl"
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var UserModelAdapter = ModelAdapter("UserModel", "v1.CreateUserPayload", func() {
	Transform("created_at", design.String, design.ISO9201Date)
	TransformFunc("created_at", func() {
		out.created_at = datetime.Parse(in.CreatedAt, "blah Format")
	})
})

var XORStorage = StorageGroup("XOR", func() {
	RelationalStore("mysql", func() {
		RelationalModel(UserModel)
	})
})

// UserModel defines the data structure used in the create user request body.
// It is also the base type for the user media type used to render users.
var UserModel = RelationalModel("UserModel", func() {
	PrimaryKey("id")
	Roler()
	Cached("60")
	Timestamps()
	Field("first_name", func() {
		SQLTag("index")
		Description("First name Description")
	})
	Field("last_name", func() {
	})
	Field("city", func() {
	})
	Field("state", func() {
	})
	Field("country", func() {
	})
	Field("email", func() {
		Format("email")
	})
	Field("bio", func() {
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
	Field("first_name", func() {
		As("person_name") // sql column name
		MinLength(2)
	})
	Field("title", func() {
		MinLength(10)
		MaxLength(200)
	})
	Field("abstract", func() {
		MinLength(50)
		MaxLength(500)
	})
	Field("detail", func() {
		MinLength(100)
		MaxLength(2000)
	})
	Field("withdrawn", Boolean)
})

// ReviewModel defines the data structure used to create a review request body
// It is also the base type for the review media type used to render reviews
var ReviewModel = Model("ReviewModel", func() {
	PrimaryKey("id")
	BelongsTo("Proposal")
	HasOne(UserModel)
	HasMany("reviewers", UserModel)
	Field("comment", func() {
		MinLength(10)
		MaxLength(200)
	})
	Field("rating", Integer, func() {
		Minimum(1)
		Maximum(5)
	})
})
