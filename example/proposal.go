package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma/example/app"
	"github.com/jinzhu/gorm"
)

// ProposalController implements the proposal resource.
type ProposalController struct {
	*goa.Controller
}

// NewProposalController creates a proposal controller.
func NewProposalController(service *goa.Service) app.ProposalController {
	return &ProposalController{Controller: service.NewController("proposal")}
}

// Create runs the create action.
func (c *ProposalController) Create(ctx *app.CreateProposalContext) error {
	return nil
}

// Delete runs the delete action.
func (c *ProposalController) Delete(ctx *app.DeleteProposalContext) error {
	return nil
}

// List runs the list action.
func (c *ProposalController) List(ctx *app.ListProposalContext) error {
	proposals := pdb.ListProposal(ctx.Context, ctx.UserID)
	return ctx.OK(proposals)
}

// Show runs the show action.
func (c *ProposalController) Show(ctx *app.ShowProposalContext) error {
	proposal, err := pdb.OneProposal(ctx.Context, ctx.ProposalID, ctx.UserID)
	if err ==  gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return goa.Response(ctx).Send(ctx, 500, err.Error)

	}
	return ctx.OK(proposal)
}

// Update runs the update action.
func (c *ProposalController) Update(ctx *app.UpdateProposalContext) error {
	return nil
}
