package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma/example/app"
)

// UserController implements theuser resource.
type UserController struct {
	goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service goa.Service) app.UserController {
	return &UserController{Controller: service.NewController("user")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	return nil
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	return nil
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	users := udb.ListAppUser(ctx.Context)
	return ctx.OK(users)
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	user := udb.OneUser(ctx.Context, ctx.UserID)
	return ctx.OK(user)
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	return nil
}
