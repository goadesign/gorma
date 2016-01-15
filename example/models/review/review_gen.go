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

package Review

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

type ReviewDB struct {
	Db gorm.DB
}

func NewReviewDB(db gorm.DB) *ReviewDB {
	return &ReviewDB{Db: db}
}

func (m *ReviewDB) DB() interface{} {
	return &m.Db
}

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
