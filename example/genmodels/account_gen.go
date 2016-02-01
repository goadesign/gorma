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
	ID        int `sql:"index" gorm:"primary_key"`
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
	One(ctx context.Context) (Account, error)
	Add(ctx context.Context, account Account) (Account, error)
	Update(ctx context.Context, account Account) error
	Delete(ctx context.Context) error
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
	rows, err := m.Db.Table(m.TableName()).Select("created_at,created_by,href,id,name").Rows()
	defer rows.Close()
	if err != nil {
		return objs
	}
	for rows.Next() {
		var iD int
		var href string
		var name string
		var createdAt time.Time
		var createdBy string

		rows.Scan(&createdAt, &createdBy, &href, &iD, &name)
		obj := app.Account{}
		obj.CreatedAt = &createdAt
		obj.Href = href
		obj.Name = name
		obj.CreatedBy = &createdBy
		obj.ID = iD
		objs = append(objs, obj)

	}
	return objs
}

// OneAccount returns an array of view: default
func (m *AccountDB) OneAccount(ctx context.Context, id int) app.Account {
	var obj app.Account
	row := m.Db.Table(m.TableName()).Select("created_at,created_by,href,id,name").Row()
	var iD int
	var href string
	var name string
	var createdAt time.Time
	var createdBy string

	row.Scan(&createdAt, &createdBy, &href, &iD, &name)
	obj.ID = iD
	obj.CreatedAt = &createdAt
	obj.Href = href
	obj.Name = name
	obj.CreatedBy = &createdBy
	return obj
}

// ListAccountLink returns an array of view: link
func (m *AccountDB) ListAccountLink(ctx context.Context) []app.AccountLink {
	var objs []app.AccountLink
	rows, err := m.Db.Table(m.TableName()).Select("href,id").Rows()
	defer rows.Close()
	if err != nil {
		return objs
	}
	for rows.Next() {
		var iD int
		var href string

		rows.Scan(&href, &iD)
		obj := app.AccountLink{}
		obj.Href = href
		obj.ID = iD
		objs = append(objs, obj)

	}
	return objs
}

// OneAccountLink returns an array of view: link
func (m *AccountDB) OneAccountLink(ctx context.Context, id int) app.AccountLink {
	var obj app.AccountLink
	row := m.Db.Table(m.TableName()).Select("href,id").Row()
	var href string
	var iD int

	row.Scan(&href, &iD)
	obj.ID = iD
	obj.Href = href
	return obj
}

// ListAccountTiny returns an array of view: tiny
func (m *AccountDB) ListAccountTiny(ctx context.Context) []app.AccountTiny {
	var objs []app.AccountTiny
	rows, err := m.Db.Table(m.TableName()).Select("href,id,name").Rows()
	defer rows.Close()
	if err != nil {
		return objs
	}
	for rows.Next() {
		var iD int
		var href string
		var name string

		rows.Scan(&href, &iD, &name)
		obj := app.AccountTiny{}
		obj.Href = href
		obj.Name = name
		obj.ID = iD
		objs = append(objs, obj)

	}
	return objs
}

// OneAccountTiny returns an array of view: tiny
func (m *AccountDB) OneAccountTiny(ctx context.Context, id int) app.AccountTiny {
	var obj app.AccountTiny
	row := m.Db.Table(m.TableName()).Select("href,id,name").Row()
	var iD int
	var href string
	var name string

	row.Scan(&href, &iD, &name)
	obj.ID = iD
	obj.Href = href
	obj.Name = name
	return obj
}

// Add creates a new record.
func (m *AccountDB) Add(ctx context.Context, model Account) (Account, error) {
	err := m.Db.Create(&model).Error
	return model, err
}

// Update modifies a single record.
func (m *AccountDB) Update(ctx context.Context, model Account) error {
	obj, err := m.One(ctx)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *AccountDB) Delete(ctx context.Context) error {
	var obj Account
	err := m.Db.Delete(&obj).Where("").Error

	if err != nil {
		return err
	}

	return nil
}
