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
	now := time.Now()

	defer goa.MeasureSince([]string{"goa", "db", "user", "listuser"}, now)
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
	testtoo.Country = m.Country
	testtoo.Email = &m.Email
	testtoo.City = m.City
	testtoo.Firstname = &m.Firstname
	testtoo.Lastname = &m.Lastname
	testtoo.State = m.State
	testtoo.Bio = m.Bio

	return testtoo
}

// OneAppUser returns an array of view: default
func (m *TestTooDB) OneUser(ctx context.Context, idone int, idtwo int) (*app.User, error) {
	now := time.Now()
	var native TestToo
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuser"}, now)
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
	now := time.Now()

	defer goa.MeasureSince([]string{"goa", "db", "user", "listuserlink"}, now)
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
	now := time.Now()
	var native TestToo
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuserlink"}, now)
	err := m.Db.Scopes().Table(m.TableName()).Where("idone = ? and idtwo = ?", idone, idtwo).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error getting TestToo", goa.KV{"error", err.Error()})
		return nil, err
	}

	view := *native.TestTooToAppUserLink()
	return &view, err

}
