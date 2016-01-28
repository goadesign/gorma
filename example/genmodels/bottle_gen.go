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
	DeletedAt     *time.Time // nullable timestamp (soft delete)
	CreatedAt     time.Time  // timestamp
	UpdatedAt     time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Bottle) TableName() string {
	return bottles

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

// ListBottle returns an array of view: default
func (m *BottleDB) ListBottle(ctx context.Context) []app.Bottle {
	var objs []app.Bottle
	rows, err := m.Db.Table(m.TableName()).Select("vineyard,myvintage,id,name,varietal").Rows()
	defer rows.Close()
	if err != nil {
		return objs
	}
	var varietal string
	var name string
	var vineyard string
	var vintage int
	var iD int

	for rows.Next() {
		rows.Scan(&vineyard, &vintage, &iD, &name, &varietal)
		obj := app.Bottle{
			Varietal: varietal,
			ID:       iD,
			Name:     name,
			Vineyard: vineyard,
			Vintage:  vintage,
		}
		objs = append(objs, obj)

	}

	return objs
}

// ListBottleViewFull returns an array of view: full
func (m *BottleDB) ListBottleViewFull(ctx context.Context) []app.BottleViewFull {
	var objs []app.BottleViewFull
	rows, err := m.Db.Table(m.TableName()).Select("vinyardcounty,myvintage,name,review,country,id,vineyard,color,sweetness,region,created_at,updated_at,varietal").Rows()
	defer rows.Close()
	if err != nil {
		return objs
	}
	var vinyardCounty string
	var vintage int
	var name string
	var review string
	var iD int
	var vineyard string
	var color string
	var country string
	var createdAt time.Time
	var updatedAt time.Time
	var varietal string
	var sweetness int
	var region string

	for rows.Next() {
		rows.Scan(&vinyardCounty, &vintage, &review, &name, &vineyard, &color, &country, &iD, &varietal, &sweetness, &region, &createdAt, &updatedAt)
		obj := app.BottleViewFull{
			ID:            iD,
			Vineyard:      vineyard,
			Color:         color,
			Country:       country,
			Varietal:      varietal,
			Sweetness:     sweetness,
			Region:        region,
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
			VinyardCounty: vinyardCounty,
			Vintage:       vintage,
			Name:          name,
			Review:        review,
		}
		objs = append(objs, obj)

	}

	return objs
}

// ListBottleViewTiny returns an array of view: tiny
func (m *BottleDB) ListBottleViewTiny(ctx context.Context) []app.BottleViewTiny {
	var objs []app.BottleViewTiny
	rows, err := m.Db.Table(m.TableName()).Select("id,name").Rows()
	defer rows.Close()
	if err != nil {
		return objs
	}
	var iD int
	var name string

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
