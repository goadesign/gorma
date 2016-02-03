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
	"time"
)

// This is the Account model
type Account struct {
	ID        int `sql:"index" gorm:"primary_key"` // primary key
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
	List(ctx goa.Context) []Account
	One(ctx goa.Context, id int) (Account, error)
	Add(ctx goa.Context, account Account) (Account, error)
	Update(ctx goa.Context, account Account) error
	Delete(ctx goa.Context, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *AccountDB) TableName() string {
	return "accounts"

}

// CRUD Functions

// ListAccount returns an array of view: default
func (m *AccountDB) ListAccount(ctx goa.Context) []app.Account {
	now := time.Now()
	defer ctx.Info("ListAccount", "duration", time.Since(now))
	var objs []app.Account
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Account", "error", err.Error())
		return objs
	}

	return objs
}

// OneAccount returns an array of view: default
func (m *AccountDB) OneAccount(ctx goa.Context, id int) app.Account {
	now := time.Now()
	defer ctx.Info("OneAccount", "duration", time.Since(now))
	var view app.Account
	var native Account

	m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native)
	fmt.Println(native)
	return view
}

// ListAccountLink returns an array of view: link
func (m *AccountDB) ListAccountLink(ctx goa.Context) []app.AccountLink {
	now := time.Now()
	defer ctx.Info("ListAccountLink", "duration", time.Since(now))
	var objs []app.AccountLink
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Account", "error", err.Error())
		return objs
	}

	return objs
}

// OneAccountLink returns an array of view: link
func (m *AccountDB) OneAccountLink(ctx goa.Context, id int) app.AccountLink {
	now := time.Now()
	defer ctx.Info("OneAccountLink", "duration", time.Since(now))
	var view app.AccountLink
	var native Account

	m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native)
	fmt.Println(native)
	return view
}

// ListAccountTiny returns an array of view: tiny
func (m *AccountDB) ListAccountTiny(ctx goa.Context) []app.AccountTiny {
	now := time.Now()
	defer ctx.Info("ListAccountTiny", "duration", time.Since(now))
	var objs []app.AccountTiny
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Account", "error", err.Error())
		return objs
	}

	return objs
}

// OneAccountTiny returns an array of view: tiny
func (m *AccountDB) OneAccountTiny(ctx goa.Context, id int) app.AccountTiny {
	now := time.Now()
	defer ctx.Info("OneAccountTiny", "duration", time.Since(now))
	var view app.AccountTiny
	var native Account

	m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native)
	fmt.Println(native)
	return view
}

// GetAccount returns a single Account as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *AccountDB) GetAccount(ctx goa.Context, id int) Account {
	now := time.Now()
	defer ctx.Info("GetAccount", "duration", time.Since(now))
	var native Account
	m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native)
	return native
}

// Add creates a new record.
func (m *AccountDB) Add(ctx goa.Context, model Account) (Account, error) {
	now := time.Now()
	defer ctx.Info("AddAccount", "duration", time.Since(now))
	err := m.Db.Create(&model).Error
	if err != nil {
		ctx.Error("error updating Account", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *AccountDB) Update(ctx goa.Context, model Account) error {
	now := time.Now()
	defer ctx.Info("UpdateAccount", "duration", time.Since(now))
	obj := m.GetAccount(ctx, model.ID)
	err := m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *AccountDB) Delete(ctx goa.Context, id int) error {
	now := time.Now()
	defer ctx.Info("DeleteAccount", "duration", time.Since(now))
	var obj Account

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		ctx.Error("error retrieving Account", "error", err.Error())
		return err
	}

	return nil
}
