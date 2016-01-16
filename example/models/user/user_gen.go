//************************************************************************//
// Generated Models
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/bketelsen/gorma/example
// --design=github.com/bketelsen/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package user

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
	"golang.org/x/net/context"
)

// User type
type User struct {
	ID        int `gorm:"primary_key"`
	City      string
	FirstName string `sql:"index"` //First name Description
	LastName  string
	Bio       string
	Country   string
	Email     string
	Role      string
	State     string
	UpdatedAt time.Time
	CreatedAt time.Time
}

// UserDB is the implementation of the storage interface for User
type UserDB struct {
	Db    gorm.DB
	cache *cache.Cache
}

// NewUserDB creates a new storage type
func NewUserDB(db gorm.DB) *UserDB {
	return &UserDB{
		Db:    db,
		cache: cache.New(5*time.Minute, 30*time.Second),
	}

}

// DB returns  the underlying database
func (m *UserDB) DB() interface{} {
	return &m.Db
}

// GetRole returns the value of the role field and satisfies the Roler interface
func (m User) GetRole() string {
	return m.Role
}

// Storage Interface
type UserStorage interface {
	DB() interface{}
	List(ctx context.Context) []User
	One(ctx context.Context, id int) (User, error)
	Add(ctx context.Context, o User) (User, error)
	Update(ctx context.Context, o User) error
	Delete(ctx context.Context, id int) error
}

// CRUD Functions

// One returns a single record by ID
func (m *UserDB) One(ctx context.Context, id int) (User, error) {
	//first attempt to retrieve from cache
	o, found := m.cache.Get(strconv.Itoa(id))
	if found {
		return o.(User), nil
	}
	// fallback to database if not found
	var obj User
	err := m.Db.Find(&obj, id).Error
	go m.cache.Set(strconv.Itoa(id), obj, cache.DefaultExpiration)
	return obj, err
}

// Add creates a new record
func (m *UserDB) Add(ctx context.Context, model User) (User, error) {
	err := m.Db.Create(&model).Error
	go m.cache.Set(strconv.Itoa(model.ID), model, cache.DefaultExpiration)
	return model, err
}

// Update modifies a single record
func (m *UserDB) Update(ctx context.Context, model User) error {
	obj, err := m.One(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	go func() {
		obj, err := m.One(ctx, model.ID)
		if err == nil {
			m.cache.Set(strconv.Itoa(model.ID), obj, cache.DefaultExpiration)
		}
	}()

	return err
}

// Delete removes a single record
func (m *UserDB) Delete(ctx context.Context, id int) error {
	var obj User

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		return err
	}
	go m.cache.Delete(strconv.Itoa(id))
	return nil
}

// LoadUser loads raw data into an instance of User
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
// func LoadUser(raw interface{}) (res *User, err error) {
func LoadUser(raw interface{}) (res *User, err error) {
	res, err = UnmarshalUser(raw, err)
	return
}
