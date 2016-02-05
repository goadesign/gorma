//************************************************************************//
// API "cellar": Model Helpers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package genmodels

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma/example/app"
	"time"
)

// MediaType Retrieval Functions
// ListAccount returns an array of view: default
func (m *AccountDB) ListAccount(ctx goa.Context) []app.Account {
	now := time.Now()
	defer ctx.Info("ListAccount", "duration", time.Since(now))
	var objs []app.Account
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Account", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Account) AccountToAccount() *app.Account {
	account := &app.Account{}
	account.Name = m.Name
	account.ID = m.ID
	account.CreatedAt = &m.CreatedAt
	account.CreatedBy = &m.CreatedBy
	account.Href = m.Href

	return account
}

// OneAccount returns an array of view: default
func (m *AccountDB) OneAccount(ctx goa.Context, id int) *app.Account {
	now := time.Now()
	defer ctx.Info("OneAccount", "duration", time.Since(now))

	var native Account

	m.Db.Table(m.TableName()).Preload("Bottle").Where("id = ?", id).Find(&native)
	view := native.AccountToAccount()
	return view

}

// MediaType Retrieval Functions
// ListAccountLink returns an array of view: link
func (m *AccountDB) ListAccountLink(ctx goa.Context) []app.AccountLink {
	now := time.Now()
	defer ctx.Info("ListAccountLink", "duration", time.Since(now))
	var objs []app.AccountLink
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Account", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Account) AccountToAccountLink() *app.AccountLink {
	account := &app.AccountLink{}
	account.ID = m.ID
	account.Href = m.Href

	return account
}

// OneAccountLink returns an array of view: link
func (m *AccountDB) OneAccountLink(ctx goa.Context, id int) *app.AccountLink {
	now := time.Now()
	defer ctx.Info("OneAccountLink", "duration", time.Since(now))

	var native Account

	m.Db.Table(m.TableName()).Preload("Bottle").Where("id = ?", id).Find(&native)
	view := native.AccountToAccountLink()
	return view

}

// MediaType Retrieval Functions
// ListAccountTiny returns an array of view: tiny
func (m *AccountDB) ListAccountTiny(ctx goa.Context) []app.AccountTiny {
	now := time.Now()
	defer ctx.Info("ListAccountTiny", "duration", time.Since(now))
	var objs []app.AccountTiny
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Account", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Account) AccountToAccountTiny() *app.AccountTiny {
	account := &app.AccountTiny{}
	account.Href = m.Href
	account.ID = m.ID
	account.Name = m.Name

	return account
}

// OneAccountTiny returns an array of view: tiny
func (m *AccountDB) OneAccountTiny(ctx goa.Context, id int) *app.AccountTiny {
	now := time.Now()
	defer ctx.Info("OneAccountTiny", "duration", time.Since(now))

	var native Account

	m.Db.Table(m.TableName()).Preload("Bottle").Where("id = ?", id).Find(&native)
	view := native.AccountToAccountTiny()
	return view

}
