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

// Review Model
type Review struct {
	ID         int `gorm:"primary_key"` // This is the Review Model PK field
	Comment    *string
	ProposalID int // Belongs To Proposal
	Rating     int
	UserID     int        // has many Review
	CreatedAt  time.Time  // timestamp
	UpdatedAt  time.Time  // timestamp
	DeletedAt  *time.Time // nullable timestamp (soft delete)
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
	log.Logger
}

// NewReviewDB creates a new storage type.
func NewReviewDB(db gorm.DB, logger log.Logger) *ReviewDB {
	glog := logger.New("db", "Review")
	return &ReviewDB{Db: db, Logger: glog}
}

// DB returns the underlying database.
func (m *ReviewDB) DB() interface{} {
	return &m.Db
}

// ReviewStorage represents the storage interface.
type ReviewStorage interface {
	DB() interface{}
	List(ctx *goa.Context) []Review
	Get(ctx *goa.Context, id int) (Review, error)
	Add(ctx *goa.Context, review *Review) (*Review, error)
	Update(ctx *goa.Context, review *Review) error
	Delete(ctx *goa.Context, id int) error

	ListAppReview(ctx *goa.Context, proposalid int, userid int) []*app.Review
	OneReview(ctx *goa.Context, id int, proposalid int, userid int) (*app.Review, error)

	ListAppReviewLink(ctx *goa.Context, proposalid int, userid int) []*app.ReviewLink
	OneReviewLink(ctx *goa.Context, id int, proposalid int, userid int) (*app.ReviewLink, error)

	UpdateFromCreateReviewPayload(ctx *goa.Context, payload *app.CreateReviewPayload, id int) error

	UpdateFromUpdateReviewPayload(ctx *goa.Context, payload *app.UpdateReviewPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *ReviewDB) TableName() string {
	return "reviews"

}

// Belongs To Relationships
// ReviewFilterByProposal is a gorm filter for a Belongs To relationship.
func ReviewFilterByProposal(proposalid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if proposalid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("proposal_id = ?", proposalid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

// Belongs To Relationships
// ReviewFilterByUser is a gorm filter for a Belongs To relationship.
func ReviewFilterByUser(userid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
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

// Get returns a single Review as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *ReviewDB) Get(ctx *goa.Context, id int) (Review, error) {
	now := time.Now()
	defer ctx.Info("Review:Get", "duration", time.Since(now))
	var native Review
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.RecordNotFound {
		return Review{}, nil
	}

	return native, err
}

// List returns an array of Review
func (m *ReviewDB) List(ctx *goa.Context) []Review {
	now := time.Now()
	defer ctx.Info("Review:List", "duration", time.Since(now))
	var objs []Review
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.RecordNotFound {
		ctx.Error("error listing Review", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *ReviewDB) Add(ctx *goa.Context, model *Review) (*Review, error) {
	now := time.Now()
	defer ctx.Info("Review:Add", "duration", time.Since(now))
	err := m.Db.Create(model).Error
	if err != nil {
		ctx.Error("error updating Review", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *ReviewDB) Update(ctx *goa.Context, model *Review) error {
	now := time.Now()
	defer ctx.Info("Review:Update", "duration", time.Since(now))
	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *ReviewDB) Delete(ctx *goa.Context, id int) error {
	now := time.Now()
	defer ctx.Info("Review:Delete", "duration", time.Since(now))
	var obj Review

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		ctx.Error("error retrieving Review", "error", err.Error())
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

// UpdateFromCreateReviewPayload applies non-nil changes from CreateReviewPayload to the model
// and saves it
func (m *ReviewDB) UpdateFromCreateReviewPayload(ctx *goa.Context, payload *app.CreateReviewPayload, id int) error {
	now := time.Now()
	defer ctx.Info("Review:Update", "duration", time.Since(now))
	var obj Review
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		ctx.Error("error retrieving Review", "error", err.Error())
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

// UpdateFromUpdateReviewPayload applies non-nil changes from UpdateReviewPayload to the model
// and saves it
func (m *ReviewDB) UpdateFromUpdateReviewPayload(ctx *goa.Context, payload *app.UpdateReviewPayload, id int) error {
	now := time.Now()
	defer ctx.Info("Review:Update", "duration", time.Since(now))
	var obj Review
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		ctx.Error("error retrieving Review", "error", err.Error())
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
