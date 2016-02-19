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
	UpdatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	CreatedAt time.Time  // timestamp
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
}

// NewProposalDB creates a new storage type.
func NewProposalDB(db gorm.DB) *ProposalDB {
	return &ProposalDB{Db: db}
}

// DB returns the underlying database.
func (m *ProposalDB) DB() interface{} {
	return &m.Db
}

// ProposalStorage represents the storage interface.
type ProposalStorage interface {
	DB() interface{}
	List(ctx context.Context) []Proposal
	Get(ctx context.Context, id int) (Proposal, error)
	Add(ctx context.Context, proposal *Proposal) (*Proposal, error)
	Update(ctx context.Context, proposal *Proposal) error
	Delete(ctx context.Context, id int) error

	ListAppProposal(ctx context.Context, userid int) []*app.Proposal
	OneProposal(ctx context.Context, id int, userid int) (*app.Proposal, error)

	ListAppProposalLink(ctx context.Context, userid int) []*app.ProposalLink
	OneProposalLink(ctx context.Context, id int, userid int) (*app.ProposalLink, error)

	UpdateFromCreateProposalPayload(ctx context.Context, payload *app.CreateProposalPayload, id int) error

	UpdateFromUpdateProposalPayload(ctx context.Context, payload *app.UpdateProposalPayload, id int) error
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
func (m *ProposalDB) Get(ctx context.Context, id int) (Proposal, error) {
	now := time.Now()
	defer goa.MeasureSince([]string{"goa", "db", "proposal", "get"}, now)
	var native Proposal
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.RecordNotFound {
		return Proposal{}, nil
	}

	return native, err
}

// List returns an array of Proposal
func (m *ProposalDB) List(ctx context.Context) []Proposal {
	now := time.Now()
	defer goa.MeasureSince([]string{"goa", "db", "proposal", "list"}, now)
	var objs []Proposal
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error listing Proposal", goa.KV{"error", err.Error()})
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *ProposalDB) Add(ctx context.Context, model *Proposal) (*Proposal, error) {
	now := time.Now()
	defer goa.MeasureSince([]string{"goa", "db", "proposal", "add"}, now)
	err := m.Db.Create(model).Error
	if err != nil {
		goa.Error(ctx, "error updating Proposal", goa.KV{"error", err.Error()})
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *ProposalDB) Update(ctx context.Context, model *Proposal) error {
	now := time.Now()
	defer goa.MeasureSince([]string{"goa", "db", "proposal", "update"}, now)
	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *ProposalDB) Delete(ctx context.Context, id int) error {
	now := time.Now()
	defer goa.MeasureSince([]string{"goa", "db", "proposal", "delete"}, now)
	var obj Proposal

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.Error(ctx, "error retrieving Proposal", goa.KV{"error", err.Error()})
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
	if payload.Withdrawn != nil {
		proposal.Withdrawn = payload.Withdrawn
	}
	proposal.Detail = payload.Detail

	return proposal
}

// UpdateFromCreateProposalPayload applies non-nil changes from CreateProposalPayload to the model
// and saves it
func (m *ProposalDB) UpdateFromCreateProposalPayload(ctx context.Context, payload *app.CreateProposalPayload, id int) error {
	now := time.Now()

	defer goa.MeasureSince([]string{"goa", "db", "proposal", "updatefromcreateProposalPayload"}, now)
	var obj Proposal
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.Error(ctx, "error retrieving Proposal", goa.KV{"error", err.Error()})
		return err
	}
	obj.Detail = payload.Detail
	obj.Abstract = payload.Abstract
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
	if payload.Withdrawn != nil {
		proposal.Withdrawn = payload.Withdrawn
	}
	if payload.Title != nil {
		proposal.Title = *payload.Title
	}
	if payload.Abstract != nil {
		proposal.Abstract = *payload.Abstract
	}
	if payload.Detail != nil {
		proposal.Detail = *payload.Detail
	}

	return proposal
}

// UpdateFromUpdateProposalPayload applies non-nil changes from UpdateProposalPayload to the model
// and saves it
func (m *ProposalDB) UpdateFromUpdateProposalPayload(ctx context.Context, payload *app.UpdateProposalPayload, id int) error {
	now := time.Now()

	defer goa.MeasureSince([]string{"goa", "db", "proposal", "updatefromupdateProposalPayload"}, now)
	var obj Proposal
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.Error(ctx, "error retrieving Proposal", goa.KV{"error", err.Error()})
		return err
	}
	if payload.Title != nil {
		obj.Title = *payload.Title
	}
	if payload.Abstract != nil {
		obj.Abstract = *payload.Abstract
	}
	if payload.Withdrawn != nil {
		obj.Withdrawn = payload.Withdrawn
	}
	if payload.Detail != nil {
		obj.Detail = *payload.Detail
	}

	err = m.Db.Save(&obj).Error
	return err
}
