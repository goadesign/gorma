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
	log "gopkg.in/inconshreveable/log15.v2"
	"time"
)

// This is the bottle model
type Bottle struct {
	ID            int `sql:"index" gorm:"primary_key"` // primary key
	AccountID     int // Belongs To Account
	Color         string
	Country       *string
	Href          string
	Name          string
	Rating        *int
	Region        *string
	Review        *string
	Sweetness     *int
	Varietal      string
	Vineyard      string
	Vintage       *string
	VinyardCounty *string    `gorm:"column:vinyardcounty"`
	UpdatedAt     *time.Time // timestamp
	DeletedAt     *time.Time // nullable timestamp (soft delete)
	CreatedAt     *time.Time // timestamp
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
}

// NewBottleDB creates a new storage type.
func NewBottleDB(db gorm.DB, logger log.Logger) *BottleDB {
	glog := logger.New("db", "Bottle")
	return &BottleDB{Db: db, Logger: glog}
}

// DB returns the underlying database.
func (m *BottleDB) DB() interface{} {
	return &m.Db
}

// BottleStorage represents the storage interface.
type BottleStorage interface {
	DB() interface{}
	List(ctx goa.Context) []Bottle
	One(ctx goa.Context, id int) (Bottle, error)
	Add(ctx goa.Context, bottle Bottle) (Bottle, error)
	Update(ctx goa.Context, bottle Bottle) error
	Delete(ctx goa.Context, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *BottleDB) TableName() string {
	return "bottles"

}

// Transformation

func BottleToBottle_v(source *Bottle) (target *app.Bottle) {
	target = new(app.Bottle)
	target.account = new(app.Account)
	target.account.created_at = source.account.created_at
	target.account.created_by = source.account.created_by
	target.account.href = source.account.href
	target.account.id = source.account.id
	target.account.name = source.account.name
	target.href = source.href
	target.id = source.id
	target.name = source.name
	target.rating = source.rating
	target.varietal = source.varietal
	target.vineyard = source.vineyard
	target.vintage = source.vintage
	return
}

// CRUD Functions
// ListBottle returns an array of view: default
func (m *BottleDB) ListBottle(ctx goa.Context) []app.Bottle {
	now := time.Now()
	defer ctx.Info("ListBottle", "duration", time.Since(now))
	var objs []app.Bottle
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Bottle", "error", err.Error())
		return objs
	}

	return objs
}

// OneBottle returns an array of view: default
func (m *BottleDB) OneBottle(ctx goa.Context, id int) app.Bottle {
	now := time.Now()
	defer ctx.Info("OneBottle", "duration", time.Since(now))
	var view app.Bottle
	var native Bottle

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	fmt.Println(native)
	return view
}

// Transformation

func BottleToBottleFull_v(source *Bottle) (target *app.BottleFull) {
	target = new(app.BottleFull)
	target.account = new(app.Account)
	target.account.created_at = source.account.created_at
	target.account.created_by = source.account.created_by
	target.account.href = source.account.href
	target.account.id = source.account.id
	target.account.name = source.account.name
	target.color = source.color
	target.country = source.country
	target.created_at = source.created_at
	target.href = source.href
	target.id = source.id
	target.name = source.name
	target.rating = source.rating
	target.region = source.region
	target.review = source.review
	target.sweetness = source.sweetness
	target.updated_at = source.updated_at
	target.varietal = source.varietal
	target.vineyard = source.vineyard
	target.vintage = source.vintage
	target.vinyard_county = source.vinyard_county
	return
}

// CRUD Functions
// ListBottleFull returns an array of view: full
func (m *BottleDB) ListBottleFull(ctx goa.Context) []app.BottleFull {
	now := time.Now()
	defer ctx.Info("ListBottleFull", "duration", time.Since(now))
	var objs []app.BottleFull
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Bottle", "error", err.Error())
		return objs
	}

	return objs
}

// OneBottleFull returns an array of view: full
func (m *BottleDB) OneBottleFull(ctx goa.Context, id int) app.BottleFull {
	now := time.Now()
	defer ctx.Info("OneBottleFull", "duration", time.Since(now))
	var view app.BottleFull
	var native Bottle

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	fmt.Println(native)
	return view
}

// Transformation

func BottleToBottleTiny_v(source *Bottle) (target *app.BottleTiny) {
	target = new(app.BottleTiny)
	target.href = source.href
	target.id = source.id
	target.name = source.name
	target.rating = source.rating
	return
}

// CRUD Functions
// ListBottleTiny returns an array of view: tiny
func (m *BottleDB) ListBottleTiny(ctx goa.Context) []app.BottleTiny {
	now := time.Now()
	defer ctx.Info("ListBottleTiny", "duration", time.Since(now))
	var objs []app.BottleTiny
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Bottle", "error", err.Error())
		return objs
	}

	return objs
}

// OneBottleTiny returns an array of view: tiny
func (m *BottleDB) OneBottleTiny(ctx goa.Context, id int) app.BottleTiny {
	now := time.Now()
	defer ctx.Info("OneBottleTiny", "duration", time.Since(now))
	var view app.BottleTiny
	var native Bottle

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	fmt.Println(native)
	return view
}

// GetBottle returns a single Bottle as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *BottleDB) GetBottle(ctx goa.Context, id int) Bottle {
	now := time.Now()
	defer ctx.Info("GetBottle", "duration", time.Since(now))
	var native Bottle
	m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native)
	return native
}

// Add creates a new record.
func (m *BottleDB) Add(ctx goa.Context, model Bottle) (Bottle, error) {
	now := time.Now()
	defer ctx.Info("AddBottle", "duration", time.Since(now))
	err := m.Db.Create(&model).Error
	if err != nil {
		ctx.Error("error updating Bottle", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *BottleDB) Update(ctx goa.Context, model Bottle) error {
	now := time.Now()
	defer ctx.Info("UpdateBottle", "duration", time.Since(now))
	obj := m.GetBottle(ctx, model.ID)
	err := m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *BottleDB) Delete(ctx goa.Context, id int) error {
	now := time.Now()
	defer ctx.Info("DeleteBottle", "duration", time.Since(now))
	var obj Bottle

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		ctx.Error("error retrieving Bottle", "error", err.Error())
		return err
	}

	return nil
}
