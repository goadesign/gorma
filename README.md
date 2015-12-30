# gorma
Gorma is a storage generator for [goa](http://goa.design)

## Purpose
Gorma uses metadata in the `goa` DSL to generate a working storage system for your API.

## Opinionated
Gorma generates Go code that uses [gorm](https://github.com/jinhzu/gorm) to access your database, therefore it is quite opinionated about how the data access layer is generated

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

### Supported Metadata Tags

```	
	Metadata("github.com/bketelsen/gorma", "Model")
```
This tag is required in your model in order for gorma to process it.


``` 
	Metadata("github.com/bketelsen/gorma#roler", "true") 
```
This tag adds a GetRole() function to the model, and returns the "Role" field of the model.  To be used with the RBAC tag.

```
	Metadata("github.com/bketelsen/gorma#hasmany", "Proposal,Review")
```
This tag denotes the model as being the parent in a "Has Many" relationship.  e.g. User "Has Many" Proposals


