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

package proposal

import (
	"strconv"
	"time"

	"github.com/bketelsen/gorma/example/models/review"
	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
	"golang.org/x/net/context"
)

// Proposal type
type Proposal struct {
	ID        int `gorm:"primary_key"`
	Withdrawn bool
	Abstract  string
	FirstName string `gorm:"column:person_name"`
	M2reviews string
	Detail    string
	Title     string
	Reviews   []review.Review
	UserID    int
	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
}

// TableName overrides the table name settings in gorm
func (m Proposal) TableName() string {
	return "proposals"
}

// ProposalDB is the implementation of the storage interface for Proposal
type ProposalDB struct {
	Db    gorm.DB
	cache *cache.Cache
}

// NewProposalDB creates a new storage type
func NewProposalDB(db gorm.DB) *ProposalDB {
	return &ProposalDB{
		Db:    db,
		cache: cache.New(5*time.Minute, 30*time.Second),
	}

}

// DB returns  the underlying database
func (m *ProposalDB) DB() interface{} {
	return &m.Db
}

// Storage Interface
type ProposalStorage interface {
	DB() interface{}
	List(ctx context.Context) []Proposal
	One(ctx context.Context, id int) (Proposal, error)
	Add(ctx context.Context, o Proposal) (Proposal, error)
	Update(ctx context.Context, o Proposal) error
	Delete(ctx context.Context, id int) error
	ListByUser(ctx context.Context, user_id int) []Proposal
	OneByUser(ctx context.Context, user_id, id int) (Proposal, error)
	ListReviews(context.Context, int) []review.Review
	AddReviews(context.Context, int, int) error
	DeleteReviews(context.Context, int, int) error
}

// stub for marshal/unmarshal functions
func (m *Proposal) Validate() error {
	return nil
}

// CRUD Functions

// One returns a single record by ID
func (m *ProposalDB) One(ctx context.Context, id int) (Proposal, error) {
	//first attempt to retrieve from cache
	o, found := m.cache.Get(strconv.Itoa(id))
	if found {
		return o.(Proposal), nil
	}
	// fallback to database if not found
	var obj Proposal
	err := m.Db.Find(&obj, id).Error
	go m.cache.Set(strconv.Itoa(id), obj, cache.DefaultExpiration)
	return obj, err
}

// Add creates a new record
func (m *ProposalDB) Add(ctx context.Context, model Proposal) (Proposal, error) {
	err := m.Db.Create(&model).Error
	go m.cache.Set(strconv.Itoa(model.ID), model, cache.DefaultExpiration)
	return model, err
}

// Update modifies a single record
func (m *ProposalDB) Update(ctx context.Context, model Proposal) error {
	obj, err := m.One(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	go func() {
		obj, err := m.One(ctx, model.ID)
		if err == nil {
			m.cache.Set(strconv.Itoa(model.ID), obj, cache.DefaultExpiration)
		}
	}()

	return err
}

// Delete removes a single record
func (m *ProposalDB) Delete(ctx context.Context, id int) error {
	var obj Proposal

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		return err
	}
	go m.cache.Delete(strconv.Itoa(id))
	return nil
}

// Belongs To Relationships

// ProposalFilterByUser is a gorm filter for a Belongs To relationship
func ProposalFilterByUser(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if parentid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id", parentid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

// ListByUser returns an array of associated User models
func (m *ProposalDB) ListByUser(ctx context.Context, parentid int) []Proposal {
	var objs []Proposal
	m.Db.Scopes(ProposalFilterByUser(parentid, &m.Db)).Find(&objs)
	return objs
}

// OneByUser returns a single associated User model
func (m *ProposalDB) OneByUser(ctx context.Context, parentid, id int) (Proposal, error) {
	//first attempt to retrieve from cache
	o, found := m.cache.Get(strconv.Itoa(id))
	if found {
		return o.(Proposal), nil
	}
	// fallback to database if not found
	var obj Proposal
	err := m.Db.Scopes(ProposalFilterByUser(parentid, &m.Db)).Find(&obj, id).Error
	go m.cache.Set(strconv.Itoa(id), obj, cache.DefaultExpiration)
	return obj, err
}

// Many To Many Relationships

// DeleteReview removes a Review/Proposal entry from the join table
func (m *ProposalDB) DeleteReview(ctx context.Context, proposalID, reviewID int) error {
	var obj Proposal
	obj.ID = proposalID
	var assoc review.Review
	var err error
	assoc.ID = reviewID
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Association("Reviews").Delete(assoc).Error
	if err != nil {
		return err
	}
	return nil
}

// AddReview creates a new Review/Proposal entry in the join table
func (m *ProposalDB) AddReview(ctx context.Context, proposalID, reviewID int) error {
	var proposal Proposal
	proposal.ID = proposalID
	var assoc review.Review
	assoc.ID = reviewID
	err := m.Db.Model(&proposal).Association("Reviews").Append(assoc).Error
	if err != nil {
		return err
	}
	return nil
}

// ListReview returns a list of the Review models related to this Proposal
func (m *ProposalDB) ListReview(ctx context.Context, proposalID int) []review.Review {
	var list []review.Review
	var obj Proposal
	obj.ID = proposalID
	m.Db.Model(&obj).Association("Reviews").Find(&list)
	return list
}

// FilterProposalByUser iterates a list and returns only those with the foreign key provided
func FilterProposalByUser(parent *int, list []Proposal) []Proposal {
	var filtered []Proposal
	for _, o := range list {
		if o.UserID == int(*parent) {
			filtered = append(filtered, o)
		}
	}
	return filtered
}
