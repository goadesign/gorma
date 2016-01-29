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
	ID            int `sql:"index" gorm:"primary_key"`
	Color         string
	Country       string
	Myvintage     int
	Name          string
	Region        string
	Review        string
	Sweetness     int
	Varietal      string
	Vineyard      string
	Vintage       int        `sql:"index"`
	VinyardCounty string     `gorm:"column:vinyardcounty"`
	UpdatedAt     time.Time  // timestamp
	DeletedAt     *time.Time // nullable timestamp (soft delete)
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
	rows, err := m.Db.Table(m.TableName()).Select("name,vineyard,varietal,myvintage,id").Rows()
	defer rows.Close()
	if err != nil {
		return objs
	}
	var varietal string
	var vintage int
	var iD int
	var vineyard string
	var name string

	for rows.Next() {
		rows.Scan(&vintage, &iD, &vineyard, &varietal, &name)
		obj := app.Bottle{
			Name:     name,
			Vineyard: vineyard,
			Varietal: varietal,
			Vintage:  vintage,
			ID:       iD,
		}
		objs = append(objs, obj)

	}

	return objs
}

// ListBottleViewFull returns an array of view: full
func (m *BottleDB) ListBottleViewFull(ctx context.Context) []app.BottleViewFull {
	var objs []app.BottleViewFull
	rows, err := m.Db.Table(Bottle.TableName()).Select("updated_at,id,vinyardcounty,myvintage,sweetness,review,name,country,created_at,vineyard,color,region,varietal").Rows()
	defer rows.Close()
	if err != nil {
		return objs
	}
	var varietal string
	var iD int
	var vinyardCounty string
	var vintage int
	var sweetness int
	var review string
	var updatedAt time.Time
	var name string
	var country string
	var createdAt time.Time
	var vineyard string
	var color string
	var region string

	for rows.Next() {
		rows.Scan(&vintage, &sweetness, &review, &updatedAt, &iD, &vinyardCounty, &name, &country, &createdAt, &region, &vineyard, &color, &varietal)
		obj := app.BottleViewFull{
			ID:            iD,
			VinyardCounty: vinyardCounty,
			Vintage:       vintage,
			Sweetness:     sweetness,
			Review:        review,
			UpdatedAt:     updatedAt,
			Name:          name,
			Country:       country,
			CreatedAt:     createdAt,
			Vineyard:      vineyard,
			Color:         color,
			Region:        region,
			Varietal:      varietal,
		}
		objs = append(objs, obj)

	}

	return objs
}

// ListBottleViewTiny returns an array of view: tiny
func (m *BottleDB) ListBottleViewTiny(ctx context.Context) []app.BottleViewTiny {
	var objs []app.BottleViewTiny
	rows, err := m.Db.Table(Bottle.TableName()).Select("id,name").Rows()
	defer rows.Close()
	if err != nil {
		return objs
	}
	var name string
	var iD int

	for rows.Next() {
		rows.Scan(&iD, &name)
		obj := app.BottleViewTiny{
			ID:   iD,
			Name: name,
		}
		objs = append(objs, obj)

	}

	return objs
}

// account
