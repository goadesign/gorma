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

import "github.com/gopheracademy/congo/models/user"

// Review type
type Review struct {
	ID         int `gorm:"primary_key"`
	Reviewers  []user.UserModel
	ProposalID int
	User       UserModel
	Comment    string
	Rating     int
}
