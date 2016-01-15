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
)

// User type
type User struct {
	ID        int    `gorm:"primary_key"`
	Firstname string `sql:"blue"` //First name Description
	Lastname  string
	Role      string
	Bio       string
	City      string
	Email     string
	State     string
	Country   string
	UpdatedAt time.Time
	CreatedAt time.Time
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
	return *m.Role
}
