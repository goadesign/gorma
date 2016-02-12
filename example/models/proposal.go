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
	log "gopkg.in/inconshreveable/log15.v2"
	"time"
)

// Proposal Model
type Proposal struct {
	ID        int `gorm:"primary_key"` // This is the Payload Model PK field
	Abstract  string
	Detail    string
	Reviews   []Review // has many Reviews
	Title     string
	UserID    int // has many Proposal
	Withdrawn *bool
	CreatedAt time.Time  // timestamp
	UpdatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	User      User
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Proposal) TableName() string {
	return "proposals"

}

// ProposalDB is the implementation of the storage interface for
// Proposal.
type ProposalDB struct {
	Db gorm.DB
	log.Logger
}

// NewProposalDB creates a new storage type.
func NewProposalDB(db gorm.DB, logger log.Logger) *ProposalDB {
	glog := logger.New("db", "Proposal")
	return &ProposalDB{Db: db, Logger: glog}
}

// DB returns the underlying database.
func (m *ProposalDB) DB() interface{} {
	return &m.Db
}

// ProposalStorage represents the storage interface.
type ProposalStorage interface {
	DB() interface{}
	List(ctx *goa.Context) []Proposal
	Get(ctx *goa.Context, id int) (Proposal, error)
	Add(ctx *goa.Context, proposal *Proposal) (*Proposal, error)
	Update(ctx *goa.Context, proposal *Proposal) error
	Delete(ctx *goa.Context, id int) error

	ListAppProposal(ctx *goa.Context, userid int) []*app.Proposal
	OneProposal(ctx *goa.Context, id int, userid int) (*app.Proposal, error)

	ListAppProposalLink(ctx *goa.Context, userid int) []*app.ProposalLink
	OneProposalLink(ctx *goa.Context, id int, userid int) (*app.ProposalLink, error)

	UpdateFromCreateProposalPayload(ctx *goa.Context, payload *app.CreateProposalPayload, id int) error

	UpdateFromUpdateProposalPayload(ctx *goa.Context, payload *app.UpdateProposalPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *ProposalDB) TableName() string {
	return "proposals"

}

// Belongs To Relationships
// ProposalFilterByUser is a gorm filter for a Belongs To relationship.
func ProposalFilterByUser(userid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if userid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

// CRUD Functions

// Get returns a single Proposal as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *ProposalDB) Get(ctx *goa.Context, id int) (Proposal, error) {
	now := time.Now()
	defer ctx.Info("Proposal:Get", "duration", time.Since(now))
	var native Proposal
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.RecordNotFound {
		return Proposal{}, nil
	}

	return native, err
}

// List returns an array of Proposal
func (m *ProposalDB) List(ctx *goa.Context) []Proposal {
	now := time.Now()
	defer ctx.Info("Proposal:List", "duration", time.Since(now))
	var objs []Proposal
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.RecordNotFound {
		ctx.Error("error listing Proposal", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *ProposalDB) Add(ctx *goa.Context, model *Proposal) (*Proposal, error) {
	now := time.Now()
	defer ctx.Info("Proposal:Add", "duration", time.Since(now))
	err := m.Db.Create(model).Error
	if err != nil {
		ctx.Error("error updating Proposal", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *ProposalDB) Update(ctx *goa.Context, model *Proposal) error {
	now := time.Now()
	defer ctx.Info("Proposal:Update", "duration", time.Since(now))
	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *ProposalDB) Delete(ctx *goa.Context, id int) error {
	now := time.Now()
	defer ctx.Info("Proposal:Delete", "duration", time.Since(now))
	var obj Proposal

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		ctx.Error("error retrieving Proposal", "error", err.Error())
		return err
	}

	return nil
}

// ProposalFromCreateProposalPayload Converts source CreateProposalPayload to target Proposal model
// only copying the non-nil fields from the source.
func ProposalFromCreateProposalPayload(payload *app.CreateProposalPayload) *Proposal {
	proposal := &Proposal{}
	proposal.Title = payload.Title
	proposal.Abstract = payload.Abstract
	proposal.Detail = payload.Detail
	if payload.Withdrawn != nil {
		proposal.Withdrawn = payload.Withdrawn
	}

	return proposal
}

// UpdateFromCreateProposalPayload applies non-nil changes from CreateProposalPayload to the model
// and saves it
func (m *ProposalDB) UpdateFromCreateProposalPayload(ctx *goa.Context, payload *app.CreateProposalPayload, id int) error {
	now := time.Now()
	defer ctx.Info("Proposal:Update", "duration", time.Since(now))
	var obj Proposal
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		ctx.Error("error retrieving Proposal", "error", err.Error())
		return err
	}
	obj.Abstract = payload.Abstract
	obj.Detail = payload.Detail
	if payload.Withdrawn != nil {
		obj.Withdrawn = payload.Withdrawn
	}
	obj.Title = payload.Title

	err = m.Db.Save(&obj).Error
	return err
}

// ProposalFromUpdateProposalPayload Converts source UpdateProposalPayload to target Proposal model
// only copying the non-nil fields from the source.
func ProposalFromUpdateProposalPayload(payload *app.UpdateProposalPayload) *Proposal {
	proposal := &Proposal{}
	if payload.Detail != nil {
		proposal.Detail = *payload.Detail
	}
	if payload.Withdrawn != nil {
		proposal.Withdrawn = payload.Withdrawn
	}
	if payload.Abstract != nil {
		proposal.Abstract = *payload.Abstract
	}
	if payload.Title != nil {
		proposal.Title = *payload.Title
	}

	return proposal
}

// UpdateFromUpdateProposalPayload applies non-nil changes from UpdateProposalPayload to the model
// and saves it
func (m *ProposalDB) UpdateFromUpdateProposalPayload(ctx *goa.Context, payload *app.UpdateProposalPayload, id int) error {
	now := time.Now()
	defer ctx.Info("Proposal:Update", "duration", time.Since(now))
	var obj Proposal
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		ctx.Error("error retrieving Proposal", "error", err.Error())
		return err
	}
	if payload.Abstract != nil {
		obj.Abstract = *payload.Abstract
	}
	if payload.Detail != nil {
		obj.Detail = *payload.Detail
	}
	if payload.Withdrawn != nil {
		obj.Withdrawn = payload.Withdrawn
	}
	if payload.Title != nil {
		obj.Title = *payload.Title
	}

	err = m.Db.Save(&obj).Error
	return err
}
