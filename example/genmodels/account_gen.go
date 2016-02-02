//************************************************************************//
// API "cellar": Models
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package genmodels

import (
	"github.com/goadesign/gorma/example/app"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// This is the Account model
type Account struct {
	ID        int      `sql:"index" gorm:"primary_key"` // primary key
	Bottles   []Bottle // has many Bottles
	CreatedBy string
	DeletedAt *time.Time
	Href      string
	Name      string
	UpdatedAt time.Time
	CreatedAt time.Time // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Account) TableName() string {
	return "accounts"

}

// AccountDB is the implementation of the storage interface for
// Account.
type AccountDB struct {
	Db gorm.DB
}

// NewAccountDB creates a new storage type.
func NewAccountDB(db gorm.DB) *AccountDB {
	return &AccountDB{Db: db}
}

// DB returns the underlying database.
func (m *AccountDB) DB() interface{} {
	return &m.Db
}

// AccountStorage represents the storage interface.
type AccountStorage interface {
	DB() interface{}
	List(ctx context.Context) []Account
	One(ctx context.Context, id int) (Account, error)
	Add(ctx context.Context, account Account) (Account, error)
	Update(ctx context.Context, account Account) error
	Delete(ctx context.Context, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *AccountDB) TableName() string {
	return "accounts"

}

// CRUD Functions

// ListAccount returns an array of view: default
func (m *AccountDB) ListAccount(ctx context.Context) []app.Account {
	var objs []app.Account
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		return objs
	}

	return objs
}

// OneAccount returns an array of view: default
func (m *AccountDB) OneAccount(ctx context.Context, id int) app.Account {
	var native Account
	m.Db.Table(m.TableName()).Find(&native).Where("id = ?", id)
	var obj app.Account

	return obj
}

// ListAccountLink returns an array of view: link
func (m *AccountDB) ListAccountLink(ctx context.Context) []app.AccountLink {
	var objs []app.AccountLink
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		return objs
	}

	return objs
}

// OneAccountLink returns an array of view: link
func (m *AccountDB) OneAccountLink(ctx context.Context, id int) app.AccountLink {
	var native Account
	m.Db.Table(m.TableName()).Find(&native).Where("id = ?", id)
	var obj app.AccountLink

	return obj
}

// ListAccountTiny returns an array of view: tiny
func (m *AccountDB) ListAccountTiny(ctx context.Context) []app.AccountTiny {
	var objs []app.AccountTiny
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		return objs
	}

	return objs
}

// OneAccountTiny returns an array of view: tiny
func (m *AccountDB) OneAccountTiny(ctx context.Context, id int) app.AccountTiny {
	var native Account
	m.Db.Table(m.TableName()).Find(&native).Where("id = ?", id)
	var obj app.AccountTiny

	return obj
}

// Add creates a new record.
func (m *AccountDB) Add(ctx context.Context, model Account) (Account, error) {
	err := m.Db.Create(&model).Error
	return model, err
}

// Update modifies a single record.
func (m *AccountDB) Update(ctx context.Context, model Account) error {
	obj := m.OneAccount(ctx, model.ID)
	err := m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *AccountDB) Delete(ctx context.Context, id int) error {
	var obj Account

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		return err
	}

	return nil
}
