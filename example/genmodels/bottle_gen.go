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
	"github.com/goadesign/gorma/example/app"
	"github.com/jinzhu/gorm"
	log "gopkg.in/inconshreveable/log15.v2"
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
	Myvintage     *string
	Name          *string
	Region        *string
	Review        *string
	Sweetness     *int
	UpdatedAt     time.Time
	Varietal      *string
	Vineyard      *string
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

func BottleToBottleApp(source *Bottle) (target *app.Bottle) {
	target = new(app.Bottle)
	target.Account = new(app.Account)
	target.Account.CreatedAt = source.Account.CreatedAt
	target.Account.CreatedBy = source.Account.CreatedBy
	target.Account.Href = source.Account.Href
	target.Account.ID = source.Account.ID
	target.Account.Name = source.Account.Name
	target.Href = source.Href
	target.ID = source.ID
	target.Name = source.Name
	target.Rating = source.Rating
	target.Varietal = source.Varietal
	target.Vineyard = source.Vineyard
	target.Vintage = source.Vintage
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

//bottle := Bottle{}
/*
 		bottle.ID = t.ID
	bottle.Name = *t.Name
	bottle.CreatedAt = *t.CreatedAt
	bottle.Region = t.Region
	bottle.Review = t.Review
	bottle.Sweetness = t.Sweetness
	bottle.Vineyard = *t.Vineyard
	bottle.UpdatedAt = *t.UpdatedAt
	bottle.Country = t.Country
	bottle.Color = *t.Color
	bottle.Varietal = *t.Varietal
	bottle.VinyardCounty = t.VinyardCounty
*/
// return bottle

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

func BottleToBottleFullApp(source *Bottle) (target *app.BottleFull) {
	target = new(app.BottleFull)
	target.Account = new(app.Account)
	target.Account.CreatedAt = source.Account.CreatedAt
	target.Account.CreatedBy = source.Account.CreatedBy
	target.Account.Href = source.Account.Href
	target.Account.ID = source.Account.ID
	target.Account.Name = source.Account.Name
	target.Color = source.Color
	target.Country = source.Country
	target.CreatedAt = source.CreatedAt
	target.Href = source.Href
	target.ID = source.ID
	target.Name = source.Name
	target.Rating = source.Rating
	target.Region = source.Region
	target.Review = source.Review
	target.Sweetness = source.Sweetness
	target.UpdatedAt = source.UpdatedAt
	target.Varietal = source.Varietal
	target.Vineyard = source.Vineyard
	target.Vintage = source.Vintage
	target.VinyardCounty = source.VinyardCounty
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

//bottle := Bottle{}
/*
 		bottle.CreatedAt = *t.CreatedAt
	bottle.Region = t.Region
	bottle.Review = t.Review
	bottle.Sweetness = t.Sweetness
	bottle.Vineyard = *t.Vineyard
	bottle.UpdatedAt = *t.UpdatedAt
	bottle.Country = t.Country
	bottle.Color = *t.Color
	bottle.Varietal = *t.Varietal
	bottle.VinyardCounty = t.VinyardCounty
	bottle.ID = t.ID
	bottle.Name = *t.Name
*/
// return bottle

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

func BottleToBottleTinyApp(source *Bottle) (target *app.BottleTiny) {
	target = new(app.BottleTiny)
	target.Href = source.Href
	target.ID = source.ID
	target.Name = source.Name
	target.Rating = source.Rating
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

//bottle := Bottle{}
/*
 		bottle.UpdatedAt = *t.UpdatedAt
	bottle.Country = t.Country
	bottle.Color = *t.Color
	bottle.Varietal = *t.Varietal
	bottle.VinyardCounty = t.VinyardCounty
	bottle.ID = t.ID
	bottle.Name = *t.Name
	bottle.Vineyard = *t.Vineyard
	bottle.CreatedAt = *t.CreatedAt
	bottle.Region = t.Region
	bottle.Review = t.Review
	bottle.Sweetness = t.Sweetness
*/
// return bottle

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
