package main

import (
	"github.com/goadesign/gorma/example/app"
	"github.com/goadesign/goa"
)

// AccountController implements theaccount resource.
type AccountController struct {
	goa.Controller
}

// NewAccountController creates a account controller.
func NewAccountController(service goa.Service) app.AccountController {
	return &AccountController{Controller: service.NewController("account")}
}

// Create runs the create action.
func (c *AccountController) Create(ctx *app.CreateAccountContext) error {
	return nil
}

// Delete runs the delete action.
func (c *AccountController) Delete(ctx *app.DeleteAccountContext) error {
	return nil
}

// Show runs the show action.
func (c *AccountController) Show(ctx *app.ShowAccountContext) error {
	res := &app.Account{}
	return ctx.OK(res, "default")
}

// Update runs the update action.
func (c *AccountController) Update(ctx *app.UpdateAccountContext) error {
	return nil
}
