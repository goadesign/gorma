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

	"github.com/gopheracademy/congo/models/review"
)

// Proposal type
type Proposal struct {
	ID        int `gorm:"primary_key"`
	Reviews   []review.ReviewModel
	UserID    int
	Abstract  string
	Withdrawn bool
	Detail    string
	Title     string
	Firstname string `gorm:"column:first_name"`
	M2reviews string
	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
}
