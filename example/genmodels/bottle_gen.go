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

// This is the Bottle model
type Bottle struct {
	ID            int `gorm:"primary_key"` // primary key
	AccountID     int // Belongs To Account
	Color         *string
	Country       *string
	CreatedAt     time.Time
	DeletedAt     *time.Time
	Name          *string
	Rating        *int
	Region        *string
	Review        *string
	Sweetness     *int
	UpdatedAt     time.Time
	Varietal      *string
	Vineyard      *string
	Vintage       *string
	VinyardCounty *string
	Account       Account
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
	log.Logger
	cache *cache.Cache
}

// NewBottleDB creates a new storage type.
func NewBottleDB(db gorm.DB, logger log.Logger) *BottleDB {
	glog := logger.New("db", "Bottle")
	return &BottleDB{
		Db:     db,
		Logger: glog,
		cache:  cache.New(5*time.Minute, 30*time.Second),
	}

}

// DB returns the underlying database.
func (m *BottleDB) DB() interface{} {
	return &m.Db
}

// BottleStorage represents the storage interface.
type BottleStorage interface {
	DB() interface{}
	List(ctx goa.Context) []Bottle
	Get(ctx goa.Context, id int) (Bottle, error)
	Add(ctx goa.Context, bottle *Bottle) (Bottle, error)
	Update(ctx goa.Context, bottle *Bottle) error
	Delete(ctx goa.Context, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *BottleDB) TableName() string {
	return "bottles"

}

// CRUD Functions

// Get returns a single Bottle as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *BottleDB) Get(ctx *goa.Context, id int) Bottle {
	now := time.Now()
	defer ctx.Info("Bottle:Get", "duration", time.Since(now))
	var native Bottle
	m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native)
	go m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)

	return native
}

// List returns an array of Bottle
func (m *BottleDB) List(ctx *goa.Context) []Bottle {
	now := time.Now()
	defer ctx.Info("Bottle:List", "duration", time.Since(now))
	var objs []Bottle
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Bottle", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.
func (m *BottleDB) Add(ctx *goa.Context, model Bottle) (Bottle, error) {
	now := time.Now()
	defer ctx.Info("Bottle:Add", "duration", time.Since(now))
	err := m.Db.Create(&model).Error
	if err != nil {
		ctx.Error("error updating Bottle", "error", err.Error())
		return model, err
	}

	go m.cache.Set(strconv.Itoa(model.ID), model, cache.DefaultExpiration)
	return model, err
}

// Update modifies a single record.
func (m *BottleDB) Update(ctx *goa.Context, model Bottle) error {
	now := time.Now()
	defer ctx.Info("Bottle:Update", "duration", time.Since(now))
	obj := m.Get(ctx, model.ID)
	err := m.Db.Model(&obj).Updates(model).Error
	go func() {
		m.cache.Set(strconv.Itoa(model.ID), obj, cache.DefaultExpiration)
	}()

	return err
}

// Delete removes a single record.
func (m *BottleDB) Delete(ctx *goa.Context, id int) error {
	now := time.Now()
	defer ctx.Info("Bottle:Delete", "duration", time.Since(now))
	var obj Bottle

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		ctx.Error("error retrieving Bottle", "error", err.Error())
		return err
	}
	go m.cache.Delete(strconv.Itoa(id))
	return nil
}
