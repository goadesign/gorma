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
func something(source *User) (target *app.User) {
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
func (m *UserDB) ListAppUser(ctx context.Context) []*app.User {
	now := time.Now()
	defer goa.Info(ctx, "ListUser", goa.KV{"duration", time.Since(now)})
	var native []*User
	var objs []*app.User
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		goa.Error(ctx, "error listing User", goa.KV{"error", err.Error()})
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToAppUser())
	}

	return objs
}

func (m *User) UserToAppUser() *app.User {
	user := &app.User{}
	user.ID = &m.ID
	user.Lastname = &m.Lastname
	user.Country = m.Country
	user.Email = &m.Email
	user.Firstname = &m.Firstname
	user.Bio = m.Bio
	user.City = m.City
	user.State = m.State

	return user
}

// OneAppUser returns an array of view: default
func (m *UserDB) OneUser(ctx context.Context, id int) (*app.User, error) {
	now := time.Now()
	var native User
	defer goa.Info(ctx, "OneUser", goa.KV{"duration", time.Since(now)})
	err := m.Db.Scopes().Table(m.TableName()).Preload("Proposals").Preload("Reviews").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error getting User", goa.KV{"error", err.Error()})
		return nil, err
	}

	view := *native.UserToAppUser()
	return &view, err

}

// MediaType Retrieval Functions
// ListUserLink returns an array of view: link
func (m *UserDB) ListAppUserLink(ctx context.Context) []*app.UserLink {
	now := time.Now()
	defer goa.Info(ctx, "ListUserLink", goa.KV{"duration", time.Since(now)})
	var native []*User
	var objs []*app.UserLink
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		goa.Error(ctx, "error listing User", goa.KV{"error", err.Error()})
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToAppUserLink())
	}

	return objs
}

func (m *User) UserToAppUserLink() *app.UserLink {
	user := &app.UserLink{}
	user.ID = &m.ID
	user.Email = &m.Email

	return user
}

// OneAppUserLink returns an array of view: link
func (m *UserDB) OneUserLink(ctx context.Context, id int) (*app.UserLink, error) {
	now := time.Now()
	var native User
	defer goa.Info(ctx, "OneUserLink", goa.KV{"duration", time.Since(now)})
	err := m.Db.Scopes().Table(m.TableName()).Preload("Proposals").Preload("Reviews").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error getting User", goa.KV{"error", err.Error()})
		return nil, err
	}

	view := *native.UserToAppUserLink()
	return &view, err

}
