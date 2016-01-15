//************************************************************************//
// Generated Models
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/bketelsen/gorma/example
// --design=github.com/bketelsen/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package review

import (
	"os/user"

	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// Review type
type Review struct {
	ID         int `gorm:"primary_key"`
	Comment    string
	Rating     int
	Reviewers  []user.User
	ProposalID int
	User       User
}

// ReviewDB is the implementation of the storage interface for Review
type ReviewDB struct {
	Db gorm.DB
}

// NewReviewDB creates a new storage type
func NewReviewDB(db gorm.DB) *ReviewDB {
	return &ReviewDB{Db: db}
}

// DB returns  the underlying database
func (m *ReviewDB) DB() interface{} {
	return &m.Db
}

// Storage Interface
type ReviewStorage interface {
	DB() interface{}
	List(ctx context.Context) []Review
	One(ctx context.Context, id int) (Review, error)
	Add(ctx context.Context, o Review) (Review, error)
	Update(ctx context.Context, o Review) error
	Delete(ctx context.Context, id int) error
	ListByProposal(ctx context.Context, proposal_id int) []Review
	OneByProposal(ctx context.Context, proposal_id, id int) (Review, error)
}

// CRUD Functions

// One returns a single record by ID
func (m *ReviewDB) One(ctx context.Context, id int) (Review, error) {

	var obj Review
	err := m.Db.Find(&obj, id).Error

	return obj, err
}

// Add creates a new record
func (m *ReviewDB) Add(ctx context.Context, model Review) (Review, error) {
	err := m.Db.Create(&model).Error

	return model, err
}

// Update modifies a single record
func (m *ReviewDB) Update(ctx context.Context, model Review) error {
	obj, err := m.One(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record
func (m *ReviewDB) Delete(ctx context.Context, id int) error {
	var obj Review

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		return err
	}

	return nil
}

// Belongs To Relationships

// ReviewFilterByProposal is a gorm filter for a Belongs To relationship
func ReviewFilterByProposal(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if parentid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", parentid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

// ListByProposal returns an array of associated Proposal models
func (m *ReviewDB) ListByProposal(ctx context.Context, parentid int) []Review {
	var objs []Review
	m.Db.Scopes(ReviewFilterByProposal(parentid, &m.Db)).Find(&objs)
	return objs
}

// OneByProposal returns a single associated Proposal model
func (m *ReviewDB) OneByProposal(ctx context.Context, parentid, id int) (Review, error) {

	var obj Review
	err := m.Db.Scopes(ReviewFilterByProposal(parentid, &m.Db)).Find(&obj, id).Error

	return obj, err
}

// FilterReviewByProposal iterates a list and returns only those with the foreign key provided
func FilterReviewByProposal(parent *int, list []Review) []Review {
	var filtered []Review
	for _, o := range list {
		if o.ProposalID == int(*parent) {
			filtered = append(filtered, o)
		}
	}
	return filtered
}
