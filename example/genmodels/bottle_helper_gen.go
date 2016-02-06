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
// ListBottle returns an array of view: default
func (m *BottleDB) ListAppBottle(ctx *goa.Context) []*app.Bottle {
	now := time.Now()
	defer ctx.Info("ListBottle", "duration", time.Since(now))
	var objs []*app.Bottle
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Bottle", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Bottle) BottleToAppBottle() *app.Bottle {
	bottle := &app.Bottle{}
	bottle.Rating = m.Rating
	bottle.ID = m.ID
	tmp1 := &m.Account
	bottle.Account = tmp1.AccountToAppAccount()
	bottle.Name = *m.Name
	bottle.Varietal = *m.Varietal
	bottle.Vineyard = *m.Vineyard
	bottle.Vintage = *m.Vintage

	return bottle
}

// OneAppBottle returns an array of view: default
func (m *BottleDB) OneBottle(ctx *goa.Context, id int) *app.Bottle {
	now := time.Now()
	var native Bottle
	defer ctx.Info("OneBottle", "duration", time.Since(now))

	idstr := strconv.Itoa(id)
	cached, ok := m.cache.Get(idstr)
	if ok {
		native = cached.(Bottle)
		view := *native.BottleToAppBottle()
		go func() {
			m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)
		}()
		return &view
	}

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	go func() {
		m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)
	}()
	view := *native.BottleToAppBottle()
	return &view

}

// MediaType Retrieval Functions
// ListBottleFull returns an array of view: full
func (m *BottleDB) ListAppBottleFull(ctx *goa.Context) []*app.BottleFull {
	now := time.Now()
	defer ctx.Info("ListBottleFull", "duration", time.Since(now))
	var objs []*app.BottleFull
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Bottle", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Bottle) BottleToAppBottleFull() *app.BottleFull {
	bottle := &app.BottleFull{}
	bottle.Name = *m.Name
	bottle.VinyardCounty = m.VinyardCounty
	bottle.Rating = m.Rating
	bottle.ID = m.ID
	bottle.UpdatedAt = &m.UpdatedAt
	tmp1 := &m.Account
	bottle.Account = tmp1.AccountToAppAccount()
	bottle.Country = m.Country
	bottle.Varietal = *m.Varietal
	bottle.Vineyard = *m.Vineyard
	bottle.CreatedAt = &m.CreatedAt
	bottle.Color = *m.Color
	bottle.Region = m.Region
	bottle.Review = m.Review
	bottle.Sweetness = m.Sweetness
	bottle.Vintage = *m.Vintage

	return bottle
}

// OneAppBottleFull returns an array of view: full
func (m *BottleDB) OneBottleFull(ctx *goa.Context, id int) *app.BottleFull {
	now := time.Now()
	var native Bottle
	defer ctx.Info("OneBottleFull", "duration", time.Since(now))

	idstr := strconv.Itoa(id)
	cached, ok := m.cache.Get(idstr)
	if ok {
		native = cached.(Bottle)
		view := *native.BottleToAppBottleFull()
		go func() {
			m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)
		}()
		return &view
	}

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	go func() {
		m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)
	}()
	view := *native.BottleToAppBottleFull()
	return &view

}

// MediaType Retrieval Functions
// ListBottleTiny returns an array of view: tiny
func (m *BottleDB) ListAppBottleTiny(ctx *goa.Context) []*app.BottleTiny {
	now := time.Now()
	defer ctx.Info("ListBottleTiny", "duration", time.Since(now))
	var objs []*app.BottleTiny
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Bottle", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Bottle) BottleToAppBottleTiny() *app.BottleTiny {
	bottle := &app.BottleTiny{}
	bottle.Rating = m.Rating
	bottle.ID = m.ID
	bottle.Name = *m.Name

	return bottle
}

// OneAppBottleTiny returns an array of view: tiny
func (m *BottleDB) OneBottleTiny(ctx *goa.Context, id int) *app.BottleTiny {
	now := time.Now()
	var native Bottle
	defer ctx.Info("OneBottleTiny", "duration", time.Since(now))

	idstr := strconv.Itoa(id)
	cached, ok := m.cache.Get(idstr)
	if ok {
		native = cached.(Bottle)
		view := *native.BottleToAppBottleTiny()
		go func() {
			m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)
		}()
		return &view
	}

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	go func() {
		m.cache.Set(strconv.Itoa(native.ID), native, cache.DefaultExpiration)
	}()
	view := *native.BottleToAppBottleTiny()
	return &view

}
