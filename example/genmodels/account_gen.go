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
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	cache "github.com/patrickmn/go-cache"
	log "gopkg.in/inconshreveable/log15.v2"
	"strconv"
	"time"
)

// This is the Account model
type Account struct {
	ID          int      `gorm:"primary_key"` // primary key
	Bottles     []Bottle // has many Bottles
	CreatedAt   time.Time
	CreatedBy   string
	DeletedAt   *time.Time
	Href        string
	Name        string
	OauthSource string
	UpdatedAt   time.Time
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
	cache *cache.Cache
}

// NewAccountDB creates a new storage type.
func NewAccountDB(db gorm.DB, logger log.Logger) *AccountDB {
	glog := logger.New("db", "Account")
	return &AccountDB{
		Db:     db,
		Logger: glog,
		cache:  cache.New(5*time.Minute, 30*time.Second),
	}

}

// DB returns the underlying database.
func (m *AccountDB) DB() interface{} {
	return &m.Db
}

// AccountStorage represents the storage interface.
type AccountStorage interface {
	DB() interface{}
	List(ctx goa.Context) []Account
	Get(ctx goa.Context, id int) (Account, error)
	Add(ctx goa.Context, account *Account) (Account, error)
	Update(ctx goa.Context, account *Account) error
	Delete(ctx goa.Context, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *AccountDB) TableName() string {
	return "accounts"

}

// CRUD Functions

// Get returns a single Account as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *AccountDB) Get(ctx *goa.Context, id int) Account {
	now := time.Now()
	defer ctx.Info("Account:Get", "duration", time.Since(now))
	var native Account
	m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native)
	go m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)

	return native
}

// List returns an array of Account
func (m *AccountDB) List(ctx *goa.Context) []Account {
	now := time.Now()
	defer ctx.Info("Account:List", "duration", time.Since(now))
	var objs []Account
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Account", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.
func (m *AccountDB) Add(ctx *goa.Context, model Account) (Account, error) {
	now := time.Now()
	defer ctx.Info("Account:Add", "duration", time.Since(now))
	err := m.Db.Create(&model).Error
	if err != nil {
		ctx.Error("error updating Account", "error", err.Error())
		return model, err
	}

	go m.cache.Set(strconv.Itoa(model.ID), model, cache.DefaultExpiration)
	return model, err
}

// Update modifies a single record.
func (m *AccountDB) Update(ctx *goa.Context, model Account) error {
	now := time.Now()
	defer ctx.Info("Account:Update", "duration", time.Since(now))
	obj := m.Get(ctx, model.ID)
	err := m.Db.Model(&obj).Updates(model).Error
	go func() {
		m.cache.Set(strconv.Itoa(model.ID), obj, cache.DefaultExpiration)
	}()

	return err
}

// Delete removes a single record.
func (m *AccountDB) Delete(ctx *goa.Context, id int) error {
	now := time.Now()
	defer ctx.Info("Account:Delete", "duration", time.Since(now))
	var obj Account

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		ctx.Error("error retrieving Account", "error", err.Error())
		return err
	}
	go m.cache.Delete(strconv.Itoa(id))
	return nil
}
