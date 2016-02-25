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
// ListReview returns an array of view: default
func (m *ReviewDB) ListAppReview(ctx context.Context, proposalid int, userid int) []*app.Review {
	defer goa.MeasureSince([]string{"goa", "db", "review", "listreview"}, time.Now())

	var native []*Review
	var objs []*app.Review
	err := m.Db.Scopes(ReviewFilterByProposal(proposalid, &m.Db), ReviewFilterByUser(userid, &m.Db)).Table(m.TableName()).Find(&native).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		goa.Error(ctx, "error listing Review", goa.KV{"error", err.Error()})
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.ReviewToAppReview())
	}

	return objs
}

func (m *Review) ReviewToAppReview() *app.Review {
	review := &app.Review{}
	review.Comment = m.Comment
	review.ID = &m.ID
	review.Rating = &m.Rating

	return review
}

// OneAppReview returns an array of view: default
func (m *ReviewDB) OneReview(ctx context.Context, id int, proposalid int, userid int) (*app.Review, error) {
	defer goa.MeasureSince([]string{"goa", "db", "review", "onereview"}, time.Now())

	var native Review
	err := m.Db.Scopes(ReviewFilterByProposal(proposalid, &m.Db), ReviewFilterByUser(userid, &m.Db)).Table(m.TableName()).Preload("Proposal").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error getting Review", goa.KV{"error", err.Error()})
		return nil, err
	}

	view := *native.ReviewToAppReview()
	return &view, err

}

// MediaType Retrieval Functions
// ListReviewLink returns an array of view: link
func (m *ReviewDB) ListAppReviewLink(ctx context.Context, proposalid int, userid int) []*app.ReviewLink {
	defer goa.MeasureSince([]string{"goa", "db", "review", "listreviewlink"}, time.Now())

	var native []*Review
	var objs []*app.ReviewLink
	err := m.Db.Scopes(ReviewFilterByProposal(proposalid, &m.Db), ReviewFilterByUser(userid, &m.Db)).Table(m.TableName()).Find(&native).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		goa.Error(ctx, "error listing Review", goa.KV{"error", err.Error()})
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.ReviewToAppReviewLink())
	}

	return objs
}

func (m *Review) ReviewToAppReviewLink() *app.ReviewLink {
	review := &app.ReviewLink{}
	review.ID = &m.ID

	return review
}

// OneAppReviewLink returns an array of view: link
func (m *ReviewDB) OneReviewLink(ctx context.Context, id int, proposalid int, userid int) (*app.ReviewLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "review", "onereviewlink"}, time.Now())

	var native Review
	err := m.Db.Scopes(ReviewFilterByProposal(proposalid, &m.Db), ReviewFilterByUser(userid, &m.Db)).Table(m.TableName()).Preload("Proposal").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		goa.Error(ctx, "error getting Review", goa.KV{"error", err.Error()})
		return nil, err
	}

	view := *native.ReviewToAppReviewLink()
	return &view, err

}
