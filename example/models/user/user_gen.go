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

package User

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
	"golang.org/x/net/context"
)

// User type
type User struct {
	ID        int `gorm:"primary_key"`
	LastName  string
	Role      string
	Bio       string
	City      string
	State     string
	Country   string
	Email     string
	FirstName string `sql:"index"` //First name Description
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDB struct {
	Db    gorm.DB
	cache *cache.Cache
}

func NewUserDB(db gorm.DB) *UserDB {
	return &UserDB{
		Db:    db,
		cache: cache.New(5*time.Minute, 30*time.Second),
	}

}

func (m *UserDB) DB() interface{} {
	return &m.Db
}

func (m User) GetRole() string {
	return m.Role
}

type UserStorage interface {
	DB() interface{}
	List(ctx context.Context) []User
	One(ctx context.Context, id int) (User, error)
	Add(ctx context.Context, o User) (User, error)
	Update(ctx context.Context, o User) error
	Delete(ctx context.Context, id int) error
}
