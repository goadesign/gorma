//************************************************************************//
// API "cellar": Models
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/bketelsen/gorma/example
// --design=github.com/bketelsen/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package bottle

import (
	"github.com/bketelsen/gorma/example/app"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// This is the bottle model
//  // Stores BottlePayload
type Bottle struct {
	ID        int             `gorm:"primary_key"` // This is the ID PK field
	Country   string          //
	Review    string          //
	Vintage   int             `sql:"index"` //
	Name      string          //
	Vineyard  string          //
	Sweetness int             //
	Varietal  string          //
	Color     string          //
	Region    string          //
	CreatedAt date.Timestamp  // timestamp
	DeletedAt *date.Timestamp // nullable timestamp (soft delete)
	UpdatedAt date.Timestamp  // timestamp
}

// BottleDB is the implementation of the storage interface for Bottle
type BottleDB struct {
	Db gorm.DB
}

// NewBottleDB creates a new storage type
func NewBottleDB(db gorm.DB) *BottleDB {
	return &BottleDB{Db: db}
}

// DB returns  the underlying database
func (m *BottleDB) DB() interface{} {
	return &m.Db
}

// Storage Interface
type BottleStorage interface {
	DB() interface{}
	List(ctx context.Context) []Bottle
	One(ctx context.Context, id int) (Bottle, error)
	Add(ctx context.Context, bottle Bottle) (Bottle, error)
	Update(ctx context.Context, bottle Bottle) error
	Delete(ctx context.Context, id int) error
}

// CRUD Functions
// One returns a single record by ID
func (m *BottleDB) One(ctx context.Context, id int) (Bottle, error) {

	var obj Bottle
	err := m.Db.Find(&obj, id).Error

	return obj, err
}

// Add creates a new record
func (m *BottleDB) Add(ctx context.Context, model Bottle) (Bottle, error) {
	err := m.Db.Create(&model).Error
	return model, err
}

// Update modifies a single record
func (m *BottleDB) Update(ctx context.Context, model Bottle) error {
	obj, err := m.One(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record
func (m *BottleDB) Delete(ctx context.Context, id int) error {
	var obj Bottle
	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		return err
	}

	return nil
}

// Useful conversion functions
func (m *BottleDB) ToBottlePayload() app.BottlePayload {
	// convert it here, please

	return &BottlePayload{}
}

// Convert from	default version CreateBottlePayload to Bottle
func BottleFromCreateBottlePayload() Bottle {

	return &Bottle{}
}

// Convert from	default version UpdateBottlePayload to Bottle
func BottleFromUpdateBottlePayload() Bottle {

	return &Bottle{}
}
