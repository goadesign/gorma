//************************************************************************//
// API "congo": Model Helpers
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
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions
// ListUser returns an array of view: default
func (m *TestTooDB) ListAppUser(ctx context.Context) []*app.User {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuser"}, time.Now())

	var native []*TestToo
	var objs []*app.User
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		goa.Error(ctx, "error listing TestToo", goa.KV{"error", err.Error()})
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.TestTooToAppUser())
	}

	return objs
}

func (m *TestToo) TestTooToAppUser() *app.User {
	testtoo := &app.User{}
	testtoo.Email = &m.Email
	testtoo.Firstname = &m.Firstname
	testtoo.State = m.State
	testtoo.Bio = m.Bio
	testtoo.City = m.City
	testtoo.Country = m.Country
	testtoo.Lastname = &m.Lastname

	return testtoo
}

// OneAppUser returns an array of view: default
func (m *TestTooDB) OneUser(ctx context.Context, idone int, idtwo int) (*app.User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuser"}, time.Now())

	var native TestToo
	err := m.Db.Scopes().Table(m.TableName()).Where("idone = ? and idtwo = ?", idone, idtwo).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error getting TestToo", goa.KV{"error", err.Error()})
		return nil, err
	}

	view := *native.TestTooToAppUser()
	return &view, err

}

// MediaType Retrieval Functions
// ListUserLink returns an array of view: link
func (m *TestTooDB) ListAppUserLink(ctx context.Context) []*app.UserLink {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuserlink"}, time.Now())

	var native []*TestToo
	var objs []*app.UserLink
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		goa.Error(ctx, "error listing TestToo", goa.KV{"error", err.Error()})
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.TestTooToAppUserLink())
	}

	return objs
}

func (m *TestToo) TestTooToAppUserLink() *app.UserLink {
	testtoo := &app.UserLink{}
	testtoo.Email = &m.Email

	return testtoo
}

// OneAppUserLink returns an array of view: link
func (m *TestTooDB) OneUserLink(ctx context.Context, idone int, idtwo int) (*app.UserLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuserlink"}, time.Now())

	var native TestToo
	err := m.Db.Scopes().Table(m.TableName()).Where("idone = ? and idtwo = ?", idone, idtwo).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error getting TestToo", goa.KV{"error", err.Error()})
		return nil, err
	}

	view := *native.TestTooToAppUserLink()
	return &view, err

}
