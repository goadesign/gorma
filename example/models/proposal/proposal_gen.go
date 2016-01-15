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

package Proposal

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
	"golang.org/x/net/context"
)

// Proposal type
type Proposal struct {
	ID        int `gorm:"primary_key"`
	Reviews   []review.Review
	UserID    int
	Abstract  string
	Title     string
	Withdrawn bool
	Detail    string
	FirstName string `gorm:"column:person_name"`
	M2reviews string
	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
}

func (m Proposal) TableName() string {
	return "proposals"
}

type ProposalDB struct {
	Db    gorm.DB
	cache *cache.Cache
}

func NewProposalDB(db gorm.DB) *ProposalDB {
	return &ProposalDB{
		Db:    db,
		cache: cache.New(5*time.Minute, 30*time.Second),
	}

}

func (m *ProposalDB) DB() interface{} {
	return &m.Db
}

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
