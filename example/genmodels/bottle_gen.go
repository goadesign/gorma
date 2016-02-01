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
	ID            int `sql:"index" gorm:"primary_key"`
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
	DeletedAt     *time.Time // nullable timestamp (soft delete)
	UpdatedAt     time.Time  // timestamp
	CreatedAt     time.Time  // timestamp
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
	One(ctx context.Context) (Bottle, error)
	Add(ctx context.Context, bottle Bottle) (Bottle, error)
	Update(ctx context.Context, bottle Bottle) error
	Delete(ctx context.Context) error
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
	rows, err := m.Db.Table(m.TableName()).Select("id,name,varietal,vineyard,vintage").Rows()
	defer rows.Close()
	if err != nil {
		return objs
	}
	for rows.Next() {
		var name string
		var varietal string
		var vintage string
		var iD int
		var vineyard string

		rows.Scan(&iD, &name, &varietal, &vineyard, &vintage)
		obj := app.Bottle{}
		obj.Vintage = vintage
		obj.ID = iD
		obj.Name = name
		obj.Varietal = varietal
		obj.Vineyard = vineyard
		objs = append(objs, obj)

	}
	return objs
}

// OneBottle returns an array of view: default
func (m *BottleDB) OneBottle(ctx context.Context, id int) app.Bottle {
	var obj app.Bottle
	row := m.Db.Table(m.TableName()).Select("id,name,varietal,vineyard,vintage").Row()
	var iD int
	var vineyard string
	var name string
	var varietal string
	var vintage string

	row.Scan(&iD, &name, &varietal, &vineyard, &vintage)
	obj.ID = iD
	obj.Name = name
	obj.Varietal = varietal
	obj.Vineyard = vineyard
	obj.Vintage = vintage
	return obj
}

// ListBottleFull returns an array of view: full
func (m *BottleDB) ListBottleFull(ctx context.Context) []app.BottleFull {
	var objs []app.BottleFull
	rows, err := m.Db.Table(m.TableName()).Select("color,country,created_at,id,name,region,review,sweetness,updated_at,varietal,vineyard,vintage,vinyardcounty").Rows()
	defer rows.Close()
	if err != nil {
		return objs
	}
	for rows.Next() {
		var country string
		var review string
		var vineyard string
		var varietal string
		var updatedAt time.Time
		var iD int
		var vinyardCounty string
		var color string
		var region string
		var vintage string
		var sweetness int
		var createdAt time.Time
		var name string

		rows.Scan(&color, &country, &createdAt, &iD, &name, &region, &review, &sweetness, &updatedAt, &varietal, &vineyard, &vintage, &vinyardCounty)
		obj := app.BottleFull{}
		obj.Vineyard = vineyard
		obj.VinyardCounty = &vinyardCounty
		obj.CreatedAt = &createdAt
		obj.Vintage = vintage
		obj.ID = iD
		obj.UpdatedAt = &updatedAt
		obj.Name = name
		obj.Sweetness = &sweetness
		obj.Varietal = varietal
		obj.Color = color
		obj.Country = &country
		obj.Region = &region
		obj.Review = &review
		objs = append(objs, obj)

	}
	return objs
}

// OneBottleFull returns an array of view: full
func (m *BottleDB) OneBottleFull(ctx context.Context, id int) app.BottleFull {
	var obj app.BottleFull
	row := m.Db.Table(m.TableName()).Select("color,country,created_at,id,name,region,review,sweetness,updated_at,varietal,vineyard,vintage,vinyardcounty").Row()
	var vineyard string
	var varietal string
	var updatedAt time.Time
	var iD int
	var color string
	var region string
	var vinyardCounty string
	var sweetness int
	var createdAt time.Time
	var name string
	var vintage string
	var review string
	var country string

	row.Scan(&color, &country, &createdAt, &iD, &name, &region, &review, &sweetness, &updatedAt, &varietal, &vineyard, &vintage, &vinyardCounty)
	obj.UpdatedAt = &updatedAt
	obj.Name = name
	obj.ID = iD
	obj.Country = &country
	obj.Region = &region
	obj.Review = &review
	obj.Sweetness = &sweetness
	obj.Varietal = varietal
	obj.Color = color
	obj.VinyardCounty = &vinyardCounty
	obj.Vineyard = vineyard
	obj.Vintage = vintage
	obj.CreatedAt = &createdAt
	return obj
}

// ListBottleTiny returns an array of view: tiny
func (m *BottleDB) ListBottleTiny(ctx context.Context) []app.BottleTiny {
	var objs []app.BottleTiny
	rows, err := m.Db.Table(m.TableName()).Select("id,name").Rows()
	defer rows.Close()
	if err != nil {
		return objs
	}
	for rows.Next() {
		var iD int
		var name string

		rows.Scan(&iD, &name)
		obj := app.BottleTiny{}
		obj.Name = name
		obj.ID = iD
		objs = append(objs, obj)

	}
	return objs
}

// OneBottleTiny returns an array of view: tiny
func (m *BottleDB) OneBottleTiny(ctx context.Context, id int) app.BottleTiny {
	var obj app.BottleTiny
	row := m.Db.Table(m.TableName()).Select("id,name").Row()
	var iD int
	var name string

	row.Scan(&iD, &name)
	obj.ID = iD
	obj.Name = name
	return obj
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

// account
