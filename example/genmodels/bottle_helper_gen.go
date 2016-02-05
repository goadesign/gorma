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
	bottle.ID = m.ID
	bottle.Varietal = *m.Varietal
	bottle.Name = *m.Name
	bottle.Vineyard = *m.Vineyard
	bottle.Account = m.Account.AccountToAccount()
	bottle.Vintage = *m.Vintage
	bottle.Rating = m.Rating

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
	bottle.Review = m.Review
	bottle.Account = m.Account.AccountToAccount()
	bottle.Vintage = *m.Vintage
	bottle.VinyardCounty = m.VinyardCounty
	bottle.Rating = m.Rating
	bottle.ID = m.ID
	bottle.UpdatedAt = &m.UpdatedAt
	bottle.Color = *m.Color
	bottle.Region = m.Region
	bottle.Sweetness = m.Sweetness
	bottle.Varietal = *m.Varietal
	bottle.Country = m.Country
	bottle.Name = *m.Name
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
	bottle.Name = *m.Name
	bottle.Account = m.Account.AccountToAccount()
	bottle.Rating = m.Rating
	bottle.ID = m.ID

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
