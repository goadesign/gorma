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
// ListProposal returns an array of view: default
func (m *ProposalDB) ListAppProposal(ctx context.Context, userid int) []*app.Proposal {
	defer goa.MeasureSince([]string{"goa", "db", "proposal", "listproposal"}, time.Now())

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
	proposal.Detail = &m.Detail
	proposal.ID = &m.ID
	for _, k := range m.Reviews {
		proposal.Reviews = append(proposal.Reviews, k.ReviewToAppReview())
	}
	proposal.Title = &m.Title
	proposal.Abstract = &m.Abstract

	return proposal
}

// OneAppProposal returns an array of view: default
func (m *ProposalDB) OneProposal(ctx context.Context, id int, userid int) (*app.Proposal, error) {
	defer goa.MeasureSince([]string{"goa", "db", "proposal", "oneproposal"}, time.Now())

	var native Proposal
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
	defer goa.MeasureSince([]string{"goa", "db", "proposal", "listproposallink"}, time.Now())

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
	proposal.ID = &m.ID
	proposal.Title = &m.Title

	return proposal
}

// OneAppProposalLink returns an array of view: link
func (m *ProposalDB) OneProposalLink(ctx context.Context, id int, userid int) (*app.ProposalLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "proposal", "oneproposallink"}, time.Now())

	var native Proposal
	err := m.Db.Scopes(ProposalFilterByUser(userid, &m.Db)).Table(m.TableName()).Preload("Reviews").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error getting Proposal", goa.KV{"error", err.Error()})
		return nil, err
	}

	view := *native.ProposalToAppProposalLink()
	return &view, err

}
