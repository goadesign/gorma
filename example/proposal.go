package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma/example/app/v1"
)

// ProposalV1Controller implements the v1 proposal resource.
type ProposalV1Controller struct {
	goa.Controller
}

// NewProposalV1Controller creates a proposal controller.
func NewProposalV1Controller(service goa.Service) v1.ProposalController {
	return &ProposalV1Controller{Controller: service.NewController("proposal v1")}
}

// Create runs the create action.
func (c *ProposalV1Controller) Create(ctx *v1.CreateProposalContext) error {
	return nil
}

// Delete runs the delete action.
func (c *ProposalV1Controller) Delete(ctx *v1.DeleteProposalContext) error {
	return nil
}

// List runs the list action.
func (c *ProposalV1Controller) List(ctx *v1.ListProposalContext) error {
	res := v1.ProposalCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *ProposalV1Controller) Show(ctx *v1.ShowProposalContext) error {
	res := &v1.Proposal{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *ProposalV1Controller) Update(ctx *v1.UpdateProposalContext) error {
	return nil
}
