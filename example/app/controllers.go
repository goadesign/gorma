//************************************************************************//
// API "congo": Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/bketelsen/gorma/example
// --design=github.com/bketelsen/gorma/example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/raphael/goa"

// AuthController is the controller interface for the Auth actions.
type AuthController interface {
	goa.Controller
	Callback(*CallbackAuthContext) error
	Oauth(*OauthAuthContext) error
	Refresh(*RefreshAuthContext) error
	Token(*TokenAuthContext) error
}

// MountAuthController "mounts" a Auth resource controller on the given service.
func MountAuthController(service goa.Service, ctrl AuthController) {
	var h goa.Handler
	mux := service.ServeMux()
	h = func(c *goa.Context) error {
		ctx, err := NewCallbackAuthContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Callback(ctx)
	}
	mux.Handle("GET", "/auth/:provider/callback", ctrl.HandleFunc("Callback", h))
	service.Info("mount", "ctrl", "Auth", "action", "Callback", "route", "GET /auth/:provider/callback")
	h = func(c *goa.Context) error {
		ctx, err := NewOauthAuthContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Oauth(ctx)
	}
	mux.Handle("GET", "/auth/:provider", ctrl.HandleFunc("Oauth", h))
	service.Info("mount", "ctrl", "Auth", "action", "Oauth", "route", "GET /auth/:provider")
	h = func(c *goa.Context) error {
		ctx, err := NewRefreshAuthContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Refresh(ctx)
	}
	mux.Handle("POST", "/auth/refresh", ctrl.HandleFunc("Refresh", h))
	service.Info("mount", "ctrl", "Auth", "action", "Refresh", "route", "POST /auth/refresh")
	h = func(c *goa.Context) error {
		ctx, err := NewTokenAuthContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Token(ctx)
	}
	mux.Handle("POST", "/auth/token", ctrl.HandleFunc("Token", h))
	service.Info("mount", "ctrl", "Auth", "action", "Token", "route", "POST /auth/token")
}

// UiController is the controller interface for the Ui actions.
type UiController interface {
	goa.Controller
	Bootstrap(*BootstrapUiContext) error
}

// MountUiController "mounts" a Ui resource controller on the given service.
func MountUiController(service goa.Service, ctrl UiController) {
	var h goa.Handler
	mux := service.ServeMux()
	h = func(c *goa.Context) error {
		ctx, err := NewBootstrapUiContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Bootstrap(ctx)
	}
	mux.Handle("GET", "/", ctrl.HandleFunc("Bootstrap", h))
	service.Info("mount", "ctrl", "Ui", "action", "Bootstrap", "route", "GET /")
}
