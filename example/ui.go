package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma/example/app"
)

// UIController implements theui resource.
type UIController struct {
	*goa.Controller
}

// NewUIController creates a ui controller.
func NewUIController(service *goa.Service) app.UIController {
	return &UIController{Controller: service.NewController("ui")}
}

// Bootstrap runs the bootstrap action.
func (c *UIController) Bootstrap(ctx *app.BootstrapUIContext) error {
	return nil
}
