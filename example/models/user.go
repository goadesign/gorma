//************************************************************************//
// API "congo": Models
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma/example/app"
	"github.com/jinzhu/gorm"
	log "gopkg.in/inconshreveable/log15.v2"
	"time"
)

// User Model Description
type User struct {
	ID        int `gorm:"primary_key"` // This is the User Model PK field
	Bio       *string
	City      *string
	Country   *string
	Email     string
	Firstname string
	Lastname  string
	Proposals []Proposal // has many Proposals
	Reviews   []Review   // has many Reviews
	State     *string
	CreatedAt time.Time  // timestamp
	UpdatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m User) TableName() string {
	return "users"

}

// UserDB is the implementation of the storage interface for
// User.
type UserDB struct {
	Db gorm.DB
	log.Logger
}

// NewUserDB creates a new storage type.
func NewUserDB(db gorm.DB, logger log.Logger) *UserDB {
	glog := logger.New("db", "User")
	return &UserDB{Db: db, Logger: glog}
}

// DB returns the underlying database.
func (m *UserDB) DB() interface{} {
	return &m.Db
}

// UserStorage represents the storage interface.
type UserStorage interface {
	DB() interface{}
	List(ctx *goa.Context) []User
	Get(ctx *goa.Context, id int) (User, error)
	Add(ctx *goa.Context, user *User) (*User, error)
	Update(ctx *goa.Context, user *User) error
	Delete(ctx *goa.Context, id int) error

	ListAppUser(ctx *goa.Context) []*app.User
	OneUser(ctx *goa.Context, id int) (*app.User, error)

	ListAppUserLink(ctx *goa.Context) []*app.UserLink
	OneUserLink(ctx *goa.Context, id int) (*app.UserLink, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserDB) TableName() string {
	return "users"

}

// CRUD Functions

// Get returns a single User as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserDB) Get(ctx *goa.Context, id int) (User, error) {
	now := time.Now()
	defer ctx.Info("User:Get", "duration", time.Since(now))
	var native User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.RecordNotFound {
		return User{}, nil
	}

	return native, err
}

// List returns an array of User
func (m *UserDB) List(ctx *goa.Context) []User {
	now := time.Now()
	defer ctx.Info("User:List", "duration", time.Since(now))
	var objs []User
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.RecordNotFound {
		ctx.Error("error listing User", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *UserDB) Add(ctx *goa.Context, model *User) (*User, error) {
	now := time.Now()
	defer ctx.Info("User:Add", "duration", time.Since(now))
	err := m.Db.Create(model).Error
	if err != nil {
		ctx.Error("error updating User", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *UserDB) Update(ctx *goa.Context, model *User) error {
	now := time.Now()
	defer ctx.Info("User:Update", "duration", time.Since(now))
	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserDB) Delete(ctx *goa.Context, id int) error {
	now := time.Now()
	defer ctx.Info("User:Delete", "duration", time.Since(now))
	var obj User

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		ctx.Error("error retrieving User", "error", err.Error())
		return err
	}

	return nil
}

func UserFromCreateUserPayload(payload *app.CreateUserPayload) *User {
	user := &User{}
	user.Bio = payload.Bio
	user.Email = payload.Email
	user.State = payload.State
	user.City = payload.City
	user.Country = payload.Country
	user.Firstname = payload.Firstname
	user.Lastname = payload.Lastname

	return user
}

func UserFromUpdateUserPayload(payload *app.UpdateUserPayload) *User {
	user := &User{}
	user.Firstname = *payload.Firstname
	user.Lastname = *payload.Lastname
	user.City = payload.City
	user.Country = payload.Country
	user.State = payload.State
	user.Bio = payload.Bio
	user.Email = payload.Email

	return user
}
