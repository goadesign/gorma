package main

import (
	"github.com/bketelsen/gorma/example/app"
	"github.com/bketelsen/gorma/example/swagger"
	"github.com/goadesign/goa"
	"github.com/goadesign/middleware"
)

func main() {
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest())
	service.Use(middleware.Recover())

	// Mount "account" controller
	c := NewAccountController(service)
	app.MountAccountController(service, c)
	// Mount "bottle" controller
	c2 := NewBottleController(service)
	app.MountBottleController(service, c2)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	// Start service, listen on port 8080
	service.ListenAndServe(":8080")
}
