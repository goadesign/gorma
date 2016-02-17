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
func something(source *Proposal) (target *app.Proposal) {
	target = new(app.Proposal)
	target.Abstract = source.Abstract
	target.Detail = source.Detail
	target.Title = source.Title
	return
}

*/

// MediaType Retrieval Functions
// ListProposal returns an array of view: default
func (m *ProposalDB) ListAppProposal(ctx context.Context, userid int) []*app.Proposal {
	now := time.Now()
	defer goa.Info(ctx, "ListProposal", goa.KV{"duration", time.Since(now)})
	var native []*Proposal
	var objs []*app.Proposal
	err := m.Db.Scopes(ProposalFilterByUser(userid, &m.Db)).Table(m.TableName()).Preload("Reviews").Find(&native).Error

	//	err := m.Db.Table(m.TableName()).Preload("Reviews").Find(&objs).Error
	if err != nil {
		goa.Error(ctx, "error listing Proposal", goa.KV{"error", err.Error()})
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.ProposalToAppProposal())
	}

	return objs
}

func (m *Proposal) ProposalToAppProposal() *app.Proposal {
	proposal := &app.Proposal{}
	var tmp1Collection app.ReviewLinkCollection
	for _, k := range m.Reviews {
		tmp1Collection = append(tmp1Collection, k.ReviewToAppReviewLink())
	}
	proposal.Links = &app.ProposalLinks{Reviews: tmp1Collection}
	for _, k := range m.Reviews {
		proposal.Reviews = append(proposal.Reviews, k.ReviewToAppReview())
	}
	proposal.ID = &m.ID
	proposal.Title = &m.Title
	proposal.Abstract = &m.Abstract
	proposal.Detail = &m.Detail

	return proposal
}

// OneAppProposal returns an array of view: default
func (m *ProposalDB) OneProposal(ctx context.Context, id int, userid int) (*app.Proposal, error) {
	now := time.Now()
	var native Proposal
	defer goa.Info(ctx, "OneProposal", goa.KV{"duration", time.Since(now)})
	err := m.Db.Scopes(ProposalFilterByUser(userid, &m.Db)).Table(m.TableName()).Preload("Reviews").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error getting Proposal", goa.KV{"error", err.Error()})
		return nil, err
	}

	view := *native.ProposalToAppProposal()
	return &view, err

}

// MediaType Retrieval Functions
// ListProposalLink returns an array of view: link
func (m *ProposalDB) ListAppProposalLink(ctx context.Context, userid int) []*app.ProposalLink {
	now := time.Now()
	defer goa.Info(ctx, "ListProposalLink", goa.KV{"duration", time.Since(now)})
	var native []*Proposal
	var objs []*app.ProposalLink
	err := m.Db.Scopes(ProposalFilterByUser(userid, &m.Db)).Table(m.TableName()).Preload("Reviews").Find(&native).Error

	//	err := m.Db.Table(m.TableName()).Preload("Reviews").Find(&objs).Error
	if err != nil {
		goa.Error(ctx, "error listing Proposal", goa.KV{"error", err.Error()})
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.ProposalToAppProposalLink())
	}

	return objs
}

func (m *Proposal) ProposalToAppProposalLink() *app.ProposalLink {
	proposal := &app.ProposalLink{}
	proposal.Title = &m.Title
	proposal.ID = &m.ID

	return proposal
}

// OneAppProposalLink returns an array of view: link
func (m *ProposalDB) OneProposalLink(ctx context.Context, id int, userid int) (*app.ProposalLink, error) {
	now := time.Now()
	var native Proposal
	defer goa.Info(ctx, "OneProposalLink", goa.KV{"duration", time.Since(now)})
	err := m.Db.Scopes(ProposalFilterByUser(userid, &m.Db)).Table(m.TableName()).Preload("Reviews").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error getting Proposal", goa.KV{"error", err.Error()})
		return nil, err
	}

	view := *native.ProposalToAppProposalLink()
	return &view, err

}
