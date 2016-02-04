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
	"github.com/jinzhu/gorm"
	log "gopkg.in/inconshreveable/log15.v2"
	"time"
)

// This is the Container model
type Container struct {
	AccountID     int // Belongs To Account
	Color         string
	Country       *string
	CreatedAt     *time.Time
	Href          string
	ID            int
	Name          string
	Rating        *int
	Region        *string
	Review        *string
	Sweetness     *int
	UpdatedAt     *time.Time
	Varietal      string
	Vineyard      string
	Vintage       string
	VinyardCounty *string
	DeletedAt     *time.Time // nullable timestamp (soft delete)
	Account       Account
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Container) TableName() string {
	return "containers"

}

// ContainerDB is the implementation of the storage interface for
// Container.
type ContainerDB struct {
	Db gorm.DB
	log.Logger
}

// NewContainerDB creates a new storage type.
func NewContainerDB(db gorm.DB, logger log.Logger) *ContainerDB {
	glog := logger.New("db", "Container")
	return &ContainerDB{Db: db, Logger: glog}
}

// DB returns the underlying database.
func (m *ContainerDB) DB() interface{} {
	return &m.Db
}

// ContainerStorage represents the storage interface.
type ContainerStorage interface {
	DB() interface{}
	List(ctx goa.Context) []Container
	One(ctx goa.Context) (Container, error)
	Add(ctx goa.Context, container Container) (Container, error)
	Update(ctx goa.Context, container Container) error
	Delete(ctx goa.Context) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *ContainerDB) TableName() string {
	return "containers"

}

// GetContainer returns a single Container as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *ContainerDB) GetContainer(ctx goa.Context, id int) Container {
	now := time.Now()
	defer ctx.Info("GetContainer", "duration", time.Since(now))
	var native Container
	m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native)
	return native
}

// Add creates a new record.
func (m *ContainerDB) Add(ctx goa.Context, model Container) (Container, error) {
	now := time.Now()
	defer ctx.Info("AddContainer", "duration", time.Since(now))
	err := m.Db.Create(&model).Error
	if err != nil {
		ctx.Error("error updating Container", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *ContainerDB) Update(ctx goa.Context, model Container) error {
	now := time.Now()
	defer ctx.Info("UpdateContainer", "duration", time.Since(now))
	obj := m.GetContainer(ctx)
	err := m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *ContainerDB) Delete(ctx goa.Context) error {
	now := time.Now()
	defer ctx.Info("DeleteContainer", "duration", time.Since(now))
	var obj Container
	err := m.Db.Delete(&obj).Where("").Error

	if err != nil {
		ctx.Error("error retrieving Container", "error", err.Error())
		return err
	}

	return nil
}
