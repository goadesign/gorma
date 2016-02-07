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
	"time"

	"github.com/goadesign/goa"

	"github.com/goadesign/gorma/example/app/v1"
)

// v1
// MediaType Retrieval Functions
// ListReview returns an array of view: default
func (m *ReviewDB) ListV1Review(ctx *goa.Context, proposalid int, userid int) []*v1.Review {
	now := time.Now()
	defer ctx.Info("ListReview", "duration", time.Since(now))
	var objs []*v1.Review
	err := m.Db.Scopes(ReviewFilterByProposal(proposalid, &m.Db), ReviewFilterByUser(userid, &m.Db)).Table(m.TableName()).Find(&objs).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Review", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Review) ReviewToV1Review() *v1.Review {
	review := &v1.Review{}
	review.Comment = m.Comment
	review.Rating = m.Rating
	review.ID = &m.ID

	return review
}

// OneV1Review returns an array of view: default
func (m *ReviewDB) OneReview(ctx *goa.Context, id int) *v1.Review {
	now := time.Now()
	var native Review
	defer ctx.Info("OneReview", "duration", time.Since(now))

	m.Db.Table(m.TableName()).Preload("Proposal").Preload("User").Where("id = ?", id).Find(&native)

	view := *native.ReviewToV1Review()
	return &view

}

// v1
// MediaType Retrieval Functions
// ListReviewLink returns an array of view: link
func (m *ReviewDB) ListV1ReviewLink(ctx *goa.Context, proposalid int, userid int) []*v1.ReviewLink {
	now := time.Now()
	defer ctx.Info("ListReviewLink", "duration", time.Since(now))
	var objs []*v1.ReviewLink
	err := m.Db.Scopes(ReviewFilterByProposal(proposalid, &m.Db), ReviewFilterByUser(userid, &m.Db)).Table(m.TableName()).Find(&objs).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Review", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Review) ReviewToV1ReviewLink() *v1.ReviewLink {
	review := &v1.ReviewLink{}
	review.ID = &m.ID

	return review
}

// OneV1ReviewLink returns an array of view: link
func (m *ReviewDB) OneReviewLink(ctx *goa.Context, id int) *v1.ReviewLink {
	now := time.Now()
	var native Review
	defer ctx.Info("OneReviewLink", "duration", time.Since(now))

	m.Db.Table(m.TableName()).Preload("Proposal").Preload("User").Where("id = ?", id).Find(&native)

	view := *native.ReviewToV1ReviewLink()
	return &view

}
