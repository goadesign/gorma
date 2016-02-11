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
	"time"
)

// MediaType Retrieval Functions
// ListProposal returns an array of view: default
func (m *ProposalDB) ListAppProposal(ctx *goa.Context, userid int) []*app.Proposal {
	now := time.Now()
	defer ctx.Info("ListProposal", "duration", time.Since(now))
	var objs []*app.Proposal
	err := m.Db.Scopes(ProposalFilterByUser(userid, &m.Db)).Table(m.TableName()).Find(&objs).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Proposal", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Proposal) ProposalToAppProposal() *app.Proposal {
	proposal := &app.Proposal{}
	proposal.Title = &m.Title
	proposal.Abstract = &m.Abstract
	proposal.Detail = &m.Detail
	proposal.ID = &m.ID

	return proposal
}

// OneAppProposal returns an array of view: default
func (m *ProposalDB) OneProposal(ctx *goa.Context, id int, userid int) (*app.Proposal, error) {
	now := time.Now()
	var native Proposal
	defer ctx.Info("OneProposal", "duration", time.Since(now))
	err := m.Db.Scopes(ProposalFilterByUser(userid, &m.Db)).Table(m.TableName()).Preload("Reviews").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		ctx.Error("error getting Proposal", "error", err.Error())
		return nil, err
	}

	view := *native.ProposalToAppProposal()
	return &view, err

}

// MediaType Retrieval Functions
// ListProposalLink returns an array of view: link
func (m *ProposalDB) ListAppProposalLink(ctx *goa.Context, userid int) []*app.ProposalLink {
	now := time.Now()
	defer ctx.Info("ListProposalLink", "duration", time.Since(now))
	var objs []*app.ProposalLink
	err := m.Db.Scopes(ProposalFilterByUser(userid, &m.Db)).Table(m.TableName()).Find(&objs).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Proposal", "error", err.Error())
		return objs
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
func (m *ProposalDB) OneProposalLink(ctx *goa.Context, id int, userid int) (*app.ProposalLink, error) {
	now := time.Now()
	var native Proposal
	defer ctx.Info("OneProposalLink", "duration", time.Since(now))
	err := m.Db.Scopes(ProposalFilterByUser(userid, &m.Db)).Table(m.TableName()).Preload("Reviews").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		ctx.Error("error getting Proposal", "error", err.Error())
		return nil, err
	}

	view := *native.ProposalToAppProposalLink()
	return &view, err

}
