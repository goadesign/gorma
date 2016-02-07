package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma/example/app/v1"
)

// ReviewV1Controller implements the v1 review resource.
type ReviewV1Controller struct {
	goa.Controller
}

// NewReviewV1Controller creates a review controller.
func NewReviewV1Controller(service goa.Service) v1.ReviewController {
	return &ReviewV1Controller{Controller: service.NewController("review v1")}
}

// Create runs the create action.
func (c *ReviewV1Controller) Create(ctx *v1.CreateReviewContext) error {
	return nil
}

// Delete runs the delete action.
func (c *ReviewV1Controller) Delete(ctx *v1.DeleteReviewContext) error {
	return nil
}

// List runs the list action.
func (c *ReviewV1Controller) List(ctx *v1.ListReviewContext) error {
	res := v1.ReviewCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *ReviewV1Controller) Show(ctx *v1.ShowReviewContext) error {
	res := &v1.Review{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *ReviewV1Controller) Update(ctx *v1.UpdateReviewContext) error {
	return nil
}
