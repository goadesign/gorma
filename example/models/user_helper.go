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
func (m *UserDB) ListAppUser(ctx context.Context) []*app.User {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuser"}, time.Now())

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
	user.City = m.City
	user.Firstname = &m.Firstname
	user.ID = &m.ID
	user.Bio = m.Bio
	user.Country = m.Country
	user.Email = &m.Email
	user.Lastname = &m.Lastname
	user.State = m.State

	return user
}

// OneAppUser returns an array of view: default
func (m *UserDB) OneUser(ctx context.Context, id int) (*app.User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuser"}, time.Now())

	var native User
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
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuserlink"}, time.Now())

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
	user.Email = &m.Email
	user.ID = &m.ID

	return user
}

// OneAppUserLink returns an array of view: link
func (m *UserDB) OneUserLink(ctx context.Context, id int) (*app.UserLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuserlink"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Preload("Proposals").Preload("Reviews").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error getting User", goa.KV{"error", err.Error()})
		return nil, err
	}

	view := *native.UserToAppUserLink()
	return &view, err

}
