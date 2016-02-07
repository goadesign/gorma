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

	"time"
)

// MediaType Retrieval Functions
// ListUser returns an array of view: default
func (m *UserDB) ListAppUser(ctx *goa.Context) []*app.User {
	now := time.Now()
	defer ctx.Info("ListUser", "duration", time.Since(now))
	var objs []*app.User
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing User", "error", err.Error())
		return objs
	}

	return objs
}

func (m *User) UserToAppUser() *app.User {
	user := &app.User{}
	user.Lastname = m.Lastname
	user.Bio = m.Bio
	user.Country = m.Country
	user.Firstname = m.Firstname
	user.ID = &m.ID
	user.City = m.City
	user.Email = m.Email
	user.State = m.State

	return user
}

// OneAppUser returns an array of view: default
func (m *UserDB) OneUser(ctx *goa.Context, id int) *app.User {
	now := time.Now()
	var native User
	defer ctx.Info("OneUser", "duration", time.Since(now))

	m.Db.Table(m.TableName()).Preload("Proposals").Preload("Reviews").Where("id = ?", id).Find(&native)

	view := *native.UserToAppUser()
	return &view

}

// MediaType Retrieval Functions
// ListUserLink returns an array of view: link
func (m *UserDB) ListAppUserLink(ctx *goa.Context) []*app.UserLink {
	now := time.Now()
	defer ctx.Info("ListUserLink", "duration", time.Since(now))
	var objs []*app.UserLink
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing User", "error", err.Error())
		return objs
	}

	return objs
}

func (m *User) UserToAppUserLink() *app.UserLink {
	user := &app.UserLink{}
	user.ID = &m.ID
	user.Email = m.Email

	return user
}

// OneAppUserLink returns an array of view: link
func (m *UserDB) OneUserLink(ctx *goa.Context, id int) *app.UserLink {
	now := time.Now()
	var native User
	defer ctx.Info("OneUserLink", "duration", time.Since(now))

	m.Db.Table(m.TableName()).Preload("Proposals").Preload("Reviews").Where("id = ?", id).Find(&native)

	view := *native.UserToAppUserLink()
	return &view

}
