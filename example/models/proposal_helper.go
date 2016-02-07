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

	"time"

	"github.com/goadesign/gorma/example/app/v1"
)

// v1
// MediaType Retrieval Functions
// ListProposal returns an array of view: default
func (m *ProposalDB) ListV1Proposal(ctx *goa.Context, userid int) []*v1.Proposal {
	now := time.Now()
	defer ctx.Info("ListProposal", "duration", time.Since(now))
	var objs []*v1.Proposal
	err := m.Db.Scopes(ProposalFilterByUser(userid, &m.Db)).Table(m.TableName()).Find(&objs).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Proposal", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Proposal) ProposalToV1Proposal() *v1.Proposal {
	proposal := &v1.Proposal{}
	proposal.ID = &m.ID
	proposal.Abstract = m.Abstract
	proposal.Title = m.Title
	proposal.Detail = m.Detail

	return proposal
}

// OneV1Proposal returns an array of view: default
func (m *ProposalDB) OneProposal(ctx *goa.Context, id int) *v1.Proposal {
	now := time.Now()
	var native Proposal
	defer ctx.Info("OneProposal", "duration", time.Since(now))

	m.Db.Table(m.TableName()).Preload("Reviews").Preload("User").Where("id = ?", id).Find(&native)

	view := *native.ProposalToV1Proposal()
	return &view

}

// v1
// MediaType Retrieval Functions
// ListProposalLink returns an array of view: link
func (m *ProposalDB) ListV1ProposalLink(ctx *goa.Context, userid int) []*v1.ProposalLink {
	now := time.Now()
	defer ctx.Info("ListProposalLink", "duration", time.Since(now))
	var objs []*v1.ProposalLink
	err := m.Db.Scopes(ProposalFilterByUser(userid, &m.Db)).Table(m.TableName()).Find(&objs).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Proposal", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Proposal) ProposalToV1ProposalLink() *v1.ProposalLink {
	proposal := &v1.ProposalLink{}
	proposal.Title = m.Title
	proposal.ID = &m.ID

	return proposal
}

// OneV1ProposalLink returns an array of view: link
func (m *ProposalDB) OneProposalLink(ctx *goa.Context, id int) *v1.ProposalLink {
	now := time.Now()
	var native Proposal
	defer ctx.Info("OneProposalLink", "duration", time.Since(now))

	m.Db.Table(m.TableName()).Preload("Reviews").Preload("User").Where("id = ?", id).Find(&native)

	view := *native.ProposalToV1ProposalLink()
	return &view

}
