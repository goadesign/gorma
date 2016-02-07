package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma/example/app"
)

// UiController implements theui resource.
type UiController struct {
	goa.Controller
}

// NewUiController creates a ui controller.
func NewUiController(service goa.Service) app.UiController {
	return &UiController{Controller: service.NewController("ui")}
}

// Bootstrap runs the bootstrap action.
func (c *UiController) Bootstrap(ctx *app.BootstrapUiContext) error {
	return nil
}
