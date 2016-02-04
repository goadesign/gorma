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
	"github.com/jinzhu/gorm"
	log "gopkg.in/inconshreveable/log15.v2"
	"time"
)

// This is the Account model
type Account struct {
	CreatedAt *time.Time
	CreatedBy *string
	DeletedAt *time.Time
	Href      string
	ID        int
	Name      string
	UpdatedAt time.Time
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
	log.Logger
}

// NewAccountDB creates a new storage type.
func NewAccountDB(db gorm.DB, logger log.Logger) *AccountDB {
	glog := logger.New("db", "Account")
	return &AccountDB{Db: db, Logger: glog}
}

// DB returns the underlying database.
func (m *AccountDB) DB() interface{} {
	return &m.Db
}

// AccountStorage represents the storage interface.
type AccountStorage interface {
	DB() interface{}
	List(ctx goa.Context) []Account
	One(ctx goa.Context) (Account, error)
	Add(ctx goa.Context, account Account) (Account, error)
	Update(ctx goa.Context, account Account) error
	Delete(ctx goa.Context) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *AccountDB) TableName() string {
	return "accounts"

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
	obj := m.GetAccount(ctx)
	err := m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *AccountDB) Delete(ctx goa.Context) error {
	now := time.Now()
	defer ctx.Info("DeleteAccount", "duration", time.Since(now))
	var obj Account
	err := m.Db.Delete(&obj).Where("").Error

	if err != nil {
		ctx.Error("error retrieving Account", "error", err.Error())
		return err
	}

	return nil
}
