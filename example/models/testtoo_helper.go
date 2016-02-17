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

/*
func something(source *TestToo) (target *app.User) {
	target = new(app.User)
	target.Bio = source.Bio
	target.City = source.City
	target.Country = source.Country
	target.Email = source.Email
	target.Firstname = source.Firstname
	target.Lastname = source.Lastname
	target.State = source.State
	return
}

*/

// MediaType Retrieval Functions
// ListUser returns an array of view: default
func (m *TestTooDB) ListAppUser(ctx context.Context) []*app.User {
	now := time.Now()
	defer goa.Info(ctx, "ListUser", goa.KV{"duration", time.Since(now)})
	var objs []*app.User
	err := m.Db.Scopes().Table(m.TableName()).Find(&objs).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		goa.Error(ctx, "error listing TestToo", goa.KV{"error", err.Error()})
		return objs
	}

	return objs
}

func (m *TestToo) TestTooToAppUser() *app.User {
	testtoo := &app.User{}
	testtoo.Country = m.Country
	testtoo.Email = &m.Email
	testtoo.Bio = m.Bio
	testtoo.City = m.City
	testtoo.Firstname = &m.Firstname
	testtoo.Lastname = &m.Lastname
	testtoo.State = m.State

	return testtoo
}

// OneAppUser returns an array of view: default
func (m *TestTooDB) OneUser(ctx context.Context, idone int, idtwo int) (*app.User, error) {
	now := time.Now()
	var native TestToo
	defer goa.Info(ctx, "OneUser", goa.KV{"duration", time.Since(now)})
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
	defer goa.Info(ctx, "ListUserLink", goa.KV{"duration", time.Since(now)})
	var objs []*app.UserLink
	err := m.Db.Scopes().Table(m.TableName()).Find(&objs).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		goa.Error(ctx, "error listing TestToo", goa.KV{"error", err.Error()})
		return objs
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
	defer goa.Info(ctx, "OneUserLink", goa.KV{"duration", time.Since(now)})
	err := m.Db.Scopes().Table(m.TableName()).Where("idone = ? and idtwo = ?", idone, idtwo).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error getting TestToo", goa.KV{"error", err.Error()})
		return nil, err
	}

	view := *native.TestTooToAppUserLink()
	return &view, err

}
