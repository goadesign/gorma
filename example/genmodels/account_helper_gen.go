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
	cache "github.com/patrickmn/go-cache"
	"strconv"
	"time"
)

// MediaType Retrieval Functions
// ListAccount returns an array of view: default
func (m *AccountDB) ListAppAccount(ctx *goa.Context) []*app.Account {
	now := time.Now()
	defer ctx.Info("ListAccount", "duration", time.Since(now))
	var objs []*app.Account
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Account", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Account) AccountToAppAccount() *app.Account {
	account := &app.Account{}
	account.CreatedAt = &m.CreatedAt
	account.CreatedBy = &m.CreatedBy
	account.Name = m.Name
	account.Href = m.Href
	account.ID = m.ID

	return account
}

// OneAppAccount returns an array of view: default
func (m *AccountDB) OneAccount(ctx *goa.Context, id int) *app.Account {
	now := time.Now()
	var native Account
	defer ctx.Info("OneAccount", "duration", time.Since(now))

	idstr := strconv.Itoa(id)
	cached, ok := m.cache.Get(idstr)
	if ok {
		native = cached.(Account)
		view := *native.AccountToAppAccount()
		go func() {
			m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)
		}()
		return &view
	}

	m.Db.Table(m.TableName()).Preload("Bottles").Where("id = ?", id).Find(&native)
	go func() {
		m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)
	}()
	view := *native.AccountToAppAccount()
	return &view

}

// MediaType Retrieval Functions
// ListAccountLink returns an array of view: link
func (m *AccountDB) ListAppAccountLink(ctx *goa.Context) []*app.AccountLink {
	now := time.Now()
	defer ctx.Info("ListAccountLink", "duration", time.Since(now))
	var objs []*app.AccountLink
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Account", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Account) AccountToAppAccountLink() *app.AccountLink {
	account := &app.AccountLink{}
	account.ID = m.ID
	account.Href = m.Href

	return account
}

// OneAppAccountLink returns an array of view: link
func (m *AccountDB) OneAccountLink(ctx *goa.Context, id int) *app.AccountLink {
	now := time.Now()
	var native Account
	defer ctx.Info("OneAccountLink", "duration", time.Since(now))

	idstr := strconv.Itoa(id)
	cached, ok := m.cache.Get(idstr)
	if ok {
		native = cached.(Account)
		view := *native.AccountToAppAccountLink()
		go func() {
			m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)
		}()
		return &view
	}

	m.Db.Table(m.TableName()).Preload("Bottles").Where("id = ?", id).Find(&native)
	go func() {
		m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)
	}()
	view := *native.AccountToAppAccountLink()
	return &view

}

// MediaType Retrieval Functions
// ListAccountTiny returns an array of view: tiny
func (m *AccountDB) ListAppAccountTiny(ctx *goa.Context) []*app.AccountTiny {
	now := time.Now()
	defer ctx.Info("ListAccountTiny", "duration", time.Since(now))
	var objs []*app.AccountTiny
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Account", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Account) AccountToAppAccountTiny() *app.AccountTiny {
	account := &app.AccountTiny{}
	account.ID = m.ID
	account.Href = m.Href
	account.Name = m.Name

	return account
}

// OneAppAccountTiny returns an array of view: tiny
func (m *AccountDB) OneAccountTiny(ctx *goa.Context, id int) *app.AccountTiny {
	now := time.Now()
	var native Account
	defer ctx.Info("OneAccountTiny", "duration", time.Since(now))

	idstr := strconv.Itoa(id)
	cached, ok := m.cache.Get(idstr)
	if ok {
		native = cached.(Account)
		view := *native.AccountToAppAccountTiny()
		go func() {
			m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)
		}()
		return &view
	}

	m.Db.Table(m.TableName()).Preload("Bottles").Where("id = ?", id).Find(&native)
	go func() {
		m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)
	}()
	view := *native.AccountToAppAccountTiny()
	return &view

}
