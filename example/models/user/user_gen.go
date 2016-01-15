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

package User

import "time"

// User type
type User struct {
	ID        int `gorm:"primary_key"`
	Bio       string
	Email     string
	State     string
	City      string
	Country   string
	Firstname string `sql:"blue"` //First name Description
	Role      string
	Lastname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
