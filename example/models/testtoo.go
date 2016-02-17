//************************************************************************//
// API "congo": Models
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma/example/app"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	log "gopkg.in/inconshreveable/log15.v2"
	"time"
)

// TestTooModel
type TestToo struct {
	Idone     int `gorm:"primary_key"` // This is one of the TestToo Model PK fields
	Idtwo     int `gorm:"primary_key"` // This is one of the TestToo Model PK fields
	Bio       *string
	City      *string
	Country   *string
	Email     string
	Firstname string
	Lastname  string
	State     *string
	DeletedAt *time.Time // nullable timestamp (soft delete)
	CreatedAt time.Time  // timestamp
	UpdatedAt time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m TestToo) TableName() string {
	return "test_toos"

}

// TestTooDB is the implementation of the storage interface for
// TestToo.
type TestTooDB struct {
	Db gorm.DB
	log.Logger
}

// NewTestTooDB creates a new storage type.
func NewTestTooDB(db gorm.DB, logger log.Logger) *TestTooDB {
	glog := logger.New("db", "TestToo")
	return &TestTooDB{Db: db, Logger: glog}
}

// DB returns the underlying database.
func (m *TestTooDB) DB() interface{} {
	return &m.Db
}

// TestTooStorage represents the storage interface.
type TestTooStorage interface {
	DB() interface{}
	List(ctx context.Context) []TestToo
	Get(ctx context.Context, idone int, idtwo int) (TestToo, error)
	Add(ctx context.Context, testtoo *TestToo) (*TestToo, error)
	Update(ctx context.Context, testtoo *TestToo) error
	Delete(ctx context.Context, idone int, idtwo int) error

	ListAppUser(ctx context.Context) []*app.User
	OneUser(ctx context.Context, idone int, idtwo int) (*app.User, error)

	ListAppUserLink(ctx context.Context) []*app.UserLink
	OneUserLink(ctx context.Context, idone int, idtwo int) (*app.UserLink, error)

	UpdateFromCreateUserPayload(ctx context.Context, payload *app.CreateUserPayload, idone int, idtwo int) error

	UpdateFromUpdateUserPayload(ctx context.Context, payload *app.UpdateUserPayload, idone int, idtwo int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *TestTooDB) TableName() string {
	return "test_toos"

}

// CRUD Functions

// Get returns a single TestToo as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *TestTooDB) Get(ctx context.Context, idone int, idtwo int) (TestToo, error) {
	now := time.Now()
	defer goa.Info(ctx, "TestToo:Get", goa.KV{"duration", time.Since(now)})
	var native TestToo
	err := m.Db.Table(m.TableName()).Where("idone = ? and idtwo = ?", idone, idtwo).Find(&native).Error
	if err == gorm.RecordNotFound {
		return TestToo{}, nil
	}

	return native, err
}

// List returns an array of TestToo
func (m *TestTooDB) ListTestToo(ctx context.Context) []TestToo {
	now := time.Now()
	defer goa.Info(ctx, "TestToo:List", goa.KV{"duration", time.Since(now)})
	var objs []TestToo
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error listing TestToo", goa.KV{"error", err.Error()})
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *TestTooDB) Add(ctx context.Context, model *TestToo) (*TestToo, error) {
	now := time.Now()
	defer goa.Info(ctx, "TestToo:Add", goa.KV{"duration", time.Since(now)})
	err := m.Db.Create(model).Error
	if err != nil {
		goa.Error(ctx, "error updating TestToo", goa.KV{"error", err.Error()})
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *TestTooDB) Update(ctx context.Context, model *TestToo) error {
	now := time.Now()
	defer goa.Info(ctx, "TestToo:Update", goa.KV{"duration", time.Since(now)})
	obj, err := m.Get(ctx, model.Idone, model.Idtwo)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *TestTooDB) Delete(ctx context.Context, idone int, idtwo int) error {
	now := time.Now()
	defer goa.Info(ctx, "TestToo:Delete", goa.KV{"duration", time.Since(now)})
	var obj TestToo
	err := m.Db.Delete(&obj).Where("idone = ? and idtwo = ?", idone, idtwo).Error

	if err != nil {
		goa.Error(ctx, "error retrieving TestToo", goa.KV{"error", err.Error()})
		return err
	}

	return nil
}

// TestTooFromCreateUserPayload Converts source CreateUserPayload to target TestToo model
// only copying the non-nil fields from the source.
func TestTooFromCreateUserPayload(payload *app.CreateUserPayload) *TestToo {
	testtoo := &TestToo{}
	testtoo.Lastname = payload.Lastname
	if payload.State != nil {
		testtoo.State = payload.State
	}
	if payload.Bio != nil {
		testtoo.Bio = payload.Bio
	}
	if payload.City != nil {
		testtoo.City = payload.City
	}
	testtoo.Firstname = payload.Firstname
	if payload.Country != nil {
		testtoo.Country = payload.Country
	}
	testtoo.Email = payload.Email

	return testtoo
}

// UpdateFromCreateUserPayload applies non-nil changes from CreateUserPayload to the model
// and saves it
func (m *TestTooDB) UpdateFromCreateUserPayload(ctx context.Context, payload *app.CreateUserPayload, idone int, idtwo int) error {
	now := time.Now()
	defer goa.Info(ctx, "TestToo:Update", goa.KV{"duration", time.Since(now)})
	var obj TestToo
	err := m.Db.Table(m.TableName()).Where("idone = ? and idtwo = ?", idone, idtwo).Find(&obj).Error
	if err != nil {
		goa.Error(ctx, "error retrieving TestToo", goa.KV{"error", err.Error()})
		return err
	}
	if payload.Country != nil {
		obj.Country = payload.Country
	}
	obj.Email = payload.Email
	if payload.Bio != nil {
		obj.Bio = payload.Bio
	}
	if payload.City != nil {
		obj.City = payload.City
	}
	obj.Firstname = payload.Firstname
	obj.Lastname = payload.Lastname
	if payload.State != nil {
		obj.State = payload.State
	}

	err = m.Db.Save(&obj).Error
	return err
}

// TestTooFromUpdateUserPayload Converts source UpdateUserPayload to target TestToo model
// only copying the non-nil fields from the source.
func TestTooFromUpdateUserPayload(payload *app.UpdateUserPayload) *TestToo {
	testtoo := &TestToo{}
	if payload.Country != nil {
		testtoo.Country = payload.Country
	}
	testtoo.Email = payload.Email
	if payload.State != nil {
		testtoo.State = payload.State
	}
	if payload.Bio != nil {
		testtoo.Bio = payload.Bio
	}
	if payload.City != nil {
		testtoo.City = payload.City
	}
	if payload.Firstname != nil {
		testtoo.Firstname = *payload.Firstname
	}
	if payload.Lastname != nil {
		testtoo.Lastname = *payload.Lastname
	}

	return testtoo
}

// UpdateFromUpdateUserPayload applies non-nil changes from UpdateUserPayload to the model
// and saves it
func (m *TestTooDB) UpdateFromUpdateUserPayload(ctx context.Context, payload *app.UpdateUserPayload, idone int, idtwo int) error {
	now := time.Now()
	defer goa.Info(ctx, "TestToo:Update", goa.KV{"duration", time.Since(now)})
	var obj TestToo
	err := m.Db.Table(m.TableName()).Where("idone = ? and idtwo = ?", idone, idtwo).Find(&obj).Error
	if err != nil {
		goa.Error(ctx, "error retrieving TestToo", goa.KV{"error", err.Error()})
		return err
	}
	if payload.Country != nil {
		obj.Country = payload.Country
	}
	obj.Email = payload.Email
	if payload.State != nil {
		obj.State = payload.State
	}
	if payload.Bio != nil {
		obj.Bio = payload.Bio
	}
	if payload.City != nil {
		obj.City = payload.City
	}
	if payload.Firstname != nil {
		obj.Firstname = *payload.Firstname
	}
	if payload.Lastname != nil {
		obj.Lastname = *payload.Lastname
	}

	err = m.Db.Save(&obj).Error
	return err
}
