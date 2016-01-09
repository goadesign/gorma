# gorma
Gorma is a storage generator for [goa](http://goa.design).

[![GoDoc](https://godoc.org/github.com/bketelsen/gorma?status.svg)](http://godoc.org/github.com/bketelsen/gorma)

## Table of Contents

- [Purpose](#purpose)
- [Opinionated](#opinionated)
- [Status](#opinionated)
- [Use](#use)
- [Supported Metadata Tags](#tags)
- [Example](#example)


## Purpose
Gorma uses metadata in the `goa` DSL to generate a working storage system for your API.


## Opinionated
Gorma generates Go code that uses [gorm](https://github.com/jinzhu/gorm) to access your database, therefore it is quite opinionated about how the data access layer is generated.


## Status
Gorma is a work in progress, but is being used in several applications that are not yet in production.  Use at your own risk.
Much of the code is in need of cleanup/rewriting/woodshed-beatdown.


## Use
From the root of your application, issue the `goagen` command as follows:

```
	goagen --design=github.com/gopheracademy/congo/design gen --pkg-path=github.com/bketelsen/gorma
```
Be sure to replace `github.com/gopheracademy/congo/design` with the design package of your `goa` application.


## Supported Metadata Tags
The following is a list of [Metadata](https://godoc.org/github.com/raphael/goa/design/dsl#Metadata) tags supported by Gorma.
Unrecognized tags will be silently ignored.

### Model
```
Metadata("github.com/bketelsen/gorma", "Model")
```
**Scope:** Model

This tag is required in your model in order for gorma to process it.

### belongsTo
```
Metadata("github.com/bketelsen/gorma#belongsTo", "User")
```
**Scope:** Model

This tag denotes that the model "belongs to" a parent, e.g. Proposal "Belongs To" User.
Multiple `belongsto` relationships can be expressed by including them as comma separated entities.

### dynTableName
```
	Metadata("github.com/bketelsen/gorma#dynTableName", "true")
```
**Scope:** Model

This tag denotes that the given model requires a dynamic table name causing
gorma to generate a `tableName` field to relevant function signatures.

### gormPKTag
```
	Metadata("github.com/bketelsen/gorma#gormPKTag", "column:users_id;primary")
```
**Scope:** Model

This tag is used in the Model scope to denote `gorm` tags that need to be used for the auto-generated ID field.

### gormTag
```
	Metadata("github.com/bketelsen/gorma#gormTag", "column:first_name")
```
**Scope:** Attribute

This tag is used in the Attribute scope to denote `gorm` tags that need to be added to the generated struct.

### hasMany
```
Metadata("github.com/bketelsen/gorma#hasMany", "Proposal,Review")
```
**Scope:** Model

This tag denotes the model as being the parent in a "Has Many" relationship, e.g. User "Has Many" Proposals.
Multiple `hasmany` relationships can be expressed by including them as comma separated entities.

### hasOne
```
Metadata("github.com/bketelsen/gorma#hasOne", "Address")
```
**Scope:** Model

This tag denotes that the model is the parent of a "belongs to" relationship, e.g. User "has one" Address.
Multiple `hasone` relationships can be expressed by including them as comma separated entities.

### many2many
```
Metadata("github.com/bketelsen/gorma#many2many", "PluralModel:SingularModel:join_table_name")
```
**Scope:** Model

This tag denotes that the model should have a join table for a relationship.
The arguments of the metadata create the struct field, the type of the struct field, and the name of the join table.
When added to a model called "Company", the below example represents a many to
many relationship between a Company and an Industry which would create a join
table called `company_industries` and a field in the Company struct called
`Industries` which is of type `[]Industry`

```
Metadata("github.com/bketelsen/gorma#many2many", "Industries:Industry:company_industries")
```

### noMedia
```
	Metadata("github.com/bketelsen/gorma#noMedia", "true")
```
**Scope:** Model

This tag informs gorma that no corresponding Media Type is defined for the given User Type definition.
This feature is useful when you want gorma to generate code for models that are not exposed in your API.

### roler
```
	Metadata("github.com/bketelsen/gorma#roler", "true")
```
**Scope:** Model

This tag adds a GetRole() function to the model, and returns the "Role" field of the model.  To be used with the RBAC tag.
Requires [github.com/mikespook/gorbac](https://github.com/mikespook/gorbac).

### sqlTag
```
	Metadata("github.com/bketelsen/gorma#sqlTag", "size:255")
```
**Scope:** Attribute

This tag is used in the Attribute scope to denote `sql` tags that need to be added to the generated struct.

### tableName
```
	Metadata("github.com/bketelsen/gorma#tableName", "example.users")
```
**Scope:** Model

This tag denotes that the underlying table name does not match gorm conventions.
The metadata argument is used as the table name in the generated model.

### skipTS
```
	Metadata("github.com/bketelsen/gorma#skipTS", "true")
```
**Scope:** Model

This tag instructs gorma to not generate the CreatedAt, UpdatedAt, and DeletedAt timestamp fields for the model.


## Example

Given this UserType DSL:

```
var UserModel = Type("UserModel", func() {
	Metadata("github.com/bketelsen/gorma", "Model")
	Metadata("github.com/bketelsen/gorma#roler", "true")
	Metadata("github.com/bketelsen/gorma#hasmany", "Proposal,Review")
	Attribute("firstname", func() {
	})
	Attribute("lastname", func() {
	})
	Attribute("city", func() {
	})
	Attribute("state", func() {
	})
	Attribute("country", func() {
	})
	Attribute("email", func() {
	})
	Attribute("bio", func() {
		MaxLength(500)
	})
	Attribute("role", func() {
	})

})

// ProposalModel defines the data structure used in the create proposal request body.
// It is also the base type for the proposal media type used to render users.
var ProposalModel = Type("ProposalModel", func() {
	Metadata("github.com/bketelsen/gorma", "Model")
	Metadata("github.com/bketelsen/gorma#belongsto", "User")
	Metadata("github.com/bketelsen/gorma#hasmany", "Review")
	Attribute("firstname", func() {
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
var ReviewModel = Type("ReviewModel", func() {
	Metadata("github.com/bketelsen/gorma", "Model")
	Metadata("github.com/bketelsen/gorma#belongsto", "Proposal,User")
	Attribute("comment", func() {
		MinLength(10)
		MaxLength(200)
	})
	Attribute("rating", Integer, func() {
		Minimum(1)
		Maximum(5)
	})
})
```
Gorma will generate models in the /models directory of your `goa` application that look like this:

```
/ app.UserModel storage type
// Identifier:
type User struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Proposals []Proposal
	Reviews   []Review
	// Auth
	Password string

	// OAuth2
	Oauth2Uid      string
	Oauth2Provider string
	Oauth2Token    string
	Oauth2Refresh  string
	Oauth2Expiry   time.Time

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Lock
	AttemptNumber int64
	AttemptTime   time.Time
	Locked        time.Time

	// Recover
	RecoverToken       string
	RecoverTokenExpiry time.Time
	Bio                string `json:"bio,omitempty"`
	City               string `json:"city,omitempty"`
	Country            string `json:"country,omitempty"`
	Email              string `json:"email,omitempty"`
	Firstname          string `json:"firstname,omitempty"`
	Lastname           string `json:"lastname,omitempty"`
	Role               string `json:"role,omitempty"`
	State              string `json:"state,omitempty"`
}

func UserFromCreatePayload(ctx *app.CreateUserContext) User {
	payload := ctx.Payload
	m := User{}
	copier.Copy(&m, payload)

	return m
}

func UserFromUpdatePayload(ctx *app.UpdateUserContext) User {
	payload := ctx.Payload
	m := User{}
	copier.Copy(&m, payload)
	return m
}

func (m User) ToApp() *app.User {
	target := app.User{}
	copier.Copy(&target, &m)
	return &target
}

func (m User) GetRole() string {
	return m.Role
}

type UserStorage interface {
	List(ctx context.Context) []User
	One(ctx context.Context, id int) (User, error)
	Add(ctx context.Context, o User) (User, error)
	Update(ctx context.Context, o User) error
	Delete(ctx context.Context, id int) error
}

type UserDB struct {
	DB gorm.DB
}

func NewUserDB(db gorm.DB) *UserDB {

	return &UserDB{DB: db}

}

func (m *UserDB) List(ctx context.Context) []User {

	var objs []User
	m.DB.Find(&objs)
	return objs
}

func (m *UserDB) One(ctx context.Context, id int) (User, error) {

	var obj User

	err := m.DB.Find(&obj, id).Error

	return obj, err
}

func (m *UserDB) Add(ctx context.Context, model User) (User, error) {
	err := m.DB.Create(&model).Error

	return model, err
}

func (m *UserDB) Update(ctx context.Context, model User) error {
	obj, err := m.One(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.DB.Model(&obj).Updates(model).Error

	return err
}

func (m *UserDB) Delete(ctx context.Context, id int) error {
	var obj User
	err := m.DB.Delete(&obj, id).Error
	if err != nil {
		return err
	}

	return nil
}
```
