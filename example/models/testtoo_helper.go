//************************************************************************//
// API "congo": Model Helpers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/goadesign/gorma/example/design
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --version=v1.0.0
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

// MediaType Retrieval Functions

// ListUser returns an array of view: default.
func (m *TestTooDB) ListUser(ctx context.Context) []*app.User {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuser"}, time.Now())

	var native []*TestToo
	var objs []*app.User
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing TestToo", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.TestTooToUser())
	}

	return objs
}

// TestTooToUser loads a TestToo and builds the default view of media type User.
func (m *TestToo) TestTooToUser() *app.User {
	testtoo := &app.User{}
	testtoo.Bio = m.Bio
	testtoo.City = m.City
	testtoo.Country = m.Country
	testtoo.Email = &m.Email
	testtoo.Firstname = &m.Firstname
	testtoo.Lastname = &m.Lastname
	testtoo.State = m.State

	return testtoo
}

// OneUser loads a TestToo and builds the default view of media type User.
func (m *TestTooDB) OneUser(ctx context.Context, idone int, idtwo int) (*app.User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuser"}, time.Now())

	var native TestToo
	err := m.Db.Scopes().Table(m.TableName()).Where("idone = ? and idtwo = ?", idone, idtwo).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting TestToo", "error", err.Error())
		return nil, err
	}

	view := *native.TestTooToUser()
	return &view, err
}

// MediaType Retrieval Functions

// ListUserLink returns an array of view: link.
func (m *TestTooDB) ListUserLink(ctx context.Context) []*app.UserLink {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuserlink"}, time.Now())

	var native []*TestToo
	var objs []*app.UserLink
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing TestToo", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.TestTooToUserLink())
	}

	return objs
}

// TestTooToUserLink loads a TestToo and builds the link view of media type User.
func (m *TestToo) TestTooToUserLink() *app.UserLink {
	testtoo := &app.UserLink{}
	testtoo.Email = &m.Email

	return testtoo
}

// OneUserLink loads a TestToo and builds the link view of media type User.
func (m *TestTooDB) OneUserLink(ctx context.Context, idone int, idtwo int) (*app.UserLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuserlink"}, time.Now())

	var native TestToo
	err := m.Db.Scopes().Table(m.TableName()).Where("idone = ? and idtwo = ?", idone, idtwo).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting TestToo", "error", err.Error())
		return nil, err
	}

	view := *native.TestTooToUserLink()
	return &view, err
}
