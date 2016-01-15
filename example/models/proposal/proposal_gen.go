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
)

// Proposal type
type Proposal struct {
	ID        int `gorm:"primary_key"`
	Title     string
	UserID    int
	Withdrawn bool
	Abstract  string
	Reviews   []review.Review
	Detail    string
	Firstname string `gorm:"column:first_name"`
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
