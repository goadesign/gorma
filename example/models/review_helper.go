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

// ListReview returns an array of view: default.
func (m *ReviewDB) ListReview(ctx context.Context, proposalID int, userID int) []*app.Review {
	defer goa.MeasureSince([]string{"goa", "db", "review", "listreview"}, time.Now())

	var native []*Review
	var objs []*app.Review
	err := m.Db.Scopes(ReviewFilterByProposal(proposalID, &m.Db), ReviewFilterByUser(userID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Review", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.ReviewToReview())
	}

	return objs
}

// ReviewToReview returns the Review representation of Review.
func (m *Review) ReviewToReview() *app.Review {
	review := &app.Review{}
	review.Comment = m.Comment
	review.ID = &m.ID
	review.Rating = &m.Rating

	return review
}

// OneReview returns an array of view: default.
func (m *ReviewDB) OneReview(ctx context.Context, id int, proposalID int, userID int) (*app.Review, error) {
	defer goa.MeasureSince([]string{"goa", "db", "review", "onereview"}, time.Now())

	var native Review
	err := m.Db.Scopes(ReviewFilterByProposal(proposalID, &m.Db), ReviewFilterByUser(userID, &m.Db)).Table(m.TableName()).Preload("Proposal").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Review", "error", err.Error())
		return nil, err
	}

	view := *native.ReviewToReview()
	return &view, err
}

// MediaType Retrieval Functions

// ListReviewLink returns an array of view: link.
func (m *ReviewDB) ListReviewLink(ctx context.Context, proposalID int, userID int) []*app.ReviewLink {
	defer goa.MeasureSince([]string{"goa", "db", "review", "listreviewlink"}, time.Now())

	var native []*Review
	var objs []*app.ReviewLink
	err := m.Db.Scopes(ReviewFilterByProposal(proposalID, &m.Db), ReviewFilterByUser(userID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Review", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.ReviewToReviewLink())
	}

	return objs
}

// ReviewToReviewLink returns the Review representation of Review.
func (m *Review) ReviewToReviewLink() *app.ReviewLink {
	review := &app.ReviewLink{}
	review.ID = &m.ID

	return review
}

// OneReviewLink returns an array of view: link.
func (m *ReviewDB) OneReviewLink(ctx context.Context, id int, proposalID int, userID int) (*app.ReviewLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "review", "onereviewlink"}, time.Now())

	var native Review
	err := m.Db.Scopes(ReviewFilterByProposal(proposalID, &m.Db), ReviewFilterByUser(userID, &m.Db)).Table(m.TableName()).Preload("Proposal").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Review", "error", err.Error())
		return nil, err
	}

	view := *native.ReviewToReviewLink()
	return &view, err
}
