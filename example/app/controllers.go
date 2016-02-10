//************************************************************************//
// API "congo": Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

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
	// Setup encoders and decoders. This is idempotent and is done by each MountXXX function.
	service.SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.ServeMux()
	h = func(c *goa.Context) error {
		ctx, err := NewCallbackAuthContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Callback(ctx)
	}
	mux.Handle("GET", "/api/auth/:provider/callback", ctrl.HandleFunc("Callback", h, nil))
	service.Info("mount", "ctrl", "Auth", "action", "Callback", "route", "GET /api/auth/:provider/callback")
	h = func(c *goa.Context) error {
		ctx, err := NewOauthAuthContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Oauth(ctx)
	}
	mux.Handle("GET", "/api/auth/:provider", ctrl.HandleFunc("Oauth", h, nil))
	service.Info("mount", "ctrl", "Auth", "action", "Oauth", "route", "GET /api/auth/:provider")
	h = func(c *goa.Context) error {
		ctx, err := NewRefreshAuthContext(c)
		if rawPayload := ctx.RawPayload(); rawPayload != nil {
			ctx.Payload = rawPayload.(*RefreshAuthPayload)
		}
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Refresh(ctx)
	}
	mux.Handle("POST", "/api/auth/refresh", ctrl.HandleFunc("Refresh", h, unmarshalRefreshAuthPayload))
	service.Info("mount", "ctrl", "Auth", "action", "Refresh", "route", "POST /api/auth/refresh")
	h = func(c *goa.Context) error {
		ctx, err := NewTokenAuthContext(c)
		if rawPayload := ctx.RawPayload(); rawPayload != nil {
			ctx.Payload = rawPayload.(*TokenAuthPayload)
		}
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Token(ctx)
	}
	mux.Handle("POST", "/api/auth/token", ctrl.HandleFunc("Token", h, unmarshalTokenAuthPayload))
	service.Info("mount", "ctrl", "Auth", "action", "Token", "route", "POST /api/auth/token")
}

// unmarshalRefreshAuthPayload unmarshals the request body.
func unmarshalRefreshAuthPayload(ctx *goa.Context) error {
	payload := &RefreshAuthPayload{}
	if err := ctx.Service().DecodeRequest(ctx, payload); err != nil {
		return err
	}
	ctx.SetPayload(payload)
	return nil
}

// unmarshalTokenAuthPayload unmarshals the request body.
func unmarshalTokenAuthPayload(ctx *goa.Context) error {
	payload := &TokenAuthPayload{}
	if err := ctx.Service().DecodeRequest(ctx, payload); err != nil {
		return err
	}
	ctx.SetPayload(payload)
	return nil
}

// UiController is the controller interface for the Ui actions.
type UiController interface {
	goa.Controller
	Bootstrap(*BootstrapUiContext) error
}

// MountUiController "mounts" a Ui resource controller on the given service.
func MountUiController(service goa.Service, ctrl UiController) {
	// Setup encoders and decoders. This is idempotent and is done by each MountXXX function.
	service.SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.ServeMux()
	h = func(c *goa.Context) error {
		ctx, err := NewBootstrapUiContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Bootstrap(ctx)
	}
	mux.Handle("GET", "/", ctrl.HandleFunc("Bootstrap", h, nil))
	service.Info("mount", "ctrl", "Ui", "action", "Bootstrap", "route", "GET /")
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Controller
	Create(*CreateUserContext) error
	Delete(*DeleteUserContext) error
	List(*ListUserContext) error
	Show(*ShowUserContext) error
	Update(*UpdateUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service goa.Service, ctrl UserController) {
	// Setup encoders and decoders. This is idempotent and is done by each MountXXX function.
	service.SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.ServeMux()
	h = func(c *goa.Context) error {
		ctx, err := NewCreateUserContext(c)
		if rawPayload := ctx.RawPayload(); rawPayload != nil {
			ctx.Payload = rawPayload.(*CreateUserPayload)
		}
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	mux.Handle("POST", "/api/users", ctrl.HandleFunc("Create", h, unmarshalCreateUserPayload))
	service.Info("mount", "ctrl", "User", "action", "Create", "route", "POST /api/users")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	mux.Handle("DELETE", "/api/users/:userID", ctrl.HandleFunc("Delete", h, nil))
	service.Info("mount", "ctrl", "User", "action", "Delete", "route", "DELETE /api/users/:userID")
	h = func(c *goa.Context) error {
		ctx, err := NewListUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	mux.Handle("GET", "/api/users", ctrl.HandleFunc("List", h, nil))
	service.Info("mount", "ctrl", "User", "action", "List", "route", "GET /api/users")
	h = func(c *goa.Context) error {
		ctx, err := NewShowUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	mux.Handle("GET", "/api/users/:userID", ctrl.HandleFunc("Show", h, nil))
	service.Info("mount", "ctrl", "User", "action", "Show", "route", "GET /api/users/:userID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateUserContext(c)
		if rawPayload := ctx.RawPayload(); rawPayload != nil {
			ctx.Payload = rawPayload.(*UpdateUserPayload)
		}
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	mux.Handle("PATCH", "/api/users/:userID", ctrl.HandleFunc("Update", h, unmarshalUpdateUserPayload))
	service.Info("mount", "ctrl", "User", "action", "Update", "route", "PATCH /api/users/:userID")
}

// unmarshalCreateUserPayload unmarshals the request body.
func unmarshalCreateUserPayload(ctx *goa.Context) error {
	payload := &CreateUserPayload{}
	if err := ctx.Service().DecodeRequest(ctx, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	ctx.SetPayload(payload)
	return nil
}

// unmarshalUpdateUserPayload unmarshals the request body.
func unmarshalUpdateUserPayload(ctx *goa.Context) error {
	payload := &UpdateUserPayload{}
	if err := ctx.Service().DecodeRequest(ctx, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	ctx.SetPayload(payload)
	return nil
}
