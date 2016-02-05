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
// ListBottle returns an array of view: default
func (m *BottleDB) ListBottle(ctx goa.Context) []app.Bottle {
	now := time.Now()
	defer ctx.Info("ListBottle", "duration", time.Since(now))
	var objs []app.Bottle
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Bottle", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Bottle) BottleToBottle() *app.Bottle {
	bottle := &app.Bottle{}
	bottle.Vintage = *m.Vintage
	bottle.ID = m.ID
	bottle.Name = *m.Name
	bottle.Rating = m.Rating
	bottle.Vineyard = *m.Vineyard
	bottle.Account = m.Account.AccountToAccount()
	bottle.Varietal = *m.Varietal

	return bottle
}

// OneBottle returns an array of view: default
func (m *BottleDB) OneBottle(ctx goa.Context, id int) *app.Bottle {
	now := time.Now()
	defer ctx.Info("OneBottle", "duration", time.Since(now))

	var native Bottle

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	view := native.BottleToBottle()
	return view

}

// MediaType Retrieval Functions
// ListBottleFull returns an array of view: full
func (m *BottleDB) ListBottleFull(ctx goa.Context) []app.BottleFull {
	now := time.Now()
	defer ctx.Info("ListBottleFull", "duration", time.Since(now))
	var objs []app.BottleFull
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Bottle", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Bottle) BottleToBottleFull() *app.BottleFull {
	bottle := &app.BottleFull{}
	bottle.CreatedAt = &m.CreatedAt
	bottle.Region = m.Region
	bottle.Vintage = *m.Vintage
	bottle.ID = m.ID
	bottle.UpdatedAt = &m.UpdatedAt
	bottle.Name = *m.Name
	bottle.Review = m.Review
	bottle.VinyardCounty = m.VinyardCounty
	bottle.Sweetness = m.Sweetness
	bottle.Rating = m.Rating
	bottle.Account = m.Account.AccountToAccount()
	bottle.Color = *m.Color
	bottle.Country = m.Country
	bottle.Varietal = *m.Varietal
	bottle.Vineyard = *m.Vineyard

	return bottle
}

// OneBottleFull returns an array of view: full
func (m *BottleDB) OneBottleFull(ctx goa.Context, id int) *app.BottleFull {
	now := time.Now()
	defer ctx.Info("OneBottleFull", "duration", time.Since(now))

	var native Bottle

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	view := native.BottleToBottleFull()
	return view

}

// MediaType Retrieval Functions
// ListBottleTiny returns an array of view: tiny
func (m *BottleDB) ListBottleTiny(ctx goa.Context) []app.BottleTiny {
	now := time.Now()
	defer ctx.Info("ListBottleTiny", "duration", time.Since(now))
	var objs []app.BottleTiny
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Bottle", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Bottle) BottleToBottleTiny() *app.BottleTiny {
	bottle := &app.BottleTiny{}
	bottle.ID = m.ID
	bottle.Name = *m.Name
	bottle.Rating = m.Rating

	return bottle
}

// OneBottleTiny returns an array of view: tiny
func (m *BottleDB) OneBottleTiny(ctx goa.Context, id int) *app.BottleTiny {
	now := time.Now()
	defer ctx.Info("OneBottleTiny", "duration", time.Since(now))

	var native Bottle

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	view := native.BottleToBottleTiny()
	return view

}
