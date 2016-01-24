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
	"time"

	"github.com/goadesign/gorma/example/app"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// This is the bottle model
type Bottle struct {
	ID            int `gorm:"primary_key"`
	Color         string
	Country       string
	Name          string
	Region        string
	Review        string
	Sweetness     int
	Varietal      string
	Vineyard      string
	Vintage       int        `sql:"index"`
	VinyardCounty string     `gorm:"column:vinyardcounty"`
	CreatedAt     time.Time  // timestamp
	DeletedAt     *time.Time // nullable timestamp (soft delete)
	UpdatedAt     time.Time  // timestamp
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
	One(ctx context.Context) (Bottle, error)
	Add(ctx context.Context, bottle Bottle) (Bottle, error)
	Update(ctx context.Context, bottle Bottle) error
	Delete(ctx context.Context) error
}

// CRUD Functions

// List returns an array of records.
func (m *BottleDB) List(ctx context.Context) []Bottle {
	var objs []Bottle
	m.Db.Find(&objs)
	return objs
}

// One returns a single record by ID.
func (m *BottleDB) One(ctx context.Context) (Bottle, error) {

	var obj Bottle
	err := m.Db.Find(&obj).Where("").Error

	return obj, err
}

// Add creates a new record.
func (m *BottleDB) Add(ctx context.Context, model Bottle) (Bottle, error) {
	err := m.Db.Create(&model).Error
	return model, err
}

// Update modifies a single record.
func (m *BottleDB) Update(ctx context.Context, model Bottle) error {
	obj, err := m.One(ctx)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *BottleDB) Delete(ctx context.Context) error {
	var obj Bottle
	err := m.Db.Delete(&obj).Where("").Error

	if err != nil {
		return err
	}

	return nil
}

// Useful conversion functions

// ToBottle converts a model Bottle to an app Bottle.
func (m *Bottle) ToBottle() app.Bottle {
	payload := app.Bottle{}
	payload.UpdatedAt = &m.UpdatedAt
	payload.Country = &m.Country
	payload.CreatedAt = &m.CreatedAt
	payload.Name = m.Name
	payload.Sweetness = &m.Sweetness
	payload.Vineyard = m.Vineyard
	payload.ID = m.ID
	payload.Color = m.Color
	payload.Vintage = m.Vintage
	payload.VinyardCounty = &m.VinyardCounty
	payload.Region = &m.Region
	payload.Review = &m.Review
	payload.Varietal = m.Varietal
	return payload
}

// Convert from	default version BottlePayload to Bottle.
func BottleFromBottlePayload(t app.BottlePayload) Bottle {
	bottle := Bottle{}
	bottle.Review = *t.Review
	bottle.Varietal = *t.Varietal
	bottle.Vintage = *t.Vintage
	bottle.VinyardCounty = *t.VinyardCounty
	bottle.Region = *t.Region
	bottle.Country = *t.Country
	bottle.Name = *t.Name
	bottle.Color = *t.Color
	bottle.Sweetness = *t.Sweetness
	bottle.Vineyard = *t.Vineyard
	return bottle
}
