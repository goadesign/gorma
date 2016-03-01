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

// Review Model
type Review struct {
	ID         int `gorm:"primary_key"` // This is the Review Model PK field
	Comment    *string
	ProposalID int // Belongs To Proposal
	Rating     int
	UserID     int        // has many Review
	CreatedAt  time.Time  // timestamp
	DeletedAt  *time.Time // nullable timestamp (soft delete)
	UpdatedAt  time.Time  // timestamp
	Proposal   Proposal
	User       User
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Review) TableName() string {
	return "reviews"

}

// ReviewDB is the implementation of the storage interface for
// Review.
type ReviewDB struct {
	Db gorm.DB
}

// NewReviewDB creates a new storage type.
func NewReviewDB(db gorm.DB) *ReviewDB {
	return &ReviewDB{Db: db}
}

// DB returns the underlying database.
func (m *ReviewDB) DB() interface{} {
	return &m.Db
}

// ReviewStorage represents the storage interface.
type ReviewStorage interface {
	DB() interface{}
	List(ctx context.Context) []Review
	Get(ctx context.Context, id int) (Review, error)
	Add(ctx context.Context, review *Review) (*Review, error)
	Update(ctx context.Context, review *Review) error
	Delete(ctx context.Context, id int) error

	ListReview(ctx context.Context, proposalID int, userID int) []*app.Review
	OneReview(ctx context.Context, id int, proposalID int, userID int) (*app.Review, error)

	ListReviewLink(ctx context.Context, proposalID int, userID int) []*app.ReviewLink
	OneReviewLink(ctx context.Context, id int, proposalID int, userID int) (*app.ReviewLink, error)

	UpdateFromCreateReviewPayload(ctx context.Context, payload *app.CreateReviewPayload, id int) error

	UpdateFromUpdateReviewPayload(ctx context.Context, payload *app.UpdateReviewPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *ReviewDB) TableName() string {
	return "reviews"

}

// Belongs To Relationships

// ReviewFilterByProposal is a gorm filter for a Belongs To relationship.
func ReviewFilterByProposal(proposalID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if proposalID > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("proposal_id = ?", proposalID)
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// Belongs To Relationships

// ReviewFilterByUser is a gorm filter for a Belongs To relationship.
func ReviewFilterByUser(userID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if userID > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userID)
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Review as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *ReviewDB) Get(ctx context.Context, id int) (Review, error) {
	defer goa.MeasureSince([]string{"goa", "db", "review", "get"}, time.Now())

	var native Review
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.RecordNotFound {
		return Review{}, nil
	}

	return native, err
}

// List returns an array of Review
func (m *ReviewDB) List(ctx context.Context) []Review {
	defer goa.MeasureSince([]string{"goa", "db", "review", "list"}, time.Now())

	var objs []Review
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error listing Review", goa.KV{"error", err.Error()})
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *ReviewDB) Add(ctx context.Context, model *Review) (*Review, error) {
	defer goa.MeasureSince([]string{"goa", "db", "review", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.Error(ctx, "error updating Review", goa.KV{"error", err.Error()})
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *ReviewDB) Update(ctx context.Context, model *Review) error {
	defer goa.MeasureSince([]string{"goa", "db", "review", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *ReviewDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "review", "delete"}, time.Now())

	var obj Review

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.Error(ctx, "error retrieving Review", goa.KV{"error", err.Error()})
		return err
	}

	return nil
}

// ReviewFromCreateReviewPayload Converts source CreateReviewPayload to target Review model
// only copying the non-nil fields from the source.
func ReviewFromCreateReviewPayload(payload *app.CreateReviewPayload) *Review {
	review := &Review{}
	if payload.Comment != nil {
		review.Comment = payload.Comment
	}
	review.Rating = payload.Rating

	return review
}

// UpdateFromCreateReviewPayload applies non-nil changes from CreateReviewPayload to the model and saves it
func (m *ReviewDB) UpdateFromCreateReviewPayload(ctx context.Context, payload *app.CreateReviewPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "review", "updatefromcreateReviewPayload"}, time.Now())

	var obj Review
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.Error(ctx, "error retrieving Review", goa.KV{"error", err.Error()})
		return err
	}
	if payload.Comment != nil {
		obj.Comment = payload.Comment
	}
	obj.Rating = payload.Rating

	err = m.Db.Save(&obj).Error
	return err
}

// ReviewFromUpdateReviewPayload Converts source UpdateReviewPayload to target Review model
// only copying the non-nil fields from the source.
func ReviewFromUpdateReviewPayload(payload *app.UpdateReviewPayload) *Review {
	review := &Review{}
	if payload.Comment != nil {
		review.Comment = payload.Comment
	}
	if payload.Rating != nil {
		review.Rating = *payload.Rating
	}

	return review
}

// UpdateFromUpdateReviewPayload applies non-nil changes from UpdateReviewPayload to the model and saves it
func (m *ReviewDB) UpdateFromUpdateReviewPayload(ctx context.Context, payload *app.UpdateReviewPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "review", "updatefromupdateReviewPayload"}, time.Now())

	var obj Review
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.Error(ctx, "error retrieving Review", goa.KV{"error", err.Error()})
		return err
	}
	if payload.Comment != nil {
		obj.Comment = payload.Comment
	}
	if payload.Rating != nil {
		obj.Rating = *payload.Rating
	}

	err = m.Db.Save(&obj).Error
	return err
}
