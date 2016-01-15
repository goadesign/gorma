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
)

// Review type
type Review struct {
	ID         int `gorm:"primary_key"`
	ProposalID int
	User       User
	Comment    string
	Rating     int
	Reviewers  []user.User
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
