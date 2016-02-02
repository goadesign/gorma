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

// This is the bottle model
type Bottle struct {
	ID            int `sql:"index" gorm:"primary_key"` // primary key
	AccountID     int // Belongs To Account
	Color         *string
	Country       *string
	Myvintage     *string
	Name          *string
	Region        *string
	Review        *string
	Sweetness     *int
	Varietal      *string
	Vineyard      *string
	Vintage       *string
	VinyardCounty *string    `gorm:"column:vinyardcounty"`
	CreatedAt     time.Time  // timestamp
	UpdatedAt     time.Time  // timestamp
	DeletedAt     *time.Time // nullable timestamp (soft delete)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Bottle) TableName() string {
	return "bottles"

}

// BottleDB is the implementation of the storage interface for
// Bottle.
type BottleDB struct {
	Db gorm.DB
}

// NewBottleDB creates a new storage type.
func NewBottleDB(db gorm.DB) *BottleDB {
	return &BottleDB{Db: db}
}

// DB returns the underlying database.
func (m *BottleDB) DB() interface{} {
	return &m.Db
}

// BottleStorage represents the storage interface.
type BottleStorage interface {
	DB() interface{}
	List(ctx context.Context) []Bottle
	One(ctx context.Context, id int) (Bottle, error)
	Add(ctx context.Context, bottle Bottle) (Bottle, error)
	Update(ctx context.Context, bottle Bottle) error
	Delete(ctx context.Context, id int) error
	ListByAccount(ctx context.Context, account_id int) []Bottle
	OneByAccount(ctx context.Context, account_id, id int) (Bottle, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *BottleDB) TableName() string {
	return "bottles"

}

// CRUD Functions

// ListBottle returns an array of view: default
func (m *BottleDB) ListBottle(ctx context.Context) []app.Bottle {
	var objs []app.Bottle
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		return objs
	}

	return objs
}

// OneBottle returns an array of view: default
func (m *BottleDB) OneBottle(ctx context.Context, id int) app.Bottle {
	var native Bottle
	m.Db.Table(m.TableName()).Find(&native).Where("id = ?", id)
	var obj app.Bottle
	account := &app.AccountLink{}
	m.Db.Table("account").Find(&account).Where("id = ?", native.AccountID)
	obj.Links = &app.BottleLinks{
		Account: account,
	}

	return obj
}

// ListBottleFull returns an array of view: full
func (m *BottleDB) ListBottleFull(ctx context.Context) []app.BottleFull {
	var objs []app.BottleFull
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		return objs
	}

	return objs
}

// OneBottleFull returns an array of view: full
func (m *BottleDB) OneBottleFull(ctx context.Context, id int) app.BottleFull {
	var native Bottle
	m.Db.Table(m.TableName()).Find(&native).Where("id = ?", id)
	var obj app.BottleFull
	account := &app.AccountLink{}
	m.Db.Table("account").Find(&account).Where("id = ?", native.AccountID)
	obj.Links = &app.BottleLinks{
		Account: account,
	}

	return obj
}

// ListBottleTiny returns an array of view: tiny
func (m *BottleDB) ListBottleTiny(ctx context.Context) []app.BottleTiny {
	var objs []app.BottleTiny
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		return objs
	}

	return objs
}

// OneBottleTiny returns an array of view: tiny
func (m *BottleDB) OneBottleTiny(ctx context.Context, id int) app.BottleTiny {
	var native Bottle
	m.Db.Table(m.TableName()).Find(&native).Where("id = ?", id)
	var obj app.BottleTiny
	account := &app.AccountLink{}
	m.Db.Table("account").Find(&account).Where("id = ?", native.AccountID)
	obj.Links = &app.BottleLinks{
		Account: account,
	}

	return obj
}

// Add creates a new record.
func (m *BottleDB) Add(ctx context.Context, model Bottle) (Bottle, error) {
	err := m.Db.Create(&model).Error
	return model, err
}

// Update modifies a single record.
func (m *BottleDB) Update(ctx context.Context, model Bottle) error {
	obj := m.OneBottle(ctx, model.ID)
	err := m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *BottleDB) Delete(ctx context.Context, id int) error {
	var obj Bottle

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		return err
	}

	return nil
}

// account
